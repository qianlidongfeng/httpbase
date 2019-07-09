package main

import (
	"crypto/tls"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func InitEcho(e *echo.Echo){
	e.DisableHTTP2=true
	//router
	e.POST("/login", login)
	e.GET("/", accessible)

	r := e.Group("/restricted")
	r.GET("", restricted)
	//recover middleware
	e.Use(G.middleware["recover"])
	//jtw auth middleware
	r.Use(G.middleware["loginJwtAuth"])
	//render
	e.Renderer=G.renderer
}

func main() {
	s := &http.Server{
		Addr:  G.cfg.Httpserver.Address,
		ReadTimeout:  G.cfg.Httpserver.ReadTimeOut*time.Millisecond,
		WriteTimeout: G.cfg.Httpserver.WriteTimeOut*time.Millisecond,
	}
	if G.cfg.Httpserver.Https{
		crt, err := tls.LoadX509KeyPair(G.cfg.Httpserver.CertFile, G.cfg.Httpserver.KeyFile)
		if err != nil {
			panic(err)
			G.log.Fatal(err)
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
	InitEcho(e)
	G.log.Fatal(e.StartServer(s))
}