package httpcontroller

import (
	"darkness8129/news-api/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Router   *gin.Engine
	Services service.Services
}

type routerOptions struct {
	RouterGroup *gin.RouterGroup
	Services    service.Services
}

func New(opt Options) {
	opt.Router.Use(gin.Logger(), gin.Recovery(), corsMiddleware)

	routerOpt := routerOptions{
		RouterGroup: opt.Router.Group("/api/v1"),
		Services:    opt.Services,
	}

	newPostController(routerOpt)
}

func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}
