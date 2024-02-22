# Customer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique Square-assigned ID for the customer profile.  If you need this ID for an API request, use the ID returned when you created the customer profile or call the [SearchCustomers](https://developer.squareup.com/reference/square_2024-01-18/customers-api/search-customers)  or [ListCustomers](https://developer.squareup.com/reference/square_2024-01-18/customers-api/list-customers) endpoint. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp when the customer profile was created, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp when the customer profile was last updated, in RFC 3339 format. | [optional] [default to null]
**Cards** | [**[]Card**](Card.md) | Payment details of the credit, debit, and gift cards stored on file for the customer profile.   DEPRECATED at version 2021-06-16. Replaced by calling [ListCards](https://developer.squareup.com/reference/square_2024-01-18/cards-api/list-cards) (for credit and debit cards on file)  or [ListGiftCards](https://developer.squareup.com/reference/square_2024-01-18/gift-cards-api/list-gift-cards) (for gift cards on file) and including the &#x60;customer_id&#x60; query parameter.  For more information, see [Migration notes](https://developer.squareup.com/docs/customers-api/what-it-does#migrate-customer-cards). | [optional] [default to null]
**GivenName** | **string** | The given name (that is, the first name) associated with the customer profile. | [optional] [default to null]
**FamilyName** | **string** | The family name (that is, the last name) associated with the customer profile. | [optional] [default to null]
**Nickname** | **string** | A nickname for the customer profile. | [optional] [default to null]
**CompanyName** | **string** | A business name associated with the customer profile. | [optional] [default to null]
**EmailAddress** | **string** | The email address associated with the customer profile. | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**PhoneNumber** | **string** | The phone number associated with the customer profile. | [optional] [default to null]
**Birthday** | **string** | The birthday associated with the customer profile, in &#x60;YYYY-MM-DD&#x60; format. For example, &#x60;1998-09-21&#x60; represents September 21, 1998, and &#x60;0000-09-21&#x60; represents September 21 (without a birth year). | [optional] [default to null]
**ReferenceId** | **string** | An optional second ID used to associate the customer profile with an entity in another system. | [optional] [default to null]
**Note** | **string** | A custom note associated with the customer profile. | [optional] [default to null]
**Preferences** | [***CustomerPreferences**](CustomerPreferences.md) |  | [optional] [default to null]
**CreationSource** | **string** | The method used to create the customer profile. | [optional] [default to null]
**GroupIds** | **[]string** | The IDs of [customer groups](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomerGroup) the customer belongs to. | [optional] [default to null]
**SegmentIds** | **[]string** | The IDs of [customer segments](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomerSegment) the customer belongs to. | [optional] [default to null]
**Version** | **int64** | The Square-assigned version number of the customer profile. The version number is incremented each time an update is committed to the customer profile, except for changes to customer segment membership and cards on file. | [optional] [default to null]
**TaxIds** | [***TaxIds**](TaxIds.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

