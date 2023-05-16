package cmd

import (
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

	if err := c.Bind(req); err != nil {
		return err
	}

	claims := user.Claims.(jwt.MapClaims)
	ownerId := int64(claims["id"].(float64))

	project, err := app.q.CreateProject(c.Request().Context(), queries.CreateProjectParams{
		Name:    req.Name,
		OwnerID: ownerId,
	})

	if err != nil {
		app.log.Println("Failed to create project", err)
		c.JSON(500, InternalServerErrorResponse)
		return err
	}

	// add to project_members table
	err = app.q.AddProjectMember(c.Request().Context(), queries.AddProjectMemberParams{
		UserID:    ownerId,
		ProjectID: int64(project.ID),
	})

	if err != nil {
		app.log.Println("Failed to add project member", err)
		c.JSON(500, InternalServerErrorResponse)
		return err
	}

	c.JSON(200, responseType{true, project, nil})
	return nil
}
