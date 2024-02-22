# {{classname}}

All URIs are relative to *https://connect.squareup.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteCatalogObjects**](CatalogApi.md#BatchDeleteCatalogObjects) | **Post** /v2/catalog/batch-delete | BatchDeleteCatalogObjects
[**BatchRetrieveCatalogObjects**](CatalogApi.md#BatchRetrieveCatalogObjects) | **Post** /v2/catalog/batch-retrieve | BatchRetrieveCatalogObjects
[**BatchUpsertCatalogObjects**](CatalogApi.md#BatchUpsertCatalogObjects) | **Post** /v2/catalog/batch-upsert | BatchUpsertCatalogObjects
[**CatalogInfo**](CatalogApi.md#CatalogInfo) | **Get** /v2/catalog/info | CatalogInfo
[**DeleteCatalogObject**](CatalogApi.md#DeleteCatalogObject) | **Delete** /v2/catalog/object/{object_id} | DeleteCatalogObject
[**ListCatalog**](CatalogApi.md#ListCatalog) | **Get** /v2/catalog/list | ListCatalog
[**RetrieveCatalogObject**](CatalogApi.md#RetrieveCatalogObject) | **Get** /v2/catalog/object/{object_id} | RetrieveCatalogObject
[**SearchCatalogItems**](CatalogApi.md#SearchCatalogItems) | **Post** /v2/catalog/search-catalog-items | SearchCatalogItems
[**SearchCatalogObjects**](CatalogApi.md#SearchCatalogObjects) | **Post** /v2/catalog/search | SearchCatalogObjects
[**UpdateItemModifierLists**](CatalogApi.md#UpdateItemModifierLists) | **Post** /v2/catalog/update-item-modifier-lists | UpdateItemModifierLists
[**UpdateItemTaxes**](CatalogApi.md#UpdateItemTaxes) | **Post** /v2/catalog/update-item-taxes | UpdateItemTaxes
[**UpsertCatalogObject**](CatalogApi.md#UpsertCatalogObject) | **Post** /v2/catalog/object | UpsertCatalogObject

# **BatchDeleteCatalogObjects**
> BatchDeleteCatalogObjectsResponse BatchDeleteCatalogObjects(ctx, body)
BatchDeleteCatalogObjects

Deletes a set of [CatalogItem](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogItem)s based on the provided list of target IDs and returns a set of successfully deleted IDs in the response. Deletion is a cascading event such that all children of the targeted object are also deleted. For example, deleting a CatalogItem will also delete all of its [CatalogItemVariation](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogItemVariation) children.  `BatchDeleteCatalogObjects` succeeds even if only a portion of the targeted IDs can be deleted. The response will only include IDs that were actually deleted.  To ensure consistency, only one delete request is processed at a time per seller account.   While one (batch or non-batch) delete request is being processed, other (batched and non-batched)  delete requests are rejected with the `429` error code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchDeleteCatalogObjectsRequest**](BatchDeleteCatalogObjectsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BatchDeleteCatalogObjectsResponse**](BatchDeleteCatalogObjectsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchRetrieveCatalogObjects**
> BatchRetrieveCatalogObjectsResponse BatchRetrieveCatalogObjects(ctx, body)
BatchRetrieveCatalogObjects

Returns a set of objects based on the provided ID. Each [CatalogItem](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogItem) returned in the set includes all of its child information including: all of its [CatalogItemVariation](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogItemVariation) objects, references to its [CatalogModifierList](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogModifierList) objects, and the ids of any [CatalogTax](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogTax) objects that apply to it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchRetrieveCatalogObjectsRequest**](BatchRetrieveCatalogObjectsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BatchRetrieveCatalogObjectsResponse**](BatchRetrieveCatalogObjectsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchUpsertCatalogObjects**
> BatchUpsertCatalogObjectsResponse BatchUpsertCatalogObjects(ctx, body)
BatchUpsertCatalogObjects

Creates or updates up to 10,000 target objects based on the provided list of objects. The target objects are grouped into batches and each batch is inserted/updated in an all-or-nothing manner. If an object within a batch is malformed in some way, or violates a database constraint, the entire batch containing that item will be disregarded. However, other batches in the same request may still succeed. Each batch may contain up to 1,000 objects, and batches will be processed in order as long as the total object count for the request (items, variations, modifier lists, discounts, and taxes) is no more than 10,000.  To ensure consistency, only one update request is processed at a time per seller account.   While one (batch or non-batch) update request is being processed, other (batched and non-batched)  update requests are rejected with the `429` error code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchUpsertCatalogObjectsRequest**](BatchUpsertCatalogObjectsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BatchUpsertCatalogObjectsResponse**](BatchUpsertCatalogObjectsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CatalogInfo**
> CatalogInfoResponse CatalogInfo(ctx, )
CatalogInfo

Retrieves information about the Square Catalog API, such as batch size limits that can be used by the `BatchUpsertCatalogObjects` endpoint.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CatalogInfoResponse**](CatalogInfoResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCatalogObject**
> DeleteCatalogObjectResponse DeleteCatalogObject(ctx, objectId)
DeleteCatalogObject

Deletes a single [CatalogObject](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogObject) based on the provided ID and returns the set of successfully deleted IDs in the response. Deletion is a cascading event such that all children of the targeted object are also deleted. For example, deleting a [CatalogItem](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogItem) will also delete all of its [CatalogItemVariation](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogItemVariation) children.  To ensure consistency, only one delete request is processed at a time per seller account.   While one (batch or non-batch) delete request is being processed, other (batched and non-batched)  delete requests are rejected with the `429` error code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The ID of the catalog object to be deleted. When an object is deleted, other objects in the graph that depend on that object will be deleted as well (for example, deleting a catalog item will delete its catalog item variations). | 

### Return type

[**DeleteCatalogObjectResponse**](DeleteCatalogObjectResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCatalog**
> ListCatalogResponse ListCatalog(ctx, optional)
ListCatalog

Returns a list of all [CatalogObject](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogObject)s of the specified types in the catalog.   The `types` parameter is specified as a comma-separated list of the [CatalogObjectType](https://developer.squareup.com/reference/square_2024-01-18/enums/CatalogObjectType) values,  for example, \"`ITEM`, `ITEM_VARIATION`, `MODIFIER`, `MODIFIER_LIST`, `CATEGORY`, `DISCOUNT`, `TAX`, `IMAGE`\".  __Important:__ ListCatalog does not return deleted catalog items. To retrieve deleted catalog items, use [SearchCatalogObjects](https://developer.squareup.com/reference/square_2024-01-18/catalog-api/search-catalog-objects) and set the `include_deleted_objects` attribute value to `true`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CatalogApiListCatalogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogApiListCatalogOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| The pagination cursor returned in the previous response. Leave unset for an initial request. The page size is currently set to be 100. See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information. | 
 **types** | **optional.String**| An optional case-insensitive, comma-separated list of object types to retrieve.  The valid values are defined in the [CatalogObjectType](https://developer.squareup.com/reference/square_2024-01-18/enums/CatalogObjectType) enum, for example, &#x60;ITEM&#x60;, &#x60;ITEM_VARIATION&#x60;, &#x60;CATEGORY&#x60;, &#x60;DISCOUNT&#x60;, &#x60;TAX&#x60;, &#x60;MODIFIER&#x60;, &#x60;MODIFIER_LIST&#x60;, &#x60;IMAGE&#x60;, etc.  If this is unspecified, the operation returns objects of all the top level types at the version of the Square API used to make the request. Object types that are nested onto other object types are not included in the defaults.  At the current API version the default object types are: ITEM, CATEGORY, TAX, DISCOUNT, MODIFIER_LIST,  PRICING_RULE, PRODUCT_SET, TIME_PERIOD, MEASUREMENT_UNIT, SUBSCRIPTION_PLAN, ITEM_OPTION, CUSTOM_ATTRIBUTE_DEFINITION, QUICK_AMOUNT_SETTINGS. | 
 **catalogVersion** | **optional.Int64**| The specific version of the catalog objects to be included in the response. This allows you to retrieve historical versions of objects. The specified version value is matched against the [CatalogObject](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogObject)s&#x27; &#x60;version&#x60; attribute.  If not included, results will be from the current version of the catalog. | 

### Return type

[**ListCatalogResponse**](ListCatalogResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCatalogObject**
> RetrieveCatalogObjectResponse RetrieveCatalogObject(ctx, objectId, optional)
RetrieveCatalogObject

Returns a single [CatalogItem](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogItem) as a [CatalogObject](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogObject) based on the provided ID. The returned object includes all of the relevant [CatalogItem](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogItem) information including: [CatalogItemVariation](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogItemVariation) children, references to its [CatalogModifierList](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogModifierList) objects, and the ids of any [CatalogTax](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogTax) objects that apply to it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The object ID of any type of catalog objects to be retrieved. | 
 **optional** | ***CatalogApiRetrieveCatalogObjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogApiRetrieveCatalogObjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeRelatedObjects** | **optional.Bool**| If &#x60;true&#x60;, the response will include additional objects that are related to the requested objects. Related objects are defined as any objects referenced by ID by the results in the &#x60;objects&#x60; field of the response. These objects are put in the &#x60;related_objects&#x60; field. Setting this to &#x60;true&#x60; is helpful when the objects are needed for immediate display to a user. This process only goes one level deep. Objects referenced by the related objects will not be included. For example,  if the &#x60;objects&#x60; field of the response contains a CatalogItem, its associated CatalogCategory objects, CatalogTax objects, CatalogImage objects and CatalogModifierLists will be returned in the &#x60;related_objects&#x60; field of the response. If the &#x60;objects&#x60; field of the response contains a CatalogItemVariation, its parent CatalogItem will be returned in the &#x60;related_objects&#x60; field of the response.  Default value: &#x60;false&#x60; | 
 **catalogVersion** | **optional.Int64**| Requests objects as of a specific version of the catalog. This allows you to retrieve historical versions of objects. The value to retrieve a specific version of an object can be found in the version field of [CatalogObject](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogObject)s. If not included, results will be from the current version of the catalog. | 
 **includeCategoryPathToRoot** | **optional.Bool**| Specifies whether or not to include the &#x60;path_to_root&#x60; list for each returned category instance. The &#x60;path_to_root&#x60; list consists of &#x60;CategoryPathToRootNode&#x60; objects and specifies the path that starts with the immediate parent category of the returned category and ends with its root category. If the returned category is a top-level category, the &#x60;path_to_root&#x60; list is empty and is not returned in the response payload. | 

### Return type

[**RetrieveCatalogObjectResponse**](RetrieveCatalogObjectResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCatalogItems**
> SearchCatalogItemsResponse SearchCatalogItems(ctx, body)
SearchCatalogItems

Searches for catalog items or item variations by matching supported search attribute values, including custom attribute values, against one or more of the specified query filters.  This (`SearchCatalogItems`) endpoint differs from the [SearchCatalogObjects](https://developer.squareup.com/reference/square_2024-01-18/catalog-api/search-catalog-objects) endpoint in the following aspects:  - `SearchCatalogItems` can only search for items or item variations, whereas `SearchCatalogObjects` can search for any type of catalog objects. - `SearchCatalogItems` supports the custom attribute query filters to return items or item variations that contain custom attribute values, where `SearchCatalogObjects` does not. - `SearchCatalogItems` does not support the `include_deleted_objects` filter to search for deleted items or item variations, whereas `SearchCatalogObjects` does. - The both endpoints use different call conventions, including the query filter formats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchCatalogItemsRequest**](SearchCatalogItemsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchCatalogItemsResponse**](SearchCatalogItemsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCatalogObjects**
> SearchCatalogObjectsResponse SearchCatalogObjects(ctx, body)
SearchCatalogObjects

Searches for [CatalogObject](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogObject) of any type by matching supported search attribute values, excluding custom attribute values on items or item variations, against one or more of the specified query filters.  This (`SearchCatalogObjects`) endpoint differs from the [SearchCatalogItems](https://developer.squareup.com/reference/square_2024-01-18/catalog-api/search-catalog-items) endpoint in the following aspects:  - `SearchCatalogItems` can only search for items or item variations, whereas `SearchCatalogObjects` can search for any type of catalog objects. - `SearchCatalogItems` supports the custom attribute query filters to return items or item variations that contain custom attribute values, where `SearchCatalogObjects` does not. - `SearchCatalogItems` does not support the `include_deleted_objects` filter to search for deleted items or item variations, whereas `SearchCatalogObjects` does. - The both endpoints have different call conventions, including the query filter formats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchCatalogObjectsRequest**](SearchCatalogObjectsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchCatalogObjectsResponse**](SearchCatalogObjectsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateItemModifierLists**
> UpdateItemModifierListsResponse UpdateItemModifierLists(ctx, body)
UpdateItemModifierLists

Updates the [CatalogModifierList](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogModifierList) objects that apply to the targeted [CatalogItem](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogItem) without having to perform an upsert on the entire item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateItemModifierListsRequest**](UpdateItemModifierListsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**UpdateItemModifierListsResponse**](UpdateItemModifierListsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateItemTaxes**
> UpdateItemTaxesResponse UpdateItemTaxes(ctx, body)
UpdateItemTaxes

Updates the [CatalogTax](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogTax) objects that apply to the targeted [CatalogItem](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogItem) without having to perform an upsert on the entire item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateItemTaxesRequest**](UpdateItemTaxesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**UpdateItemTaxesResponse**](UpdateItemTaxesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpsertCatalogObject**
> UpsertCatalogObjectResponse UpsertCatalogObject(ctx, body)
UpsertCatalogObject

Creates a new or updates the specified [CatalogObject](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogObject).  To ensure consistency, only one update request is processed at a time per seller account.   While one (batch or non-batch) update request is being processed, other (batched and non-batched)  update requests are rejected with the `429` error code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpsertCatalogObjectRequest**](UpsertCatalogObjectRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**UpsertCatalogObjectResponse**](UpsertCatalogObjectResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

