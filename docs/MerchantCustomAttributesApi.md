# {{classname}}

All URIs are relative to *https://connect.squareup.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteMerchantCustomAttributes**](MerchantCustomAttributesApi.md#BulkDeleteMerchantCustomAttributes) | **Post** /v2/merchants/custom-attributes/bulk-delete | BulkDeleteMerchantCustomAttributes
[**BulkUpsertMerchantCustomAttributes**](MerchantCustomAttributesApi.md#BulkUpsertMerchantCustomAttributes) | **Post** /v2/merchants/custom-attributes/bulk-upsert | BulkUpsertMerchantCustomAttributes
[**CreateMerchantCustomAttributeDefinition**](MerchantCustomAttributesApi.md#CreateMerchantCustomAttributeDefinition) | **Post** /v2/merchants/custom-attribute-definitions | CreateMerchantCustomAttributeDefinition
[**DeleteMerchantCustomAttribute**](MerchantCustomAttributesApi.md#DeleteMerchantCustomAttribute) | **Delete** /v2/merchants/{merchant_id}/custom-attributes/{key} | DeleteMerchantCustomAttribute
[**DeleteMerchantCustomAttributeDefinition**](MerchantCustomAttributesApi.md#DeleteMerchantCustomAttributeDefinition) | **Delete** /v2/merchants/custom-attribute-definitions/{key} | DeleteMerchantCustomAttributeDefinition
[**ListMerchantCustomAttributeDefinitions**](MerchantCustomAttributesApi.md#ListMerchantCustomAttributeDefinitions) | **Get** /v2/merchants/custom-attribute-definitions | ListMerchantCustomAttributeDefinitions
[**ListMerchantCustomAttributes**](MerchantCustomAttributesApi.md#ListMerchantCustomAttributes) | **Get** /v2/merchants/{merchant_id}/custom-attributes | ListMerchantCustomAttributes
[**RetrieveMerchantCustomAttribute**](MerchantCustomAttributesApi.md#RetrieveMerchantCustomAttribute) | **Get** /v2/merchants/{merchant_id}/custom-attributes/{key} | RetrieveMerchantCustomAttribute
[**RetrieveMerchantCustomAttributeDefinition**](MerchantCustomAttributesApi.md#RetrieveMerchantCustomAttributeDefinition) | **Get** /v2/merchants/custom-attribute-definitions/{key} | RetrieveMerchantCustomAttributeDefinition
[**UpdateMerchantCustomAttributeDefinition**](MerchantCustomAttributesApi.md#UpdateMerchantCustomAttributeDefinition) | **Put** /v2/merchants/custom-attribute-definitions/{key} | UpdateMerchantCustomAttributeDefinition
[**UpsertMerchantCustomAttribute**](MerchantCustomAttributesApi.md#UpsertMerchantCustomAttribute) | **Post** /v2/merchants/{merchant_id}/custom-attributes/{key} | UpsertMerchantCustomAttribute

# **BulkDeleteMerchantCustomAttributes**
> BulkDeleteMerchantCustomAttributesResponse BulkDeleteMerchantCustomAttributes(ctx, body)
BulkDeleteMerchantCustomAttributes

Deletes [custom attributes](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) for a merchant as a bulk operation. To delete a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkDeleteMerchantCustomAttributesRequest**](BulkDeleteMerchantCustomAttributesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkDeleteMerchantCustomAttributesResponse**](BulkDeleteMerchantCustomAttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUpsertMerchantCustomAttributes**
> BulkUpsertMerchantCustomAttributesResponse BulkUpsertMerchantCustomAttributes(ctx, body)
BulkUpsertMerchantCustomAttributes

Creates or updates [custom attributes](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) for a merchant as a bulk operation. Use this endpoint to set the value of one or more custom attributes for a merchant. A custom attribute is based on a custom attribute definition in a Square seller account, which is created using the [CreateMerchantCustomAttributeDefinition](https://developer.squareup.com/reference/square_2024-01-18/merchant-custom-attributes-api/create-merchant-custom-attribute-definition) endpoint. This `BulkUpsertMerchantCustomAttributes` endpoint accepts a map of 1 to 25 individual upsert requests and returns a map of individual upsert responses. Each upsert request has a unique ID and provides a merchant ID and custom attribute. Each upsert response is returned with the ID of the corresponding request. To create or update a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkUpsertMerchantCustomAttributesRequest**](BulkUpsertMerchantCustomAttributesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkUpsertMerchantCustomAttributesResponse**](BulkUpsertMerchantCustomAttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMerchantCustomAttributeDefinition**
> CreateMerchantCustomAttributeDefinitionResponse CreateMerchantCustomAttributeDefinition(ctx, body)
CreateMerchantCustomAttributeDefinition

Creates a merchant-related [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) for a Square seller account. Use this endpoint to define a custom attribute that can be associated with a merchant connecting to your application. A custom attribute definition specifies the `key`, `visibility`, `schema`, and other properties for a custom attribute. After the definition is created, you can call [UpsertMerchantCustomAttribute](https://developer.squareup.com/reference/square_2024-01-18/merchant-custom-attributes-api/upsert-merchant-custom-attribute) or [BulkUpsertMerchantCustomAttributes](https://developer.squareup.com/reference/square_2024-01-18/merchant-custom-attributes-api/bulk-upsert-merchant-custom-attributes) to set the custom attribute for a merchant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateMerchantCustomAttributeDefinitionRequest**](CreateMerchantCustomAttributeDefinitionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateMerchantCustomAttributeDefinitionResponse**](CreateMerchantCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMerchantCustomAttribute**
> DeleteMerchantCustomAttributeResponse DeleteMerchantCustomAttribute(ctx, merchantId, key)
DeleteMerchantCustomAttribute

Deletes a [custom attribute](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) associated with a merchant. To delete a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **merchantId** | **string**| The ID of the target [merchant](https://developer.squareup.com/reference/square_2024-01-18/objects/Merchant). | 
  **key** | **string**| The key of the custom attribute to delete. This key must match the &#x60;key&#x60; of a custom attribute definition in the Square seller account. If the requesting application is not the definition owner, you must use the qualified key. | 

### Return type

[**DeleteMerchantCustomAttributeResponse**](DeleteMerchantCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMerchantCustomAttributeDefinition**
> DeleteMerchantCustomAttributeDefinitionResponse DeleteMerchantCustomAttributeDefinition(ctx, key)
DeleteMerchantCustomAttributeDefinition

Deletes a merchant-related [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) from a Square seller account. Deleting a custom attribute definition also deletes the corresponding custom attribute from the merchant. Only the definition owner can delete a custom attribute definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute definition to delete. | 

### Return type

[**DeleteMerchantCustomAttributeDefinitionResponse**](DeleteMerchantCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMerchantCustomAttributeDefinitions**
> ListMerchantCustomAttributeDefinitionsResponse ListMerchantCustomAttributeDefinitions(ctx, optional)
ListMerchantCustomAttributeDefinitions

Lists the merchant-related [custom attribute definitions](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) that belong to a Square seller account. When all response pages are retrieved, the results include all custom attribute definitions that are visible to the requesting application, including those that are created by other applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MerchantCustomAttributesApiListMerchantCustomAttributeDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantCustomAttributesApiListMerchantCustomAttributeDefinitionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **visibilityFilter** | **optional.String**| Filters the &#x60;CustomAttributeDefinition&#x60; results by their &#x60;visibility&#x60; values. | 
 **limit** | **optional.Int32**| The maximum number of results to return in a single paged response. This limit is advisory. The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100. The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **cursor** | **optional.String**| The cursor returned in the paged response from the previous call to this endpoint. Provide this cursor to retrieve the next page of results for your original request. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 

### Return type

[**ListMerchantCustomAttributeDefinitionsResponse**](ListMerchantCustomAttributeDefinitionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMerchantCustomAttributes**
> ListMerchantCustomAttributesResponse ListMerchantCustomAttributes(ctx, merchantId, optional)
ListMerchantCustomAttributes

Lists the [custom attributes](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) associated with a merchant. You can use the `with_definitions` query parameter to also retrieve custom attribute definitions in the same call. When all response pages are retrieved, the results include all custom attributes that are visible to the requesting application, including those that are owned by other applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **merchantId** | **string**| The ID of the target [merchant](https://developer.squareup.com/reference/square_2024-01-18/objects/Merchant). | 
 **optional** | ***MerchantCustomAttributesApiListMerchantCustomAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantCustomAttributesApiListMerchantCustomAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **visibilityFilter** | **optional.String**| Filters the &#x60;CustomAttributeDefinition&#x60; results by their &#x60;visibility&#x60; values. | 
 **limit** | **optional.Int32**| The maximum number of results to return in a single paged response. This limit is advisory. The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100. The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **cursor** | **optional.String**| The cursor returned in the paged response from the previous call to this endpoint. Provide this cursor to retrieve the next page of results for your original request. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **withDefinitions** | **optional.Bool**| Indicates whether to return the [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) in the &#x60;definition&#x60; field of each custom attribute. Set this parameter to &#x60;true&#x60; to get the name and description of each custom attribute, information about the data type, or other definition details. The default value is &#x60;false&#x60;. | 

### Return type

[**ListMerchantCustomAttributesResponse**](ListMerchantCustomAttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveMerchantCustomAttribute**
> RetrieveMerchantCustomAttributeResponse RetrieveMerchantCustomAttribute(ctx, merchantId, key, optional)
RetrieveMerchantCustomAttribute

Retrieves a [custom attribute](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) associated with a merchant. You can use the `with_definition` query parameter to also retrieve the custom attribute definition in the same call. To retrieve a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **merchantId** | **string**| The ID of the target [merchant](https://developer.squareup.com/reference/square_2024-01-18/objects/Merchant). | 
  **key** | **string**| The key of the custom attribute to retrieve. This key must match the &#x60;key&#x60; of a custom attribute definition in the Square seller account. If the requesting application is not the definition owner, you must use the qualified key. | 
 **optional** | ***MerchantCustomAttributesApiRetrieveMerchantCustomAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantCustomAttributesApiRetrieveMerchantCustomAttributeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withDefinition** | **optional.Bool**| Indicates whether to return the [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) in the &#x60;definition&#x60; field of the custom attribute. Set this parameter to &#x60;true&#x60; to get the name and description of the custom attribute, information about the data type, or other definition details. The default value is &#x60;false&#x60;. | 
 **version** | **optional.Int32**| The current version of the custom attribute, which is used for strongly consistent reads to guarantee that you receive the most up-to-date data. When included in the request, Square returns the specified version or a higher version if one exists. If the specified version is higher than the current version, Square returns a &#x60;BAD_REQUEST&#x60; error. | 

### Return type

[**RetrieveMerchantCustomAttributeResponse**](RetrieveMerchantCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveMerchantCustomAttributeDefinition**
> RetrieveMerchantCustomAttributeDefinitionResponse RetrieveMerchantCustomAttributeDefinition(ctx, key, optional)
RetrieveMerchantCustomAttributeDefinition

Retrieves a merchant-related [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) from a Square seller account. To retrieve a custom attribute definition created by another application, the `visibility` setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute definition to retrieve. If the requesting application is not the definition owner, you must use the qualified key. | 
 **optional** | ***MerchantCustomAttributesApiRetrieveMerchantCustomAttributeDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantCustomAttributesApiRetrieveMerchantCustomAttributeDefinitionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.Int32**| The current version of the custom attribute definition, which is used for strongly consistent reads to guarantee that you receive the most up-to-date data. When included in the request, Square returns the specified version or a higher version if one exists. If the specified version is higher than the current version, Square returns a &#x60;BAD_REQUEST&#x60; error. | 

### Return type

[**RetrieveMerchantCustomAttributeDefinitionResponse**](RetrieveMerchantCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMerchantCustomAttributeDefinition**
> UpdateMerchantCustomAttributeDefinitionResponse UpdateMerchantCustomAttributeDefinition(ctx, body, key)
UpdateMerchantCustomAttributeDefinition

Updates a merchant-related [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) for a Square seller account. Use this endpoint to update the following fields: `name`, `description`, `visibility`, or the `schema` for a `Selection` data type. Only the definition owner can update a custom attribute definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateMerchantCustomAttributeDefinitionRequest**](UpdateMerchantCustomAttributeDefinitionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **key** | **string**| The key of the custom attribute definition to update. | 

### Return type

[**UpdateMerchantCustomAttributeDefinitionResponse**](UpdateMerchantCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpsertMerchantCustomAttribute**
> UpsertMerchantCustomAttributeResponse UpsertMerchantCustomAttribute(ctx, body, merchantId, key)
UpsertMerchantCustomAttribute

Creates or updates a [custom attribute](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) for a merchant. Use this endpoint to set the value of a custom attribute for a specified merchant. A custom attribute is based on a custom attribute definition in a Square seller account, which is created using the [CreateMerchantCustomAttributeDefinition](https://developer.squareup.com/reference/square_2024-01-18/merchant-custom-attributes-api/create-merchant-custom-attribute-definition) endpoint. To create or update a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpsertMerchantCustomAttributeRequest**](UpsertMerchantCustomAttributeRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **merchantId** | **string**| The ID of the target [merchant](https://developer.squareup.com/reference/square_2024-01-18/objects/Merchant). | 
  **key** | **string**| The key of the custom attribute to create or update. This key must match the &#x60;key&#x60; of a custom attribute definition in the Square seller account. If the requesting application is not the definition owner, you must use the qualified key. | 

### Return type

[**UpsertMerchantCustomAttributeResponse**](UpsertMerchantCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

