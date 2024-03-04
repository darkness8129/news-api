package service

import (
	"context"
	"darkness8129/news-api/app/entity"
	"darkness8129/news-api/app/service/mocks"
	"darkness8129/news-api/packages/logging"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPostService_Create(t *testing.T) {
	t.Parallel()

	logger, err := logging.NewZapLogger()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	testCases := []struct {
		name      string
		mock      func(m *mocks.PostStorage)
		input     CreatePostOpt
		expected  *entity.Post
		expectErr bool
	}{
		{
			name: "Create",
			mock: func(m *mocks.PostStorage) {
				m.On("Create", context.Background(), &entity.Post{
					Title:   "title",
					Content: "content",
				}).Return(&entity.Post{
					ID:      uuid.NewString(),
					Title:   "title",
					Content: "content",
				}, nil)
			},
			input: CreatePostOpt{
				Title:   "title",
				Content: "content",
			},
			expected: &entity.Post{
				Title:   "title",
				Content: "content",
			},
		},
		{
			name: "Create with unexpected error in storage",
			mock: func(m *mocks.PostStorage) {
				m.On("Create", context.Background(), &entity.Post{
					Title:   "title",
					Content: "content",
				}).Return(nil, errors.New("error!"))
			},
			input: CreatePostOpt{
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

			postStorageMock := mocks.NewPostStorage(t)
			tc.mock(postStorageMock)
			storages := Storages{Post: postStorageMock}

			postService := NewPostService(storages, logger)
			actual, err := postService.Create(context.Background(), tc.input)
			if !tc.expectErr {
				require.NoError(t, err, "failed to create post")
				require.NotEmpty(t, actual, "post is empty")
			} else {
				require.Error(t, err, "no error")
				require.Nil(t, actual, "post is not nil")
			}
		})
	}
}

func TestPostService_List(t *testing.T) {
	t.Parallel()

	logger, err := logging.NewZapLogger()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	testCases := []struct {
		name        string
		mock        func(m *mocks.PostStorage)
		expectedLen int
		expectErr   bool
	}{
		{
			name: "List",
			mock: func(m *mocks.PostStorage) {
				m.On("List", context.Background()).
					Return([]entity.Post{
						{
							Title:   "title",
							Content: "content",
						},
						{
							Title:   "title",
							Content: "content",
						},
					}, nil)
			},
			expectedLen: 2,
		},
		{
			name: "Create with unexpected error in storage",
			mock: func(m *mocks.PostStorage) {
				m.On("List", context.Background()).Return(nil, errors.New("error!"))
			},
			expectedLen: 0,
			expectErr:   true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			postStorageMock := mocks.NewPostStorage(t)
			tc.mock(postStorageMock)
			storages := Storages{Post: postStorageMock}

			postService := NewPostService(storages, logger)
			actual, err := postService.List(context.Background())
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

func TestPostService_Get(t *testing.T) {
	t.Parallel()

	logger, err := logging.NewZapLogger()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	postID := uuid.NewString()
	wrongID := uuid.NewString()
	post := &entity.Post{
		ID:      postID,
		Title:   "title",
		Content: "content",
	}

	testCases := []struct {
		name      string
		mock      func(m *mocks.PostStorage)
		inputID   string
		expected  *entity.Post
		expectErr bool
	}{
		{
			name: "Get",
			mock: func(m *mocks.PostStorage) {
				m.On("Get", context.Background(), postID).Return(post, nil)
			},
			inputID:  postID,
			expected: post,
		},
		{
			name: "Get with wrong ID",
			mock: func(m *mocks.PostStorage) {
				m.On("Get", context.Background(), wrongID).Return(nil, ErrGetPostNotFound)
			},
			inputID:   wrongID,
			expectErr: true,
		},
		{
			name: "Get with unexpected error in storage",
			mock: func(m *mocks.PostStorage) {
				m.On("Get", context.Background(), postID).Return(nil, errors.New("error!"))
			},
			inputID:   postID,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			postStorageMock := mocks.NewPostStorage(t)
			tc.mock(postStorageMock)
			storages := Storages{Post: postStorageMock}

			postService := NewPostService(storages, logger)
			actual, err := postService.Get(context.Background(), tc.inputID)
			if !tc.expectErr {
				require.NoError(t, err, "failed to get post")
				require.Equal(t, tc.expected.ID, actual.ID, "IDs are not equal")
			} else {
				require.Error(t, err, "no error")
				require.Nil(t, actual, "post is not nil")
			}
		})
	}
}

func TestPostService_Update(t *testing.T) {
	t.Parallel()

	logger, err := logging.NewZapLogger()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	postID := uuid.NewString()
	invalidID := "invalid"
	input := UpdatePostOpt{
		Title:   "updated title",
		Content: "updated content",
	}
	updatedPost := &entity.Post{
		ID:      postID,
		Title:   "updated title",
		Content: "updated content",
	}

	testCases := []struct {
		name      string
		mock      func(m *mocks.PostStorage)
		input     UpdatePostOpt
		inputID   string
		expected  *entity.Post
		expectErr bool
	}{
		{
			name: "Update",
			mock: func(m *mocks.PostStorage) {
				m.On("Update", context.Background(), postID, &entity.Post{
					Title:   "updated title",
					Content: "updated content",
				}).Return(updatedPost, nil)
			},
			input:    input,
			expected: updatedPost,
			inputID:  postID,
		},
		{
			name: "Update with invalid ID",
			mock: func(m *mocks.PostStorage) {
				m.On("Update", context.Background(), invalidID, &entity.Post{
					Title:   "updated title",
					Content: "updated content",
				}).Return(nil, errors.New("invalid id"))
			},
			input:     input,
			inputID:   invalidID,
			expectErr: true,
		},
		{
			name: "Update with unexpected error in storage",
			mock: func(m *mocks.PostStorage) {
				m.On("Update", context.Background(), postID, &entity.Post{
					Title:   "updated title",
					Content: "updated content",
				}).Return(nil, errors.New("error!"))
			},
			input:     input,
			inputID:   postID,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			postStorageMock := mocks.NewPostStorage(t)
			tc.mock(postStorageMock)
			storages := Storages{Post: postStorageMock}

			postService := NewPostService(storages, logger)
			actual, err := postService.Update(context.Background(), tc.inputID, tc.input)
			if !tc.expectErr {
				require.NoError(t, err, "failed to update post")
				require.Equal(t, tc.expected.ID, actual.ID, "IDs are not equal")
				require.Equal(t, tc.expected.Title, actual.Title, "titles are not equal")
				require.Equal(t, tc.expected.Content, actual.Content, "content is not equal")
			} else {
				require.Error(t, err, "no error")
				require.Nil(t, actual, "post is not nil")
			}
		})
	}
}

func TestPostService_Delete(t *testing.T) {
	t.Parallel()

	logger, err := logging.NewZapLogger()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	postID := uuid.NewString()

	testCases := []struct {
		name      string
		mock      func(m *mocks.PostStorage)
		inputID   string
		expectErr bool
	}{
		{
			name: "Delete",
			mock: func(m *mocks.PostStorage) {
				m.On("Delete", context.Background(), postID).Return(nil)
			},
			inputID: postID,
		},
		{
			name: "Update with unexpected error in storage",
			mock: func(m *mocks.PostStorage) {
				m.On("Delete", context.Background(), postID).Return(errors.New("error!"))
			},
			inputID:   postID,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			postStorageMock := mocks.NewPostStorage(t)
			tc.mock(postStorageMock)
			storages := Storages{Post: postStorageMock}

			postService := NewPostService(storages, logger)
			err = postService.Delete(context.Background(), tc.inputID)
			if !tc.expectErr {
				require.NoError(t, err, "failed to delete post")
			} else {
				require.Error(t, err, "no error")
			}
		})
	}
}
