package tabletests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Putri",
			request:  "Putri",
			expected: "Hello, Putri",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello, Joko",
		},
		{
			name:     "Raju",
			request:  "Raju",
			expected: "Hello, Raju",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}
