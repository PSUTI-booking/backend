package main

import (
	"booking/config"
	"booking/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})

	// Connect To Database
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	bookRoute := e.Group("/book")
	bookRoute.GET("/:id", controller.GetBook)
	bookRoute.PUT("/:id", controller.UpdateBook)

	e.Logger.Fatal(e.Start(":8080"))
}
