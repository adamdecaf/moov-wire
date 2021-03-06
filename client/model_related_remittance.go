/*
 * Wire API
 *
 * Moov Wire implements an HTTP API for creating, parsing, and validating Fedwire messages.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// RelatedRemittance struct for RelatedRemittance
type RelatedRemittance struct {
	// RemittanceIdentification
	RemittanceIdentification string `json:"remittanceIdentification,omitempty"`
	// RemittanceLocationMethod  * `EDIC` - Electronic Data Interchange * `EMAL` - Email * `FAXI` - Fax * `POST` - Postal services * `SMS` - Short Message Service (text) * `URI` - Uniform Resource Identifier
	RemittanceLocationMethod string `json:"remittanceLocationMethod,omitempty"`
	// RemittanceLocationElectronicAddress (E-mail or URL address)
	RemittanceLocationElectronicAddress string         `json:"remittanceLocationElectronicAddress,omitempty"`
	RemittanceData                      RemittanceData `json:"remittanceData,omitempty"`
}
