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

// Represents a [BulkDeleteMerchantCustomAttributes](https://developer.squareup.com/reference/square_2024-01-18/merchant-custom-attributes-api/bulk-delete-merchant-custom-attributes) response, which contains a map of responses that each corresponds to an individual delete request.
type BulkDeleteMerchantCustomAttributesResponse struct {
	// A map of responses that correspond to individual delete requests. Each response has the same key as the corresponding request.
	Values map[string]MerchantCustomAttributeDeleteResponse `json:"values"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
