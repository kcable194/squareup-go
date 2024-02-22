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

// Describes when the loyalty program expires.
type LoyaltyProgramExpirationPolicy struct {
	// The number of months before points expire, in `P[n]M` RFC 3339 duration format. For example, a value of `P12M` represents a duration of 12 months.  Points are valid through the last day of the month in which they are scheduled to expire. For example, with a  `P12M` duration, points earned on July 6, 2020 expire on August 1, 2021.
	ExpirationDuration string `json:"expiration_duration"`
}