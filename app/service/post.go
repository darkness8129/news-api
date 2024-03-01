package service

import (
	"context"
	"darkness8129/news-api/app/entity"
	"fmt"
)

var _ PostService = (*postService)(nil)

type postService struct {
	storages Storages
}

func NewPostService(storages Storages) *postService {
	return &postService{storages}
}

func (s *postService) Create(ctx context.Context, opt CreatePostOpt) (*entity.Post, error) {
	createdPost, err := s.storages.Post.Create(ctx, &entity.Post{
		Title:   opt.Title,
		Content: opt.Content,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	return createdPost, nil
}

func (s *postService) List(ctx context.Context) ([]entity.Post, error) {
	posts, err := s.storages.Post.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list posts: %w", err)
	}

	return posts, nil
}

func (s *postService) Get(ctx context.Context, id string) (*entity.Post, error) {
	post, err := s.storages.Post.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}
	if post == nil { // TODO: handle as expected
		return nil, fmt.Errorf("post not found")
	}

	return post, nil
}

func (s *postService) Update(ctx context.Context, id string, opt UpdatePostOpt) (*entity.Post, error) {
	updatedPost, err := s.storages.Post.Update(ctx, id, &entity.Post{
		Title:   opt.Title,
		Content: opt.Content,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	return updatedPost, nil
}

func (s *postService) Delete(ctx context.Context, id string) error {
	err := s.storages.Post.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	return nil
}
