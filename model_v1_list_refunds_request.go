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

type V1ListRefundsRequest struct {
	// The order in which payments are listed in the response.
	Order string `json:"order,omitempty"`
	// The beginning of the requested reporting period, in ISO 8601 format. If this value is before January 1, 2013 (2013-01-01T00:00:00Z), this endpoint returns an error. Default value: The current time minus one year.
	BeginTime string `json:"begin_time,omitempty"`
	// The end of the requested reporting period, in ISO 8601 format. If this value is more than one year greater than begin_time, this endpoint returns an error. Default value: The current time.
	EndTime string `json:"end_time,omitempty"`
	// The approximate number of refunds to return in a single response. Default: 100. Max: 200. Response may contain more results than the prescribed limit when refunds are made simultaneously to multiple tenders in a payment or when refunds are generated in an exchange to account for the value of returned goods.
	Limit int32 `json:"limit,omitempty"`
	// A pagination cursor to retrieve the next set of results for your original query to the endpoint.
	BatchToken string `json:"batch_token,omitempty"`
}
