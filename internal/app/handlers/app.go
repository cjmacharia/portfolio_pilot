package handlers

import (
	"fmt"
	"github.com/cjmacharia/portfolioPilot/internal/app/service"
	"github.com/cjmacharia/portfolioPilot/internal/infra/database"
	"github.com/cjmacharia/portfolioPilot/internal/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitApp() {
	dbClient := database.GetDbClient()
	u := UserHandler{
		UserService: service.NewUserService(database.NewUserRepositoryDB(dbClient)),
	}
	s := StockHandler{stockService: service.NewStockService(database.NewStockRepositoryDB(dbClient))}
	t := TransactionsHandler{
		transactionService: service.NewTransactionService(database.NewTransactionRepositoryDB(dbClient)),
		stockService:       s.stockService,
		userService:        u.UserService,
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/stock/{id:[0-9]+}/transaction", middlewares.AuthMiddleware(t.HandleTransaction)).Methods("POST")
	router.HandleFunc("/api/stock", s.PostStocksHandler).Methods("POST")
	router.HandleFunc("/api/userSignup", u.UserSignUpHandler).Methods("POST")
	router.HandleFunc("/api/login", u.LoginHandler).Methods("POST")
	router.HandleFunc("/api/user/{id:[0-9]+}/portfolio", t.GetUserPortfolio).Methods("GET")

	http.Handle("/", router)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("server not started on 8080", err)
		return
	}
	fmt.Println("server started on 8080")

}
