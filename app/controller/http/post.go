package httpcontroller

import (
	"darkness8129/news-api/app/service"

	"github.com/gin-gonic/gin"
)

type postController struct {
	services service.Services
}

func newPostController(opt routerOptions) {
	c := postController{
		services: opt.Services,
	}

	group := opt.RouterGroup.Group("/posts")
	group.POST("", c.create)
}

func (ctrl *postController) create(c *gin.Context) {
}
