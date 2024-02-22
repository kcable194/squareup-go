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

// Represents an individual delete response in a [BulkDeleteLocationCustomAttributes](https://developer.squareup.com/reference/square_2024-01-18/location-custom-attributes-api/bulk-delete-location-custom-attributes) request.
type BulkDeleteLocationCustomAttributesResponseLocationCustomAttributeDeleteResponse struct {
	// The ID of the location associated with the custom attribute.
	LocationId string `json:"location_id,omitempty"`
	// Errors that occurred while processing the individual LocationCustomAttributeDeleteRequest request
	Errors []ModelError `json:"errors,omitempty"`
}
