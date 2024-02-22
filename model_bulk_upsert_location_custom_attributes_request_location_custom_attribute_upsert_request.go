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

// Represents an individual upsert request in a [BulkUpsertLocationCustomAttributes](https://developer.squareup.com/reference/square_2024-01-18/location-custom-attributes-api/bulk-upsert-location-custom-attributes) request. An individual request contains a location ID, the custom attribute to create or update, and an optional idempotency key.
type BulkUpsertLocationCustomAttributesRequestLocationCustomAttributeUpsertRequest struct {
	// The ID of the target [location](https://developer.squareup.com/reference/square_2024-01-18/objects/Location).
	LocationId      string           `json:"location_id"`
	CustomAttribute *CustomAttribute `json:"custom_attribute"`
	// A unique identifier for this individual upsert request, used to ensure idempotency. For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}
