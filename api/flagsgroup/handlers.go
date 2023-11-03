package flagsgroup

import (
	"fmt"
	"net/http"

	"github.com/chtushar/toggler/api/app"
	"github.com/chtushar/toggler/api/responses"
	"github.com/chtushar/toggler/db"
	"github.com/chtushar/toggler/db/queries"
	"github.com/chtushar/toggler/utils"
	"github.com/jackc/pgtype"
	"github.com/labstack/echo/v4"
)

func handleCreateFlagsGroup(c echo.Context) error {
	var (
		app     = c.Get("app").(*app.App)
		orgUUID = c.Param("orgUUID")
		req     = &struct {
			Name       string `json:"name" validate:"required,min=3"`
			FolderUUID string `json:"folder_uuid" validate:"required"`
		}{}
	)
	_, err := utils.IsValidUUID(orgUUID)
	if err != nil {
		app.Log.Println("Unable to parse the org uuid")
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	fg, err := db.WithDBTransaction[queries.FlagsGroup](app, c.Request().Context(), func(q *queries.Queries) (*queries.FlagsGroup, error) {
		org, err := q.GetOrganizationByUUID(c.Request().Context(), orgUUID)

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		folder, err := q.GetFolderByUUID(c.Request().Context(), req.FolderUUID)

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		fg, err := q.CreateFlagsGroup(c.Request().Context(), queries.CreateFlagsGroupParams{
			Name:     req.Name,
			OrgID:    org.ID,
			FolderID: folder.ID,
		})

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		return &fg, nil
	})

	if err != nil {
		return err
	}

	db.WithDBTransaction[bool](app, c.Request().Context(), func(q *queries.Queries) (*bool, error) {
		envs, err := q.GetOrganizationEnvironments(c.Request().Context(), fg.OrgID)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		var jsonbValue pgtype.JSONB

		err = jsonbValue.Set("{}")

		if err != nil {
			fmt.Println("Error:", err)
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		for _, e := range envs {
			_, err := q.CreateFlagsGroupState(c.Request().Context(), queries.CreateFlagsGroupStateParams{
				FlagsGroupID:  fg.ID,
				EnvironmentID: e.ID,
				Json:          jsonbValue,
			})

			if err != nil {
				app.Log.Println(err)
				return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
			}
		}

		ok := true
		return &ok, nil
	})

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data:    fg,
		Error:   nil,
	})
	return nil
}

func handleGetFlagsGroup(c echo.Context) error {
	return nil
}

func handleDeleteFlagsGroup(c echo.Context) error {
	return nil
}
