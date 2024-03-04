package storage

import (
	"context"
	"darkness8129/news-api/app/entity"
	"darkness8129/news-api/app/service"
	"darkness8129/news-api/config"
	"darkness8129/news-api/packages/database"
	"darkness8129/news-api/packages/logging"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	storage service.PostStorage
)

func init() {
	logger, err := logging.NewZapLogger()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cfg, err := config.New()
	if err != nil {
		logger.Fatal("failed to get config", "err", err)
	}

	sql, err := database.NewPostgreSQLDatabase(database.Options{
		User:     cfg.Test.PostgreSQLUser,
		Password: cfg.Test.PostgreSQLPassword,
		Database: cfg.Test.PostgreSQLDatabase,
		Port:     cfg.Test.PostgreSQLPort,
		Host:     cfg.Test.PostgreSQLHost,
		Logger:   logger,
	})
	if err != nil {
		logger.Fatal("failed to init postgresql db", "err", err)
	}

	DB, ok := sql.DB().(*gorm.DB)
	if !ok {
		logger.Fatal("failed type assertion for db")
	}

	err = DB.AutoMigrate(&entity.Post{})
	if err != nil {
		logger.Fatal("automigration failed", "err", err)
	}

	storage = NewPostStorage(DB, logger)
	db = DB
}

func TestPostStorage_Create(t *testing.T) {
	testCases := []struct {
		name      string
		input     *entity.Post
		expected  *entity.Post
		expectErr bool
	}{
		{
			name: "Create",
			input: &entity.Post{
				Title:   "title",
				Content: "content",
			},
			expected: &entity.Post{
				Title:   "title",
				Content: "content",
			},
		},
		{
			name: "Create with ID",
			input: &entity.Post{
				ID:      uuid.NewString(),
				Title:   "title",
				Content: "content",
			},
			expected: &entity.Post{
				Title:   "title",
				Content: "content",
			},
		},
		{
			name:      "Create without post",
			input:     nil,
			expectErr: true,
		},
		{
			name: "Create with invalid ID",
			input: &entity.Post{
				ID:      "invalid",
				Title:   "title",
				Content: "content",
			},
			expectErr: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			t.Cleanup(func() {
				err := db.Exec("DELETE FROM posts;").Error
				require.NoError(t, err, "failed to delete posts")
			})

			actual, err := storage.Create(context.Background(), tc.input)
			if !tc.expectErr {
				require.NoError(t, err, "failed to create post")
				require.NotEmpty(t, actual.ID, "id is empty")
				require.Equal(t, tc.expected.Title, actual.Title, "titles are not equal")
				require.Equal(t, tc.expected.Content, actual.Content, "content is not equal")
				require.NotEmpty(t, actual.CreatedAt, "createdAt is empty")
				require.NotEmpty(t, actual.UpdatedAt, "updatedAt is empty")
				require.Empty(t, actual.DeletedAt, "deletedAt is not empty")
			} else {
				require.Error(t, err, "no error")
				require.Nil(t, actual, "post is not nil")
			}
		})
	}
}

func TestPostStorage_List(t *testing.T) {
	testCases := []struct {
		name          string
		postsToCreate []entity.Post
		expectedLen   int
		expectErr     bool
	}{
		{
			name:          "List 0 posts",
			postsToCreate: []entity.Post{},
			expectedLen:   0,
		},
		{
			name: "List 2 posts",
			postsToCreate: []entity.Post{
				entity.Post{
					Title:   "title",
					Content: "content",
				},
				entity.Post{
					Title:   "title",
					Content: "content",
				},
			},
			expectedLen: 2,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Cleanup(func() {
				err := db.Exec("DELETE FROM posts;").Error
				require.NoError(t, err)
			})

			for _, p := range tc.postsToCreate {
				_, err := storage.Create(context.Background(), &p)
				require.NoError(t, err, "failed to create post")
			}

			actual, err := storage.List(context.Background())
			if !tc.expectErr {
				require.NoError(t, err, "failed to list posts")
				require.Equal(t, tc.expectedLen, len(actual), "len is not equal")
			} else {
				require.Error(t, err, "no error")
				require.Empty(t, actual, "slice is not empty")
			}
		})
	}
}
