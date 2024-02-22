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

// Defines the fields that are included in requests to the [UpdateOrder](https://developer.squareup.com/reference/square_2024-01-18/orders-api/update-order) endpoint.
type UpdateOrderRequest struct {
	Order *Order `json:"order,omitempty"`
	// The [dot notation paths](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders#identifying-fields-to-delete) fields to clear. For example, `line_items[uid].note`. For more information, see [Deleting fields](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders#deleting-fields).
	FieldsToClear []string `json:"fields_to_clear,omitempty"`
	// A value you specify that uniquely identifies this update request.  If you are unsure whether a particular update was applied to an order successfully, you can reattempt it with the same idempotency key without worrying about creating duplicate updates to the order. The latest order version is returned.  For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}
