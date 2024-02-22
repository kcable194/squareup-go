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

// GiftCardGanSource : Indicates the source that generated the gift card  account number (GAN).
type GiftCardGanSource string

// List of GiftCardGANSource
const (
	SQUARE_GiftCardGanSource GiftCardGanSource = "SQUARE"
	OTHER_GiftCardGanSource  GiftCardGanSource = "OTHER"
)
