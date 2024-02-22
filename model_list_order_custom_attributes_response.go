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

// Represents a response from listing order custom attributes.
type ListOrderCustomAttributesResponse struct {
	// The retrieved custom attributes. If no custom attribute are found, Square returns an empty object (`{}`).
	CustomAttributes []CustomAttribute `json:"custom_attributes,omitempty"`
	// The cursor to provide in your next call to this endpoint to retrieve the next page of results for your original request.  This field is present only if the request succeeded and additional results are available. For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Cursor string `json:"cursor,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
