package httpcontroller

import (
	// This import is necessary for loading documentation on API endpoints
	// https://github.com/swaggo/swag/issues/830
	_ "darkness8129/news-api/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func newDocsController(options controllerOptions) {
	group := options.RouterGroup.Group("/docs")
	{
		group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}
