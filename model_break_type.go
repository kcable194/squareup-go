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

// A defined break template that sets an expectation for possible `Break` instances on a `Shift`.
type BreakType struct {
	// The UUID for this object.
	Id string `json:"id,omitempty"`
	// The ID of the business location this type of break applies to.
	LocationId string `json:"location_id"`
	// A human-readable name for this type of break. The name is displayed to employees in Square products.
	BreakName string `json:"break_name"`
	// Format: RFC-3339 P[n]Y[n]M[n]DT[n]H[n]M[n]S. The expected length of this break. Precision less than minutes is truncated.  Example for break expected duration of 15 minutes: T15M
	ExpectedDuration string `json:"expected_duration"`
	// Whether this break counts towards time worked for compensation purposes.
	IsPaid bool `json:"is_paid"`
	// Used for resolving concurrency issues. The request fails if the version provided does not match the server version at the time of the request. If a value is not provided, Square's servers execute a \"blind\" write; potentially overwriting another writer's data.
	Version int32 `json:"version,omitempty"`
	// A read-only timestamp in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// A read-only timestamp in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
}
