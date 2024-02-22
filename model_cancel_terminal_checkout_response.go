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

type CancelTerminalCheckoutResponse struct {
	// Information about errors encountered during the request.
	Errors   []ModelError      `json:"errors,omitempty"`
	Checkout *TerminalCheckout `json:"checkout,omitempty"`
}
