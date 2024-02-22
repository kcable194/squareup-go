# {{classname}}

All URIs are relative to *https://connect.squareup.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteLocationCustomAttributes**](LocationCustomAttributesApi.md#BulkDeleteLocationCustomAttributes) | **Post** /v2/locations/custom-attributes/bulk-delete | BulkDeleteLocationCustomAttributes
[**BulkUpsertLocationCustomAttributes**](LocationCustomAttributesApi.md#BulkUpsertLocationCustomAttributes) | **Post** /v2/locations/custom-attributes/bulk-upsert | BulkUpsertLocationCustomAttributes
[**CreateLocationCustomAttributeDefinition**](LocationCustomAttributesApi.md#CreateLocationCustomAttributeDefinition) | **Post** /v2/locations/custom-attribute-definitions | CreateLocationCustomAttributeDefinition
[**DeleteLocationCustomAttribute**](LocationCustomAttributesApi.md#DeleteLocationCustomAttribute) | **Delete** /v2/locations/{location_id}/custom-attributes/{key} | DeleteLocationCustomAttribute
[**DeleteLocationCustomAttributeDefinition**](LocationCustomAttributesApi.md#DeleteLocationCustomAttributeDefinition) | **Delete** /v2/locations/custom-attribute-definitions/{key} | DeleteLocationCustomAttributeDefinition
[**ListLocationCustomAttributeDefinitions**](LocationCustomAttributesApi.md#ListLocationCustomAttributeDefinitions) | **Get** /v2/locations/custom-attribute-definitions | ListLocationCustomAttributeDefinitions
[**ListLocationCustomAttributes**](LocationCustomAttributesApi.md#ListLocationCustomAttributes) | **Get** /v2/locations/{location_id}/custom-attributes | ListLocationCustomAttributes
[**RetrieveLocationCustomAttribute**](LocationCustomAttributesApi.md#RetrieveLocationCustomAttribute) | **Get** /v2/locations/{location_id}/custom-attributes/{key} | RetrieveLocationCustomAttribute
[**RetrieveLocationCustomAttributeDefinition**](LocationCustomAttributesApi.md#RetrieveLocationCustomAttributeDefinition) | **Get** /v2/locations/custom-attribute-definitions/{key} | RetrieveLocationCustomAttributeDefinition
[**UpdateLocationCustomAttributeDefinition**](LocationCustomAttributesApi.md#UpdateLocationCustomAttributeDefinition) | **Put** /v2/locations/custom-attribute-definitions/{key} | UpdateLocationCustomAttributeDefinition
[**UpsertLocationCustomAttribute**](LocationCustomAttributesApi.md#UpsertLocationCustomAttribute) | **Post** /v2/locations/{location_id}/custom-attributes/{key} | UpsertLocationCustomAttribute

# **BulkDeleteLocationCustomAttributes**
> BulkDeleteLocationCustomAttributesResponse BulkDeleteLocationCustomAttributes(ctx, body)
BulkDeleteLocationCustomAttributes

Deletes [custom attributes](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) for locations as a bulk operation. To delete a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkDeleteLocationCustomAttributesRequest**](BulkDeleteLocationCustomAttributesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkDeleteLocationCustomAttributesResponse**](BulkDeleteLocationCustomAttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUpsertLocationCustomAttributes**
> BulkUpsertLocationCustomAttributesResponse BulkUpsertLocationCustomAttributes(ctx, body)
BulkUpsertLocationCustomAttributes

Creates or updates [custom attributes](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) for locations as a bulk operation. Use this endpoint to set the value of one or more custom attributes for one or more locations. A custom attribute is based on a custom attribute definition in a Square seller account, which is created using the [CreateLocationCustomAttributeDefinition](https://developer.squareup.com/reference/square_2024-01-18/location-custom-attributes-api/create-location-custom-attribute-definition) endpoint. This `BulkUpsertLocationCustomAttributes` endpoint accepts a map of 1 to 25 individual upsert requests and returns a map of individual upsert responses. Each upsert request has a unique ID and provides a location ID and custom attribute. Each upsert response is returned with the ID of the corresponding request. To create or update a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkUpsertLocationCustomAttributesRequest**](BulkUpsertLocationCustomAttributesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkUpsertLocationCustomAttributesResponse**](BulkUpsertLocationCustomAttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLocationCustomAttributeDefinition**
> CreateLocationCustomAttributeDefinitionResponse CreateLocationCustomAttributeDefinition(ctx, body)
CreateLocationCustomAttributeDefinition

Creates a location-related [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) for a Square seller account. Use this endpoint to define a custom attribute that can be associated with locations. A custom attribute definition specifies the `key`, `visibility`, `schema`, and other properties for a custom attribute. After the definition is created, you can call [UpsertLocationCustomAttribute](https://developer.squareup.com/reference/square_2024-01-18/location-custom-attributes-api/upsert-location-custom-attribute) or [BulkUpsertLocationCustomAttributes](https://developer.squareup.com/reference/square_2024-01-18/location-custom-attributes-api/bulk-upsert-location-custom-attributes) to set the custom attribute for locations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateLocationCustomAttributeDefinitionRequest**](CreateLocationCustomAttributeDefinitionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateLocationCustomAttributeDefinitionResponse**](CreateLocationCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLocationCustomAttribute**
> DeleteLocationCustomAttributeResponse DeleteLocationCustomAttribute(ctx, locationId, key)
DeleteLocationCustomAttribute

Deletes a [custom attribute](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) associated with a location. To delete a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the target [location](https://developer.squareup.com/reference/square_2024-01-18/objects/Location). | 
  **key** | **string**| The key of the custom attribute to delete. This key must match the &#x60;key&#x60; of a custom attribute definition in the Square seller account. If the requesting application is not the definition owner, you must use the qualified key. | 

### Return type

[**DeleteLocationCustomAttributeResponse**](DeleteLocationCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLocationCustomAttributeDefinition**
> DeleteLocationCustomAttributeDefinitionResponse DeleteLocationCustomAttributeDefinition(ctx, key)
DeleteLocationCustomAttributeDefinition

Deletes a location-related [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) from a Square seller account. Deleting a custom attribute definition also deletes the corresponding custom attribute from all locations. Only the definition owner can delete a custom attribute definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute definition to delete. | 

### Return type

[**DeleteLocationCustomAttributeDefinitionResponse**](DeleteLocationCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLocationCustomAttributeDefinitions**
> ListLocationCustomAttributeDefinitionsResponse ListLocationCustomAttributeDefinitions(ctx, optional)
ListLocationCustomAttributeDefinitions

Lists the location-related [custom attribute definitions](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) that belong to a Square seller account. When all response pages are retrieved, the results include all custom attribute definitions that are visible to the requesting application, including those that are created by other applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocationCustomAttributesApiListLocationCustomAttributeDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationCustomAttributesApiListLocationCustomAttributeDefinitionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **visibilityFilter** | **optional.String**| Filters the &#x60;CustomAttributeDefinition&#x60; results by their &#x60;visibility&#x60; values. | 
 **limit** | **optional.Int32**| The maximum number of results to return in a single paged response. This limit is advisory. The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100. The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **cursor** | **optional.String**| The cursor returned in the paged response from the previous call to this endpoint. Provide this cursor to retrieve the next page of results for your original request. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 

### Return type

[**ListLocationCustomAttributeDefinitionsResponse**](ListLocationCustomAttributeDefinitionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLocationCustomAttributes**
> ListLocationCustomAttributesResponse ListLocationCustomAttributes(ctx, locationId, optional)
ListLocationCustomAttributes

Lists the [custom attributes](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) associated with a location. You can use the `with_definitions` query parameter to also retrieve custom attribute definitions in the same call. When all response pages are retrieved, the results include all custom attributes that are visible to the requesting application, including those that are owned by other applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the target [location](https://developer.squareup.com/reference/square_2024-01-18/objects/Location). | 
 **optional** | ***LocationCustomAttributesApiListLocationCustomAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationCustomAttributesApiListLocationCustomAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **visibilityFilter** | **optional.String**| Filters the &#x60;CustomAttributeDefinition&#x60; results by their &#x60;visibility&#x60; values. | 
 **limit** | **optional.Int32**| The maximum number of results to return in a single paged response. This limit is advisory. The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100. The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **cursor** | **optional.String**| The cursor returned in the paged response from the previous call to this endpoint. Provide this cursor to retrieve the next page of results for your original request. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **withDefinitions** | **optional.Bool**| Indicates whether to return the [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) in the &#x60;definition&#x60; field of each custom attribute. Set this parameter to &#x60;true&#x60; to get the name and description of each custom attribute, information about the data type, or other definition details. The default value is &#x60;false&#x60;. | 

### Return type

[**ListLocationCustomAttributesResponse**](ListLocationCustomAttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveLocationCustomAttribute**
> RetrieveLocationCustomAttributeResponse RetrieveLocationCustomAttribute(ctx, locationId, key, optional)
RetrieveLocationCustomAttribute

Retrieves a [custom attribute](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) associated with a location. You can use the `with_definition` query parameter to also retrieve the custom attribute definition in the same call. To retrieve a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the target [location](https://developer.squareup.com/reference/square_2024-01-18/objects/Location). | 
  **key** | **string**| The key of the custom attribute to retrieve. This key must match the &#x60;key&#x60; of a custom attribute definition in the Square seller account. If the requesting application is not the definition owner, you must use the qualified key. | 
 **optional** | ***LocationCustomAttributesApiRetrieveLocationCustomAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationCustomAttributesApiRetrieveLocationCustomAttributeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withDefinition** | **optional.Bool**| Indicates whether to return the [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) in the &#x60;definition&#x60; field of the custom attribute. Set this parameter to &#x60;true&#x60; to get the name and description of the custom attribute, information about the data type, or other definition details. The default value is &#x60;false&#x60;. | 
 **version** | **optional.Int32**| The current version of the custom attribute, which is used for strongly consistent reads to guarantee that you receive the most up-to-date data. When included in the request, Square returns the specified version or a higher version if one exists. If the specified version is higher than the current version, Square returns a &#x60;BAD_REQUEST&#x60; error. | 

### Return type

[**RetrieveLocationCustomAttributeResponse**](RetrieveLocationCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveLocationCustomAttributeDefinition**
> RetrieveLocationCustomAttributeDefinitionResponse RetrieveLocationCustomAttributeDefinition(ctx, key, optional)
RetrieveLocationCustomAttributeDefinition

Retrieves a location-related [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) from a Square seller account. To retrieve a custom attribute definition created by another application, the `visibility` setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute definition to retrieve. If the requesting application is not the definition owner, you must use the qualified key. | 
 **optional** | ***LocationCustomAttributesApiRetrieveLocationCustomAttributeDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationCustomAttributesApiRetrieveLocationCustomAttributeDefinitionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.Int32**| The current version of the custom attribute definition, which is used for strongly consistent reads to guarantee that you receive the most up-to-date data. When included in the request, Square returns the specified version or a higher version if one exists. If the specified version is higher than the current version, Square returns a &#x60;BAD_REQUEST&#x60; error. | 

### Return type

[**RetrieveLocationCustomAttributeDefinitionResponse**](RetrieveLocationCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLocationCustomAttributeDefinition**
> UpdateLocationCustomAttributeDefinitionResponse UpdateLocationCustomAttributeDefinition(ctx, body, key)
UpdateLocationCustomAttributeDefinition

Updates a location-related [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) for a Square seller account. Use this endpoint to update the following fields: `name`, `description`, `visibility`, or the `schema` for a `Selection` data type. Only the definition owner can update a custom attribute definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateLocationCustomAttributeDefinitionRequest**](UpdateLocationCustomAttributeDefinitionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **key** | **string**| The key of the custom attribute definition to update. | 

### Return type

[**UpdateLocationCustomAttributeDefinitionResponse**](UpdateLocationCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpsertLocationCustomAttribute**
> UpsertLocationCustomAttributeResponse UpsertLocationCustomAttribute(ctx, body, locationId, key)
UpsertLocationCustomAttribute

Creates or updates a [custom attribute](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttribute) for a location. Use this endpoint to set the value of a custom attribute for a specified location. A custom attribute is based on a custom attribute definition in a Square seller account, which is created using the [CreateLocationCustomAttributeDefinition](https://developer.squareup.com/reference/square_2024-01-18/location-custom-attributes-api/create-location-custom-attribute-definition) endpoint. To create or update a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpsertLocationCustomAttributeRequest**](UpsertLocationCustomAttributeRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the target [location](https://developer.squareup.com/reference/square_2024-01-18/objects/Location). | 
  **key** | **string**| The key of the custom attribute to create or update. This key must match the &#x60;key&#x60; of a custom attribute definition in the Square seller account. If the requesting application is not the definition owner, you must use the qualified key. | 

### Return type

[**UpsertLocationCustomAttributeResponse**](UpsertLocationCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

