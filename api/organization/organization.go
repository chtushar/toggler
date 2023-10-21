package organization

import (
	"net/http"
	"strings"

	"github.com/chtushar/toggler/api/app"
	"github.com/chtushar/toggler/api/responses"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)


func CheckOrgAccessMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
		var (
			app = c.Get("app").(*app.App)
			orgUUID = c.Param("orgUUID")
		)
		
		cookie, err := c.Cookie("auth_token")

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, responses.UnauthorizedResponse)
		}
		
		claims := jwt.MapClaims{}

		if token, err := jwt.ParseWithClaims(cookie.Value, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(app.Jwt), err
		}); err == nil && token.Valid {
			orgs := claims["orgs"].(string)
			if orgs == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, responses.UnauthorizedResponse)
			}

			orgsUUIDs := strings.Split(orgs, ",")

			for _, o := range orgsUUIDs {
				if o == orgUUID {
					return next(c)
				}
			}
		}
		
		return echo.NewHTTPError(http.StatusUnauthorized, responses.UnauthorizedResponse)
	}
}