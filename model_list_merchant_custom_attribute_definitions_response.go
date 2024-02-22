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

// Represents a [ListMerchantCustomAttributeDefinitions](https://developer.squareup.com/reference/square_2024-01-18/merchant-custom-attributes-api/list-merchant-custom-attribute-definitions) response. Either `custom_attribute_definitions`, an empty object, or `errors` is present in the response. If additional results are available, the `cursor` field is also present along with `custom_attribute_definitions`.
type ListMerchantCustomAttributeDefinitionsResponse struct {
	// The retrieved custom attribute definitions. If no custom attribute definitions are found, Square returns an empty object (`{}`).
	CustomAttributeDefinitions []CustomAttributeDefinition `json:"custom_attribute_definitions,omitempty"`
	// The cursor to provide in your next call to this endpoint to retrieve the next page of results for your original request. This field is present only if the request succeeded and additional results are available. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
