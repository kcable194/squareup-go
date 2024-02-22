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

// Defines output parameters in a response from the [ListSubscriptionEvents](https://developer.squareup.com/reference/square_2024-01-18/subscriptions-api/list-subscription-events).
type ListSubscriptionEventsResponse struct {
	// Errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The retrieved subscription events.
	SubscriptionEvents []SubscriptionEvent `json:"subscription_events,omitempty"`
	// When the total number of resulting subscription events exceeds the limit of a paged response,  the response includes a cursor for you to use in a subsequent request to fetch the next set of events. If the cursor is unset, the response contains the last page of the results.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
}
