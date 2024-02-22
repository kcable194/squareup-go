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

// Represents a bulk delete request for one or more order custom attributes.
type BulkDeleteOrderCustomAttributesRequest struct {
	// A map of requests that correspond to individual delete operations for custom attributes.
	Values map[string]DeleteCustomAttribute `json:"values"`
}