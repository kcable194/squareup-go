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

// Configuration associated with Custom Attribute Definitions of type `STRING`.
type CatalogCustomAttributeDefinitionStringConfig struct {
	// If true, each Custom Attribute instance associated with this Custom Attribute Definition must have a unique value within the seller's catalog. For example, this may be used for a value like a SKU that should not be duplicated within a seller's catalog. May not be modified after the definition has been created.
	EnforceUniqueness bool `json:"enforce_uniqueness,omitempty"`
}
