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

type PaymentBalanceActivityDepositFeeDetail struct {
	// The ID of the payout that triggered this deposit fee activity.
	PayoutId string `json:"payout_id,omitempty"`
}
