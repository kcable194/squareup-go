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

// Creates a [Subscription](https://developer.squareup.com/reference/square_2024-01-18/objects/WebhookSubscription).
type CreateWebhookSubscriptionRequest struct {
	// A unique string that identifies the [CreateWebhookSubscription](https://developer.squareup.com/reference/square_2024-01-18/webhook-subscriptions-api/create-webhook-subscription) request.
	IdempotencyKey string `json:"idempotency_key,omitempty"`
	Subscription *WebhookSubscription `json:"subscription"`
}