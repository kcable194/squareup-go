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

// Represents a response from an update request containing the updated `WageSetting` object or error messages.
type UpdateWageSettingResponse struct {
	WageSetting *WageSetting `json:"wage_setting,omitempty"`
	// The errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
