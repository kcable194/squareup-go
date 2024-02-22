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

// A response that includes the `LoyaltyEvent` published for redeeming the reward.
type RedeemLoyaltyRewardResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError  `json:"errors,omitempty"`
	Event  *LoyaltyEvent `json:"event,omitempty"`
}
