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
// CatalogPricingType : Indicates whether the price of a CatalogItemVariation should be entered manually at the time of sale.
type CatalogPricingType string

// List of CatalogPricingType
const (
	FIXED_PRICING_CatalogPricingType CatalogPricingType = "FIXED_PRICING"
	VARIABLE_PRICING_CatalogPricingType CatalogPricingType = "VARIABLE_PRICING"
)