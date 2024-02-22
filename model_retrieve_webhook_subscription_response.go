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

// Defines the fields that are included in the response body of a request to the [RetrieveWebhookSubscription](https://developer.squareup.com/reference/square_2024-01-18/webhook-subscriptions-api/retrieve-webhook-subscription) endpoint.  Note: if there are errors processing the request, the [Subscription](https://developer.squareup.com/reference/square_2024-01-18/objects/WebhookSubscription) will not be present.
type RetrieveWebhookSubscriptionResponse struct {
	// Information on errors encountered during the request.
	Errors       []ModelError         `json:"errors,omitempty"`
	Subscription *WebhookSubscription `json:"subscription,omitempty"`
}
