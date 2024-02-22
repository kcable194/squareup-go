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
// FulfillmentType : The type of fulfillment.
type FulfillmentType string

// List of FulfillmentType
const (
	PICKUP_FulfillmentType FulfillmentType = "PICKUP"
	SHIPMENT_FulfillmentType FulfillmentType = "SHIPMENT"
	DELIVERY_FulfillmentType FulfillmentType = "DELIVERY"
)