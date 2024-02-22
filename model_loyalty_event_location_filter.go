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

// Filter events by location.
type LoyaltyEventLocationFilter struct {
	// The [location](https://developer.squareup.com/reference/square_2024-01-18/objects/Location) IDs for loyalty events to query. If multiple values are specified, the endpoint uses  a logical OR to combine them.
	LocationIds []string `json:"location_ids"`
}
