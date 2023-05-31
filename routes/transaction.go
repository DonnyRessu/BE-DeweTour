package routes

import (
	"week2/handlers"
	"week2/pkg/middleware"
	"week2/pkg/mysql"
	"week2/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	e.GET("/transactions", h.FindTransactions)
	e.GET("/transaction/:id", h.GetTransaction)
	e.POST("/transaction", middleware.Auth(middleware.UploadFile(h.CreateTransaction)))
	e.PATCH("/transaction/:id", middleware.Auth(middleware.UploadFile(h.UpdateTransaction)))
}
