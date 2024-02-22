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

// Contains the name and abbreviation for standard measurement unit.
type StandardUnitDescription struct {
	Unit *MeasurementUnit `json:"unit,omitempty"`
	// UI display name of the measurement unit. For example, 'Pound'.
	Name string `json:"name,omitempty"`
	// UI display abbreviation for the measurement unit. For example, 'lb'.
	Abbreviation string `json:"abbreviation,omitempty"`
}