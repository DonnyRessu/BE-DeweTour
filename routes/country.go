package routes

import (
	"week2/handlers"
	"week2/pkg/middleware"
	"week2/pkg/mysql"
	"week2/repositories"

	"github.com/labstack/echo/v4"
)

func CountryRoutes(e *echo.Group) {
	countryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handlers.HandlersCountry(countryRepository)

	e.GET("/countries", h.FindCountry)
	e.GET("/country/:id", h.GetCountry)
	e.POST("/country", middleware.Auth(h.CreateCountry))
	e.PATCH("/country/:id", middleware.Auth(h.UpdateCountry))
	e.DELETE("/country/:id", middleware.Auth(h.DeleteCountry))
}
