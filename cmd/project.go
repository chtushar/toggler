package cmd

import (
	"net/http"

	"github.com/chtushar/toggler/db/queries"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func handleCreateProject(c echo.Context) error {
	var (
		app  = c.Get("app").(*App)
		user = c.Get("user").(*jwt.Token)
		req  = &struct {
			Name string `json:"name"`
		}{}
	)

	type resType struct {
		ID           int32                 `json:"id"`
		Name         string                `json:"name"`
		Enviornments []queries.Environment `json:"enviornments"`
	}

	if err := c.Bind(req); err != nil {
		return err
	}

	claims := user.Claims.(jwt.MapClaims)
	ownerId := int64(claims["id"].(float64))

	tx, err := app.dbConn.Begin(c.Request().Context())

	if err != nil {
		app.log.Println("Failed to create project", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	defer tx.Rollback(c.Request().Context())

	qtx := app.q.WithTx(tx)

	// create project
	project, err := qtx.CreateProject(c.Request().Context(), queries.CreateProjectParams{
		Name:    req.Name,
		OwnerID: ownerId,
	})

	if err != nil {
		app.log.Println("Failed to create project", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	// add to project_members table
	err = qtx.AddProjectMember(c.Request().Context(), queries.AddProjectMemberParams{
		UserID:    ownerId,
		ProjectID: int64(project.ID),
	})

	if err != nil {
		app.log.Println("Failed to add project member", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	// add default environments
	envs, err := qtx.CreateProdAndDevEnvironments(c.Request().Context())

	if err != nil {
		app.log.Println("Failed to create environments", err)
		c.JSON(500, InternalServerErrorResponse)
		return err
	}

	// add environments to project
	err = qtx.AddProdAndDevProjectEnviornments(c.Request().Context(), queries.AddProdAndDevProjectEnviornmentsParams{
		ProjectID:       int64(project.ID),
		EnvironmentID:   int64(envs[0].ID),
		EnvironmentID_2: int64(envs[1].ID),
	})

	if err != nil {
		app.log.Println("Failed to add environments to project", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	tx.Commit(c.Request().Context())

	response := resType{
		ID:           project.ID,
		Name:         project.Name,
		Enviornments: envs,
	}

	c.JSON(http.StatusOK, responseType{true, response, nil})
	return nil
}

func handleGetUserProjects(c echo.Context) error {
	var (
		app  = c.Get("app").(*App)
		user = c.Get("user").(*jwt.Token)
	)

	userId := int64(user.Claims.(jwt.MapClaims)["id"].(float64))

	projects, err := app.q.GetUserProjects(c.Request().Context(), userId)

	if err != nil {
		app.log.Println("Failed to get user projects", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	c.JSON(http.StatusOK, responseType{true, projects, nil})

	return nil
}
