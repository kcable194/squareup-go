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

// Represents an update request for the `WageSetting` object describing a `TeamMember`.
type UpdateWageSettingRequest struct {
	WageSetting *WageSetting `json:"wage_setting"`
}