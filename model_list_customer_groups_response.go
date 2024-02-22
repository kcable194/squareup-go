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

// Defines the fields that are included in the response body of a request to the [ListCustomerGroups](https://developer.squareup.com/reference/square_2024-01-18/customer-groups-api/list-customer-groups) endpoint.  Either `errors` or `groups` is present in a given response (never both).
type ListCustomerGroupsResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// A list of customer groups belonging to the current seller.
	Groups []CustomerGroup `json:"groups,omitempty"`
	// A pagination cursor to retrieve the next set of results for your original query to the endpoint. This value is present only if the request succeeded and additional results are available.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
}
