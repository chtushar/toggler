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

	// Create Project transaction
	tx, err := app.dbConn.Begin(c.Request().Context())

	if err != nil {
		app.log.Println("Failed to create tx", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	defer tx.Rollback(c.Request().Context())

	qtx := app.q.WithTx(tx)

	p, err := qtx.CreateProject(c.Request().Context(), queries.CreateProjectParams{
		Name: req.Name,
		OwnerID: ownerId,
		OrgID: int64(orgId),
	})

	if err != nil {
		app.log.Println("Failed to create project", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}
	
	apiKeyProd, _ := utils.GenerateAPIKey()
	apiKeyDev, _ := utils.GenerateAPIKey()


	_, err = qtx.CreateEnvironment(c.Request().Context(), queries.CreateEnvironmentParams{
		Name: "Production",
		ProjectID: int64(p.ID),
		ApiKeys: []string{apiKeyProd},
	})

	if err != nil {
		app.log.Println("Failed to create Prod env", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	_, err = qtx.CreateEnvironment(c.Request().Context(), queries.CreateEnvironmentParams{
		Name: "Development",
		ProjectID: int64(p.ID),
		ApiKeys: []string{apiKeyDev},
	})

	if err != nil {
		app.log.Println("Failed to create Dev env", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	tx.Commit(c.Request().Context())

	c.JSON(http.StatusOK, responseType{true, p, nil})
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
