package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	e := echo.New()
	e.GET("/:num", func(c echo.Context) error {
		num := c.Param("num")
		num1, _ := strconv.Atoi(num)
		f := strconv.Itoa(fib(num1))
		return c.String(http.StatusOK, f)
	})

	e.Start(":")
}
