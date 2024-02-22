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

type UpsertCatalogObjectResponse struct {
	// Any errors that occurred during the request.
	Errors        []ModelError   `json:"errors,omitempty"`
	CatalogObject *CatalogObject `json:"catalog_object,omitempty"`
	// The mapping between client and server IDs for this upsert.
	IdMappings []CatalogIdMapping `json:"id_mappings,omitempty"`
}
