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

// GiftCardActivityAdjustDecrementReason : Indicates the reason for deducting money from a [gift card](https://developer.squareup.com/reference/square_2024-01-18/objects/GiftCard).
type GiftCardActivityAdjustDecrementReason string

// List of GiftCardActivityAdjustDecrementReason
const (
	SUSPICIOUS_ACTIVITY_GiftCardActivityAdjustDecrementReason            GiftCardActivityAdjustDecrementReason = "SUSPICIOUS_ACTIVITY"
	BALANCE_ACCIDENTALLY_INCREASED_GiftCardActivityAdjustDecrementReason GiftCardActivityAdjustDecrementReason = "BALANCE_ACCIDENTALLY_INCREASED"
	SUPPORT_ISSUE_GiftCardActivityAdjustDecrementReason                  GiftCardActivityAdjustDecrementReason = "SUPPORT_ISSUE"
	PURCHASE_WAS_REFUNDED_GiftCardActivityAdjustDecrementReason          GiftCardActivityAdjustDecrementReason = "PURCHASE_WAS_REFUNDED"
)
