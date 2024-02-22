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
// OrderServiceChargeTreatmentType : Indicates whether the service charge will be treated as a value-holding line item or apportioned toward a line item.
type OrderServiceChargeTreatmentType string

// List of OrderServiceChargeTreatmentType
const (
	LINE_ITEM_TREATMENT_OrderServiceChargeTreatmentType OrderServiceChargeTreatmentType = "LINE_ITEM_TREATMENT"
	APPORTIONED_TREATMENT_OrderServiceChargeTreatmentType OrderServiceChargeTreatmentType = "APPORTIONED_TREATMENT"
)
