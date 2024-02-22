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

type DataCollectionOptions struct {
	// The title text to display in the data collection flow on the Terminal.
	Title string `json:"title"`
	// The body text to display under the title in the data collection screen flow on the Terminal.
	Body string `json:"body"`
	// Represents the type of the input text.
	InputType string `json:"input_type"`
	CollectedData *CollectedData `json:"collected_data,omitempty"`
}