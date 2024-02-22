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

type SelectOption struct {
	// The reference id for the option.
	ReferenceId string `json:"reference_id"`
	// The title text that displays in the select option button.
	Title string `json:"title"`
}