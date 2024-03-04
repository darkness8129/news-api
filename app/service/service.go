package service

import (
	"context"
	"darkness8129/news-api/app/entity"
	"darkness8129/news-api/packages/errs"
)

const (
	postNotFoundErrCode = "post_not_found"
	// other err codes should be here
)

type Services struct {
	Post PostService
	// other services should be here
}

type PostService interface {
	Create(ctx context.Context, opt CreatePostOpt) (*entity.Post, error)
	List(ctx context.Context) ([]entity.Post, error)
	Get(ctx context.Context, id string) (*entity.Post, error)
	Update(ctx context.Context, id string, opt UpdatePostOpt) (*entity.Post, error)
	Delete(ctx context.Context, id string) error
}

// expected errors for this service should be here

type CreatePostOpt struct {
	Title   string
	Content string
}

type UpdatePostOpt struct {
	Title   string
	Content string
}

type Storages struct {
	Post PostStorage
	// other storages should be here
}

type PostStorage interface {
	Create(ctx context.Context, post *entity.Post) (*entity.Post, error)
	List(ctx context.Context) ([]entity.Post, error)
	Get(ctx context.Context, id string) (*entity.Post, error)
	Update(ctx context.Context, id string, post *entity.Post) (*entity.Post, error)
	Delete(ctx context.Context, id string) error
}

var (
	ErrGetPostNotFound = errs.New(errs.Options{Message: "post not found", Code: postNotFoundErrCode})
	ErrGetPostEmptyID  = errs.New(errs.Options{Message: "empty id"})
	// other expected errors for this storage should be here
)
