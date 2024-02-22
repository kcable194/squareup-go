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

// A request for a set of `TeamMemberWage` objects.
type ListTeamMemberWagesRequest struct {
	// Filter the returned wages to only those that are associated with the specified team member.
	TeamMemberId string `json:"team_member_id,omitempty"`
	// The maximum number of `TeamMemberWage` results to return per page. The number can range between 1 and 200. The default is 200.
	Limit int32 `json:"limit,omitempty"`
	// A pointer to the next page of `EmployeeWage` results to fetch.
	Cursor string `json:"cursor,omitempty"`
}
