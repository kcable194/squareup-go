# UpdateCustomerRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | **string** | The given name (that is, the first name) associated with the customer profile.  The maximum length for this value is 300 characters. | [optional] [default to null]
**FamilyName** | **string** | The family name (that is, the last name) associated with the customer profile.  The maximum length for this value is 300 characters. | [optional] [default to null]
**CompanyName** | **string** | A business name associated with the customer profile.  The maximum length for this value is 500 characters. | [optional] [default to null]
**Nickname** | **string** | A nickname for the customer profile.  The maximum length for this value is 100 characters. | [optional] [default to null]
**EmailAddress** | **string** | The email address associated with the customer profile.  The maximum length for this value is 254 characters. | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**PhoneNumber** | **string** | The phone number associated with the customer profile. The phone number must be valid and can contain 9–16 digits, with an optional &#x60;+&#x60; prefix and country code. For more information, see [Customer phone numbers](https://developer.squareup.com/docs/customers-api/use-the-api/keep-records#phone-number). | [optional] [default to null]
**ReferenceId** | **string** | An optional second ID used to associate the customer profile with an entity in another system.  The maximum length for this value is 100 characters. | [optional] [default to null]
**Note** | **string** | A custom note associated with the customer profile. | [optional] [default to null]
**Birthday** | **string** | The birthday associated with the customer profile, in &#x60;YYYY-MM-DD&#x60; or &#x60;MM-DD&#x60; format. For example, specify &#x60;1998-09-21&#x60; for September 21, 1998, or &#x60;09-21&#x60; for September 21. Birthdays are returned in &#x60;YYYY-MM-DD&#x60; format, where &#x60;YYYY&#x60; is the specified birth year or &#x60;0000&#x60; if a birth year is not specified. | [optional] [default to null]
**Version** | **int64** | The current version of the customer profile.  As a best practice, you should include this field to enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency) control. For more information, see [Update a customer profile](https://developer.squareup.com/docs/customers-api/use-the-api/keep-records#update-a-customer-profile). | [optional] [default to null]
**TaxIds** | [***CustomerTaxIds**](CustomerTaxIds.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

