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

// The response to a request for a set of `EmployeeWage` objects. The response contains a set of `EmployeeWage` objects.
type ListEmployeeWagesResponse struct {
	// A page of `EmployeeWage` results.
	EmployeeWages []EmployeeWage `json:"employee_wages,omitempty"`
	// The value supplied in the subsequent request to fetch the next page of `EmployeeWage` results.
	Cursor string `json:"cursor,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
