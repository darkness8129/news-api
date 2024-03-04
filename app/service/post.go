package service

import (
	"context"
	"darkness8129/news-api/app/entity"
	"darkness8129/news-api/packages/errs"
	"darkness8129/news-api/packages/logging"
	"fmt"
)

var _ PostService = (*postService)(nil)

type postService struct {
	storages Storages
	logger   logging.Logger
}

func NewPostService(storages Storages, logger logging.Logger) *postService {
	return &postService{storages, logger.Named("postService")}
}

func (s *postService) Create(ctx context.Context, opt CreatePostOpt) (*entity.Post, error) {
	logger := s.logger.Named("Create")

	createdPost, err := s.storages.Post.Create(ctx, &entity.Post{
		Title:   opt.Title,
		Content: opt.Content,
	})
	if err != nil {
		if errs.IsCustom(err) {
			logger.Info(err.Error())
			return nil, err
		}

		logger.Error("failed to create post", "err", err)
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	logger.Info("successfully created post", "createdPost", createdPost)
	return createdPost, nil
}

func (s *postService) List(ctx context.Context) ([]entity.Post, error) {
	logger := s.logger.Named("List")

	posts, err := s.storages.Post.List(ctx)
	if err != nil {
		if errs.IsCustom(err) {
			logger.Info(err.Error())
			return nil, err
		}

		logger.Error("failed to list posts", "err", err)
		return nil, fmt.Errorf("failed to list posts: %w", err)
	}

	logger.Info("successfully listed posts", "posts", posts)
	return posts, nil
}

func (s *postService) Get(ctx context.Context, id string) (*entity.Post, error) {
	logger := s.logger.Named("Get")

	post, err := s.storages.Post.Get(ctx, id)
	if err != nil {
		if errs.IsCustom(err) {
			logger.Info(err.Error())
			return nil, err
		}

		logger.Error("failed to get post", "err", err)
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	logger.Info("successfully got post", "post", post)
	return post, nil
}

func (s *postService) Update(ctx context.Context, id string, opt UpdatePostOpt) (*entity.Post, error) {
	logger := s.logger.Named("Update")

	updatedPost, err := s.storages.Post.Update(ctx, id, &entity.Post{
		Title:   opt.Title,
		Content: opt.Content,
	})
	if err != nil {
		if errs.IsCustom(err) {
			logger.Info(err.Error())
			return nil, err
		}

		logger.Error("failed to update post", "err", err)
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	logger.Info("successfully updated post", "updatedPost", updatedPost)
	return updatedPost, nil
}

func (s *postService) Delete(ctx context.Context, id string) error {
	logger := s.logger.Named("Delete")

	err := s.storages.Post.Delete(ctx, id)
	if err != nil {
		if errs.IsCustom(err) {
			logger.Info(err.Error())
			return err
		}

		logger.Error("failed to delete post", "err", err)
		return fmt.Errorf("failed to delete post: %w", err)
	}

	logger.Info("successfully deleted post", "id", id)
	return nil
}
