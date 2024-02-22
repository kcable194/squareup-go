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

// A tip being returned.
type OrderReturnTip struct {
	// A unique ID that identifies the tip only within this order.
	Uid          string `json:"uid,omitempty"`
	AppliedMoney *Money `json:"applied_money,omitempty"`
	// The tender `uid` from the order that contains the original application of this tip.
	SourceTenderUid string `json:"source_tender_uid,omitempty"`
	// The tender `id` from the order that contains the original application of this tip.
	SourceTenderId string `json:"source_tender_id,omitempty"`
}
