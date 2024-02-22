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

// Represents an input to a call to [BulkUpdateVendors](https://developer.squareup.com/reference/square_2024-01-18/vendors-api/bulk-update-vendors).
type BulkUpdateVendorsRequest struct {
	// A set of [UpdateVendorRequest](https://developer.squareup.com/reference/square_2024-01-18/objects/UpdateVendorRequest) objects encapsulating to-be-updated [Vendor](https://developer.squareup.com/reference/square_2024-01-18/objects/Vendor) objects. The set is represented by  a collection of `Vendor`-ID/`UpdateVendorRequest`-object pairs.
	Vendors map[string]UpdateVendorRequest `json:"vendors"`
}
