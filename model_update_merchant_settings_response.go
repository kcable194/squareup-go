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

type UpdateMerchantSettingsResponse struct {
	// Any errors that occurred when updating the merchant settings.
	Errors []ModelError `json:"errors,omitempty"`
	MerchantSettings *CheckoutMerchantSettings `json:"merchant_settings,omitempty"`
}