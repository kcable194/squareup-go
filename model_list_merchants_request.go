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

// Request object for the [ListMerchant](https://developer.squareup.com/reference/square_2024-01-18/merchants-api/list-merchants) endpoint.
type ListMerchantsRequest struct {
	// The cursor generated by the previous response.
	Cursor int32 `json:"cursor,omitempty"`
}
