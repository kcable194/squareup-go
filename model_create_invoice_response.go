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

// The response returned by the `CreateInvoice` request.
type CreateInvoiceResponse struct {
	Invoice *Invoice `json:"invoice,omitempty"`
	// Information about errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
