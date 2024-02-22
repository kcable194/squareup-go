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

// Provides metadata when the event `type` is `EXPIRE_POINTS`.
type LoyaltyEventExpirePoints struct {
	// The Square-assigned ID of the [loyalty program](https://developer.squareup.com/reference/square_2024-01-18/objects/LoyaltyProgram).
	LoyaltyProgramId string `json:"loyalty_program_id"`
	// The number of points expired.
	Points int32 `json:"points"`
}