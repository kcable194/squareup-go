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

// Represents a discount that applies to one or more line items in an order.  Fixed-amount, order-scoped discounts are distributed across all non-zero line item totals. The amount distributed to each line item is relative to the amount contributed by the item to the order subtotal.
type OrderLineItemDiscount struct {
	// A unique ID that identifies the discount only within this order.
	Uid string `json:"uid,omitempty"`
	// The catalog object ID referencing [CatalogDiscount](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogDiscount).
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
	// The version of the catalog object that this discount references.
	CatalogVersion int64 `json:"catalog_version,omitempty"`
	// The discount's name.
	Name string `json:"name,omitempty"`
	// The type of the discount.  Discounts that do not reference a catalog object ID must have a type of `FIXED_PERCENTAGE` or `FIXED_AMOUNT`.
	Type_ string `json:"type,omitempty"`
	// The percentage of the discount, as a string representation of a decimal number. A value of `7.25` corresponds to a percentage of 7.25%.  `percentage` is not set for amount-based discounts.
	Percentage   string `json:"percentage,omitempty"`
	AmountMoney  *Money `json:"amount_money,omitempty"`
	AppliedMoney *Money `json:"applied_money,omitempty"`
	// Application-defined data attached to this discount. Metadata fields are intended to store descriptive references or associations with an entity in another system or store brief information about the object. Square does not process this field; it only stores and returns it in relevant API calls. Do not use metadata to store any sensitive information (such as personally identifiable information or card details).  Keys written by applications must be 60 characters or less and must be in the character set `[a-zA-Z0-9_-]`. Entries can also include metadata generated by Square. These keys are prefixed with a namespace, separated from the key with a ':' character.  Values have a maximum length of 255 characters.  An application can have up to 10 entries per metadata field.  Entries written by applications are private and can only be read or modified by the same application.  For more information, see [Metadata](https://developer.squareup.com/docs/build-basics/metadata).
	Metadata map[string]string `json:"metadata,omitempty"`
	// Indicates the level at which the discount applies. For `ORDER` scoped discounts, Square generates references in `applied_discounts` on all order line items that do not have them. For `LINE_ITEM` scoped discounts, the discount only applies to line items with a discount reference in their `applied_discounts` field.  This field is immutable. To change the scope of a discount, you must delete the discount and re-add it as a new discount.
	Scope string `json:"scope,omitempty"`
	// The reward IDs corresponding to this discount. The application and specification of discounts that have `reward_ids` are completely controlled by the backing criteria corresponding to the reward tiers of the rewards that are added to the order through the Loyalty API. To manually unapply discounts that are the result of added rewards, the rewards must be removed from the order through the Loyalty API.
	RewardIds []string `json:"reward_ids,omitempty"`
	// The object ID of a [pricing rule](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogPricingRule) to be applied automatically to this discount. The specification and application of the discounts, to which a `pricing_rule_id` is assigned, are completely controlled by the corresponding pricing rule.
	PricingRuleId string `json:"pricing_rule_id,omitempty"`
}
