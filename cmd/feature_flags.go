package cmd

import (
	"net/http"

	"github.com/chtushar/toggler/db/queries"
	"github.com/chtushar/toggler/utils"
	"github.com/labstack/echo/v4"
)

func handleCreateFeatureFlag (c echo.Context) error {
	var (
		app  = c.Get("app").(*App)
		req = &struct {
			Name string `json:"name"`;
			ProjectId int32 `json:"project_id"`;
			FlagType queries.FeatureFlagType `json:"flag_type"`
		}{}
	)

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, responseType{
			Success: false,
			Data: nil,
			Error: &errorWrap{
				Code: http.StatusBadGateway,
				Data: nil,
				Message: "Bad Request. Please check the payload.",	
			},
		})
		return err
	}

	tx, err := app.dbConn.Begin(c.Request().Context())

	if err != nil {
		app.log.Println("Failed to create tx", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	defer tx.Rollback(c.Request().Context())

	qtx := app.q.WithTx(tx)

	ff, err := qtx.CreateFeatureFlag(c.Request().Context(), queries.CreateFeatureFlagParams{
		Name: req.Name,
		ProjectID: int64(req.ProjectId),
		FlagType: req.FlagType,
	})

	if err != nil {
		app.log.Println("Failed to create feature flag", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	envs, err := qtx.GetProjectEnviornments(c.Request().Context(), int64(req.ProjectId))
	
	if err != nil {
		app.log.Println("Failed to get project envs", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	jsonb := utils.CreateEmptyJSONB()

	for _, env := range envs {
		_, err := qtx.CreateFeatureState(c.Request().Context(), queries.CreateFeatureStateParams{
			FeatureFlagID: int64(ff.ID),
			EnvironmentID: int64(env.ID),
			Enabled: false,
			Value: jsonb,
		})

		if err != nil {
			app.log.Println("Failed to create create feature flag state", err)
			c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
			return err
		}
	}

	tx.Commit(c.Request().Context())

	c.JSON(http.StatusOK, responseType{true, ff, nil})
	return nil
}