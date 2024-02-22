# {{classname}}

All URIs are relative to *https://connect.squareup.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteOrderCustomAttributes**](OrderCustomAttributesApi.md#BulkDeleteOrderCustomAttributes) | **Post** /v2/orders/custom-attributes/bulk-delete | BulkDeleteOrderCustomAttributes
[**BulkUpsertOrderCustomAttributes**](OrderCustomAttributesApi.md#BulkUpsertOrderCustomAttributes) | **Post** /v2/orders/custom-attributes/bulk-upsert | BulkUpsertOrderCustomAttributes
[**CreateOrderCustomAttributeDefinition**](OrderCustomAttributesApi.md#CreateOrderCustomAttributeDefinition) | **Post** /v2/orders/custom-attribute-definitions | CreateOrderCustomAttributeDefinition
[**DeleteOrderCustomAttribute**](OrderCustomAttributesApi.md#DeleteOrderCustomAttribute) | **Delete** /v2/orders/{order_id}/custom-attributes/{custom_attribute_key} | DeleteOrderCustomAttribute
[**DeleteOrderCustomAttributeDefinition**](OrderCustomAttributesApi.md#DeleteOrderCustomAttributeDefinition) | **Delete** /v2/orders/custom-attribute-definitions/{key} | DeleteOrderCustomAttributeDefinition
[**ListOrderCustomAttributeDefinitions**](OrderCustomAttributesApi.md#ListOrderCustomAttributeDefinitions) | **Get** /v2/orders/custom-attribute-definitions | ListOrderCustomAttributeDefinitions
[**ListOrderCustomAttributes**](OrderCustomAttributesApi.md#ListOrderCustomAttributes) | **Get** /v2/orders/{order_id}/custom-attributes | ListOrderCustomAttributes
[**RetrieveOrderCustomAttribute**](OrderCustomAttributesApi.md#RetrieveOrderCustomAttribute) | **Get** /v2/orders/{order_id}/custom-attributes/{custom_attribute_key} | RetrieveOrderCustomAttribute
[**RetrieveOrderCustomAttributeDefinition**](OrderCustomAttributesApi.md#RetrieveOrderCustomAttributeDefinition) | **Get** /v2/orders/custom-attribute-definitions/{key} | RetrieveOrderCustomAttributeDefinition
[**UpdateOrderCustomAttributeDefinition**](OrderCustomAttributesApi.md#UpdateOrderCustomAttributeDefinition) | **Put** /v2/orders/custom-attribute-definitions/{key} | UpdateOrderCustomAttributeDefinition
[**UpsertOrderCustomAttribute**](OrderCustomAttributesApi.md#UpsertOrderCustomAttribute) | **Post** /v2/orders/{order_id}/custom-attributes/{custom_attribute_key} | UpsertOrderCustomAttribute

# **BulkDeleteOrderCustomAttributes**
> BulkDeleteOrderCustomAttributesResponse BulkDeleteOrderCustomAttributes(ctx, body)
BulkDeleteOrderCustomAttributes

Deletes order [custom attributes](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) as a bulk operation.  Use this endpoint to delete one or more custom attributes from one or more orders. A custom attribute is based on a custom attribute definition in a Square seller account.  (To create a custom attribute definition, use the [CreateOrderCustomAttributeDefinition](https://developer.squareup.com/reference/square_2024-01-18/order-custom-attributes-api/create-order-custom-attribute-definition) endpoint.)  This `BulkDeleteOrderCustomAttributes` endpoint accepts a map of 1 to 25 individual delete requests and returns a map of individual delete responses. Each delete request has a unique ID and provides an order ID and custom attribute. Each delete response is returned with the ID of the corresponding request.  To delete a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkDeleteOrderCustomAttributesRequest**](BulkDeleteOrderCustomAttributesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkDeleteOrderCustomAttributesResponse**](BulkDeleteOrderCustomAttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUpsertOrderCustomAttributes**
> BulkUpsertOrderCustomAttributesResponse BulkUpsertOrderCustomAttributes(ctx, body)
BulkUpsertOrderCustomAttributes

Creates or updates order [custom attributes](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) as a bulk operation.  Use this endpoint to delete one or more custom attributes from one or more orders. A custom attribute is based on a custom attribute definition in a Square seller account.  (To create a custom attribute definition, use the [CreateOrderCustomAttributeDefinition](https://developer.squareup.com/reference/square_2024-01-18/order-custom-attributes-api/create-order-custom-attribute-definition) endpoint.)  This `BulkUpsertOrderCustomAttributes` endpoint accepts a map of 1 to 25 individual upsert requests and returns a map of individual upsert responses. Each upsert request has a unique ID and provides an order ID and custom attribute. Each upsert response is returned with the ID of the corresponding request.  To create or update a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkUpsertOrderCustomAttributesRequest**](BulkUpsertOrderCustomAttributesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkUpsertOrderCustomAttributesResponse**](BulkUpsertOrderCustomAttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrderCustomAttributeDefinition**
> CreateOrderCustomAttributeDefinitionResponse CreateOrderCustomAttributeDefinition(ctx, body)
CreateOrderCustomAttributeDefinition

Creates an order-related custom attribute definition.  Use this endpoint to define a custom attribute that can be associated with orders.  After creating a custom attribute definition, you can set the custom attribute for orders in the Square seller account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOrderCustomAttributeDefinitionRequest**](CreateOrderCustomAttributeDefinitionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateOrderCustomAttributeDefinitionResponse**](CreateOrderCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrderCustomAttribute**
> DeleteOrderCustomAttributeResponse DeleteOrderCustomAttribute(ctx, orderId, customAttributeKey)
DeleteOrderCustomAttribute

Deletes a [custom attribute](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) associated with a customer profile.  To delete a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| The ID of the target [order](https://developer.squareup.com/reference/square_2024-01-18/objects/Order). | 
  **customAttributeKey** | **string**| The key of the custom attribute to delete.  This key must match the key of an existing custom attribute definition. | 

### Return type

[**DeleteOrderCustomAttributeResponse**](DeleteOrderCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrderCustomAttributeDefinition**
> DeleteOrderCustomAttributeDefinitionResponse DeleteOrderCustomAttributeDefinition(ctx, key)
DeleteOrderCustomAttributeDefinition

Deletes an order-related [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) from a Square seller account.  Only the definition owner can delete a custom attribute definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute definition to delete. | 

### Return type

[**DeleteOrderCustomAttributeDefinitionResponse**](DeleteOrderCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrderCustomAttributeDefinitions**
> ListOrderCustomAttributeDefinitionsResponse ListOrderCustomAttributeDefinitions(ctx, optional)
ListOrderCustomAttributeDefinitions

Lists the order-related [custom attribute definitions](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) that belong to a Square seller account.  When all response pages are retrieved, the results include all custom attribute definitions that are visible to the requesting application, including those that are created by other applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderCustomAttributesApiListOrderCustomAttributeDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderCustomAttributesApiListOrderCustomAttributeDefinitionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **visibilityFilter** | **optional.String**| Requests that all of the custom attributes be returned, or only those that are read-only or read-write. | 
 **cursor** | **optional.String**| The cursor returned in the paged response from the previous call to this endpoint.  Provide this cursor to retrieve the next page of results for your original request.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | 
 **limit** | **optional.Int32**| The maximum number of results to return in a single paged response. This limit is advisory.  The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.  The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | 

### Return type

[**ListOrderCustomAttributeDefinitionsResponse**](ListOrderCustomAttributeDefinitionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrderCustomAttributes**
> ListOrderCustomAttributesResponse ListOrderCustomAttributes(ctx, orderId, optional)
ListOrderCustomAttributes

Lists the [custom attributes](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) associated with an order.  You can use the `with_definitions` query parameter to also retrieve custom attribute definitions in the same call.  When all response pages are retrieved, the results include all custom attributes that are visible to the requesting application, including those that are owned by other applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| The ID of the target [order](https://developer.squareup.com/reference/square_2024-01-18/objects/Order). | 
 **optional** | ***OrderCustomAttributesApiListOrderCustomAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderCustomAttributesApiListOrderCustomAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **visibilityFilter** | **optional.String**| Requests that all of the custom attributes be returned, or only those that are read-only or read-write. | 
 **cursor** | **optional.String**| The cursor returned in the paged response from the previous call to this endpoint.  Provide this cursor to retrieve the next page of results for your original request.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | 
 **limit** | **optional.Int32**| The maximum number of results to return in a single paged response. This limit is advisory.  The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.  The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | 
 **withDefinitions** | **optional.Bool**| Indicates whether to return the [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) in the &#x60;definition&#x60; field of each custom attribute. Set this parameter to &#x60;true&#x60; to get the name and description of each custom attribute,  information about the data type, or other definition details. The default value is &#x60;false&#x60;. | 

### Return type

[**ListOrderCustomAttributesResponse**](ListOrderCustomAttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveOrderCustomAttribute**
> RetrieveOrderCustomAttributeResponse RetrieveOrderCustomAttribute(ctx, orderId, customAttributeKey, optional)
RetrieveOrderCustomAttribute

Retrieves a [custom attribute](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) associated with an order.  You can use the `with_definition` query parameter to also retrieve the custom attribute definition in the same call.  To retrieve a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| The ID of the target [order](https://developer.squareup.com/reference/square_2024-01-18/objects/Order). | 
  **customAttributeKey** | **string**| The key of the custom attribute to retrieve.  This key must match the key of an existing custom attribute definition. | 
 **optional** | ***OrderCustomAttributesApiRetrieveOrderCustomAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderCustomAttributesApiRetrieveOrderCustomAttributeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.Int32**| To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency) control, include this optional field and specify the current version of the custom attribute. | 
 **withDefinition** | **optional.Bool**| Indicates whether to return the [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) in the &#x60;definition&#x60; field of each  custom attribute. Set this parameter to &#x60;true&#x60; to get the name and description of each custom attribute,  information about the data type, or other definition details. The default value is &#x60;false&#x60;. | 

### Return type

[**RetrieveOrderCustomAttributeResponse**](RetrieveOrderCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveOrderCustomAttributeDefinition**
> RetrieveOrderCustomAttributeDefinitionResponse RetrieveOrderCustomAttributeDefinition(ctx, key, optional)
RetrieveOrderCustomAttributeDefinition

Retrieves an order-related [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) from a Square seller account.  To retrieve a custom attribute definition created by another application, the `visibility` setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute definition to retrieve. | 
 **optional** | ***OrderCustomAttributesApiRetrieveOrderCustomAttributeDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderCustomAttributesApiRetrieveOrderCustomAttributeDefinitionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.Int32**| To enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency) control, include this optional field and specify the current version of the custom attribute. | 

### Return type

[**RetrieveOrderCustomAttributeDefinitionResponse**](RetrieveOrderCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrderCustomAttributeDefinition**
> UpdateOrderCustomAttributeDefinitionResponse UpdateOrderCustomAttributeDefinition(ctx, body, key)
UpdateOrderCustomAttributeDefinition

Updates an order-related custom attribute definition for a Square seller account.  Only the definition owner can update a custom attribute definition. Note that sellers can view all custom attributes in exported customer data, including those set to `VISIBILITY_HIDDEN`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateOrderCustomAttributeDefinitionRequest**](UpdateOrderCustomAttributeDefinitionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **key** | **string**| The key of the custom attribute definition to update. | 

### Return type

[**UpdateOrderCustomAttributeDefinitionResponse**](UpdateOrderCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpsertOrderCustomAttribute**
> UpsertOrderCustomAttributeResponse UpsertOrderCustomAttribute(ctx, body, orderId, customAttributeKey)
UpsertOrderCustomAttribute

Creates or updates a [custom attribute](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) for an order.  Use this endpoint to set the value of a custom attribute for a specific order. A custom attribute is based on a custom attribute definition in a Square seller account. (To create a custom attribute definition, use the [CreateOrderCustomAttributeDefinition](https://developer.squareup.com/reference/square_2024-01-18/order-custom-attributes-api/create-order-custom-attribute-definition) endpoint.)  To create or update a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpsertOrderCustomAttributeRequest**](UpsertOrderCustomAttributeRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **orderId** | **string**| The ID of the target [order](https://developer.squareup.com/reference/square_2024-01-18/objects/Order). | 
  **customAttributeKey** | **string**| The key of the custom attribute to create or update.  This key must match the key  of an existing custom attribute definition. | 

### Return type

[**UpsertOrderCustomAttributeResponse**](UpsertOrderCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

