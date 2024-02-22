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

// The response to the request to create a `BreakType`. The response contains the created `BreakType` object and might contain a set of `Error` objects if the request resulted in errors.
type CreateBreakTypeResponse struct {
	BreakType *BreakType `json:"break_type,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
