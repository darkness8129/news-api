package errs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
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
			require.Equal(t, tc.expected, actual, "messages are not equal")
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
			require.Equal(t, tc.expected, actual, "isCustom not equal")
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
			require.Equal(t, tc.expected, actual, "codes are not equal")
		})
	}
}
