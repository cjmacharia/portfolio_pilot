package handlers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/cjmacharia/portfolioPilot/internal/app/service"
	"github.com/cjmacharia/portfolioPilot/internal/domain"
	"github.com/cjmacharia/portfolioPilot/internal/dto"
	"github.com/cjmacharia/portfolioPilot/internal/utils"
	"github.com/gocolly/colly"
	"net/http"
	"os"
	"strconv"
	"time"
)

type StockHandler struct {
	stockService service.StockService
}

func (sh *StockHandler) PostStocksHandler(w http.ResponseWriter, r *http.Request) {
	stocks, err := scrapNasdaq()
	if err != nil {
		http.Error(w, utils.ErrInvalidPayload.Error(), http.StatusBadRequest)
		return
	}

	for _, stock := range stocks {
		_, err = sh.stockService.PostStock(&stock)
		if err != nil {
			http.Error(w, utils.ErrInvalidStockID.Error(), http.StatusInternalServerError)
			return
		}
	}
}
func (sh *StockHandler) getAllStocks(w http.ResponseWriter, r *http.Request) {
	s, err := sh.stockService.GetAllStocks()
	if err != nil {
		fmt.Println(err, "post static")
		http.Error(w, utils.ErrInternalServerError.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(s)
}

func (sh *StockHandler) UpdateStockPrices(w http.ResponseWriter, r *http.Request) {
	s, err := sh.stockService.GetAllStocks()
	stocks, err := scrapNasdaq()
	if err != nil {
		http.Error(w, utils.ErrInvalidPayload.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, utils.ErrInvalidPayload.Error(), http.StatusBadRequest)
		return
	}
	dbStocksMap := make(map[string]dto.StockResponse)
	for _, dbStock := range s {
		dbStocksMap[dbStock.Symbol] = dbStock
	}
	for _, scrapedStock := range stocks {
		dbStock, ok := dbStocksMap[scrapedStock.Symbol]
		if !ok {
			// If the stock is not found in the database, continue to the next one
			continue
		}
		if scrapedStock.Price != dbStock.Price {
			// If the price has changed, update the database
			err = sh.stockService.UpdateStockPrice(scrapedStock.Symbol, scrapedStock.Price)
			if err != nil {
				http.Error(w, utils.ErrInternalServerError.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func scrapNasdaq() ([]domain.Stock, error) {
	httpClient := &http.Client{
		Timeout: 100 * time.Minute,
		Transport: &http.Transport{
			// Disable HTTP/2
			TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
		},
	}
	c := colly.NewCollector(colly.MaxDepth(1), colly.DetectCharset())
	c.WithTransport(httpClient.Transport)
	var stocks []domain.Stock
	c.OnHTML("div[class='t'] > table", func(e *colly.HTMLElement) {
		e.ForEach("tbody tr", func(_ int, row *colly.HTMLElement) {

			priceStr := row.ChildText("td:nth-child(4)") // Get the price as a string
			price, _ := strconv.ParseFloat(priceStr, 64)
			stock := domain.Stock{
				Name:     row.ChildText("td:nth-child(2)"),
				Symbol:   row.ChildText("td:nth-child(1)"),
				Exchange: "Nairobi Securities Exchange",
				Price:    price,
			}
			stocks = append(stocks, stock)
		})
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnError(func(resp *colly.Response, err error) {
		url := resp.Request.URL.String()
		fmt.Fprintf(
			os.Stdout, "ERR on URL: %s (from: %s), error: %s\n", url,
			resp.Request.Ctx.Get("Referrer"), err,
		)

	})
	err := c.Visit("https://afx.kwayisi.org/nse/")
	if err != nil {
		return nil, err
	}
	return stocks, nil
}
