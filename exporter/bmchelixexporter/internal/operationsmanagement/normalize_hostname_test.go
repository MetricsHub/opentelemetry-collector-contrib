// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package operationsmanagement

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalizeHostname(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "valid hostname",
			input:    "server-01.example.com",
			expected: "server-01.example.com",
		},
		{
			name:     "hostname with underscores",
			input:    "server_01",
			expected: "server-01",
		},
		{
			name:     "hostname with spaces",
			input:    "server 01",
			expected: "server-01",
		},
		{
			name:     "hostname with colons",
			input:    "server:01:test",
			expected: "server-01-test",
		},
		{
			name:     "hostname with special chars",
			input:    "server@01#test",
			expected: "server-01-test",
		},
		{
			name:     "hostname with consecutive invalid chars",
			input:    "server__01",
			expected: "server-01",
		},
		{
			name:     "hostname with leading/trailing hyphens",
			input:    "-server-01-",
			expected: "server-01",
		},
		{
			name:     "simple hostname",
			input:    "localhost",
			expected: "localhost",
		},
		{
			name:     "IP address style hostname",
			input:    "192.168.1.1",
			expected: "192.168.1.1",
		},
		{
			name:     "hostname with brackets",
			input:    "server[01]",
			expected: "server-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeHostname(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestSanitizeHostnameRune(t *testing.T) {
	// Valid runes should be returned as-is
	require.Equal(t, 'a', sanitizeHostnameRune('a'))
	require.Equal(t, 'Z', sanitizeHostnameRune('Z'))
	require.Equal(t, '5', sanitizeHostnameRune('5'))
	require.Equal(t, '-', sanitizeHostnameRune('-'))
	require.Equal(t, '.', sanitizeHostnameRune('.'))

	// Invalid runes should be replaced with '-'
	require.Equal(t, '-', sanitizeHostnameRune('_'))
	require.Equal(t, '-', sanitizeHostnameRune(' '))
	require.Equal(t, '-', sanitizeHostnameRune(':'))
	require.Equal(t, '-', sanitizeHostnameRune('/'))
	require.Equal(t, '-', sanitizeHostnameRune('@'))
	require.Equal(t, '-', sanitizeHostnameRune('#'))
}
