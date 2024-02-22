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

// GiftCardActivityAdjustIncrementReason : Indicates the reason for adding money to a [gift card](https://developer.squareup.com/reference/square_2024-01-18/objects/GiftCard).
type GiftCardActivityAdjustIncrementReason string

// List of GiftCardActivityAdjustIncrementReason
const (
	COMPLIMENTARY_GiftCardActivityAdjustIncrementReason      GiftCardActivityAdjustIncrementReason = "COMPLIMENTARY"
	SUPPORT_ISSUE_GiftCardActivityAdjustIncrementReason      GiftCardActivityAdjustIncrementReason = "SUPPORT_ISSUE"
	TRANSACTION_VOIDED_GiftCardActivityAdjustIncrementReason GiftCardActivityAdjustIncrementReason = "TRANSACTION_VOIDED"
)
