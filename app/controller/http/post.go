package httpcontroller

import (
	"darkness8129/news-api/app/entity"
	"darkness8129/news-api/app/service"
	"fmt"
	"net/http"

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
	group.GET("", c.list)
	group.GET(":id", c.get)
	group.PUT(":id", c.update)
	group.DELETE(":id", c.delete)
}

type createPostBody struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type createPostResponse struct {
	Post *entity.Post `json:"post"`
}

func (ctrl *postController) create(c *gin.Context) {
	var body createPostBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid request body: %w", err).Error()})
		return
	}

	post, err := ctrl.services.Post.Create(c, service.CreatePostOpt{
		Title:   body.Title,
		Content: body.Content,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": fmt.Errorf("failed to create post: %w", err).Error()})
		return
	}

	c.JSON(http.StatusOK, createPostResponse{post})
}

type listPostsResponse struct {
	Posts []entity.Post `json:"posts"`
}

func (ctrl *postController) list(c *gin.Context) {
	posts, err := ctrl.services.Post.List(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": fmt.Errorf("failed to list posts: %w", err).Error()})
		return
	}

	c.JSON(http.StatusOK, listPostsResponse{posts})
}

type getPostPathParams struct {
	ID string `uri:"id" json:"id" binding:"required"`
}

type getPostResponse struct {
	Post *entity.Post `json:"post"`
}

func (ctrl *postController) get(c *gin.Context) {
	var pathParams getPostPathParams
	err := c.ShouldBindUri(&pathParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid path params: %w", err).Error()})
	}

	post, err := ctrl.services.Post.Get(c, pathParams.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": fmt.Errorf("failed to get post: %w", err).Error()})
		return
	}

	c.JSON(http.StatusOK, getPostResponse{post})
}

type updatePostPathParams struct {
	ID string `uri:"id" json:"id" binding:"required"`
}

type updatePostBody struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type updatePostResponse struct {
	Post *entity.Post `json:"post"`
}

func (ctrl *postController) update(c *gin.Context) {
	var pathParams updatePostPathParams
	err := c.ShouldBindUri(&pathParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid path params: %w", err).Error()})
	}

	var body updatePostBody
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid request body: %w", err).Error()})
		return
	}

	post, err := ctrl.services.Post.Update(c, pathParams.ID, service.UpdatePostOpt{
		Title:   body.Title,
		Content: body.Content,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": fmt.Errorf("failed to create post: %w", err).Error()})
		return
	}

	c.JSON(http.StatusOK, createPostResponse{post})
}

type deletePostPathParams struct {
	ID string `uri:"id" json:"id" binding:"required"`
}

type deletePostResponse struct {
}

func (ctrl *postController) delete(c *gin.Context) {
	var pathParams updatePostPathParams
	err := c.ShouldBindUri(&pathParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid path params: %w", err).Error()})
	}

	err = ctrl.services.Post.Delete(c, pathParams.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": fmt.Errorf("failed to delete post: %w", err).Error()})
		return
	}

	c.JSON(http.StatusOK, deletePostResponse{})
}
