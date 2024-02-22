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
// GiftCardActivityClearBalanceReason : Indicates the reason for clearing the balance of a [gift card](https://developer.squareup.com/reference/square_2024-01-18/objects/GiftCard).
type GiftCardActivityClearBalanceReason string

// List of GiftCardActivityClearBalanceReason
const (
	SUSPICIOUS_ACTIVITY_GiftCardActivityClearBalanceReason GiftCardActivityClearBalanceReason = "SUSPICIOUS_ACTIVITY"
	REUSE_GIFTCARD_GiftCardActivityClearBalanceReason GiftCardActivityClearBalanceReason = "REUSE_GIFTCARD"
	UNKNOWN_REASON_GiftCardActivityClearBalanceReason GiftCardActivityClearBalanceReason = "UNKNOWN_REASON"
)
