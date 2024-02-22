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

// Represents an [AccumulateLoyaltyPoints](https://developer.squareup.com/reference/square_2024-01-18/loyalty-api/accumulate-loyalty-points) response.
type AccumulateLoyaltyPointsResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	Event *LoyaltyEvent `json:"event,omitempty"`
	// The resulting loyalty events. If the purchase qualifies for points, the `ACCUMULATE_POINTS` event is always included. When using the Orders API, the `ACCUMULATE_PROMOTION_POINTS` event is included if the purchase also qualifies for a loyalty promotion.
	Events []LoyaltyEvent `json:"events,omitempty"`
}