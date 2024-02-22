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

// The query filter to return the items containing the specified item option IDs.
type CatalogQueryItemsForItemOptions struct {
	// A set of `CatalogItemOption` IDs to be used to find associated `CatalogItem`s. All Items that contain all of the given Item Options (in any order) will be returned.
	ItemOptionIds []string `json:"item_option_ids,omitempty"`
}
