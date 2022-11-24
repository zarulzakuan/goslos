package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		RunDiskIO("1G", "2G", "10G", 100, 50)
		RunCPULoad(8, 60, 70, 100)
		RunMemLoad(60, "5G")
		return c.String(http.StatusOK, "Hello, World!")
	})
	// RunCPULoad(0, 60, 0, 0)
	RunDiskIO("1M", "20M", "1s", 100, 75)

	e.Logger.Fatal(e.Start(":8080"))
}
