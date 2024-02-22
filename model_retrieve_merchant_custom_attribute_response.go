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

// Represents a [RetrieveMerchantCustomAttribute](https://developer.squareup.com/reference/square_2024-01-18/merchant-custom-attributes-api/retrieve-merchant-custom-attribute) response. Either `custom_attribute_definition` or `errors` is present in the response.
type RetrieveMerchantCustomAttributeResponse struct {
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}