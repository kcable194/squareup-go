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

// Pricing options for an order. The options affect how the order's price is calculated. They can be used, for example, to apply automatic price adjustments that are based on preconfigured [pricing rules](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogPricingRule).
type OrderPricingOptions struct {
	// The option to determine whether pricing rule-based discounts are automatically applied to an order.
	AutoApplyDiscounts bool `json:"auto_apply_discounts,omitempty"`
	// The option to determine whether rule-based taxes are automatically applied to an order when the criteria of the corresponding rules are met.
	AutoApplyTaxes bool `json:"auto_apply_taxes,omitempty"`
}
