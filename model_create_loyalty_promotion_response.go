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

// Represents a [CreateLoyaltyPromotion](https://developer.squareup.com/reference/square_2024-01-18/loyalty-api/create-loyalty-promotion) response. Either `loyalty_promotion` or `errors` is present in the response.
type CreateLoyaltyPromotionResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	LoyaltyPromotion *LoyaltyPromotion `json:"loyalty_promotion,omitempty"`
}