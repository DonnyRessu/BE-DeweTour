package routes

import (
	"week2/handlers"
	"week2/pkg/middleware"
	"week2/pkg/mysql"
	"week2/repositories"

	"github.com/labstack/echo/v4"
)

func TripRoutes(e *echo.Group) {
	tripRepository := repositories.RepositoryTrip(mysql.DB)
	h := handlers.HandlerTrip(tripRepository)

	e.GET("/trips", h.FindTrip)
	e.GET("/trip/:id", h.GetTrip)
	e.POST("/trip", middleware.Auth(middleware.UploadFile(h.CreateTrip)))
	e.PATCH("/trip/:id", middleware.Auth(middleware.UploadFile(h.UpdateTrip)))
	e.DELETE("/trip/:id", middleware.Auth(h.DeleteTrip))
}