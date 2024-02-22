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

type UpdateItemTaxesRequest struct {
	// IDs for the CatalogItems associated with the CatalogTax objects being updated. No more than 1,000 IDs may be provided.
	ItemIds []string `json:"item_ids"`
	// IDs of the CatalogTax objects to enable. At least one of `taxes_to_enable` or `taxes_to_disable` must be specified.
	TaxesToEnable []string `json:"taxes_to_enable,omitempty"`
	// IDs of the CatalogTax objects to disable. At least one of `taxes_to_enable` or `taxes_to_disable` must be specified.
	TaxesToDisable []string `json:"taxes_to_disable,omitempty"`
}