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

type CashDrawerShiftEvent struct {
	// The unique ID of the event.
	Id string `json:"id,omitempty"`
	// The type of cash drawer shift event.
	EventType string `json:"event_type,omitempty"`
	EventMoney *Money `json:"event_money,omitempty"`
	// The event time in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// An optional description of the event, entered by the employee that created the event.
	Description string `json:"description,omitempty"`
	// The ID of the team member that created the event.
	TeamMemberId string `json:"team_member_id,omitempty"`
}