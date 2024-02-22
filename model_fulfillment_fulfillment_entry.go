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

// Links an order line item to a fulfillment. Each entry must reference a valid `uid` for an order line item in the `line_item_uid` field, as well as a `quantity` to fulfill.
type FulfillmentFulfillmentEntry struct {
	// A unique ID that identifies the fulfillment entry only within this order.
	Uid string `json:"uid,omitempty"`
	// The `uid` from the order line item.
	LineItemUid string `json:"line_item_uid"`
	// The quantity of the line item being fulfilled, formatted as a decimal number. For example, `\"3\"`.  Fulfillments for line items with a `quantity_unit` can have non-integer quantities. For example, `\"1.70000\"`.
	Quantity string `json:"quantity"`
	// Application-defined data attached to this fulfillment entry. Metadata fields are intended to store descriptive references or associations with an entity in another system or store brief information about the object. Square does not process this field; it only stores and returns it in relevant API calls. Do not use metadata to store any sensitive information (such as personally identifiable information or card details).  Keys written by applications must be 60 characters or less and must be in the character set `[a-zA-Z0-9_-]`. Entries can also include metadata generated by Square. These keys are prefixed with a namespace, separated from the key with a ':' character.  Values have a maximum length of 255 characters.  An application can have up to 10 entries per metadata field.  Entries written by applications are private and can only be read or modified by the same application.  For more information, see [Metadata](https://developer.squareup.com/docs/build-basics/metadata).
	Metadata map[string]string `json:"metadata,omitempty"`
}
