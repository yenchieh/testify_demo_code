package main

import (
	"bowery-golang_testing/route"

	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	route.SetupRoute(e)

	e.Logger.Fatal(e.Start(":1232"))
}

func addItUp(number int) int {
	var count int
	for i := 0; i <= number; i++ {
		count += i
	}

	return count
}

func looping() int {
	number := 1
	for i := 1; i < 100000; i++ {
		number += i
	}

	return number
}

type Hello interface {
	YourName(name string) (string, error)
}

func YourName(name string) (string, error) {
	return fmt.Sprintf("%s %s", "Hello", name), nil
}
