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

type DeviceCode struct {
	// The unique id for this device code.
	Id string `json:"id,omitempty"`
	// An optional user-defined name for the device code.
	Name string `json:"name,omitempty"`
	// The unique code that can be used to login.
	Code string `json:"code,omitempty"`
	// The unique id of the device that used this code. Populated when the device is paired up.
	DeviceId string `json:"device_id,omitempty"`
	// The targeting product type of the device code.
	ProductType string `json:"product_type"`
	// The location assigned to this code.
	LocationId string `json:"location_id,omitempty"`
	// The pairing status of the device code.
	Status string `json:"status,omitempty"`
	// When this DeviceCode will expire and no longer login. Timestamp in RFC 3339 format.
	PairBy string `json:"pair_by,omitempty"`
	// When this DeviceCode was created. Timestamp in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// When this DeviceCode's status was last changed. Timestamp in RFC 3339 format.
	StatusChangedAt string `json:"status_changed_at,omitempty"`
	// When this DeviceCode was paired. Timestamp in RFC 3339 format.
	PairedAt string `json:"paired_at,omitempty"`
}
