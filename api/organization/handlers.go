package organization

import (
	"net/http"

	"github.com/chtushar/toggler/api/app"
	"github.com/chtushar/toggler/api/responses"
	"github.com/chtushar/toggler/db"
	"github.com/chtushar/toggler/db/queries"
	"github.com/chtushar/toggler/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func handleCreateOrganization (c echo.Context) error {
	var (
		token = c.Get("user").(*jwt.Token)
		app = c.Get("app").(*app.App)
		req = &struct{
			Name string `json:"name" validate:"required,min=3"`
		}{}
	)

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	userUuidStr := token.Claims.(jwt.MapClaims)["uuid"].(string)
	ok, err := utils.IsValidUUID(userUuidStr)

	if !ok {
		app.Log.Println("Failed to parse the user uuid", err)
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}	

	org, err := db.WithDBTransaction[queries.Organization](app, c.Request().Context(), func (q *queries.Queries) (*queries.Organization, error) {
		org, err := q.CreateOrganization(c.Request().Context(), req.Name)

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		user, err := q.GetUserByUUID(c.Request().Context(), userUuidStr)

		if err!= nil {
			app.Log.Println("Failed to get the user", err)
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		q.AddOrganizationMember(c.Request().Context(), queries.AddOrganizationMemberParams{
			UserID: user.ID,
			OrgID: org.ID,
		})

		return &org, nil
	})

	if err != nil {
		app.Log.Println("Failed to create an organization", err)
		return err
	}

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data: *org,
		Error: nil,
	})

	return nil
}

func handleAddOrganizationMembers (c echo.Context) error {
	var (
		app = c.Get("app").(*app.App)
		orgUUIDStr = c.Param("orgUUID")
		req = &struct{
			Emails []string `json:"emails" validate:"required,dive,email"`
		}{}
	)

	ok, err := utils.IsValidUUID(orgUUIDStr)
	
	if !ok {
		app.Log.Println("Failed to parse the uuid", err, orgUUIDStr)
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	if err != nil {
		app.Log.Println("Failed to parse the uuid", err, orgUUIDStr)
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	_, err = db.WithDBTransaction[bool](app, c.Request().Context(), func(q *queries.Queries) (*bool, error){
		org, err := q.GetOrganizationByUUID(c.Request().Context(), orgUUIDStr)
		
		if err != nil {
			app.Log.Println("Can't find the org", err)
			return nil, echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
		}

		for _, e := range req.Emails {
			u, err := q.GetUserByEmail(c.Request().Context(), e)
			
			if err != nil {
				app.Log.Println("Can't find user with email", e)
				continue
			}

			err = q.AddOrganizationMember(c.Request().Context(), queries.AddOrganizationMemberParams{
				OrgID: org.ID,
				UserID: u.ID,
			})
			if err != nil {
				return nil, echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
			}
		}
		ok := true
		return &ok, nil
	})

	if err != nil {
		app.Log.Println("Unable to add members", err)
		return err
	}

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data: "Members added",
		Error: nil,
	})
	return nil
}

func handleGetOrganization (c echo.Context) error {
	var (
		app = c.Get("app").(*app.App)
		orgUUID = c.Param("orgUUID")
	)

	ok, err := utils.IsValidUUID(orgUUID)
	
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}
	
	if err != nil {
		app.Log.Println("Failed to parse the uuid", err, orgUUID)
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	org, err := app.Q.GetOrganizationByUUID(c.Request().Context(), orgUUID)
	if err != nil {
		app.Log.Println("Failed to Get organization ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data: org,
		Error: nil,
	})
	return nil
}

func handleGetUserOrganizations(c echo.Context) error {
	var (
		app = c.Get("app").(*app.App)
		token = c.Get("user").(*jwt.Token)
	)

	userUuid := token.Claims.(jwt.MapClaims)["uuid"].(string)
	ok, _ := utils.IsValidUUID(userUuid)
	
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	orgs, err := app.Q.GetUserOrganizations(c.Request().Context(), userUuid)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data: orgs,
		Error: nil,
	})
	return nil
}
