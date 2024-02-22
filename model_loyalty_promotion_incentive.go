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

// Represents how points for a [loyalty promotion](https://developer.squareup.com/reference/square_2024-01-18/objects/LoyaltyPromotion) are calculated, either by multiplying the points earned from the base program or by adding a specified number of points to the points earned from the base program.
type LoyaltyPromotionIncentive struct {
	// The type of points incentive.
	Type_ string `json:"type"`
	PointsMultiplierData *LoyaltyPromotionIncentivePointsMultiplierData `json:"points_multiplier_data,omitempty"`
	PointsAdditionData *LoyaltyPromotionIncentivePointsAdditionData `json:"points_addition_data,omitempty"`
}
