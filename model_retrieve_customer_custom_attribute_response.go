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

// Represents a [RetrieveCustomerCustomAttribute](https://developer.squareup.com/reference/square_2024-01-18/customer-custom-attributes-api/retrieve-customer-custom-attribute) response. Either `custom_attribute_definition` or `errors` is present in the response.
type RetrieveCustomerCustomAttributeResponse struct {
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}