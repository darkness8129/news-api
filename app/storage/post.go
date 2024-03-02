package storage

import (
	"context"
	"darkness8129/news-api/app/entity"
	"darkness8129/news-api/app/service"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var _ service.PostStorage = (*postStorage)(nil)

type postStorage struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *postStorage {
	return &postStorage{db}
}

func (s *postStorage) Create(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	err := s.db.Create(post).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	return post, nil
}

func (s *postStorage) List(ctx context.Context) ([]entity.Post, error) {
	var posts []entity.Post
	err := s.db.Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list posts: %w", err)
	}

	return posts, nil
}

func (s *postStorage) Get(ctx context.Context, id string) (*entity.Post, error) {
	var post entity.Post
	err := s.db.
		Where(entity.Post{ID: id}).
		First(&post).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	return &post, nil
}

func (s *postStorage) Update(ctx context.Context, id string, post *entity.Post) (*entity.Post, error) {
	err := s.db.
		Where(entity.Post{ID: id}).
		Updates(post).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	updatedPost, err := s.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get updated post: %w", err)
	}

	return updatedPost, nil
}

func (r *postStorage) Delete(ctx context.Context, id string) error {
	err := r.db.
		Delete(&entity.Post{ID: id}).Error
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	return nil
}
