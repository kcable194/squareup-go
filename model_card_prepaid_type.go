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

// CardPrepaidType : Indicates a card's prepaid type, such as `NOT_PREPAID` or `PREPAID`.
type CardPrepaidType string

// List of CardPrepaidType
const (
	UNKNOWN_PREPAID_TYPE_CardPrepaidType CardPrepaidType = "UNKNOWN_PREPAID_TYPE"
	NOT_PREPAID_CardPrepaidType          CardPrepaidType = "NOT_PREPAID"
	PREPAID_CardPrepaidType              CardPrepaidType = "PREPAID"
)
