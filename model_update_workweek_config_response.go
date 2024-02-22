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

// The response to a request to update a `WorkweekConfig` object. The response contains the updated `WorkweekConfig` object and might contain a set of `Error` objects if the request resulted in errors.
type UpdateWorkweekConfigResponse struct {
	WorkweekConfig *WorkweekConfig `json:"workweek_config,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
