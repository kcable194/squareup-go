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

// Represents a [ListLoyaltyPromotions](https://developer.squareup.com/reference/square_2024-01-18/loyalty-api/list-loyalty-promotions) response. One of `loyalty_promotions`, an empty object, or `errors` is present in the response. If additional results are available, the `cursor` field is also present along with `loyalty_promotions`.
type ListLoyaltyPromotionsResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The retrieved loyalty promotions.
	LoyaltyPromotions []LoyaltyPromotion `json:"loyalty_promotions,omitempty"`
	// The cursor to use in your next call to this endpoint to retrieve the next page of results for your original request. This field is present only if the request succeeded and additional results are available. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
}
