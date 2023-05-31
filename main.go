package main

import (
	"fmt"
	"week2/database"
	"week2/pkg/mysql"
	"week2/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mysql.DatabaseInit()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))
	e.Static("/uploads", "/uploads")

	fmt.Println("server running on localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
