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

func newPostController(opt controllerOptions) {
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
} // @name Post

func ToPostDTO(p *entity.Post) *postDTO {
	return &postDTO{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
	}
}

type createPostBody struct {
	Title   string `json:"title" binding:"required,max=50"`
	Content string `json:"content" binding:"required,max=200"`
} // @name createPostBody

type createPostResponse struct {
	Post *postDTO `json:"post"`
} // @name createPostResponse

// @ID           CreatePost
// @Summary      CreatePost provides the logic for creating a post with passed data.
// @Accept       application/json
// @Produce      application/json
// @Param        fields body createPostBody true "data"
// @Success      200 {object} createPostResponse
// @Failure      422,500 {object} httpErr
// @Router       /posts [POST]
func (ctrl *postController) create(c *gin.Context) (interface{}, *httpErr) {
	logger := ctrl.logger.Named("create")

	var body createPostBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		logger.Info("invalid request body", "err", err)
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid request body", Details: err}
	}
	logger.Debug("parsed request body", "body", body)

	post, err := ctrl.services.Post.Create(c, service.CreatePostOpt{
		Title:   body.Title,
		Content: body.Content,
	})
	if err != nil {
		if errs.IsCustom(err) {
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
} // @name listPostsResponse

// @ID           ListPosts
// @Summary      ListPosts provides the logic for retrieving all posts.
// @Produce      application/json
// @Success      200 {object} listPostsResponse
// @Failure      422,500 {object} httpErr
// @Router       /posts [GET]
func (ctrl *postController) list(c *gin.Context) (interface{}, *httpErr) {
	logger := ctrl.logger.Named("list")

	posts, err := ctrl.services.Post.List(c)
	if err != nil {
		if errs.IsCustom(err) {
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
	ID string `uri:"id" json:"id" binding:"required,uuid"`
} // @name getPostPathParams

type getPostResponse struct {
	Post *postDTO `json:"post"`
} // @name getPostResponse

// @ID           GetPost
// @Summary      GetPost provides the logic for retrieving a post by its ID.
// @Produce      application/json
// @Param        id path string true "Post ID"
// @Success      200 {object} getPostResponse
// @Failure      422,500 {object} httpErr
// @Router       /posts/{id} [GET]
func (ctrl *postController) get(c *gin.Context) (interface{}, *httpErr) {
	logger := ctrl.logger.Named("get")

	var pathParams getPostPathParams
	err := c.ShouldBindUri(&pathParams)
	if err != nil {
		logger.Info("invalid path params", "err", err)
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid path params", Details: err}
	}
	logger.Debug("parsed path params", "pathParams", pathParams)

	post, err := ctrl.services.Post.Get(c, pathParams.ID)
	if err != nil {
		if errs.IsCustom(err) {
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
	ID string `uri:"id" json:"id" binding:"required,uuid"`
} // @name updatePostPathParams

type updatePostBody struct {
	Title   string `json:"title" binding:"required,max=50"`
	Content string `json:"content" binding:"required,max=200"`
} // @name updatePostBody

type updatePostResponse struct {
	Post *postDTO `json:"post"`
} // @name updatePostResponse

// @ID           UpdatePost
// @Summary      UpdatePost provides the logic for updating a post with passed data by its ID.
// @Accept       application/json
// @Produce      application/json
// @Param        id path string true "Post ID"
// @Param        fields body updatePostBody true "data"
// @Success      200 {object} updatePostResponse
// @Failure      422,500 {object} httpErr
// @Router       /posts/{id} [PUT]
func (ctrl *postController) update(c *gin.Context) (interface{}, *httpErr) {
	logger := ctrl.logger.Named("update")

	var pathParams updatePostPathParams
	err := c.ShouldBindUri(&pathParams)
	if err != nil {
		logger.Info("invalid path params", "err", err)
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid path params", Details: err}
	}
	logger.Debug("parsed path params", "pathParams", pathParams)

	var body updatePostBody
	err = c.ShouldBindJSON(&body)
	if err != nil {
		logger.Info("invalid request body", "err", err)
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid request body", Details: err}
	}
	logger.Debug("parsed request body", "body", body)

	updatedPost, err := ctrl.services.Post.Update(c, pathParams.ID, service.UpdatePostOpt{
		Title:   body.Title,
		Content: body.Content,
	})
	if err != nil {
		if errs.IsCustom(err) {
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
	ID string `uri:"id" json:"id" binding:"required,uuid"`
} // @name deletePostPathParams

type deletePostResponse struct {
} // @name deletePostResponse

// @ID           DeletePost
// @Summary      DeletePost provides the logic for deleting a post by its ID. If wrong ID is passed, an error will not be returned.
// @Produce      application/json
// @Param        id path string true "Post ID"
// @Success      200 {object} deletePostResponse
// @Failure      422,500 {object} httpErr
// @Router       /posts/{id} [DELETE]
func (ctrl *postController) delete(c *gin.Context) (interface{}, *httpErr) {
	logger := ctrl.logger.Named("delete")

	var pathParams updatePostPathParams
	err := c.ShouldBindUri(&pathParams)
	if err != nil {
		logger.Info("invalid path params", "err", err)
		return nil, &httpErr{Type: httpErrTypeClient, Message: "invalid path params", Details: err}
	}
	logger.Debug("parsed path params", "pathParams", pathParams)

	err = ctrl.services.Post.Delete(c, pathParams.ID)
	if err != nil {
		if errs.IsCustom(err) {
			logger.Info(err.Error())
			return nil, &httpErr{Type: httpErrTypeClient, Message: err.Error(), Code: errs.Code(err)}
		}

		logger.Error("failed to delete post", "err", err)
		return nil, &httpErr{Type: httpErrTypeServer, Message: "failed to delete post"}
	}

	logger.Info("successfully deleted post", "id", pathParams.ID)
	return deletePostResponse{}, nil
}
