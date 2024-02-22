# RetrieveTokenStatusResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | **[]string** | The list of scopes associated with an access token. | [optional] [default to null]
**ExpiresAt** | **string** | The date and time when the &#x60;access_token&#x60; expires, in RFC 3339 format. Empty if the token never expires. | [optional] [default to null]
**ClientId** | **string** | The Square-issued application ID associated with the access token. This is the same application ID used to obtain the token. | [optional] [default to null]
**MerchantId** | **string** | The ID of the authorizing merchant&#x27;s business. | [optional] [default to null]
**Errors** | [**[]ModelError**](Error.md) |  Any errors that occurred during the request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

