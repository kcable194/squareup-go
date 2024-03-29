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

// Represents additional data for rules with the `ITEM_VARIATION` accrual type.
type LoyaltyProgramAccrualRuleItemVariationData struct {
	// The ID of the `ITEM_VARIATION` [catalog object](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogObject) that buyers can purchase to earn points.
	ItemVariationId string `json:"item_variation_id"`
}
