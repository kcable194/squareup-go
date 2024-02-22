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

type SelectOptions struct {
	// The title text to display in the select flow on the Terminal.
	Title string `json:"title"`
	// The body text to display in the select flow on the Terminal.
	Body string `json:"body"`
	// Represents the buttons/options that should be displayed in the select flow on the Terminal.
	Options []SelectOption `json:"options"`
	SelectedOption *SelectOption `json:"selected_option,omitempty"`
}
