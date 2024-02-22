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

// Represents a discount being returned that applies to one or more return line items in an order.  Fixed-amount, order-scoped discounts are distributed across all non-zero return line item totals. The amount distributed to each return line item is relative to that item’s contribution to the order subtotal.
type OrderReturnDiscount struct {
	// A unique ID that identifies the returned discount only within this order.
	Uid string `json:"uid,omitempty"`
	// The discount `uid` from the order that contains the original application of this discount.
	SourceDiscountUid string `json:"source_discount_uid,omitempty"`
	// The catalog object ID referencing [CatalogDiscount](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogDiscount).
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
	// The version of the catalog object that this discount references.
	CatalogVersion int64 `json:"catalog_version,omitempty"`
	// The discount's name.
	Name string `json:"name,omitempty"`
	// The type of the discount. If it is created by the API, it is `FIXED_PERCENTAGE` or `FIXED_AMOUNT`.  Discounts that do not reference a catalog object ID must have a type of `FIXED_PERCENTAGE` or `FIXED_AMOUNT`.
	Type_ string `json:"type,omitempty"`
	// The percentage of the tax, as a string representation of a decimal number. A value of `\"7.25\"` corresponds to a percentage of 7.25%.  `percentage` is not set for amount-based discounts.
	Percentage   string `json:"percentage,omitempty"`
	AmountMoney  *Money `json:"amount_money,omitempty"`
	AppliedMoney *Money `json:"applied_money,omitempty"`
	// Indicates the level at which the `OrderReturnDiscount` applies. For `ORDER` scoped discounts, the server generates references in `applied_discounts` on all `OrderReturnLineItem`s. For `LINE_ITEM` scoped discounts, the discount is only applied to `OrderReturnLineItem`s with references in their `applied_discounts` field.
	Scope string `json:"scope,omitempty"`
}
