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

// SubscriptionEventInfoCode : Supported info codes of a subscription event.
type SubscriptionEventInfoCode string

// List of SubscriptionEventInfoCode
const (
	LOCATION_NOT_ACTIVE_SubscriptionEventInfoCode            SubscriptionEventInfoCode = "LOCATION_NOT_ACTIVE"
	LOCATION_CANNOT_ACCEPT_PAYMENT_SubscriptionEventInfoCode SubscriptionEventInfoCode = "LOCATION_CANNOT_ACCEPT_PAYMENT"
	CUSTOMER_DELETED_SubscriptionEventInfoCode               SubscriptionEventInfoCode = "CUSTOMER_DELETED"
	CUSTOMER_NO_EMAIL_SubscriptionEventInfoCode              SubscriptionEventInfoCode = "CUSTOMER_NO_EMAIL"
	CUSTOMER_NO_NAME_SubscriptionEventInfoCode               SubscriptionEventInfoCode = "CUSTOMER_NO_NAME"
	USER_PROVIDED_SubscriptionEventInfoCode                  SubscriptionEventInfoCode = "USER_PROVIDED"
)
