package server

import (
	"app/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func NewRestapi() *gin.Engine {
	config.App.CheckPort()
	var r *gin.Engine
	if config.App.Env != "production" {
		r = gin.Default()
	} else {
		r = gin.New()
	}
	//
	//r.ForwardedByClientIP = true
	//r.SetTrustedProxies([]string{"*"})

	//r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(addCorsMiddleware())

	return r
}

func addCorsMiddleware() gin.HandlerFunc {
	configCors := cors.DefaultConfig()

	configCors.AllowAllOrigins = true
	configCors.AllowWildcard = true
	configCors.AllowCredentials = true
	configCors.MaxAge = 12 * time.Hour
	configCors.AllowHeaders = append(configCors.AllowHeaders, "Authorization", "Accept-Encoding", "Origin", "Content-Type")
	return cors.New(configCors)
}
