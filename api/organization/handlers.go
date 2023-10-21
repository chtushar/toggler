package organization

import (
	"net/http"

	"github.com/chtushar/toggler/api/app"
	"github.com/chtushar/toggler/api/responses"
	"github.com/chtushar/toggler/db"
	"github.com/chtushar/toggler/db/queries"
	"github.com/chtushar/toggler/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func handleCreateOrganization (c echo.Context) error {
	var (
		user = c.Get("user").(*jwt.Token)
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

	userUuidStr := user.Claims.(jwt.MapClaims)["uuid"].(string)
	userUuid, err := uuid.Parse(userUuidStr)

	if err != nil {
		app.Log.Println("Failed to parse the user uuid", err)
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}	

	org, err := db.WithDBTransaction[queries.Organization](app, c.Request().Context(), func (q *queries.Queries) (*queries.Organization, error) {
		org, err := q.CreateOrganization(c.Request().Context(), req.Name)

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		orgUuid, err := uuid.Parse(org.Uuid)

		if err!= nil {
			app.Log.Println("Failed to parse the uuid", err)
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		q.AddOrganizationMember(c.Request().Context(), queries.AddOrganizationMemberParams{
			UserUuid: &userUuid,
			OrgUuid: &orgUuid,
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

	orgUUID, err := uuid.Parse(orgUUIDStr)
	
	if err != nil {
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
		for _, e := range req.Emails {
			u, err := q.GetUserByEmail(c.Request().Context(), e)
			if err != nil {
				app.Log.Println("Can't find user with email", e)
				continue
			}
			userUUID, err := uuid.Parse(u.Uuid)
			if err != nil {
				app.Log.Println("Can't parse user uuid", u.Uuid)
				continue
			}
			err = q.AddOrganizationMember(c.Request().Context(), queries.AddOrganizationMemberParams{
				OrgUuid: &orgUUID,
				UserUuid: &userUUID,
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
		app.Log.Println("Failed to add ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data: org,
		Error: nil,
	})
	return nil
}