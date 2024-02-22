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

// Represents an error encountered during a request to the Connect API.  See [Handling errors](https://developer.squareup.com/docs/build-basics/handling-errors) for more information.
type ModelError struct {
	// The high-level category for the error.
	Category string `json:"category"`
	// The specific code of the error.
	Code string `json:"code"`
	// A human-readable description of the error for debugging purposes.
	Detail string `json:"detail,omitempty"`
	// The name of the field provided in the original request (if any) that the error pertains to.
	Field string `json:"field,omitempty"`
}
