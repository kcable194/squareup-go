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

// Options to control how to override the default behavior of the specified modifier.
type CatalogModifierOverride struct {
	// The ID of the `CatalogModifier` whose default behavior is being overridden.
	ModifierId string `json:"modifier_id"`
	// If `true`, this `CatalogModifier` should be selected by default for this `CatalogItem`.
	OnByDefault bool `json:"on_by_default,omitempty"`
}
