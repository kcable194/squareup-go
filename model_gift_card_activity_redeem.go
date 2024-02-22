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

// Represents details about a `REDEEM` [gift card activity type](https://developer.squareup.com/reference/square_2024-01-18/objects/GiftCardActivityType).
type GiftCardActivityRedeem struct {
	AmountMoney *Money `json:"amount_money"`
	// The ID of the payment that represents the gift card redemption. Square populates this field  if the payment was processed by Square.
	PaymentId string `json:"payment_id,omitempty"`
	// A client-specified ID that associates the gift card activity with an entity in another system.   Applications that use a custom payment processing system can use this field to track information related to an order or payment.
	ReferenceId string `json:"reference_id,omitempty"`
	// The status of the gift card redemption. Gift cards redeemed from Square Point of Sale or the  Square Seller Dashboard use a two-state process: `PENDING`  to `COMPLETED` or `PENDING` to  `CANCELED`. Gift cards redeemed using the Gift Card Activities API  always have a `COMPLETED` status.
	Status string `json:"status,omitempty"`
}