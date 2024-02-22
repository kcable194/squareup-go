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

// Supported custom attribute query expressions for calling the [SearchCatalogItems](https://developer.squareup.com/reference/square_2024-01-18/catalog-api/search-catalog-items) endpoint to search for items or item variations.
type CustomAttributeFilter struct {
	// A query expression to filter items or item variations by matching their custom attributes' `custom_attribute_definition_id` property value against the the specified id. Exactly one of `custom_attribute_definition_id` or `key` must be specified.
	CustomAttributeDefinitionId string `json:"custom_attribute_definition_id,omitempty"`
	// A query expression to filter items or item variations by matching their custom attributes' `key` property value against the specified key. Exactly one of `custom_attribute_definition_id` or `key` must be specified.
	Key string `json:"key,omitempty"`
	// A query expression to filter items or item variations by matching their custom attributes' `string_value`  property value against the specified text. Exactly one of `string_filter`, `number_filter`, `selection_uids_filter`, or `bool_filter` must be specified.
	StringFilter string `json:"string_filter,omitempty"`
	NumberFilter *ModelRange `json:"number_filter,omitempty"`
	// A query expression to filter items or item variations by matching  their custom attributes' `selection_uid_values` values against the specified selection uids. Exactly one of `string_filter`, `number_filter`, `selection_uids_filter`, or `bool_filter` must be specified.
	SelectionUidsFilter []string `json:"selection_uids_filter,omitempty"`
	// A query expression to filter items or item variations by matching their custom attributes' `boolean_value` property values against the specified Boolean expression. Exactly one of `string_filter`, `number_filter`, `selection_uids_filter`, or `bool_filter` must be specified.
	BoolFilter bool `json:"bool_filter,omitempty"`
}