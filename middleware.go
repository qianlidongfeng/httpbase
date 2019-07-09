package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/qianlidongfeng/execho"
	"github.com/qianlidongfeng/loger"
)

type Mw struct{

}

func (this *Mw) recover(loger loger.Loger)echo.MiddlewareFunc{
	exMiddleWare:=execho.NewMiddleWare()
	return exMiddleWare.Recover(loger)
}

func (this *Mw) loginJwtAuth(jtwKey []byte) echo.MiddlewareFunc{
	c := middleware.DefaultJWTConfig
	c.SigningKey = jtwKey
	c.TokenLookup = "cookie:session"
	return middleware.JWTWithConfig(c)
}