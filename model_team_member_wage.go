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

// The hourly wage rate that a team member earns on a `Shift` for doing the job specified by the `title` property of this object.
type TeamMemberWage struct {
	// The UUID for this object.
	Id string `json:"id,omitempty"`
	// The `TeamMember` that this wage is assigned to.
	TeamMemberId string `json:"team_member_id,omitempty"`
	// The job title that this wage relates to.
	Title      string `json:"title,omitempty"`
	HourlyRate *Money `json:"hourly_rate,omitempty"`
	// An identifier for the job that this wage relates to. This cannot be used to retrieve the job.
	JobId string `json:"job_id,omitempty"`
	// Whether team members are eligible for tips when working this job.
	TipEligible bool `json:"tip_eligible,omitempty"`
}
