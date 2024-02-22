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

// V1OrderHistoryEntry
type V1OrderHistoryEntry struct {
	// The type of action performed on the order.
	Action string `json:"action,omitempty"`
	// The time when the action was performed, in ISO 8601 format.
	CreatedAt string `json:"created_at,omitempty"`
}
