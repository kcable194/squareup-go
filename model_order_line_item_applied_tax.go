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

// Represents an applied portion of a tax to a line item in an order.  Order-scoped taxes automatically include the applied taxes in each line item. Line item taxes must be referenced from any applicable line items. The corresponding applied money is automatically computed, based on the set of participating line items.
type OrderLineItemAppliedTax struct {
	// A unique ID that identifies the applied tax only within this order.
	Uid string `json:"uid,omitempty"`
	// The `uid` of the tax for which this applied tax represents. It must reference a tax present in the `order.taxes` field.  This field is immutable. To change which taxes apply to a line item, delete and add a new `OrderLineItemAppliedTax`.
	TaxUid       string `json:"tax_uid"`
	AppliedMoney *Money `json:"applied_money,omitempty"`
}
