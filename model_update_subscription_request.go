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

// Defines input parameters in a request to the  [UpdateSubscription](https://developer.squareup.com/reference/square_2024-01-18/subscriptions-api/update-subscription) endpoint.
type UpdateSubscriptionRequest struct {
	Subscription *Subscription `json:"subscription,omitempty"`
}