package main

import (
	"net/http"
	"strconv"

	// go get으로 설치 후 import
	"github.com/jdk829355/learn_golang/myprime"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/:number", func(c echo.Context) error {
		nstr := c.Param("number")
		n, err := strconv.Atoi(nstr)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusOK, strconv.FormatBool(myprime.IsPrime(n)))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
