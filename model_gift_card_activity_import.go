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

// Represents details about an `IMPORT` [gift card activity type](https://developer.squareup.com/reference/square_2024-01-18/objects/GiftCardActivityType). This activity type is used when Square imports a third-party gift card, in which case the  `gan_source` of the gift card is set to `OTHER`.
type GiftCardActivityImport struct {
	AmountMoney *Money `json:"amount_money"`
}
