package cmd

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func handleGetProjectEnvironments (c echo.Context) error {
	var (
		app  = c.Get("app").(*App)
	)

	projectIdParam := c.Param("projectId")
	projectId, err := strconv.Atoi(projectIdParam)

	if err != nil {
		app.log.Println(err)
		c.JSON(http.StatusBadRequest, BadRequestResponse)
		return err
	}

	envs, err := app.q.GetProjectEnvironments(c.Request().Context(), int64(projectId))

	if err != nil {
		app.log.Println("Failed to get project environments")
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	c.JSON(http.StatusOK, responseType{true, envs, nil})

	return nil
}