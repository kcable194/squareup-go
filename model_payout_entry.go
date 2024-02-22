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

// One or more PayoutEntries that make up a Payout. Each one has a date, amount, and type of activity. The total amount of the payout will equal the sum of the payout entries for a batch payout
type PayoutEntry struct {
	// A unique ID for the payout entry.
	Id string `json:"id"`
	// The ID of the payout entries’ associated payout.
	PayoutId string `json:"payout_id"`
	// The timestamp of when the payout entry affected the balance, in RFC 3339 format.
	EffectiveAt string `json:"effective_at,omitempty"`
	// The type of activity associated with this payout entry.
	Type_                                   string                                                    `json:"type,omitempty"`
	GrossAmountMoney                        *Money                                                    `json:"gross_amount_money,omitempty"`
	FeeAmountMoney                          *Money                                                    `json:"fee_amount_money,omitempty"`
	NetAmountMoney                          *Money                                                    `json:"net_amount_money,omitempty"`
	TypeAppFeeRevenueDetails                *PaymentBalanceActivityAppFeeRevenueDetail                `json:"type_app_fee_revenue_details,omitempty"`
	TypeAppFeeRefundDetails                 *PaymentBalanceActivityAppFeeRefundDetail                 `json:"type_app_fee_refund_details,omitempty"`
	TypeAutomaticSavingsDetails             *PaymentBalanceActivityAutomaticSavingsDetail             `json:"type_automatic_savings_details,omitempty"`
	TypeAutomaticSavingsReversedDetails     *PaymentBalanceActivityAutomaticSavingsReversedDetail     `json:"type_automatic_savings_reversed_details,omitempty"`
	TypeChargeDetails                       *PaymentBalanceActivityChargeDetail                       `json:"type_charge_details,omitempty"`
	TypeDepositFeeDetails                   *PaymentBalanceActivityDepositFeeDetail                   `json:"type_deposit_fee_details,omitempty"`
	TypeDisputeDetails                      *PaymentBalanceActivityDisputeDetail                      `json:"type_dispute_details,omitempty"`
	TypeFeeDetails                          *PaymentBalanceActivityFeeDetail                          `json:"type_fee_details,omitempty"`
	TypeFreeProcessingDetails               *PaymentBalanceActivityFreeProcessingDetail               `json:"type_free_processing_details,omitempty"`
	TypeHoldAdjustmentDetails               *PaymentBalanceActivityHoldAdjustmentDetail               `json:"type_hold_adjustment_details,omitempty"`
	TypeOpenDisputeDetails                  *PaymentBalanceActivityOpenDisputeDetail                  `json:"type_open_dispute_details,omitempty"`
	TypeOtherDetails                        *PaymentBalanceActivityOtherDetail                        `json:"type_other_details,omitempty"`
	TypeOtherAdjustmentDetails              *PaymentBalanceActivityOtherAdjustmentDetail              `json:"type_other_adjustment_details,omitempty"`
	TypeRefundDetails                       *PaymentBalanceActivityRefundDetail                       `json:"type_refund_details,omitempty"`
	TypeReleaseAdjustmentDetails            *PaymentBalanceActivityReleaseAdjustmentDetail            `json:"type_release_adjustment_details,omitempty"`
	TypeReserveHoldDetails                  *PaymentBalanceActivityReserveHoldDetail                  `json:"type_reserve_hold_details,omitempty"`
	TypeReserveReleaseDetails               *PaymentBalanceActivityReserveReleaseDetail               `json:"type_reserve_release_details,omitempty"`
	TypeSquareCapitalPaymentDetails         *PaymentBalanceActivitySquareCapitalPaymentDetail         `json:"type_square_capital_payment_details,omitempty"`
	TypeSquareCapitalReversedPaymentDetails *PaymentBalanceActivitySquareCapitalReversedPaymentDetail `json:"type_square_capital_reversed_payment_details,omitempty"`
	TypeTaxOnFeeDetails                     *PaymentBalanceActivityTaxOnFeeDetail                     `json:"type_tax_on_fee_details,omitempty"`
	TypeThirdPartyFeeDetails                *PaymentBalanceActivityThirdPartyFeeDetail                `json:"type_third_party_fee_details,omitempty"`
	TypeThirdPartyFeeRefundDetails          *PaymentBalanceActivityThirdPartyFeeRefundDetail          `json:"type_third_party_fee_refund_details,omitempty"`
}
