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

// Represents a business that sells with Square.
type Merchant struct {
	// The Square-issued ID of the merchant.
	Id string `json:"id,omitempty"`
	// The name of the merchant's overall business.
	BusinessName string `json:"business_name,omitempty"`
	// The country code associated with the merchant, in the two-letter format of ISO 3166. For example, `US` or `JP`.
	Country string `json:"country"`
	// The code indicating the [language preferences](https://developer.squareup.com/docs/build-basics/general-considerations/language-preferences) of the merchant, in [BCP 47 format](https://tools.ietf.org/html/bcp47#appendix-A). For example, `en-US` or `fr-CA`.
	LanguageCode string `json:"language_code,omitempty"`
	// The currency associated with the merchant, in ISO 4217 format. For example, the currency code for US dollars is `USD`.
	Currency string `json:"currency,omitempty"`
	// The merchant's status.
	Status string `json:"status,omitempty"`
	// The ID of the [main `Location`](https://developer.squareup.com/docs/locations-api#about-the-main-location) for this merchant.
	MainLocationId string `json:"main_location_id,omitempty"`
	// The time when the merchant was created, in RFC 3339 format.    For more information, see [Working with Dates](https://developer.squareup.com/docs/build-basics/working-with-dates).
	CreatedAt string `json:"created_at,omitempty"`
}
