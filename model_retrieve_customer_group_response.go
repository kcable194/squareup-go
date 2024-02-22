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

// Defines the fields that are included in the response body of a request to the [RetrieveCustomerGroup](https://developer.squareup.com/reference/square_2024-01-18/customer-groups-api/retrieve-customer-group) endpoint.  Either `errors` or `group` is present in a given response (never both).
type RetrieveCustomerGroupResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError   `json:"errors,omitempty"`
	Group  *CustomerGroup `json:"group,omitempty"`
}
