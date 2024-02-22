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

// Defines the fields that are included in the response body of a request to the [UpdateWebhookSubscription](https://developer.squareup.com/reference/square_2024-01-18/webhook-subscriptions-api/update-webhook-subscription) endpoint.  Note: If there are errors processing the request, the [Subscription](https://developer.squareup.com/reference/square_2024-01-18/objects/WebhookSubscription) is not present.
type UpdateWebhookSubscriptionResponse struct {
	// Information on errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	Subscription *WebhookSubscription `json:"subscription,omitempty"`
}
