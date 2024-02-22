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

// Provides information about a loyalty event.  For more information, see [Search for Balance-Changing Loyalty Events](https://developer.squareup.com/docs/loyalty-api/loyalty-events).
type LoyaltyEvent struct {
	// The Square-assigned ID of the loyalty event.
	Id string `json:"id"`
	// The type of the loyalty event.
	Type_ string `json:"type"`
	// The timestamp when the event was created, in RFC 3339 format.
	CreatedAt        string                        `json:"created_at"`
	AccumulatePoints *LoyaltyEventAccumulatePoints `json:"accumulate_points,omitempty"`
	CreateReward     *LoyaltyEventCreateReward     `json:"create_reward,omitempty"`
	RedeemReward     *LoyaltyEventRedeemReward     `json:"redeem_reward,omitempty"`
	DeleteReward     *LoyaltyEventDeleteReward     `json:"delete_reward,omitempty"`
	AdjustPoints     *LoyaltyEventAdjustPoints     `json:"adjust_points,omitempty"`
	// The ID of the [loyalty account](https://developer.squareup.com/reference/square_2024-01-18/objects/LoyaltyAccount) associated with the event.
	LoyaltyAccountId string `json:"loyalty_account_id"`
	// The ID of the [location](https://developer.squareup.com/reference/square_2024-01-18/objects/Location) where the event occurred.
	LocationId string `json:"location_id,omitempty"`
	// Defines whether the event was generated by the Square Point of Sale.
	Source                    string                                 `json:"source"`
	ExpirePoints              *LoyaltyEventExpirePoints              `json:"expire_points,omitempty"`
	OtherEvent                *LoyaltyEventOther                     `json:"other_event,omitempty"`
	AccumulatePromotionPoints *LoyaltyEventAccumulatePromotionPoints `json:"accumulate_promotion_points,omitempty"`
}
