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

type CreateDeviceCodeRequest struct {
	// A unique string that identifies this CreateDeviceCode request. Keys can be any valid string but must be unique for every CreateDeviceCode request.  See [Idempotency keys](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) for more information.
	IdempotencyKey string `json:"idempotency_key"`
	DeviceCode *DeviceCode `json:"device_code"`
}
