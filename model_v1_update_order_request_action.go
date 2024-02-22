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

type V1UpdateOrderRequestAction string

// List of V1UpdateOrderRequestAction
const (
	COMPLETE_V1UpdateOrderRequestAction V1UpdateOrderRequestAction = "COMPLETE"
	CANCEL_V1UpdateOrderRequestAction V1UpdateOrderRequestAction = "CANCEL"
	REFUND_V1UpdateOrderRequestAction V1UpdateOrderRequestAction = "REFUND"
)
