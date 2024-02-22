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

// A response that contains a `GiftCard`. The response might contain a set of `Error` objects if the request resulted in errors.
type CreateGiftCardResponse struct {
	// Any errors that occurred during the request.
	Errors   []ModelError `json:"errors,omitempty"`
	GiftCard *GiftCard    `json:"gift_card,omitempty"`
}
