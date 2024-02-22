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

// A rounding adjustment of the money being returned. Commonly used to apply cash rounding when the minimum unit of the account is smaller than the lowest physical denomination of the currency.
type OrderRoundingAdjustment struct {
	// A unique ID that identifies the rounding adjustment only within this order.
	Uid string `json:"uid,omitempty"`
	// The name of the rounding adjustment from the original sale order.
	Name string `json:"name,omitempty"`
	AmountMoney *Money `json:"amount_money,omitempty"`
}