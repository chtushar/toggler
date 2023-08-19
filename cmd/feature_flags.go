package cmd

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/chtushar/toggler/db/queries"
	"github.com/chtushar/toggler/utils"
	"github.com/jackc/pgtype"
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

	envs, err := qtx.GetProjectEnvironments(c.Request().Context(), int64(req.ProjectId))
	
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

func handleGetProjectFeatureFlags (c echo.Context) error {
	var (
		app  = c.Get("app").(*App)
	)

	type resType struct {
		ID        int32           `json:"id"`
		Uuid      string   `json:"uuid"`
		Name      string          `json:"name"`
		FlagType  queries.FeatureFlagType `json:"flag_type"`
		Enabled   bool    `json:"enabled"`
		Value     pgtype.JSONB    `json:"value"`
		UpdatedAt string    `json:"updated_at"`
	}

	projectIdParam := c.Param("projectId")
	environmentIdParam := c.Param("environmentId")
	
	projectId, err := strconv.Atoi(projectIdParam)
	
	if err != nil {
		app.log.Println("Failed to parse project id", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err	
	}

	environmentId, err := strconv.Atoi(environmentIdParam)

	if err != nil {
		app.log.Println("Failed to parse project id", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
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

	ffs, err := qtx.GetProjectFeatureFlags(c.Request().Context(), queries.GetProjectFeatureFlagsParams{
		ProjectID: int64(projectId),
		EnvironmentID: int64(environmentId),
	})

	if err != nil {
		app.log.Println("Failed to get project feature flags", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	tx.Commit(c.Request().Context())

	res := make([]resType, 0)

	for _, f := range ffs {
		res = append(res, resType{
			ID: f.ID,
			Uuid: f.Uuid,
			Name: f.Name,
			FlagType: f.FlagType,
			Enabled: *f.Enabled,
			Value: f.Value,
			UpdatedAt: f.UpdatedAt.String(),
		})
	}

	c.JSON(http.StatusOK, responseType{
		Success: true,
		Data: res,
		Error: nil,
	})

	return nil
}

func handleGetFlags (c echo.Context) error {
	var (
		app  = c.Get("app").(*App)
	)

	projectUuid := c.Request().Header.Get("X-project-uuid")
	apiKey := c.Request().Header.Get("X-api-key")

	if apiKey == "" || projectUuid == "" {
		app.log.Println("Check for the headers")
		c.JSON(http.StatusBadRequest, BadRequestResponse)
		return fmt.Errorf("there was an error")	
	}

	// tx, err := app.dbConn.Begin(c.Request().Context())

	// if err != nil {
	// 	app.log.Println("Failed to create tx", err)
	// 	c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
	// 	return err
	// }

	// defer tx.Rollback(c.Request().Context())

	// qtx := app.q.WithTx(tx)

	ffs, err := app.q.GetFeatureFlags(c.Request().Context(), queries.GetFeatureFlagsParams{
		Uuid: projectUuid,
		Column2: apiKey,
	})

	if err != nil {
		app.log.Println("Failed to get the feature flags", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	// tx.Commit(c.Request().Context())

	c.JSON(http.StatusOK, responseType{
		Success: true,
		Data: ffs,
		Error: nil,
	})
	return nil
}

func handleToggleFeatureFlag (c echo.Context) error {
	var (
		app  = c.Get("app").(*App)
	)
	ffIdParam := c.Param("ffId")
	ffId, err := strconv.Atoi(ffIdParam)

	if err != nil {
		app.log.Println(err)
		c.JSON(http.StatusBadRequest, BadRequestResponse)
		return err
	}

	ff, err := app.q.ToggleFeatureFlag(c.Request().Context(), int64(ffId))

	if err != nil {
		app.log.Println("Failed to toggle the feature flag", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	c.JSON(http.StatusOK, responseType{
		Success: true,
		Data: ff,
		Error: nil,
	})

	return nil
}