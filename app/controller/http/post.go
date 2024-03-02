package httpcontroller

import (
	"darkness8129/news-api/app/entity"
	"darkness8129/news-api/app/service"
	"darkness8129/news-api/packages/errs"

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
	group.POST("", errorHandler(c.create))
	group.GET("", errorHandler(c.list))
	group.GET(":id", errorHandler(c.get))
	group.PUT(":id", errorHandler(c.update))
	group.DELETE(":id", errorHandler(c.delete))
}

type createPostBody struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type createPostResponse struct {
	Post *entity.Post `json:"post"`
}

func (ctrl *postController) create(c *gin.Context) (interface{}, *httpErr) {
	var body createPostBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid request body"}
	}

	post, err := ctrl.services.Post.Create(c, service.CreatePostOpt{
		Title:   body.Title,
		Content: body.Content,
	})
	if err != nil {
		if errs.IsExpected(err) {
			return nil, &httpErr{Type: httpErrTypeClient, Message: err.Error(), Code: errs.Code(err)}
		}

		return nil, &httpErr{Type: httpErrTypeServer, Message: "failed to create post"}
	}

	return createPostResponse{post}, nil
}

type listPostsResponse struct {
	Posts []entity.Post `json:"posts"`
}

func (ctrl *postController) list(c *gin.Context) (interface{}, *httpErr) {
	posts, err := ctrl.services.Post.List(c)
	if err != nil {
		if errs.IsExpected(err) {
			return nil, &httpErr{Type: httpErrTypeClient, Message: err.Error(), Code: errs.Code(err)}
		}

		return nil, &httpErr{Type: httpErrTypeServer, Message: "failed to list posts"}
	}

	return listPostsResponse{posts}, nil
}

type getPostPathParams struct {
	ID string `uri:"id" json:"id" binding:"required"`
}

type getPostResponse struct {
	Post *entity.Post `json:"post"`
}

func (ctrl *postController) get(c *gin.Context) (interface{}, *httpErr) {
	var pathParams getPostPathParams
	err := c.ShouldBindUri(&pathParams)
	if err != nil {
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid path params"}
	}

	post, err := ctrl.services.Post.Get(c, pathParams.ID)
	if err != nil {
		if errs.IsExpected(err) {
			return nil, &httpErr{Type: httpErrTypeClient, Message: err.Error(), Code: errs.Code(err)}
		}

		return nil, &httpErr{Type: httpErrTypeServer, Message: "failed to get post"}
	}

	return getPostResponse{post}, nil
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

func (ctrl *postController) update(c *gin.Context) (interface{}, *httpErr) {
	var pathParams updatePostPathParams
	err := c.ShouldBindUri(&pathParams)
	if err != nil {
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid path params"}
	}

	var body updatePostBody
	err = c.ShouldBindJSON(&body)
	if err != nil {
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid request body"}
	}

	post, err := ctrl.services.Post.Update(c, pathParams.ID, service.UpdatePostOpt{
		Title:   body.Title,
		Content: body.Content,
	})
	if err != nil {
		if errs.IsExpected(err) {
			return nil, &httpErr{Type: httpErrTypeClient, Message: err.Error(), Code: errs.Code(err)}
		}

		return nil, &httpErr{Type: httpErrTypeServer, Message: "failed to update post"}
	}

	return createPostResponse{post}, nil
}

type deletePostPathParams struct {
	ID string `uri:"id" json:"id" binding:"required"`
}

type deletePostResponse struct {
}

func (ctrl *postController) delete(c *gin.Context) (interface{}, *httpErr) {
	var pathParams updatePostPathParams
	err := c.ShouldBindUri(&pathParams)
	if err != nil {
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid path params"}
	}

	err = ctrl.services.Post.Delete(c, pathParams.ID)
	if err != nil {
		if errs.IsExpected(err) {
			return nil, &httpErr{Type: httpErrTypeClient, Message: err.Error(), Code: errs.Code(err)}
		}

		return nil, &httpErr{Type: httpErrTypeServer, Message: "failed to delete post"}
	}

	return deletePostResponse{}, nil
}
