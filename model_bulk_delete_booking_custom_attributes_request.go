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

// Represents a [BulkDeleteBookingCustomAttributes](https://developer.squareup.com/reference/square_2024-01-18/booking-custom-attributes-api/bulk-delete-booking-custom-attributes) request.
type BulkDeleteBookingCustomAttributesRequest struct {
	// A map containing 1 to 25 individual Delete requests. For each request, provide an arbitrary ID that is unique for this `BulkDeleteBookingCustomAttributes` request and the information needed to delete a custom attribute.
	Values map[string]BookingCustomAttributeDeleteRequest `json:"values"`
}
