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

// Request payload for the [BulkRetrieveTeamMemberBookingProfiles](https://developer.squareup.com/reference/square_2024-01-18/bookings-api/bulk-retrieve-team-member-booking-profiles) endpoint.
type BulkRetrieveTeamMemberBookingProfilesRequest struct {
	// A non-empty list of IDs of team members whose booking profiles you want to retrieve.
	TeamMemberIds []string `json:"team_member_ids"`
}
