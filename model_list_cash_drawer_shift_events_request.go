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

type ListCashDrawerShiftEventsRequest struct {
	// The ID of the location to list cash drawer shifts for.
	LocationId string `json:"location_id"`
	// Number of resources to be returned in a page of results (200 by default, 1000 max).
	Limit int32 `json:"limit,omitempty"`
	// Opaque cursor for fetching the next page of results.
	Cursor string `json:"cursor,omitempty"`
}
