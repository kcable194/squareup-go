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

// Provides metadata when the event `type` is `CREATE_REWARD`.
type LoyaltyEventCreateReward struct {
	// The ID of the [loyalty program](https://developer.squareup.com/reference/square_2024-01-18/objects/LoyaltyProgram).
	LoyaltyProgramId string `json:"loyalty_program_id"`
	// The Square-assigned ID of the created [loyalty reward](https://developer.squareup.com/reference/square_2024-01-18/objects/LoyaltyReward). This field is returned only if the event source is `LOYALTY_API`.
	RewardId string `json:"reward_id,omitempty"`
	// The loyalty points used to create the reward.
	Points int32 `json:"points"`
}
