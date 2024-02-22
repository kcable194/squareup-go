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

// A query filter to search for buyer-accessible appointment segments by.
type SegmentFilter struct {
	// The ID of the [CatalogItemVariation](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogItemVariation) object representing the service booked in this segment.
	ServiceVariationId string `json:"service_variation_id"`
	TeamMemberIdFilter *FilterValue `json:"team_member_id_filter,omitempty"`
}