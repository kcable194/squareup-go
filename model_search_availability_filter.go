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

// A query filter to search for buyer-accessible availabilities by.
type SearchAvailabilityFilter struct {
	StartAtRange *TimeRange `json:"start_at_range"`
	// The query expression to search for buyer-accessible availabilities with their location IDs matching the specified location ID. This query expression cannot be set if `booking_id` is set.
	LocationId string `json:"location_id,omitempty"`
	// The query expression to search for buyer-accessible availabilities matching the specified list of segment filters. If the size of the `segment_filters` list is `n`, the search returns availabilities with `n` segments per availability.  This query expression cannot be set if `booking_id` is set.
	SegmentFilters []SegmentFilter `json:"segment_filters,omitempty"`
	// The query expression to search for buyer-accessible availabilities for an existing booking by matching the specified `booking_id` value. This is commonly used to reschedule an appointment. If this expression is set, the `location_id` and `segment_filters` expressions cannot be set.
	BookingId string `json:"booking_id,omitempty"`
}
