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

// Represents a query, consisting of specified query expressions, used to search for subscriptions.
type SearchSubscriptionsQuery struct {
	Filter *SearchSubscriptionsFilter `json:"filter,omitempty"`
}
