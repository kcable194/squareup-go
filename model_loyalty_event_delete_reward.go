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

// Provides metadata when the event `type` is `DELETE_REWARD`.
type LoyaltyEventDeleteReward struct {
	// The ID of the [loyalty program](https://developer.squareup.com/reference/square_2024-01-18/objects/LoyaltyProgram).
	LoyaltyProgramId string `json:"loyalty_program_id"`
	// The ID of the deleted [loyalty reward](https://developer.squareup.com/reference/square_2024-01-18/objects/LoyaltyReward). This field is returned only if the event source is `LOYALTY_API`.
	RewardId string `json:"reward_id,omitempty"`
	// The number of points returned to the loyalty account.
	Points int32 `json:"points"`
}