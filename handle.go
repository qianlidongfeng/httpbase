package main

import (
	"github.com/casbin/casbin"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	realPasswd := G.DB.GetUserPasswordByName(username)
	// Throws unauthorized error
	if password != realPasswd {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["exp"] = time.Now().Add(time.Hour*12).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(G.cfg.Secret.LoginJwt))
	if err != nil {
		return err
	}
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = t
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly=true
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func casbinAuth(c echo.Context) error{
	ce:=casbin.NewEnforcer("casbin_auth_model.conf", "casbin_auth_policy.csv")
	re:=ce.Enforce("alice","data1","read")
	_=re
	return c.String(http.StatusOK, "Welcome "+"!")
}