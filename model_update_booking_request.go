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

type UpdateBookingRequest struct {
	// A unique key to make this request an idempotent operation.
	IdempotencyKey string `json:"idempotency_key,omitempty"`
	Booking *Booking `json:"booking"`
}
