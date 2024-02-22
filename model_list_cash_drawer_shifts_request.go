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

type ListCashDrawerShiftsRequest struct {
	// The ID of the location to query for a list of cash drawer shifts.
	LocationId string `json:"location_id"`
	// The order in which cash drawer shifts are listed in the response, based on their opened_at field. Default value: ASC
	SortOrder string `json:"sort_order,omitempty"`
	// The inclusive start time of the query on opened_at, in ISO 8601 format.
	BeginTime string `json:"begin_time,omitempty"`
	// The exclusive end date of the query on opened_at, in ISO 8601 format.
	EndTime string `json:"end_time,omitempty"`
	// Number of cash drawer shift events in a page of results (200 by default, 1000 max).
	Limit int32 `json:"limit,omitempty"`
	// Opaque cursor for fetching the next page of results.
	Cursor string `json:"cursor,omitempty"`
}
