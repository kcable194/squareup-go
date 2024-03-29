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

// Represents an input to a call to [BulkCreateVendors](https://developer.squareup.com/reference/square_2024-01-18/vendors-api/bulk-create-vendors).
type BulkCreateVendorsRequest struct {
	// Specifies a set of new [Vendor](https://developer.squareup.com/reference/square_2024-01-18/objects/Vendor) objects as represented by a collection of idempotency-key/`Vendor`-object pairs.
	Vendors map[string]Vendor `json:"vendors"`
}
