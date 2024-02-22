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
// GiftCardActivityDeactivateReason : Indicates the reason for deactivating a [gift card](https://developer.squareup.com/reference/square_2024-01-18/objects/GiftCard).
type GiftCardActivityDeactivateReason string

// List of GiftCardActivityDeactivateReason
const (
	SUSPICIOUS_ACTIVITY_GiftCardActivityDeactivateReason GiftCardActivityDeactivateReason = "SUSPICIOUS_ACTIVITY"
	UNKNOWN_REASON_GiftCardActivityDeactivateReason GiftCardActivityDeactivateReason = "UNKNOWN_REASON"
	CHARGEBACK_DEACTIVATE_GiftCardActivityDeactivateReason GiftCardActivityDeactivateReason = "CHARGEBACK_DEACTIVATE"
)