# BulkDeleteOrderCustomAttributesResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Any errors that occurred during the request. | [optional] [default to null]
**Values** | [**map[string]DeleteOrderCustomAttributeResponse**](DeleteOrderCustomAttributeResponse.md) |  A map of responses that correspond to individual delete requests. Each response has the same ID  as the corresponding request and contains either a &#x60;custom_attribute&#x60; or an &#x60;errors&#x60; field. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

