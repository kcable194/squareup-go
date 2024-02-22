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

// LoyaltyPromotionIncentiveType : Indicates the type of points incentive for a [loyalty promotion](https://developer.squareup.com/reference/square_2024-01-18/objects/LoyaltyPromotion), which is used to determine how buyers can earn points from the promotion.
type LoyaltyPromotionIncentiveType string

// List of LoyaltyPromotionIncentiveType
const (
	MULTIPLIER_LoyaltyPromotionIncentiveType LoyaltyPromotionIncentiveType = "POINTS_MULTIPLIER"
	ADDITION_LoyaltyPromotionIncentiveType   LoyaltyPromotionIncentiveType = "POINTS_ADDITION"
)
