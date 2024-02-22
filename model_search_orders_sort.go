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

// Sorting criteria for a `SearchOrders` request. Results can only be sorted by a timestamp field.
type SearchOrdersSort struct {
	// The field to sort by.  __Important:__ When using a [DateTimeFilter](https://developer.squareup.com/reference/square_2024-01-18/objects/SearchOrdersFilter), `sort_field` must match the timestamp field that the `DateTimeFilter` uses to filter. For example, if you set your `sort_field` to `CLOSED_AT` and you use a `DateTimeFilter`, your `DateTimeFilter` must filter for orders by their `CLOSED_AT` date. If this field does not match the timestamp field in `DateTimeFilter`, `SearchOrders` returns an error.  Default: `CREATED_AT`.
	SortField string `json:"sort_field"`
	// The chronological order in which results are returned. Defaults to `DESC`.
	SortOrder string `json:"sort_order,omitempty"`
}
