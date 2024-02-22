# SearchCatalogItemsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TextFilter** | **string** | The text filter expression to return items or item variations containing specified text in the &#x60;name&#x60;, &#x60;description&#x60;, or &#x60;abbreviation&#x60; attribute value of an item, or in the &#x60;name&#x60;, &#x60;sku&#x60;, or &#x60;upc&#x60; attribute value of an item variation. | [optional] [default to null]
**CategoryIds** | **[]string** | The category id query expression to return items containing the specified category IDs. | [optional] [default to null]
**StockLevels** | **[]string** | The stock-level query expression to return item variations with the specified stock levels. | [optional] [default to null]
**EnabledLocationIds** | **[]string** | The enabled-location query expression to return items and item variations having specified enabled locations. | [optional] [default to null]
**Cursor** | **string** | The pagination token, returned in the previous response, used to fetch the next batch of pending results. | [optional] [default to null]
**Limit** | **int32** | The maximum number of results to return per page. The default value is 100. | [optional] [default to null]
**SortOrder** | **string** | The order to sort the results by item names. The default sort order is ascending (&#x60;ASC&#x60;). | [optional] [default to null]
**ProductTypes** | **[]string** | The product types query expression to return items or item variations having the specified product types. | [optional] [default to null]
**CustomAttributeFilters** | [**[]CustomAttributeFilter**](CustomAttributeFilter.md) | The customer-attribute filter to return items or item variations matching the specified custom attribute expressions. A maximum number of 10 custom attribute expressions are supported in a single call to the [SearchCatalogItems](https://developer.squareup.com/reference/square_2024-01-18/catalog-api/search-catalog-items) endpoint. | [optional] [default to null]
**ArchivedState** | **string** | The query filter to return not archived (&#x60;ARCHIVED_STATE_NOT_ARCHIVED&#x60;), archived (&#x60;ARCHIVED_STATE_ARCHIVED&#x60;), or either type (&#x60;ARCHIVED_STATE_ALL&#x60;) of items. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

