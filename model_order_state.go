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

// OrderState : The state of the order.
type OrderState string

// List of OrderState
const (
	OPEN_OrderState      OrderState = "OPEN"
	COMPLETED_OrderState OrderState = "COMPLETED"
	CANCELED_OrderState  OrderState = "CANCELED"
	DRAFT_OrderState     OrderState = "DRAFT"
)
