# BulkUpdateVendorsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Errors encountered when the request fails. | [optional] [default to null]
**Responses** | [**map[string]UpdateVendorResponse**](UpdateVendorResponse.md) | A set of [UpdateVendorResponse](https://developer.squareup.com/reference/square_2024-01-18/objects/UpdateVendorResponse) objects encapsulating successfully created [Vendor](https://developer.squareup.com/reference/square_2024-01-18/objects/Vendor) objects or error responses for failed attempts. The set is represented by a collection of &#x60;Vendor&#x60;-ID/&#x60;UpdateVendorResponse&#x60;-object or  &#x60;Vendor&#x60;-ID/error-object pairs. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

