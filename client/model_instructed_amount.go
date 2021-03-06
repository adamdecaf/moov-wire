/*
 * Wire API
 *
 * Moov Wire implements an HTTP API for creating, parsing, and validating Fedwire messages.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// InstructedAmount struct for InstructedAmount
type InstructedAmount struct {
	// CurrencyCode
	CurrencyCode string `json:"currencyCode,omitempty"`
	// Amount  Must begin with at least one numeric character (0-9) and contain only one decimal comma marker (e.g., $1,234.56 should be entered as 1234,56 and $0.99 should be entered as 0,99).
	Amount string `json:"amount,omitempty"`
}
