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

// The hourly wage rate used to compensate an employee for this shift.
type ShiftWage struct {
	// The name of the job performed during this shift.
	Title string `json:"title,omitempty"`
	HourlyRate *Money `json:"hourly_rate,omitempty"`
	// The id of the job performed during this shift. Square labor-reporting UIs might group shifts together by id. This cannot be used to retrieve the job.
	JobId string `json:"job_id,omitempty"`
	// Whether team members are eligible for tips when working this job.
	TipEligible bool `json:"tip_eligible,omitempty"`
}
