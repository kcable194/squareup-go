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

type DeletePaymentLinkResponse struct {
	Errors []ModelError `json:"errors,omitempty"`
	// The ID of the link that is deleted.
	Id string `json:"id,omitempty"`
	// The ID of the order that is canceled. When a payment link is deleted, Square updates the the `state` (of the order that the checkout link created) to CANCELED.
	CancelledOrderId string `json:"cancelled_order_id,omitempty"`
}