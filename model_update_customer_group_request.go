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

// Defines the body parameters that can be included in a request to the [UpdateCustomerGroup](https://developer.squareup.com/reference/square_2024-01-18/customer-groups-api/update-customer-group) endpoint.
type UpdateCustomerGroupRequest struct {
	Group *CustomerGroup `json:"group"`
}
