package interceptor

import (
    "github.com/labstack/echo/middleware"
    "github.com/labstack/echo"
    "github.com/abbot/go-http-auth"
    "net/http"
)

func authenticate(user, realm string) string {
    if user == "user" {
        // password is "hello"
        return "password"
    }
    return ""
}

var authenticator * auth.DigestAuth = authwrapNewDigestAuthenticator()

func authwrapNewDigestAuthenticator() *auth.DigestAuth {
    res := auth.NewDigestAuthenticator("example.com", authenticate)
    return res
}

func DigestAuthenticate() echo.HandlerFunc {
    authenticator.PlainTextSecrets=true
    return func(c echo.Context) error {
        r := c.Request()
        w := c.Response().Writer
        if username, authinfo := authenticator.CheckAuth(r); username == "" {
            authenticator.RequireAuth(w, r)
            return echo.NewHTTPError(http.StatusUnauthorized, "Please write collect username and password")
        } else {
            ar := &auth.AuthenticatedRequest{Request: *r, Username: username}
            if authinfo != nil {
                w.Header().Set(authenticator.Headers.V().AuthInfo, *authinfo)
            }
            return HandleIndex(c, ar)
        }
    }
}

func NoAuthenticate(input func(c echo.Context, r *auth.AuthenticatedRequest) error) echo.HandlerFunc {
    return func(c echo.Context) error {
        return input(c, nil);
    }
}

func HandleIndex(c echo.Context, r *auth.AuthenticatedRequest) error{
    return c.JSON(http.StatusOK, map[string]interface{}{"hello": "world"})
}

func BasicAuth() echo.MiddlewareFunc  {
    return middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
        if username == "user" && password == "password" {
            return true,nil
        }
        return false,nil
    })
}
