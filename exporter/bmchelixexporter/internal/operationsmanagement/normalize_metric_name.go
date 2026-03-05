// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package operationsmanagement // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/bmchelixexporter/internal/operationsmanagement"

import (
	"strings"
	"unicode"
)

// NormalizeMetricName normalizes the metric name so that it starts with a Unicode letter or underscore,
// followed by Unicode letters, digits, underscores, colons, or dots.
// Unsafe characters are replaced with underscores.
// If the metric name starts with a digit, it is prefixed with an underscore.
func NormalizeMetricName(name string) string {
	if name == "" {
		return name
	}

	// Replace all unsafe characters with underscores
	name = strings.Map(sanitizeMetricNameRune, name)

	// Collapse consecutive underscores
	for strings.Contains(name, "__") {
		name = strings.ReplaceAll(name, "__", "_")
	}

	// Trim leading and trailing underscores
	name = strings.Trim(name, "_")

	// Metric name cannot start with a digit, so prefix it with "_" in this case
	if name != "" && unicode.IsDigit(rune(name[0])) {
		name = "_" + name
	}

	return name
}

// sanitizeMetricNameRune returns the rune if it's valid for a metric name,
// otherwise returns '_'.
// Valid characters: letters, digits, underscore, colon, dot
func sanitizeMetricNameRune(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == ':' || r == '.' {
		return r
	}
	return '_'
}
