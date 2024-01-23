package handlers

import (
	"fmt"
	"github.com/cjmacharia/portfolioPilot/internal/app/service"
	"github.com/cjmacharia/portfolioPilot/internal/infra/database"
	"net/http"

	"github.com/gorilla/mux"
)

func InitApp() {
	dbClient := database.GetDbClient()
	fmt.Println(dbClient, ">vsd>>>")

	fmt.Println("server st")

	sh := StockHandler{service: service.NewStockService(database.NewStockRepositoryDB(dbClient))}
	fmt.Println("server start")

	router := mux.NewRouter()
	router.HandleFunc("/api/stock", sh.PostStocksHandler).Methods("POST")
	http.Handle("/", router)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("server not started on 8080", err)
		return
	}
	fmt.Println("server started on 8080")
}
