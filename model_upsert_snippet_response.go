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

// Represents an `UpsertSnippet` response. The response can include either `snippet` or `errors`.
type UpsertSnippetResponse struct {
	// Any errors that occurred during the request.
	Errors  []ModelError `json:"errors,omitempty"`
	Snippet *Snippet     `json:"snippet,omitempty"`
}
