package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func getMoviesHandler(c echo.Context) error {
	return c.JSON(200, []string{"movie1", "movie2"})
}

func main() {
	e := echo.New()

	e.GET("/movies", getMoviesHandler)

	port := ":8080"
	log.Fatal(e.Start(port))
}
