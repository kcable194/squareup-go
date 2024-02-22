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

// The query filter to return the search result(s) by exact match of the specified `attribute_name` and any of the `attribute_values`.
type CatalogQuerySet struct {
	// The name of the attribute to be searched. Matching of the attribute name is exact.
	AttributeName string `json:"attribute_name"`
	// The desired values of the search attribute. Matching of the attribute values is exact and case insensitive. A maximum of 250 values may be searched in a request.
	AttributeValues []string `json:"attribute_values"`
}
