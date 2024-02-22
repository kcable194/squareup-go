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

// The response object returned by the [UpdateLocation](https://developer.squareup.com/reference/square_2024-01-18/locations-api/update-location) endpoint.
type UpdateLocationResponse struct {
	// Information about errors encountered during the request.
	Errors   []ModelError `json:"errors,omitempty"`
	Location *Location    `json:"location,omitempty"`
}
