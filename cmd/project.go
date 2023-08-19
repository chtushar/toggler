package cmd

import (
	"net/http"

	"github.com/chtushar/toggler/db/queries"
	"github.com/chtushar/toggler/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func handleCreateProject(c echo.Context) error {
	var (
		app  = c.Get("app").(*App)
		user = c.Get("user").(*jwt.Token)
		orgId = c.Get("orgId").(int)
		req  = &struct {
			Name string `json:"name"`
		}{}
	)

	if err := c.Bind(req); err != nil {
		app.log.Println("Failed to create project", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
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
		OrgID: int64(orgId),
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

	apiKey1, _ := utils.GenerateAPIKey()
	apiKey2, _ := utils.GenerateAPIKey()

	// add default environments
	envs, err := qtx.CreateProdAndDevEnvironments(c.Request().Context(), queries.CreateProdAndDevEnvironmentsParams{
		ApiKeys: []string{apiKey1},
		ApiKeys_2: []string{apiKey2},
	})

	if err != nil {
		app.log.Println("Failed to create environments", err)
		c.JSON(500, InternalServerErrorResponse)
		return err
	}

	// add environments to project
	err = qtx.AddProdAndDevProjectEnvironments(c.Request().Context(), queries.AddProdAndDevProjectEnvironmentsParams{
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

	c.JSON(http.StatusOK, responseType{true, project, nil})
	return nil
}

func handleGetOrgProjects(c echo.Context) error {
	var (
		app  = c.Get("app").(*App)
		user = c.Get("user").(*jwt.Token)
		orgId = c.Get("orgId").(int)
	)

	userId := int64(user.Claims.(jwt.MapClaims)["id"].(float64))
	projects, err := app.q.GetUserOrgProjects(c.Request().Context(), queries.GetUserOrgProjectsParams{
		UserID: userId,
		OrgID: int64(orgId),
	})

	if err != nil {
		app.log.Println("Failed to get user projects", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	c.JSON(http.StatusOK, responseType{true, projects, nil})
	return nil
}
