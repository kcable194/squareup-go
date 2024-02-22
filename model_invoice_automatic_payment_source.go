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
// InvoiceAutomaticPaymentSource : Indicates the automatic payment method for an [invoice payment request](https://developer.squareup.com/reference/square_2024-01-18/objects/InvoicePaymentRequest).
type InvoiceAutomaticPaymentSource string

// List of InvoiceAutomaticPaymentSource
const (
	NONE_InvoiceAutomaticPaymentSource InvoiceAutomaticPaymentSource = "NONE"
	CARD_ON_FILE_InvoiceAutomaticPaymentSource InvoiceAutomaticPaymentSource = "CARD_ON_FILE"
	BANK_ON_FILE_InvoiceAutomaticPaymentSource InvoiceAutomaticPaymentSource = "BANK_ON_FILE"
)
