# ListCashDrawerShiftsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | **string** | The ID of the location to query for a list of cash drawer shifts. | [default to null]
**SortOrder** | **string** | The order in which cash drawer shifts are listed in the response, based on their opened_at field. Default value: ASC | [optional] [default to null]
**BeginTime** | **string** | The inclusive start time of the query on opened_at, in ISO 8601 format. | [optional] [default to null]
**EndTime** | **string** | The exclusive end date of the query on opened_at, in ISO 8601 format. | [optional] [default to null]
**Limit** | **int32** | Number of cash drawer shift events in a page of results (200 by default, 1000 max). | [optional] [default to null]
**Cursor** | **string** | Opaque cursor for fetching the next page of results. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

