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

// Represents details about a `BLOCK` [gift card activity type](https://developer.squareup.com/reference/square_2024-01-18/objects/GiftCardActivityType).
type GiftCardActivityBlock struct {
	// The reason the gift card was blocked.
	Reason string `json:"reason"`
}
