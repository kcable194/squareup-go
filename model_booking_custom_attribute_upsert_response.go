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

// Represents a response for an individual upsert request in a [BulkUpsertBookingCustomAttributes](https://developer.squareup.com/reference/square_2024-01-18/booking-custom-attributes-api/bulk-upsert-booking-custom-attributes) operation.
type BookingCustomAttributeUpsertResponse struct {
	// The ID of the [booking](https://developer.squareup.com/reference/square_2024-01-18/objects/Booking) associated with the custom attribute.
	BookingId string `json:"booking_id,omitempty"`
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty"`
	// Any errors that occurred while processing the individual request.
	Errors []ModelError `json:"errors,omitempty"`
}
