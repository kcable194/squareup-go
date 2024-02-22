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

// Represents an [UpdateCustomerCustomAttributeDefinition](https://developer.squareup.com/reference/square_2024-01-18/customer-custom-attributes-api/update-customer-custom-attribute-definition) response. Either `custom_attribute_definition` or `errors` is present in the response.
type UpdateCustomerCustomAttributeDefinitionResponse struct {
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
