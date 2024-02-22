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

type ListDevicesResponse struct {
	// Information about errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The requested list of `Device` objects.
	Devices []Device `json:"devices,omitempty"`
	// The pagination cursor to be used in a subsequent request. If empty, this is the final response. See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information.
	Cursor string `json:"cursor,omitempty"`
}