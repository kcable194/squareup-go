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

// Represents a checkout processed by the Square Terminal.
type TerminalCheckout struct {
	// A unique ID for this `TerminalCheckout`.
	Id string `json:"id,omitempty"`
	AmountMoney *Money `json:"amount_money"`
	// An optional user-defined reference ID that can be used to associate this `TerminalCheckout` to another entity in an external system. For example, an order ID generated by a third-party shopping cart. The ID is also associated with any payments used to complete the checkout.
	ReferenceId string `json:"reference_id,omitempty"`
	// An optional note to associate with the checkout, as well as with any payments used to complete the checkout. Note: maximum 500 characters
	Note string `json:"note,omitempty"`
	// The reference to the Square order ID for the checkout request. Supported only in the US.
	OrderId string `json:"order_id,omitempty"`
	PaymentOptions *PaymentOptions `json:"payment_options,omitempty"`
	DeviceOptions *DeviceCheckoutOptions `json:"device_options"`
	// An RFC 3339 duration, after which the checkout is automatically canceled. A `TerminalCheckout` that is `PENDING` is automatically `CANCELED` and has a cancellation reason of `TIMED_OUT`.  Default: 5 minutes from creation  Maximum: 5 minutes
	DeadlineDuration string `json:"deadline_duration,omitempty"`
	// The status of the `TerminalCheckout`. Options: `PENDING`, `IN_PROGRESS`, `CANCEL_REQUESTED`, `CANCELED`, `COMPLETED`
	Status string `json:"status,omitempty"`
	// The reason why `TerminalCheckout` is canceled. Present if the status is `CANCELED`.
	CancelReason string `json:"cancel_reason,omitempty"`
	// A list of IDs for payments created by this `TerminalCheckout`.
	PaymentIds []string `json:"payment_ids,omitempty"`
	// The time when the `TerminalCheckout` was created, as an RFC 3339 timestamp.
	CreatedAt string `json:"created_at,omitempty"`
	// The time when the `TerminalCheckout` was last updated, as an RFC 3339 timestamp.
	UpdatedAt string `json:"updated_at,omitempty"`
	// The ID of the application that created the checkout.
	AppId string `json:"app_id,omitempty"`
	// The location of the device where the `TerminalCheckout` was directed.
	LocationId string `json:"location_id,omitempty"`
	// The type of payment the terminal should attempt to capture from. Defaults to `CARD_PRESENT`.
	PaymentType string `json:"payment_type,omitempty"`
	// An optional ID of the team member associated with creating the checkout.
	TeamMemberId string `json:"team_member_id,omitempty"`
	// An optional ID of the customer associated with the checkout.
	CustomerId string `json:"customer_id,omitempty"`
	AppFeeMoney *Money `json:"app_fee_money,omitempty"`
	// Optional additional payment information to include on the customer's card statement as part of the statement description. This can be, for example, an invoice number, ticket number, or short description that uniquely identifies the purchase. Supported only in the US.
	StatementDescriptionIdentifier string `json:"statement_description_identifier,omitempty"`
	TipMoney *Money `json:"tip_money,omitempty"`
}