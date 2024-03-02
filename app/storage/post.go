package storage

import (
	"context"
	"darkness8129/news-api/app/entity"
	"darkness8129/news-api/app/service"
	"darkness8129/news-api/packages/logging"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var _ service.PostStorage = (*postStorage)(nil)

type postStorage struct {
	db     *gorm.DB
	logger logging.Logger
}

func NewPostStorage(db *gorm.DB, logger logging.Logger) *postStorage {
	return &postStorage{db, logger.Named("postStorage")}
}

func (s *postStorage) Create(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	logger := s.logger.Named("Create")

	err := s.db.Create(post).Error
	if err != nil {
		logger.Error("failed to create post", "err", err)
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	logger.Info("successfully created post", "post", post)
	return post, nil
}

func (s *postStorage) List(ctx context.Context) ([]entity.Post, error) {
	logger := s.logger.Named("List")

	var posts []entity.Post
	err := s.db.Find(&posts).Error
	if err != nil {
		logger.Error("failed to list posts", "err", err)
		return nil, fmt.Errorf("failed to list posts: %w", err)
	}

	logger.Info("successfully listed posts", "posts", posts)
	return posts, nil
}

func (s *postStorage) Get(ctx context.Context, id string) (*entity.Post, error) {
	logger := s.logger.Named("Get")

	var post entity.Post
	err := s.db.
		Where(entity.Post{ID: id}).
		First(&post).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Info("post not found", "id", id)
		return nil, nil
	}
	if err != nil {
		logger.Error("failed to get post", "err", err)
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	logger.Info("successfully got post", "post", post)
	return &post, nil
}

func (s *postStorage) Update(ctx context.Context, id string, post *entity.Post) (*entity.Post, error) {
	logger := s.logger.Named("Update")

	err := s.db.
		Where(entity.Post{ID: id}).
		Updates(post).Error
	if err != nil {
		logger.Error("failed to update post", "err", err)
		return nil, fmt.Errorf("failed to update post: %w", err)
	}
	logger.Debug("updated post")

	updatedPost, err := s.Get(ctx, id)
	if err != nil {
		logger.Error("failed to get updated post", "err", err)
		return nil, fmt.Errorf("failed to get updated post: %w", err)
	}

	logger.Info("successfully updated post", "updatedPost", updatedPost)
	return updatedPost, nil
}

func (s *postStorage) Delete(ctx context.Context, id string) error {
	logger := s.logger.Named("Delete")

	err := s.db.
		Delete(&entity.Post{ID: id}).Error
	if err != nil {
		logger.Error("failed to delete post", "err", err)
		return fmt.Errorf("failed to delete post: %w", err)
	}

	logger.Info("successfully deleted post", "id", id)
	return nil
}
