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

// OrderFulfillmentType : The type of fulfillment.
type OrderFulfillmentType string

// List of OrderFulfillmentType
const (
	PICKUP_OrderFulfillmentType   OrderFulfillmentType = "PICKUP"
	SHIPMENT_OrderFulfillmentType OrderFulfillmentType = "SHIPMENT"
	DELIVERY_OrderFulfillmentType OrderFulfillmentType = "DELIVERY"
)
