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

// Defines the fields that the [RetrieveLocation](https://developer.squareup.com/reference/square_2024-01-18/locations-api/retrieve-location) endpoint returns in a response.
type RetrieveLocationResponse struct {
	// Information about errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	Location *Location `json:"location,omitempty"`
}
