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

// Filter events by date time range.
type LoyaltyEventDateTimeFilter struct {
	CreatedAt *TimeRange `json:"created_at"`
}
