/*
 * Square Connect API
 *
 * Client library for accessing the Square Connect APIs
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// V1PaymentTax
type V1PaymentTax struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The merchant-defined name of the tax.
	Name string `json:"name,omitempty"`
	AppliedMoney *V1Money `json:"applied_money,omitempty"`
	// The rate of the tax, as a string representation of a decimal number. A value of 0.07 corresponds to a rate of 7%.
	Rate string `json:"rate,omitempty"`
	// Whether the tax is an ADDITIVE tax or an INCLUSIVE tax.
	InclusionType string `json:"inclusion_type,omitempty"`
	// The ID of the tax, if available. Taxes applied in older versions of Square Register might not have an ID.
	FeeId string `json:"fee_id,omitempty"`
}
