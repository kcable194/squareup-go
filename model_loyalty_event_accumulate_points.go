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

// Provides metadata when the event `type` is `ACCUMULATE_POINTS`.
type LoyaltyEventAccumulatePoints struct {
	// The ID of the [loyalty program](https://developer.squareup.com/reference/square_2024-01-18/objects/LoyaltyProgram).
	LoyaltyProgramId string `json:"loyalty_program_id,omitempty"`
	// The number of points accumulated by the event.
	Points int32 `json:"points,omitempty"`
	// The ID of the [order](https://developer.squareup.com/reference/square_2024-01-18/objects/Order) for which the buyer accumulated the points. This field is returned only if the Orders API is used to process orders.
	OrderId string `json:"order_id,omitempty"`
}