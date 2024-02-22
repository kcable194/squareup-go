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

// Represents a [BulkUpsertCustomerCustomAttributes](https://developer.squareup.com/reference/square_2024-01-18/customer-custom-attributes-api/bulk-upsert-customer-custom-attributes) request.
type BulkUpsertCustomerCustomAttributesRequest struct {
	// A map containing 1 to 25 individual upsert requests. For each request, provide an arbitrary ID that is unique for this `BulkUpsertCustomerCustomAttributes` request and the information needed to create or update a custom attribute.
	Values map[string]CustomerCustomAttributeUpsertRequest `json:"values"`
}