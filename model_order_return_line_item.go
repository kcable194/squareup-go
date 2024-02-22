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

// The line item being returned in an order.
type OrderReturnLineItem struct {
	// A unique ID for this return line-item entry.
	Uid string `json:"uid,omitempty"`
	// The `uid` of the line item in the original sale order.
	SourceLineItemUid string `json:"source_line_item_uid,omitempty"`
	// The name of the line item.
	Name string `json:"name,omitempty"`
	// The quantity returned, formatted as a decimal number. For example, `\"3\"`.  Line items with a `quantity_unit` can have non-integer quantities. For example, `\"1.70000\"`.
	Quantity string `json:"quantity"`
	QuantityUnit *OrderQuantityUnit `json:"quantity_unit,omitempty"`
	// The note of the return line item.
	Note string `json:"note,omitempty"`
	// The [CatalogItemVariation](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogItemVariation) ID applied to this return line item.
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
	// The version of the catalog object that this line item references.
	CatalogVersion int64 `json:"catalog_version,omitempty"`
	// The name of the variation applied to this return line item.
	VariationName string `json:"variation_name,omitempty"`
	// The type of line item: an itemized return, a non-itemized return (custom amount), or the return of an unactivated gift card sale.
	ItemType string `json:"item_type,omitempty"`
	// The [CatalogModifier](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogModifier)s applied to this line item.
	ReturnModifiers []OrderReturnLineItemModifier `json:"return_modifiers,omitempty"`
	// The list of references to `OrderReturnTax` entities applied to the return line item. Each `OrderLineItemAppliedTax` has a `tax_uid` that references the `uid` of a top-level `OrderReturnTax` applied to the return line item. On reads, the applied amount is populated.
	AppliedTaxes []OrderLineItemAppliedTax `json:"applied_taxes,omitempty"`
	// The list of references to `OrderReturnDiscount` entities applied to the return line item. Each `OrderLineItemAppliedDiscount` has a `discount_uid` that references the `uid` of a top-level `OrderReturnDiscount` applied to the return line item. On reads, the applied amount is populated.
	AppliedDiscounts []OrderLineItemAppliedDiscount `json:"applied_discounts,omitempty"`
	BasePriceMoney *Money `json:"base_price_money,omitempty"`
	VariationTotalPriceMoney *Money `json:"variation_total_price_money,omitempty"`
	GrossReturnMoney *Money `json:"gross_return_money,omitempty"`
	TotalTaxMoney *Money `json:"total_tax_money,omitempty"`
	TotalDiscountMoney *Money `json:"total_discount_money,omitempty"`
	TotalMoney *Money `json:"total_money,omitempty"`
	// The list of references to `OrderReturnServiceCharge` entities applied to the return line item. Each `OrderLineItemAppliedServiceCharge` has a `service_charge_uid` that references the `uid` of a top-level `OrderReturnServiceCharge` applied to the return line item. On reads, the applied amount is populated.
	AppliedServiceCharges []OrderLineItemAppliedServiceCharge `json:"applied_service_charges,omitempty"`
	TotalServiceChargeMoney *Money `json:"total_service_charge_money,omitempty"`
}
