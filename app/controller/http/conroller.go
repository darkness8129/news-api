package httpcontroller

import (
	"darkness8129/news-api/app/service"
	"fmt"
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
	// other controllers should be here
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

// httpErr provides a base error type for all http controller errors
type httpErr struct {
	Type    httpErrType `json:"-"`
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message"`
}

type httpErrType string

const (
	httpErrTypeServer httpErrType = "server"
	httpErrTypeClient httpErrType = "client"
)

// errorHandler provides unified error handling for all http controllers
func errorHandler(handler func(c *gin.Context) (interface{}, *httpErr)) gin.HandlerFunc {
	return func(c *gin.Context) {
		// handle panics
		defer func() {
			err := recover()
			if err != nil {
				err := c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w", err))
				if err != nil {
					fmt.Println("failed to abort with error", "err", err)
				}
			}
		}()

		body, err := handler(c)
		if err != nil {
			if err.Type == httpErrTypeServer {
				c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			} else {
				c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
			}

			return
		}

		c.JSON(http.StatusOK, body)
	}
}
