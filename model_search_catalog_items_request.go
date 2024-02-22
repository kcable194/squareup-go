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

// Defines the request body for the [SearchCatalogItems](https://developer.squareup.com/reference/square_2024-01-18/catalog-api/search-catalog-items) endpoint.
type SearchCatalogItemsRequest struct {
	// The text filter expression to return items or item variations containing specified text in the `name`, `description`, or `abbreviation` attribute value of an item, or in the `name`, `sku`, or `upc` attribute value of an item variation.
	TextFilter string `json:"text_filter,omitempty"`
	// The category id query expression to return items containing the specified category IDs.
	CategoryIds []string `json:"category_ids,omitempty"`
	// The stock-level query expression to return item variations with the specified stock levels.
	StockLevels []string `json:"stock_levels,omitempty"`
	// The enabled-location query expression to return items and item variations having specified enabled locations.
	EnabledLocationIds []string `json:"enabled_location_ids,omitempty"`
	// The pagination token, returned in the previous response, used to fetch the next batch of pending results.
	Cursor string `json:"cursor,omitempty"`
	// The maximum number of results to return per page. The default value is 100.
	Limit int32 `json:"limit,omitempty"`
	// The order to sort the results by item names. The default sort order is ascending (`ASC`).
	SortOrder string `json:"sort_order,omitempty"`
	// The product types query expression to return items or item variations having the specified product types.
	ProductTypes []string `json:"product_types,omitempty"`
	// The customer-attribute filter to return items or item variations matching the specified custom attribute expressions. A maximum number of 10 custom attribute expressions are supported in a single call to the [SearchCatalogItems](https://developer.squareup.com/reference/square_2024-01-18/catalog-api/search-catalog-items) endpoint.
	CustomAttributeFilters []CustomAttributeFilter `json:"custom_attribute_filters,omitempty"`
	// The query filter to return not archived (`ARCHIVED_STATE_NOT_ARCHIVED`), archived (`ARCHIVED_STATE_ARCHIVED`), or either type (`ARCHIVED_STATE_ALL`) of items.
	ArchivedState string `json:"archived_state,omitempty"`
}
