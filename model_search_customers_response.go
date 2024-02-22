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

// Defines the fields that are included in the response body of a request to the `SearchCustomers` endpoint.  Either `errors` or `customers` is present in a given response (never both).
type SearchCustomersResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The customer profiles that match the search query. If any search condition is not met, the result is an empty object (`{}`). Only customer profiles with public information (`given_name`, `family_name`, `company_name`, `email_address`, or `phone_number`) are included in the response.
	Customers []Customer `json:"customers,omitempty"`
	// A pagination cursor that can be used during subsequent calls to `SearchCustomers` to retrieve the next set of results associated with the original query. Pagination cursors are only present when a request succeeds and additional results are available.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
	// The total count of customers associated with the Square account that match the search query. Only customer profiles with public information (`given_name`, `family_name`, `company_name`, `email_address`, or `phone_number`) are counted. This field is present only if `count` is set to `true` in the request.
	Count int64 `json:"count,omitempty"`
}
