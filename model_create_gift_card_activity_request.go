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

// A request to create a gift card activity.
type CreateGiftCardActivityRequest struct {
	// A unique string that identifies the `CreateGiftCardActivity` request.
	IdempotencyKey   string            `json:"idempotency_key"`
	GiftCardActivity *GiftCardActivity `json:"gift_card_activity"`
}
