package cmd

import (
	"net/http"
	"time"

	"github.com/chtushar/toggler/db/queries"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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

func handleUpdateOrganization (c echo.Context) error {
	var (
		app  = c.Get("app").(*App)
		req = &struct {
			Name string `json:"name"`;
			OrgId	int32 `json:"orgId"`  
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

	org, err := app.q.UpdateOrganization(c.Request().Context(), queries.UpdateOrganizationParams{
		ID: req.OrgId,
		Name: req.Name,
	})

	if err != nil {
		app.log.Println("Failed to update the organization", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	c.JSON(http.StatusOK, responseType{true, org, nil})

	return nil
}

func handleGetOrganizationMembers (c echo.Context) error {
	var (
		app  = c.Get("app").(*App)
		orgId = c.Get("orgId").(int)
	)

	type member struct {
		Id            int32            	`json:"id"`
		Uuid		  uuid.NullUUID	   	`json:"uuid"`
		Name          string           	`json:"name"`
		Email         string  			`json:"email"`
		EmailVerified bool              `json:"email_verified"`
		CreatedAt     time.Time      	`json:"created_at"`
	}

	members, err := app.q.GetOrganizationMembers(c.Request().Context(), int64(orgId))

	if err != nil {
		app.log.Println("Failed to get organization members", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	var res []member

	for _, m := range members {
		res = append(res, member{
			Id: m.ID,
			Name: *m.Name,
			Email: *m.Email,
			EmailVerified: m.EmailVerified,
			CreatedAt: m.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, responseType{
		true,
		res,
		nil,
	})

	return nil
}

func handleAddTeamMember(c echo.Context) error {
	var (
		app  = c.Get("app").(*App)
		orgId = c.Get("orgId").(int)
		req = &struct {
			Email string `json:"email"`;  
		}{}
	)

	tx, err := app.dbConn.Begin(c.Request().Context())

	if err != nil {
		app.log.Println("Failed to add member", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	defer tx.Rollback(c.Request().Context())
	qtx := app.q.WithTx(tx)

	// Check if the email-id already exists
	ok, err := qtx.CheckIfUserExists(c.Request().Context(), &req.Email)

	if err != nil {
		app.log.Println("Failed to add member", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
		return err
	}

	if ok {
		// If yes add the user as an organization member
		user, err := qtx.GetUserByEmail(c.Request().Context(), &req.Email)

		if err != nil {
			app.log.Println("Failed to add member", err)
			c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
			return err
		}

		qtx.AddOrganizationMember(c.Request().Context(), queries.AddOrganizationMemberParams{
			OrgID: int64(orgId),
			UserID: int64(user.ID),
		})
	} else {
		// If not create an account with just the email, email_veridfied=false
		user, err := qtx.CreateUser(c.Request().Context(), queries.CreateUserParams{
			Email: &req.Email,
			EmailVerified: false,
		})

		if err != nil {
			app.log.Println("Failed to add member", err)
			c.JSON(http.StatusInternalServerError, InternalServerErrorResponse)
			return err
		}

		qtx.AddOrganizationMember(c.Request().Context(), queries.AddOrganizationMemberParams{
			UserID: int64(user.ID),
			OrgID: int64(orgId),
		})
	}

	tx.Commit(c.Request().Context())

	c.JSON(http.StatusOK, responseType{
		Success: true,
		Data: nil,
		Error: nil,
	})

	return nil
}