package service

import (
	"context"
	"darkness8129/news-api/app/entity"
)

type Services struct {
	Post PostService
}

type PostService interface {
	Create(ctx context.Context, opt CreatePostOpt) (*entity.Post, error)
	List(ctx context.Context) ([]entity.Post, error)
	Get(ctx context.Context, id string) (*entity.Post, error)
	Update(ctx context.Context, id string, opt UpdatePostOpt) (*entity.Post, error)
	Delete(ctx context.Context, id string) error
}

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
}

type PostStorage interface {
	Create(ctx context.Context, post *entity.Post) (*entity.Post, error)
	List(ctx context.Context) ([]entity.Post, error)
	Get(ctx context.Context, id string) (*entity.Post, error)
	Update(ctx context.Context, id string, post *entity.Post) (*entity.Post, error)
	Delete(ctx context.Context, id string) error
}
