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

// Defines the fields that are included in the response body of a request to the `CreateCustomerCard` endpoint.  Either `errors` or `card` is present in a given response (never both).
type CreateCustomerCardResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	Card   *Card        `json:"card,omitempty"`
}
