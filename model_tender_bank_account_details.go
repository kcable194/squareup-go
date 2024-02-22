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

// Represents the details of a tender with `type` `BANK_ACCOUNT`.  See [BankAccountPaymentDetails](https://developer.squareup.com/reference/square_2024-01-18/objects/BankAccountPaymentDetails) for more exposed details of a bank account payment.
type TenderBankAccountDetails struct {
	// The bank account payment's current state.  See [TenderBankAccountPaymentDetailsStatus](https://developer.squareup.com/reference/square_2024-01-18/enums/TenderBankAccountDetailsStatus) for possible values.
	Status string `json:"status,omitempty"`
}
