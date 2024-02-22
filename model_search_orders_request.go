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

type SearchOrdersRequest struct {
	// The location IDs for the orders to query. All locations must belong to the same merchant.  Max: 10 location IDs.
	LocationIds []string `json:"location_ids,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for your original query. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
	Query *SearchOrdersQuery `json:"query,omitempty"`
	// The maximum number of results to be returned in a single page.  Default: `500` Max: `1000`
	Limit int32 `json:"limit,omitempty"`
	// A Boolean that controls the format of the search results. If `true`, `SearchOrders` returns [OrderEntry](https://developer.squareup.com/reference/square_2024-01-18/objects/OrderEntry) objects. If `false`, `SearchOrders` returns complete order objects.  Default: `false`.
	ReturnEntries bool `json:"return_entries,omitempty"`
}
