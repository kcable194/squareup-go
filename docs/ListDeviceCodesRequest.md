# ListDeviceCodesRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query.  See [Paginating results](https://developer.squareup.com/docs/working-with-apis/pagination) for more information. | [optional] [default to null]
**LocationId** | **string** | If specified, only returns DeviceCodes of the specified location. Returns DeviceCodes of all locations if empty. | [optional] [default to null]
**ProductType** | **string** | If specified, only returns DeviceCodes targeting the specified product type. Returns DeviceCodes of all product types if empty. | [optional] [default to null]
**Status** | **[]string** | If specified, returns DeviceCodes with the specified statuses. Returns DeviceCodes of status &#x60;PAIRED&#x60; and &#x60;UNPAIRED&#x60; if empty. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

