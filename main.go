package main

import (
	"estate/controllers"
	"estate/storage"
	"net/http"

	// "github.com/labstack/echo/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Connect To Database
	storage.DatabaseInit()
	gorm := storage.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	g := e.Group("/api")
	// g.GET("/estates", func(c echo.Context) error {
	// 	u := new(Estate)
	// 	if err := c.Bind(u); err != nil {
	// 		return err
	// 	}
	// 	return c.JSON(http.StatusCreated, u)
	// 	// or
	// 	// return c.XML(http.StatusCreated, u)
	// })
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	g.GET("/estates", controllers.GetEstates)

	e.Logger.Fatal(e.Start(":8181"))
}
