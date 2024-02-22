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
// CashDrawerEventType : The types of events on a CashDrawerShift. Each event type represents an employee action on the actual cash drawer represented by a CashDrawerShift.
type CashDrawerEventType string

// List of CashDrawerEventType
const (
	NO_SALE_CashDrawerEventType CashDrawerEventType = "NO_SALE"
	CASH_TENDER_PAYMENT_CashDrawerEventType CashDrawerEventType = "CASH_TENDER_PAYMENT"
	OTHER_TENDER_PAYMENT_CashDrawerEventType CashDrawerEventType = "OTHER_TENDER_PAYMENT"
	CASH_TENDER_CANCELLED_PAYMENT_CashDrawerEventType CashDrawerEventType = "CASH_TENDER_CANCELLED_PAYMENT"
	OTHER_TENDER_CANCELLED_PAYMENT_CashDrawerEventType CashDrawerEventType = "OTHER_TENDER_CANCELLED_PAYMENT"
	CASH_TENDER_REFUND_CashDrawerEventType CashDrawerEventType = "CASH_TENDER_REFUND"
	OTHER_TENDER_REFUND_CashDrawerEventType CashDrawerEventType = "OTHER_TENDER_REFUND"
	PAID_IN_CashDrawerEventType CashDrawerEventType = "PAID_IN"
	PAID_OUT_CashDrawerEventType CashDrawerEventType = "PAID_OUT"
)
