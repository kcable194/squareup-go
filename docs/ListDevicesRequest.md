# ListDevicesRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information. | [optional] [default to null]
**SortOrder** | **string** | The order in which results are listed. - &#x60;ASC&#x60; - Oldest to newest. - &#x60;DESC&#x60; - Newest to oldest (default). | [optional] [default to null]
**Limit** | **int32** | The number of results to return in a single page. | [optional] [default to null]
**LocationId** | **string** | If present, only returns devices at the target location. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

