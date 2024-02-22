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

// V1PaymentDiscount
type V1PaymentDiscount struct {
	// The discount's name.
	Name         string   `json:"name,omitempty"`
	AppliedMoney *V1Money `json:"applied_money,omitempty"`
	// The ID of the applied discount, if available. Discounts applied in older versions of Square Register might not have an ID.
	DiscountId string `json:"discount_id,omitempty"`
}
