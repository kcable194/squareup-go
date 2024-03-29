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

// A batch of catalog objects.
type CatalogObjectBatch struct {
	// A list of CatalogObjects belonging to this batch.
	Objects []CatalogObject `json:"objects"`
}
