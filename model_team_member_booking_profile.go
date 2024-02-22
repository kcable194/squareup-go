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

// The booking profile of a seller's team member, including the team member's ID, display name, description and whether the team member can be booked as a service provider.
type TeamMemberBookingProfile struct {
	// The ID of the [TeamMember](https://developer.squareup.com/reference/square_2024-01-18/objects/TeamMember) object for the team member associated with the booking profile.
	TeamMemberId string `json:"team_member_id,omitempty"`
	// The description of the team member.
	Description string `json:"description,omitempty"`
	// The display name of the team member.
	DisplayName string `json:"display_name,omitempty"`
	// Indicates whether the team member can be booked through the Bookings API or the seller's online booking channel or site (`true`) or not (`false`).
	IsBookable bool `json:"is_bookable,omitempty"`
	// The URL of the team member's image for the bookings profile.
	ProfileImageUrl string `json:"profile_image_url,omitempty"`
}
