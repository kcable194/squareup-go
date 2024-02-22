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

// Represents a payment refund processed by the Square Terminal. Only supports Interac (Canadian debit network) payment refunds.
type TerminalRefund struct {
	// A unique ID for this `TerminalRefund`.
	Id string `json:"id,omitempty"`
	// The reference to the payment refund created by completing this `TerminalRefund`.
	RefundId string `json:"refund_id,omitempty"`
	// The unique ID of the payment being refunded.
	PaymentId string `json:"payment_id"`
	// The reference to the Square order ID for the payment identified by the `payment_id`.
	OrderId     string `json:"order_id,omitempty"`
	AmountMoney *Money `json:"amount_money"`
	// A description of the reason for the refund.
	Reason string `json:"reason"`
	// The unique ID of the device intended for this `TerminalRefund`. The Id can be retrieved from /v2/devices api.
	DeviceId string `json:"device_id"`
	// The RFC 3339 duration, after which the refund is automatically canceled. A `TerminalRefund` that is `PENDING` is automatically `CANCELED` and has a cancellation reason of `TIMED_OUT`.  Default: 5 minutes from creation.  Maximum: 5 minutes
	DeadlineDuration string `json:"deadline_duration,omitempty"`
	// The status of the `TerminalRefund`. Options: `PENDING`, `IN_PROGRESS`, `CANCEL_REQUESTED`, `CANCELED`, or `COMPLETED`.
	Status string `json:"status,omitempty"`
	// Present if the status is `CANCELED`.
	CancelReason string `json:"cancel_reason,omitempty"`
	// The time when the `TerminalRefund` was created, as an RFC 3339 timestamp.
	CreatedAt string `json:"created_at,omitempty"`
	// The time when the `TerminalRefund` was last updated, as an RFC 3339 timestamp.
	UpdatedAt string `json:"updated_at,omitempty"`
	// The ID of the application that created the refund.
	AppId string `json:"app_id,omitempty"`
	// The location of the device where the `TerminalRefund` was directed.
	LocationId string `json:"location_id,omitempty"`
}
