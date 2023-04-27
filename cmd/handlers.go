package cmd

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type okResp struct {
	Data interface{} `json:"data"`
}

// handleHealthCheck is a healthcheck endpoint that returns a 200 response.
func handleHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, okResp{true})
}

func initHTTPHandler(e *echo.Echo, app *App) {
	fmt.Println("Initializing HTTP handlers")
	// ...
	e.GET("/api/healthcheck", handleHealthCheck)

	fmt.Println("Initialized HTTP handlers")
}
