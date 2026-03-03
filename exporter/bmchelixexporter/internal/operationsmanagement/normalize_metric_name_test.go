// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package operationsmanagement

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalizeMetricName(t *testing.T) {
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
			name:     "valid metric name",
			input:    "system.cpu.usage",
			expected: "system.cpu.usage",
		},
		{
			name:     "metric name with colons",
			input:    "system:cpu:usage",
			expected: "system:cpu:usage",
		},
		{
			name:     "metric name with underscores",
			input:    "system_cpu_usage",
			expected: "system_cpu_usage",
		},
		{
			name:     "metric name starting with digit",
			input:    "123metric",
			expected: "_123metric",
		},
		{
			name:     "metric name with spaces",
			input:    "system cpu usage",
			expected: "system_cpu_usage",
		},
		{
			name:     "metric name with special chars",
			input:    "system/cpu\\usage",
			expected: "system_cpu_usage",
		},
		{
			name:     "metric name with brackets",
			input:    "system[cpu](usage)",
			expected: "system_cpu_usage",
		},
		{
			name:     "metric name with consecutive special chars",
			input:    "system..cpu//usage",
			expected: "system..cpu_usage",
		},
		{
			name:     "metric name with leading/trailing underscores",
			input:    "_system_cpu_usage_",
			expected: "system_cpu_usage",
		},
		{
			name:     "metric name with equals and semicolons",
			input:    "metric=value;test",
			expected: "metric_value_test",
		},
		{
			name:     "metric name from enriched attribute - eth0 receive",
			input:    "system.network.io.eth0.receive",
			expected: "system.network.io.eth0.receive",
		},
		{
			name:     "metric name with curly braces",
			input:    "metric{label}",
			expected: "metric_label",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeMetricName(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestSanitizeMetricNameRune(t *testing.T) {
	// Valid runes should be returned as-is
	require.Equal(t, 'a', sanitizeMetricNameRune('a'))
	require.Equal(t, 'Z', sanitizeMetricNameRune('Z'))
	require.Equal(t, '5', sanitizeMetricNameRune('5'))
	require.Equal(t, '_', sanitizeMetricNameRune('_'))
	require.Equal(t, ':', sanitizeMetricNameRune(':'))
	require.Equal(t, '.', sanitizeMetricNameRune('.'))

	// Invalid runes should be replaced with '_'
	require.Equal(t, '_', sanitizeMetricNameRune(' '))
	require.Equal(t, '_', sanitizeMetricNameRune('/'))
	require.Equal(t, '_', sanitizeMetricNameRune('\\'))
	require.Equal(t, '_', sanitizeMetricNameRune('['))
	require.Equal(t, '_', sanitizeMetricNameRune(']'))
	require.Equal(t, '_', sanitizeMetricNameRune('{'))
	require.Equal(t, '_', sanitizeMetricNameRune('}'))
}
