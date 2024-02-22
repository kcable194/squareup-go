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

// Represents a response from listing order custom attribute definitions.
type ListOrderCustomAttributeDefinitionsResponse struct {
	// The retrieved custom attribute definitions. If no custom attribute definitions are found, Square returns an empty object (`{}`).
	CustomAttributeDefinitions []CustomAttributeDefinition `json:"custom_attribute_definitions"`
	// The cursor to provide in your next call to this endpoint to retrieve the next page of results for your original request.  This field is present only if the request succeeded and additional results are available. For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Cursor string `json:"cursor,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
