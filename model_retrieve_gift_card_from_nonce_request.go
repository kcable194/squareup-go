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

// A request to retrieve a gift card by using a payment token.
type RetrieveGiftCardFromNonceRequest struct {
	// The payment token of the gift card to retrieve. Payment tokens are generated by the  Web Payments SDK or In-App Payments SDK.
	Nonce string `json:"nonce"`
}
