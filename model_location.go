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

// Represents one of a business' [locations](https://developer.squareup.com/docs/locations-api).
type Location struct {
	// A short generated string of letters and numbers that uniquely identifies this location instance.
	Id string `json:"id,omitempty"`
	// The name of the location. This information appears in the Seller Dashboard as the nickname. A location name must be unique within a seller account.
	Name string `json:"name,omitempty"`
	Address *Address `json:"address,omitempty"`
	// The [IANA time zone](https://www.iana.org/time-zones) identifier for the time zone of the location. For example, `America/Los_Angeles`.
	Timezone string `json:"timezone,omitempty"`
	// The Square features that are enabled for the location. See [LocationCapability](https://developer.squareup.com/reference/square_2024-01-18/enums/LocationCapability) for possible values.
	Capabilities []string `json:"capabilities,omitempty"`
	// The status of the location.
	Status string `json:"status,omitempty"`
	// The time when the location was created, in RFC 3339 format. For more information, see [Working with Dates](https://developer.squareup.com/docs/build-basics/working-with-dates).
	CreatedAt string `json:"created_at,omitempty"`
	// The ID of the merchant that owns the location.
	MerchantId string `json:"merchant_id,omitempty"`
	// The country of the location, in the two-letter format of ISO 3166. For example, `US` or `JP`.  See [Country](https://developer.squareup.com/reference/square_2024-01-18/enums/Country) for possible values.
	Country string `json:"country,omitempty"`
	// The language associated with the location, in [BCP 47 format](https://tools.ietf.org/html/bcp47#appendix-A). For more information, see [Language Preferences](https://developer.squareup.com/docs/build-basics/general-considerations/language-preferences).
	LanguageCode string `json:"language_code,omitempty"`
	// The currency used for all transactions at this location, in ISO 4217 format. For example, the currency code for US dollars is `USD`. See [Currency](https://developer.squareup.com/reference/square_2024-01-18/enums/Currency) for possible values.
	Currency string `json:"currency,omitempty"`
	// The phone number of the location. For example, `+1 855-700-6000`.
	PhoneNumber string `json:"phone_number,omitempty"`
	// The name of the location's overall business. This name is present on receipts and other customer-facing branding, and can be changed no more than three times in a twelve-month period.
	BusinessName string `json:"business_name,omitempty"`
	// The type of the location.
	Type_ string `json:"type,omitempty"`
	// The website URL of the location.  For example, `https://squareup.com`.
	WebsiteUrl string `json:"website_url,omitempty"`
	BusinessHours *BusinessHours `json:"business_hours,omitempty"`
	// The email address of the location. This can be unique to the location and is not always the email address for the business owner or administrator.
	BusinessEmail string `json:"business_email,omitempty"`
	// The description of the location. For example, `Main Street location`.
	Description string `json:"description,omitempty"`
	// The Twitter username of the location without the '&#64;' symbol. For example, `Square`.
	TwitterUsername string `json:"twitter_username,omitempty"`
	// The Instagram username of the location without the '&#64;' symbol. For example, `square`.
	InstagramUsername string `json:"instagram_username,omitempty"`
	// The Facebook profile URL of the location. The URL should begin with 'facebook.com/'. For example, `https://www.facebook.com/square`.
	FacebookUrl string `json:"facebook_url,omitempty"`
	Coordinates *Coordinates `json:"coordinates,omitempty"`
	// The URL of the logo image for the location. When configured in the Seller Dashboard (Receipts section), the logo appears on transactions (such as receipts and invoices) that Square generates on behalf of the seller. This image should have a roughly square (1:1) aspect ratio and should be at least 200x200 pixels.
	LogoUrl string `json:"logo_url,omitempty"`
	// The URL of the Point of Sale background image for the location.
	PosBackgroundUrl string `json:"pos_background_url,omitempty"`
	// A four-digit number that describes the kind of goods or services sold at the location. The [merchant category code (MCC)](https://developer.squareup.com/docs/locations-api#initialize-a-merchant-category-code) of the location as standardized by ISO 18245. For example, `5045`, for a location that sells computer goods and software.
	Mcc string `json:"mcc,omitempty"`
	// The URL of a full-format logo image for the location. When configured in the Seller Dashboard (Receipts section), the logo appears on transactions (such as receipts and invoices) that Square generates on behalf of the seller. This image can be wider than it is tall and should be at least 1280x648 pixels.
	FullFormatLogoUrl string `json:"full_format_logo_url,omitempty"`
	TaxIds *TaxIds `json:"tax_ids,omitempty"`
}
