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

// Defines input parameters in a call to the [BulkSwapPlan](https://developer.squareup.com/reference/square_2024-01-18/subscriptions-api/bulk-swap-plan) endpoint.
type BulkSwapPlanRequest struct {
	// The ID of the new subscription plan variation.  This field is required.
	NewPlanVariationId string `json:"new_plan_variation_id"`
	// The ID of the plan variation whose subscriptions should be swapped. Active subscriptions using this plan variation will be subscribed to the new plan variation on their next billing day.
	OldPlanVariationId string `json:"old_plan_variation_id"`
	// The ID of the location to associate with the swapped subscriptions.
	LocationId string `json:"location_id"`
}
