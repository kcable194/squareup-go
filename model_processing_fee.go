/*
 * Square Connect API
 *
 * Client library for accessing the Square Connect APIs
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package squareup

// Represents the Square processing fee.
type ProcessingFee struct {
	// The timestamp of when the fee takes effect, in RFC 3339 format.
	EffectiveAt string `json:"effective_at,omitempty"`
	// The type of fee assessed or adjusted. The fee type can be `INITIAL` or `ADJUSTMENT`.
	Type_       string `json:"type,omitempty"`
	AmountMoney *Money `json:"amount_money,omitempty"`
}
