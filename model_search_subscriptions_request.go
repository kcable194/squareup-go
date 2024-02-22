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

// Defines input parameters in a request to the  [SearchSubscriptions](https://developer.squareup.com/reference/square_2024-01-18/subscriptions-api/search-subscriptions) endpoint.
type SearchSubscriptionsRequest struct {
	// When the total number of resulting subscriptions exceeds the limit of a paged response,  specify the cursor returned from a preceding response here to fetch the next set of results. If the cursor is unset, the response contains the last page of the results.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
	// The upper limit on the number of subscriptions to return in a paged response.
	Limit int32 `json:"limit,omitempty"`
	Query *SearchSubscriptionsQuery `json:"query,omitempty"`
	// An option to include related information in the response.   The supported values are:   - `actions`: to include scheduled actions on the targeted subscriptions.
	Include []string `json:"include,omitempty"`
}
