package cmd

import (
	"net/http"

	"github.com/chtushar/toggler/db/queries"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func handleCreateOrganization(c echo.Context) error {
	var (
		user = c.Get("user").(*jwt.Token)
		app  = c.Get("app").(*App)
		req = &struct {
			Name string `json:"name"`;
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
		app.log.Println("Failed to create organization", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	defer tx.Rollback(c.Request().Context())

	qtx := app.q.WithTx(tx)

	// Create Organization
	org, err := qtx.CreateOrganization(c.Request().Context(), req.Name)

	if err != nil {
		app.log.Println("Failed to create organization", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	claims := user.Claims.(jwt.MapClaims)
	userId := int64(claims["id"].(float64))

	// add to organization_members table
	err = qtx.AddOrganizationMember(c.Request().Context(), queries.AddOrganizationMemberParams{
		UserID: userId,
		OrgID: int64(org.ID),
	})

	if err != nil {
		app.log.Println("Failed to add organization member", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	tx.Commit(c.Request().Context())

	c.JSON(http.StatusOK, responseType{
		Success: true,
		Data: org,
		Error: nil,
	})
	return nil
}

func handleGetUserOrganizations (c echo.Context) error {
	var (
		user = c.Get("user").(*jwt.Token)
		app  = c.Get("app").(*App)
	)

	claims := user.Claims.(jwt.MapClaims)
	userId := int64(claims["id"].(float64))

	orgs, err := app.q.GetUserOrganizations(c.Request().Context(), userId)
	if err != nil {
		app.log.Println("Failed to get user organizations", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	c.JSON(http.StatusOK, responseType{
		Success: true,
		Data: orgs,
		Error: nil,
	})

	return nil
}
