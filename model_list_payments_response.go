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

// Defines the response returned by [ListPayments](https://developer.squareup.com/reference/square_2024-01-18/payments-api/list-payments).
type ListPaymentsResponse struct {
	// Information about errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The requested list of payments.
	Payments []Payment `json:"payments,omitempty"`
	// The pagination cursor to be used in a subsequent request. If empty, this is the final response.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
}
