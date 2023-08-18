package unpackstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	input    string
	expected string
	err      error
}

func TestUnpack(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abcd",
			expected: "abcd",
		},
		{
			input:    "45",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "aaa10b",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "a2o3f0",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "a2c000",
			expected: "",
			err:      ErrInvalidString,
		},
	} {
		t.Run(tst.input, func(t *testing.T) {
			result, err := Unpack(tst.input)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.expected, result)
		})
	}
}
