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

type CatalogInfoResponse struct {
	// Any errors that occurred during the request.
	Errors                       []ModelError                  `json:"errors,omitempty"`
	Limits                       *CatalogInfoResponseLimits    `json:"limits,omitempty"`
	StandardUnitDescriptionGroup *StandardUnitDescriptionGroup `json:"standard_unit_description_group,omitempty"`
}
