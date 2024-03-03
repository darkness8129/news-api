package httpcontroller

import (
	"darkness8129/news-api/app/service"
	"darkness8129/news-api/packages/logging"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Options struct {
	Router   *gin.Engine
	Services service.Services
	Logger   logging.Logger
}

type controllerOptions struct {
	RouterGroup *gin.RouterGroup
	Services    service.Services
	Logger      logging.Logger
}

func New(opt Options) {
	opt.Router.Use(gin.Logger(), gin.Recovery(), corsMiddleware)

	controllerOpt := controllerOptions{
		RouterGroup: opt.Router.Group("/api/v1"),
		Services:    opt.Services,
		Logger:      opt.Logger.Named("httpController"),
	}

	newPostController(controllerOpt)
	newDocsController(controllerOpt)
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
	Type             httpErrType            `json:"-"`
	Code             string                 `json:"code,omitempty"`
	Message          string                 `json:"message"`
	Details          interface{}            `json:"details,omitempty"`
	ValidationErrors map[string]interface{} `json:"validationErrors,omitempty"`
} // @name httpErr

type httpErrType string

const (
	httpErrTypeServer httpErrType = "server"
	httpErrTypeClient httpErrType = "client"
)

// errorDecorator provides unified error handling for all http controllers
func errorDecorator(logger logging.Logger, handler func(c *gin.Context) (interface{}, *httpErr)) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logger.Named("errorHandler")

		// handle panics
		defer func() {
			err := recover()
			if err != nil {
				err := c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w", err))
				if err != nil {
					logger.Error("failed to abort with error", "err", err)
				}
			}
		}()

		body, err := handler(c)
		if err != nil {
			if err.Type == httpErrTypeServer {
				logger.Error("internal server error", "err", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			} else {
				handleValidationErrors(err)

				logger.Info("expected client error", "err", err)
				c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
			}

			return
		}

		logger.Info("successfully handled request")
		c.JSON(http.StatusOK, body)
	}
}

func handleValidationErrors(err *httpErr) {
	// checking whether validation errors exist
	validationErrors, ok := err.Details.(validator.ValidationErrors)
	if !ok {
		return
	}

	err.Details = nil
	err.ValidationErrors = make(map[string]interface{})
	for _, e := range validationErrors {
		fieldName := e.Field()
		switch e.Tag() {
		case "required":
			err.ValidationErrors[fieldName] = "field is required"
		case "max":
			err.ValidationErrors[fieldName] = "maximum allowed characters exceeded"
		case "uuid":
			err.ValidationErrors[fieldName] = "invalid ID"
		default:
			err.ValidationErrors[fieldName] = "unknown validation error"
		}
	}
}
