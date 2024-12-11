package config

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
	"backend/internal/repository"
	"net/http"
)

func CreateRouter(queries *repository.Queries) *http.ServeMux {
	router := http.NewServeMux()
	userHandler := handlers.NewUserHandler(queries)
	ttHandler := handlers.NewTransactionTypesHandler(queries)
	tHandler := handlers.NewTransactionsHandler(queries)

	router.HandleFunc("/", handlers.GetStatus)

	// user
	router.HandleFunc("POST /user/register", userHandler.CreateUser)
	router.HandleFunc("POST /user/login", userHandler.LoginUser)
	router.HandleFunc("GET /user/{id}", userHandler.GetUserById)
	router.HandleFunc("POST /user", userHandler.CreateUser)

	// transaction types
	router.HandleFunc("GET /transaction-type/{id}", middleware.IsAuthenticated(ttHandler.GetTransactionTypeById))
	router.HandleFunc("POST /transaction-type", middleware.IsAuthenticated(ttHandler.CreateTransactionType))

	// transactions
	router.HandleFunc("GET /transaction/{id}", middleware.IsAuthenticated(tHandler.GetTransactionById))
	router.HandleFunc("POST /transaction", middleware.IsAuthenticated(tHandler.CreateTransaction))
	router.HandleFunc("PUT /transaction", middleware.IsAuthenticated(tHandler.UpdateTransaction))
	router.HandleFunc("DELETE /transaction/{id}", middleware.IsAuthenticated(tHandler.DeleteTransaction))

	// user transactions
	router.HandleFunc("GET /user/transactions", middleware.IsAuthenticated(tHandler.GetTransactionsForUser))

	return router
}
