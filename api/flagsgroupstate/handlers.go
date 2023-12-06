package flagsgroupstate

import (
	"net/http"

	"github.com/chtushar/toggler/api/app"
	"github.com/chtushar/toggler/api/responses"
	"github.com/chtushar/toggler/db"
	"github.com/chtushar/toggler/db/queries"
	"github.com/chtushar/toggler/utils"
	"github.com/labstack/echo/v4"
)

func handleGetFlagsGroupState(c echo.Context) error {
	var (
		app     = c.Get("app").(*app.App)
		fgUUID  = c.Param("fgUUID")
	)
	ok, err := utils.IsValidUUID(fgUUID)

	if !ok {
		app.Log.Println("Failed to parse flags group uuid", err)
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	fgs, err := db.WithDBTransaction[queries.FlagsGroupState](app, c.Request().Context(), func(q *queries.Queries) (*queries.FlagsGroupState, error) {
		fg, err := q.GetFlagsGroupByUUID(c.Request().Context(), fgUUID)

		if err != nil {
			app.Log.Println(" Couldn't get the flags group", err)
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		fgs, err := q.GetFlagsGroupState(c.Request().Context(), fg.ID)

		if err != nil {
			app.Log.Println(" Couldn't get the Flags group state", err)
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		return &fgs, nil
	})

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data:    fgs,
		Error:   nil,
	})
	return nil
}

func handleUpdateFlagsGroupStateJS(c echo.Context) error {
	var (
		app    = c.Get("app").(*app.App)
		fgUUID = c.Param("fgUUID")
		req    = &struct {
			Value   string `json:"value"`
		}{}
	)

	ok, err := utils.IsValidUUID(fgUUID)

	if !ok {
		app.Log.Println("Failed to parse flags group uuid", err)
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	if err := c.Bind(req); err != nil {
		app.Log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}
	
	if err := c.Validate(req); err != nil {
		app.Log.Println("Failed to validate request body", err)
		return err
	}

	fgs, err := db.WithDBTransaction[queries.FlagsGroupState](app, c.Request().Context(), func(q *queries.Queries) (*queries.FlagsGroupState, error) {
		fg, err := q.GetFlagsGroupByUUID(c.Request().Context(), fgUUID)

		if err != nil {
			app.Log.Println("Couldn't get the flag group", err)
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		fgs, err := q.UpdateFlagGroupsStateJS(c.Request().Context(), queries.UpdateFlagGroupsStateJSParams{
			FlagsGroupID:  fg.ID,
			Js: &req.Value,
		})

		if err != nil {
			app.Log.Println("Couldn't set the flag groups state js code", err)
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		return &fgs, nil
	})

	if err != nil {
		return nil
	}

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data:    fgs,
		Error:   nil,
	})
	return nil
}
