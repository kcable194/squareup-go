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

// Contains all information related to a single order to process with Square, including line items that specify the products to purchase. `Order` objects also include information about any associated tenders, refunds, and returns.  All Connect V2 Transactions have all been converted to Orders including all associated itemization data.
type Order struct {
	// The order's unique ID.
	Id string `json:"id,omitempty"`
	// The ID of the seller location that this order is associated with.
	LocationId string `json:"location_id"`
	// A client-specified ID to associate an entity in another system with this order.
	ReferenceId string `json:"reference_id,omitempty"`
	Source *OrderSource `json:"source,omitempty"`
	// The ID of the [customer](https://developer.squareup.com/reference/square_2024-01-18/objects/Customer) associated with the order.  You should specify a `customer_id` on the order (or the payment) to ensure that transactions are reliably linked to customers. Omitting this field might result in the creation of new [instant profiles](https://developer.squareup.com/docs/customers-api/what-it-does#instant-profiles).
	CustomerId string `json:"customer_id,omitempty"`
	// The line items included in the order.
	LineItems []OrderLineItem `json:"line_items,omitempty"`
	// The list of all taxes associated with the order.  Taxes can be scoped to either `ORDER` or `LINE_ITEM`. For taxes with `LINE_ITEM` scope, an `OrderLineItemAppliedTax` must be added to each line item that the tax applies to. For taxes with `ORDER` scope, the server generates an `OrderLineItemAppliedTax` for every line item.  On reads, each tax in the list includes the total amount of that tax applied to the order.  __IMPORTANT__: If `LINE_ITEM` scope is set on any taxes in this field, using the deprecated `line_items.taxes` field results in an error. Use `line_items.applied_taxes` instead.
	Taxes []OrderLineItemTax `json:"taxes,omitempty"`
	// The list of all discounts associated with the order.  Discounts can be scoped to either `ORDER` or `LINE_ITEM`. For discounts scoped to `LINE_ITEM`, an `OrderLineItemAppliedDiscount` must be added to each line item that the discount applies to. For discounts with `ORDER` scope, the server generates an `OrderLineItemAppliedDiscount` for every line item.  __IMPORTANT__: If `LINE_ITEM` scope is set on any discounts in this field, using the deprecated `line_items.discounts` field results in an error. Use `line_items.applied_discounts` instead.
	Discounts []OrderLineItemDiscount `json:"discounts,omitempty"`
	// A list of service charges applied to the order.
	ServiceCharges []OrderServiceCharge `json:"service_charges,omitempty"`
	// Details about order fulfillment.  Orders can only be created with at most one fulfillment. However, orders returned by the API might contain multiple fulfillments.
	Fulfillments []Fulfillment `json:"fulfillments,omitempty"`
	// A collection of items from sale orders being returned in this one. Normally part of an itemized return or exchange. There is exactly one `Return` object per sale `Order` being referenced.
	Returns []OrderReturn `json:"returns,omitempty"`
	ReturnAmounts *OrderMoneyAmounts `json:"return_amounts,omitempty"`
	NetAmounts *OrderMoneyAmounts `json:"net_amounts,omitempty"`
	RoundingAdjustment *OrderRoundingAdjustment `json:"rounding_adjustment,omitempty"`
	// The tenders that were used to pay for the order.
	Tenders []Tender `json:"tenders,omitempty"`
	// The refunds that are part of this order.
	Refunds []Refund `json:"refunds,omitempty"`
	// Application-defined data attached to this order. Metadata fields are intended to store descriptive references or associations with an entity in another system or store brief information about the object. Square does not process this field; it only stores and returns it in relevant API calls. Do not use metadata to store any sensitive information (such as personally identifiable information or card details).  Keys written by applications must be 60 characters or less and must be in the character set `[a-zA-Z0-9_-]`. Entries can also include metadata generated by Square. These keys are prefixed with a namespace, separated from the key with a ':' character.  Values have a maximum length of 255 characters.  An application can have up to 10 entries per metadata field.  Entries written by applications are private and can only be read or modified by the same application.  For more information, see  [Metadata](https://developer.squareup.com/docs/build-basics/metadata).
	Metadata map[string]string `json:"metadata,omitempty"`
	// The timestamp for when the order was created, in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp for when the order was last updated, in RFC 3339 format (for example, \"2016-09-04T23:59:33.123Z\").
	UpdatedAt string `json:"updated_at,omitempty"`
	// The timestamp for when the order reached a terminal [state](https://developer.squareup.com/reference/square_2024-01-18/enums/OrderState), in RFC 3339 format (for example \"2016-09-04T23:59:33.123Z\").
	ClosedAt string `json:"closed_at,omitempty"`
	// The current state of the order.
	State string `json:"state,omitempty"`
	// The version number, which is incremented each time an update is committed to the order. Orders not created through the API do not include a version number and therefore cannot be updated.  [Read more about working with versions](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders).
	Version int32 `json:"version,omitempty"`
	TotalMoney *Money `json:"total_money,omitempty"`
	TotalTaxMoney *Money `json:"total_tax_money,omitempty"`
	TotalDiscountMoney *Money `json:"total_discount_money,omitempty"`
	TotalTipMoney *Money `json:"total_tip_money,omitempty"`
	TotalServiceChargeMoney *Money `json:"total_service_charge_money,omitempty"`
	// A short-term identifier for the order (such as a customer first name, table number, or auto-generated order number that resets daily).
	TicketName string `json:"ticket_name,omitempty"`
	PricingOptions *OrderPricingOptions `json:"pricing_options,omitempty"`
	// A set-like list of Rewards that have been added to the Order.
	Rewards []OrderReward `json:"rewards,omitempty"`
	NetAmountDueMoney *Money `json:"net_amount_due_money,omitempty"`
}