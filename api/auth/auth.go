package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func writeAuthTokenToCookie(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = token
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	c.SetCookie(cookie)
}
