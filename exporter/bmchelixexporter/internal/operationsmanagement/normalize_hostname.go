// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package operationsmanagement // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/bmchelixexporter/internal/operationsmanagement"

import (
	"strings"
	"unicode"
)

// NormalizeHostname normalizes the hostname to contain only valid characters.
// Allowed characters: Unicode letters, Unicode digits, hyphen (-), and dot (.).
// Invalid characters are replaced with hyphens (DNS convention).
func NormalizeHostname(hostname string) string {
	if hostname == "" {
		return hostname
	}

	// Replace all invalid characters with hyphens
	hostname = strings.Map(sanitizeHostnameRune, hostname)

	// Collapse consecutive hyphens
	for strings.Contains(hostname, "--") {
		hostname = strings.ReplaceAll(hostname, "--", "-")
	}

	// Trim leading and trailing hyphens
	hostname = strings.Trim(hostname, "-")

	return hostname
}

// sanitizeHostnameRune returns the rune if it's valid for a hostname,
// otherwise returns '-'.
// Valid characters: Unicode letters, Unicode digits, hyphen, dot
func sanitizeHostnameRune(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '.' {
		return r
	}
	return '-'
}
