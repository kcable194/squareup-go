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

// Represents a [dispute](https://developer.squareup.com/docs/disputes-api/overview) a cardholder initiated with their bank.
type Dispute struct {
	// The unique ID for this `Dispute`, generated by Square.
	DisputeId string `json:"dispute_id,omitempty"`
	// The unique ID for this `Dispute`, generated by Square.
	Id string `json:"id,omitempty"`
	AmountMoney *Money `json:"amount_money,omitempty"`
	// The reason why the cardholder initiated the dispute.
	Reason string `json:"reason,omitempty"`
	// The current state of this dispute.
	State string `json:"state,omitempty"`
	// The deadline by which the seller must respond to the dispute, in [RFC 3339 format](https://developer.squareup.com/docs/build-basics/common-data-types/working-with-dates).
	DueAt string `json:"due_at,omitempty"`
	DisputedPayment *DisputedPayment `json:"disputed_payment,omitempty"`
	// The IDs of the evidence associated with the dispute.
	EvidenceIds []string `json:"evidence_ids,omitempty"`
	// The card brand used in the disputed payment.
	CardBrand string `json:"card_brand,omitempty"`
	// The timestamp when the dispute was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp when the dispute was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
	// The ID of the dispute in the card brand system, generated by the card brand.
	BrandDisputeId string `json:"brand_dispute_id,omitempty"`
	// The timestamp when the dispute was reported, in RFC 3339 format.
	ReportedDate string `json:"reported_date,omitempty"`
	// The timestamp when the dispute was reported, in RFC 3339 format.
	ReportedAt string `json:"reported_at,omitempty"`
	// The current version of the `Dispute`.
	Version int32 `json:"version,omitempty"`
	// The ID of the location where the dispute originated.
	LocationId string `json:"location_id,omitempty"`
}