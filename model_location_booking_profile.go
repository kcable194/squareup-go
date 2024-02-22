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

// The booking profile of a seller's location, including the location's ID and whether the location is enabled for online booking.
type LocationBookingProfile struct {
	// The ID of the [location](https://developer.squareup.com/reference/square_2024-01-18/objects/Location).
	LocationId string `json:"location_id,omitempty"`
	// Url for the online booking site for this location.
	BookingSiteUrl string `json:"booking_site_url,omitempty"`
	// Indicates whether the location is enabled for online booking.
	OnlineBookingEnabled bool `json:"online_booking_enabled,omitempty"`
}