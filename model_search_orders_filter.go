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

// Filtering criteria to use for a `SearchOrders` request. Multiple filters are ANDed together.
type SearchOrdersFilter struct {
	StateFilter *SearchOrdersStateFilter `json:"state_filter,omitempty"`
	DateTimeFilter *SearchOrdersDateTimeFilter `json:"date_time_filter,omitempty"`
	FulfillmentFilter *SearchOrdersFulfillmentFilter `json:"fulfillment_filter,omitempty"`
	SourceFilter *SearchOrdersSourceFilter `json:"source_filter,omitempty"`
	CustomerFilter *SearchOrdersCustomerFilter `json:"customer_filter,omitempty"`
}
