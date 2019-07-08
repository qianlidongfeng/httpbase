package main

import (
	"crypto/tls"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/qianlidongfeng/execho"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Addr:  global.cfg.Httpserver.Address,
		ReadTimeout:  global.cfg.Httpserver.ReadTimeOut*time.Millisecond,
		WriteTimeout: global.cfg.Httpserver.WriteTimeOut*time.Millisecond,
	}
	if global.cfg.Httpserver.Https{
		crt, err := tls.LoadX509KeyPair(global.cfg.Httpserver.CertFile, global.cfg.Httpserver.KeyFile)
		if err != nil {
			panic(err)
			global.log.Fatal(err)
			return
		}
		s.TLSConfig= &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
			Certificates: []tls.Certificate{crt},
		}
		s.TLSNextProto=make(map[string]func(*http.Server, *tls.Conn, http.Handler))
	}
	e := echo.New()
	//set
	e.DisableHTTP2=true
	//router
	e.POST("/login", login)
	e.GET("/", accessible)
	//recover middleware
	exMiddleWare:=execho.NewMiddleWare()
	e.Use(exMiddleWare.Recover(global.log))
	r := e.Group("/restricted")
	c := middleware.DefaultJWTConfig
	c.SigningKey = []byte("secret")
	c.TokenLookup = "cookie:sid"
	//jtw auth middleware
	r.Use(middleware.JWTWithConfig(c))
	r.GET("", restricted)
	global.log.Fatal(e.StartServer(s))
}