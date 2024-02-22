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

// A request for a filtered and sorted set of `Shift` objects.
type SearchShiftsRequest struct {
	Query *ShiftQuery `json:"query,omitempty"`
	// The number of resources in a page (200 by default).
	Limit int32 `json:"limit,omitempty"`
	// An opaque cursor for fetching the next page.
	Cursor string `json:"cursor,omitempty"`
}
