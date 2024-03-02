package httpcontroller

import (
	"darkness8129/news-api/app/entity"
	"darkness8129/news-api/app/service"
	"darkness8129/news-api/packages/errs"
	"darkness8129/news-api/packages/logging"

	"github.com/gin-gonic/gin"
)

type postController struct {
	services service.Services
	logger   logging.Logger
}

func newPostController(opt routerOptions) {
	logger := opt.Logger.Named("postController")

	c := postController{
		services: opt.Services,
		logger:   logger,
	}

	group := opt.RouterGroup.Group("/posts")
	group.POST("", errorDecorator(logger, c.create))
	group.GET("", errorDecorator(logger, c.list))
	group.GET(":id", errorDecorator(logger, c.get))
	group.PUT(":id", errorDecorator(logger, c.update))
	group.DELETE(":id", errorDecorator(logger, c.delete))
}

type postDTO struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func ToPostDTO(p *entity.Post) *postDTO {
	return &postDTO{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
	}
}

type createPostBody struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type createPostResponse struct {
	Post *postDTO `json:"post"`
}

func (ctrl *postController) create(c *gin.Context) (interface{}, *httpErr) {
	logger := ctrl.logger.Named("create")

	var body createPostBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		logger.Info("invalid request body", "err", err)
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid request body"}
	}
	logger.Debug("parsed request body", "body", body)

	post, err := ctrl.services.Post.Create(c, service.CreatePostOpt{
		Title:   body.Title,
		Content: body.Content,
	})
	if err != nil {
		if errs.IsExpected(err) {
			logger.Info(err.Error())
			return nil, &httpErr{Type: httpErrTypeClient, Message: err.Error(), Code: errs.Code(err)}
		}

		logger.Error("failed to create post", "err", err)
		return nil, &httpErr{Type: httpErrTypeServer, Message: "failed to create post"}
	}

	logger.Info("successfully created post", "post", post)
	return createPostResponse{ToPostDTO(post)}, nil
}

type listPostsResponse struct {
	Posts []*postDTO `json:"posts"`
}

func (ctrl *postController) list(c *gin.Context) (interface{}, *httpErr) {
	logger := ctrl.logger.Named("list")

	posts, err := ctrl.services.Post.List(c)
	if err != nil {
		if errs.IsExpected(err) {
			logger.Info(err.Error())
			return nil, &httpErr{Type: httpErrTypeClient, Message: err.Error(), Code: errs.Code(err)}
		}

		logger.Error("failed to list posts", "err", err)
		return nil, &httpErr{Type: httpErrTypeServer, Message: "failed to list posts"}
	}

	var postsDTO []*postDTO
	for _, p := range posts {
		postsDTO = append(postsDTO, ToPostDTO(&p))
	}

	logger.Info("successfully listed posts", "posts", posts)
	return listPostsResponse{postsDTO}, nil
}

type getPostPathParams struct {
	ID string `uri:"id" json:"id" binding:"required"`
}

type getPostResponse struct {
	Post *postDTO `json:"post"`
}

func (ctrl *postController) get(c *gin.Context) (interface{}, *httpErr) {
	logger := ctrl.logger.Named("get")

	var pathParams getPostPathParams
	err := c.ShouldBindUri(&pathParams)
	if err != nil {
		logger.Info("invalid path params", "err", err)
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid path params"}
	}
	logger.Debug("parsed path params", "pathParams", pathParams)

	post, err := ctrl.services.Post.Get(c, pathParams.ID)
	if err != nil {
		if errs.IsExpected(err) {
			logger.Info(err.Error())
			return nil, &httpErr{Type: httpErrTypeClient, Message: err.Error(), Code: errs.Code(err)}
		}

		logger.Error("failed to get post", "err", err)
		return nil, &httpErr{Type: httpErrTypeServer, Message: "failed to get post"}
	}

	logger.Info("successfully got post", "post", post)
	return getPostResponse{ToPostDTO(post)}, nil
}

type updatePostPathParams struct {
	ID string `uri:"id" json:"id" binding:"required"`
}

type updatePostBody struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type updatePostResponse struct {
	Post *postDTO `json:"post"`
}

func (ctrl *postController) update(c *gin.Context) (interface{}, *httpErr) {
	logger := ctrl.logger.Named("update")

	var pathParams updatePostPathParams
	err := c.ShouldBindUri(&pathParams)
	if err != nil {
		logger.Info("invalid path params", "err", err)
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid path params"}
	}
	logger.Debug("parsed path params", "pathParams", pathParams)

	var body updatePostBody
	err = c.ShouldBindJSON(&body)
	if err != nil {
		logger.Info("invalid request body", "err", err)
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid request body"}
	}
	logger.Debug("parsed request body", "body", body)

	updatedPost, err := ctrl.services.Post.Update(c, pathParams.ID, service.UpdatePostOpt{
		Title:   body.Title,
		Content: body.Content,
	})
	if err != nil {
		if errs.IsExpected(err) {
			logger.Info(err.Error())
			return nil, &httpErr{Type: httpErrTypeClient, Message: err.Error(), Code: errs.Code(err)}
		}

		logger.Error("failed to update post", "err", err)
		return nil, &httpErr{Type: httpErrTypeServer, Message: "failed to update post"}
	}

	logger.Info("successfully updated post", "updatedPost", updatedPost)
	return updatePostResponse{ToPostDTO(updatedPost)}, nil
}

type deletePostPathParams struct {
	ID string `uri:"id" json:"id" binding:"required"`
}

type deletePostResponse struct {
}

func (ctrl *postController) delete(c *gin.Context) (interface{}, *httpErr) {
	logger := ctrl.logger.Named("delete")

	var pathParams updatePostPathParams
	err := c.ShouldBindUri(&pathParams)
	if err != nil {
		logger.Info("invalid path params", "err", err)
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid path params"}
	}
	logger.Debug("parsed path params", "pathParams", pathParams)

	err = ctrl.services.Post.Delete(c, pathParams.ID)
	if err != nil {
		if errs.IsExpected(err) {
			return nil, &httpErr{Type: httpErrTypeClient, Message: err.Error(), Code: errs.Code(err)}
		}

		logger.Error("failed to delete post", "err", err)
		return nil, &httpErr{Type: httpErrTypeServer, Message: "failed to delete post"}
	}

	logger.Info("successfully deleted post", "id", pathParams.ID)
	return deletePostResponse{}, nil
}
