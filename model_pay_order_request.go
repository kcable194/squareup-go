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

// Defines the fields that are included in requests to the [PayOrder](https://developer.squareup.com/reference/square_2024-01-18/orders-api/pay-order) endpoint.
type PayOrderRequest struct {
	// A value you specify that uniquely identifies this request among requests you have sent. If you are unsure whether a particular payment request was completed successfully, you can reattempt it with the same idempotency key without worrying about duplicate payments.  For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
	IdempotencyKey string `json:"idempotency_key"`
	// The version of the order being paid. If not supplied, the latest version will be paid.
	OrderVersion int32 `json:"order_version,omitempty"`
	// The IDs of the [payments](https://developer.squareup.com/reference/square_2024-01-18/objects/Payment) to collect. The payment total must match the order total.
	PaymentIds []string `json:"payment_ids,omitempty"`
}