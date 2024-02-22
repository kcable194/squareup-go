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

// Represents details about a `TRANSFER_BALANCE_TO` [gift card activity type](https://developer.squareup.com/reference/square_2024-01-18/objects/GiftCardActivityType).
type GiftCardActivityTransferBalanceTo struct {
	// The ID of the gift card from which the specified amount was transferred.
	TransferFromGiftCardId string `json:"transfer_from_gift_card_id"`
	AmountMoney *Money `json:"amount_money"`
}
