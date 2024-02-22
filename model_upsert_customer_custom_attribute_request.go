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

// Represents an [UpsertCustomerCustomAttribute](https://developer.squareup.com/reference/square_2024-01-18/customer-custom-attributes-api/upsert-customer-custom-attribute) request.
type UpsertCustomerCustomAttributeRequest struct {
	CustomAttribute *CustomAttribute `json:"custom_attribute"`
	// A unique identifier for this request, used to ensure idempotency. For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}
