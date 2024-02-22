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

type OrderUpdated struct {
	// The order's unique ID.
	OrderId string `json:"order_id,omitempty"`
	// The version number, which is incremented each time an update is committed to the order. Orders that were not created through the API do not include a version number and therefore cannot be updated.  [Read more about working with versions.](https://developer.squareup.com/docs/orders-api/manage-orders/update-orders)
	Version int32 `json:"version,omitempty"`
	// The ID of the seller location that this order is associated with.
	LocationId string `json:"location_id,omitempty"`
	// The state of the order.
	State string `json:"state,omitempty"`
	// The timestamp for when the order was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp for when the order was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
}
