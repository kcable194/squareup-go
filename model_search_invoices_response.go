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

// Describes a `SearchInvoices` response.
type SearchInvoicesResponse struct {
	// The list of invoices returned by the search.
	Invoices []Invoice `json:"invoices,omitempty"`
	// When a response is truncated, it includes a cursor that you can use in a  subsequent request to fetch the next set of invoices. If empty, this is the final  response.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
	// Information about errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
