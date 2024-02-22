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

// Represents a Square customer profile in the Customer Directory of a Square seller.
type Customer struct {
	// A unique Square-assigned ID for the customer profile.  If you need this ID for an API request, use the ID returned when you created the customer profile or call the [SearchCustomers](https://developer.squareup.com/reference/square_2024-01-18/customers-api/search-customers)  or [ListCustomers](https://developer.squareup.com/reference/square_2024-01-18/customers-api/list-customers) endpoint.
	Id string `json:"id,omitempty"`
	// The timestamp when the customer profile was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp when the customer profile was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
	// Payment details of the credit, debit, and gift cards stored on file for the customer profile.   DEPRECATED at version 2021-06-16. Replaced by calling [ListCards](https://developer.squareup.com/reference/square_2024-01-18/cards-api/list-cards) (for credit and debit cards on file)  or [ListGiftCards](https://developer.squareup.com/reference/square_2024-01-18/gift-cards-api/list-gift-cards) (for gift cards on file) and including the `customer_id` query parameter.  For more information, see [Migration notes](https://developer.squareup.com/docs/customers-api/what-it-does#migrate-customer-cards).
	Cards []Card `json:"cards,omitempty"`
	// The given name (that is, the first name) associated with the customer profile.
	GivenName string `json:"given_name,omitempty"`
	// The family name (that is, the last name) associated with the customer profile.
	FamilyName string `json:"family_name,omitempty"`
	// A nickname for the customer profile.
	Nickname string `json:"nickname,omitempty"`
	// A business name associated with the customer profile.
	CompanyName string `json:"company_name,omitempty"`
	// The email address associated with the customer profile.
	EmailAddress string   `json:"email_address,omitempty"`
	Address      *Address `json:"address,omitempty"`
	// The phone number associated with the customer profile.
	PhoneNumber string `json:"phone_number,omitempty"`
	// The birthday associated with the customer profile, in `YYYY-MM-DD` format. For example, `1998-09-21` represents September 21, 1998, and `0000-09-21` represents September 21 (without a birth year).
	Birthday string `json:"birthday,omitempty"`
	// An optional second ID used to associate the customer profile with an entity in another system.
	ReferenceId string `json:"reference_id,omitempty"`
	// A custom note associated with the customer profile.
	Note        string               `json:"note,omitempty"`
	Preferences *CustomerPreferences `json:"preferences,omitempty"`
	// The method used to create the customer profile.
	CreationSource string `json:"creation_source,omitempty"`
	// The IDs of [customer groups](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomerGroup) the customer belongs to.
	GroupIds []string `json:"group_ids,omitempty"`
	// The IDs of [customer segments](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomerSegment) the customer belongs to.
	SegmentIds []string `json:"segment_ids,omitempty"`
	// The Square-assigned version number of the customer profile. The version number is incremented each time an update is committed to the customer profile, except for changes to customer segment membership and cards on file.
	Version int64   `json:"version,omitempty"`
	TaxIds  *TaxIds `json:"tax_ids,omitempty"`
}
