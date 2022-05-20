package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/zepyrshut/tpv-api/internal/config"
)

var app *config.Application

func NewMiddleware(a *config.Application) {
	app = a
}

func Sessions(name string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(app.Config.Session.Secret))
	return sessions.Sessions(name, store)
}

// Cross Origin Resource Sharing
func CORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"}
	config.AllowMethods = []string{"POST", "OPTIONS", "GET", "PUT"}

	return cors.New(config)
}
