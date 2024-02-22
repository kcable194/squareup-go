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

type ListBookingsResponse struct {
	// The list of targeted bookings.
	Bookings []Booking `json:"bookings,omitempty"`
	// The pagination cursor to be used in the subsequent request to get the next page of the results. Stop retrieving the next page of the results when the cursor is not set.
	Cursor string `json:"cursor,omitempty"`
	// Errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}