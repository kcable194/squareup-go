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

// Represents a [BulkDeleteMerchantCustomAttributes](https://developer.squareup.com/reference/square_2024-01-18/merchant-custom-attributes-api/bulk-delete-merchant-custom-attributes) request.
type BulkDeleteMerchantCustomAttributesRequest struct {
	// The data used to update the `CustomAttribute` objects. The keys must be unique and are used to map to the corresponding response.
	Values map[string]MerchantCustomAttributeDeleteRequest `json:"values"`
}