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
// SubscriptionPricingType : Determines the pricing of a [Subscription](https://developer.squareup.com/reference/square_2024-01-18/objects/Subscription)
type SubscriptionPricingType string

// List of SubscriptionPricingType
const (
	STATIC_SubscriptionPricingType SubscriptionPricingType = "STATIC"
	RELATIVE_SubscriptionPricingType SubscriptionPricingType = "RELATIVE"
)