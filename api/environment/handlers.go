package environment

import (
	"net/http"

	"github.com/chtushar/toggler/api/app"
	"github.com/chtushar/toggler/api/responses"
	"github.com/chtushar/toggler/db"
	"github.com/chtushar/toggler/db/queries"
	"github.com/chtushar/toggler/utils"
	"github.com/labstack/echo/v4"
)

func handleCreateEnvironments (c echo.Context) error {
	type Env struct {
		Name string `json:"name" validate:"required,min=3"`
		Color string `json:"color" validate:"required,min=7,max=7"`
	}
	var (
		app = c.Get("app").(*app.App)
		orgUUID = c.Param("orgUUID")
		req = &struct{
			Envs []Env `json:"envs" validate:"required,dive"`
		}{}
	)

	ok, err := utils.IsValidUUID(orgUUID)

	if !ok {
		app.Log.Println("Can't parse the org uuid", err)
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	db.WithDBTransaction[bool](app, c.Request().Context(), func(q *queries.Queries) (*bool, error) {
		org, err := q.GetOrganizationByUUID(c.Request().Context(), orgUUID)

		if err != nil {
			app.Log.Println("Couldn't get the org", err)
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		for _, e := range req.Envs {
			err = q.CreateEnvironment(c.Request().Context(), queries.CreateEnvironmentParams{
				OrgID: org.ID,
				Name: e.Name,
				Color: &e.Color,
			})

			if err != nil {
				return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
			}
		}

		ok := true
		return &ok, err
	})

	c.JSON(http.StatusOK, responses.ResponseType{Success: true, Data: "Envs Added", Error: nil})
	return nil
}

func handleGetEnvironments (c echo.Context) error {
	var (
		app = c.Get("app").(*app.App)
		orgUUID = c.Param("orgUUID")
	)

	ok, err := utils.IsValidUUID(orgUUID)

	if !ok {
		app.Log.Println("Failed to get Organization", err)
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	org, err := app.Q.GetOrganizationByUUID(c.Request().Context(), orgUUID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	envs, err := app.Q.GetOrganizationEnvironments(c.Request().Context(), org.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data: envs,
		Error: nil,
	})
	return nil
}

func handleUpdateEnvironment (c echo.Context) error {
	return nil
}