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
// LoyaltyPromotionStatus : Indicates the status of a [loyalty promotion](https://developer.squareup.com/reference/square_2024-01-18/objects/LoyaltyPromotion).
type LoyaltyPromotionStatus string

// List of LoyaltyPromotionStatus
const (
	ACTIVE_LoyaltyPromotionStatus LoyaltyPromotionStatus = "ACTIVE"
	ENDED_LoyaltyPromotionStatus LoyaltyPromotionStatus = "ENDED"
	CANCELED_LoyaltyPromotionStatus LoyaltyPromotionStatus = "CANCELED"
	SCHEDULED_LoyaltyPromotionStatus LoyaltyPromotionStatus = "SCHEDULED"
)