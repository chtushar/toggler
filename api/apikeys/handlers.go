package apikeys

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

func handleCreateAPIKey(c echo.Context) error {
	var (
		app = c.Get("app").(*app.App)
		orgUUID = c.Param("orgUUID")
		token = c.Get("user").(*jwt.Token)
		req = &struct{
			Name string `json:"name" validate:"required"`
			Allowed_domains []string `json:"allowed_domains" validate:"required"`
		}{}
	)

	userUuid := token.Claims.(jwt.MapClaims)["uuid"].(string)
	_, err := utils.IsValidUUID(userUuid)

	if err != nil{
		app.Log.Println("Unable to parse the user uuid")
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	_, err = utils.IsValidUUID(orgUUID)
	if err != nil{
		app.Log.Println("Unable to parse the org uuid")
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequestResponse)
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	apiKey, err := generateAPIKey(GenerateAPIKeyParams{
		allowed_domains: req.Allowed_domains,
	})

	if err != nil {
		app.Log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	apiKeyRes, err := db.WithDBTransaction[queries.ApiKey](app, c.Request().Context(), func(q *queries.Queries) (*queries.ApiKey, error) {
		u, err := q.GetUserByUUID(c.Request().Context(), userUuid)
		
		if err != nil {
			app.Log.Println("Unable to get user from db", err)
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		o, err := q.GetOrganizationByUUID(c.Request().Context(), orgUUID)

		if err != nil {
			app.Log.Println("Unable to get org from db", err)
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		apiKeyRes, err := q.CreateAPIKey(c.Request().Context(), queries.CreateAPIKeyParams{
			ApiKey: *apiKey,
			OrgID: o.ID,
			UserID: u.ID,
			AllowedDomains: req.Allowed_domains,
			Name: req.Name,
		})

		if err != nil {
			app.Log.Println("Unable to add api key to db", err)
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}

		return &apiKeyRes, nil
	})

	if err != nil {
		return err
	}


	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data: apiKeyRes,
		Error: nil,
	})
	return nil
}

func handleGetOrgAPIKeys(c echo.Context) error {
	var (
		app = c.Get("app").(*app.App)
		orgUUID = c.Param("orgUUID")
	)

	_, err := utils.IsValidUUID(orgUUID)
	if err != nil{
		app.Log.Println("Unable to parse the org uuid")
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
	}

	apiKeys, err := db.WithDBTransaction[[]queries.ApiKey](app, c.Request().Context(), func(q *queries.Queries) (*[]queries.ApiKey, error) {
		org, err := q.GetOrganizationByUUID(c.Request().Context(), orgUUID)
		
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}
	
		apiKeys, err := q.GetOrganizationAPIKeys(c.Request().Context(), org.ID)
		
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, responses.InternalServerErrorResponse)
		}
		
		return &apiKeys, err
	})

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, responses.ResponseType{
		Success: true,
		Data: apiKeys,
		Error: nil,
	})
	return nil
}

func handleDeleteAPIKey(c echo.Context) error {
	return nil
}