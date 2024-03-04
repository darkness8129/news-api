package errs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrs_Error(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		inputErr error
		expected string
	}{
		{
			name: "Error",
			inputErr: New(Options{
				Message: "msg",
			}),
			expected: "msg",
		},
		{
			name:     "Error with no msg",
			inputErr: New(Options{}),
			expected: "",
		},
		{
			name: "Error with empty msg",
			inputErr: New(Options{
				Message: "",
			}),
			expected: "",
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := tc.inputErr.Error()
			assert.Equal(t, tc.expected, actual, "messages are not equal; got: %s, want: %s", actual, tc.expected)
		})
	}
}

func TestErrs_IsCustom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		inputErr error
		expected bool
	}{
		{
			name: "IsCustom",
			inputErr: New(Options{
				Message: "msg",
				Code:    "code",
			}),
			expected: true,
		},
		{
			name:     "IsCustom with not custom error",
			inputErr: errors.New("not custom"),
			expected: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := IsCustom(tc.inputErr)
			assert.Equal(t, tc.expected, actual, "not equal; got: %t, want: %t", actual, tc.expected)
		})
	}
}

func TestErrs_GetCode(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		inputErr error
		expected string
	}{
		{
			name: "GetCode",
			inputErr: New(Options{
				Message: "msg",
				Code:    "code",
			}),
			expected: "code",
		},
		{
			name: "GetCode with no code",
			inputErr: New(Options{
				Message: "msg",
			}),
			expected: "",
		},
		{
			name: "GetCode with empty code",
			inputErr: New(Options{
				Message: "msg",
				Code:    "",
			}),
			expected: "",
		},
		{
			name:     "GetCode with not custom error",
			inputErr: errors.New("not custom"),
			expected: "",
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := Code(tc.inputErr)
			assert.Equal(t, tc.expected, actual, "codes are not equal; got: %s, want: %s", actual, tc.expected)
		})
	}
}
