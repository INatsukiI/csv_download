package main

import (
	"bytes"
	"encoding/csv"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {
	e.GET("/download_csv", downloadCsv)

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())

	return e
}

func downloadCsv(c echo.Context) error {
	csvData := [][]string{
		{"test1", "30", "aaa"},
		{"test2", "25", "bbb"},
		{"test3", "35", "ccc"},
	}
	csvBuffer := new(bytes.Buffer)
	csvWriter := csv.NewWriter(csvBuffer)
	csvWriter.WriteAll(csvData)

	if csvWriter.Error() != nil {
		return c.String(http.StatusInternalServerError, "Failed to write CSV data")
	}
	return c.Blob(http.StatusOK, "text/csv", csvBuffer.Bytes())
}