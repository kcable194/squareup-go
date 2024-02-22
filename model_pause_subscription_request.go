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

// Defines input parameters in a request to the [PauseSubscription](https://developer.squareup.com/reference/square_2024-01-18/subscriptions-api/pause-subscription) endpoint.
type PauseSubscriptionRequest struct {
	// The `YYYY-MM-DD`-formatted date when the scheduled `PAUSE` action takes place on the subscription.  When this date is unspecified or falls within the current billing cycle, the subscription is paused on the starting date of the next billing cycle.
	PauseEffectiveDate string `json:"pause_effective_date,omitempty"`
	// The number of billing cycles the subscription will be paused before it is reactivated.   When this is set, a `RESUME` action is also scheduled to take place on the subscription at  the end of the specified pause cycle duration. In this case, neither `resume_effective_date`  nor `resume_change_timing` may be specified.
	PauseCycleDuration int64 `json:"pause_cycle_duration,omitempty"`
	// The date when the subscription is reactivated by a scheduled `RESUME` action.  This date must be at least one billing cycle ahead of `pause_effective_date`.
	ResumeEffectiveDate string `json:"resume_effective_date,omitempty"`
	// The timing whether the subscription is reactivated immediately or at the end of the billing cycle, relative to  `resume_effective_date`.
	ResumeChangeTiming string `json:"resume_change_timing,omitempty"`
	// The user-provided reason to pause the subscription.
	PauseReason string `json:"pause_reason,omitempty"`
}