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

// The hours of operation for a location.
type BusinessHours struct {
	// The list of time periods during which the business is open. There can be at most 10 periods per day.
	Periods []BusinessHoursPeriod `json:"periods,omitempty"`
}
