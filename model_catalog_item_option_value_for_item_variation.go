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

// A `CatalogItemOptionValue` links an item variation to an item option as an item option value. For example, a t-shirt item may offer a color option and a size option. An item option value would represent each variation of t-shirt: For example, \"Color:Red, Size:Small\" or \"Color:Blue, Size:Medium\".
type CatalogItemOptionValueForItemVariation struct {
	// The unique id of an item option.
	ItemOptionId string `json:"item_option_id,omitempty"`
	// The unique id of the selected value for the item option.
	ItemOptionValueId string `json:"item_option_value_id,omitempty"`
}
