# SearchOrdersRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationIds** | **[]string** | The location IDs for the orders to query. All locations must belong to the same merchant.  Max: 10 location IDs. | [optional] [default to null]
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for your original query. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | [optional] [default to null]
**Query** | [***SearchOrdersQuery**](SearchOrdersQuery.md) |  | [optional] [default to null]
**Limit** | **int32** | The maximum number of results to be returned in a single page.  Default: &#x60;500&#x60; Max: &#x60;1000&#x60; | [optional] [default to null]
**ReturnEntries** | **bool** | A Boolean that controls the format of the search results. If &#x60;true&#x60;, &#x60;SearchOrders&#x60; returns [OrderEntry](https://developer.squareup.com/reference/square_2024-01-18/objects/OrderEntry) objects. If &#x60;false&#x60;, &#x60;SearchOrders&#x60; returns complete order objects.  Default: &#x60;false&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

