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

// The response object returned by the [RetrieveMerchant](https://developer.squareup.com/reference/square_2024-01-18/merchants-api/retrieve-merchant) endpoint.
type RetrieveMerchantResponse struct {
	// Information on errors encountered during the request.
	Errors   []ModelError `json:"errors,omitempty"`
	Merchant *Merchant    `json:"merchant,omitempty"`
}
