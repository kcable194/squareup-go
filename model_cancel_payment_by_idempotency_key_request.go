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

// Describes a request to cancel a payment using  [CancelPaymentByIdempotencyKey](https://developer.squareup.com/reference/square_2024-01-18/payments-api/cancel-payment-by-idempotency-key).
type CancelPaymentByIdempotencyKeyRequest struct {
	// The `idempotency_key` identifying the payment to be canceled.
	IdempotencyKey string `json:"idempotency_key"`
}
