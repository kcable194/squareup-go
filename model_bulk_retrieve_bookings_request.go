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

// Request payload for bulk retrieval of bookings.
type BulkRetrieveBookingsRequest struct {
	// A non-empty list of [Booking](https://developer.squareup.com/reference/square_2024-01-18/objects/Booking) IDs specifying bookings to retrieve.
	BookingIds []string `json:"booking_ids"`
}
