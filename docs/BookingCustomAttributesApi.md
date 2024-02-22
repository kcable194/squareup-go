# {{classname}}

All URIs are relative to *https://connect.squareup.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteBookingCustomAttributes**](BookingCustomAttributesApi.md#BulkDeleteBookingCustomAttributes) | **Post** /v2/bookings/custom-attributes/bulk-delete | BulkDeleteBookingCustomAttributes
[**BulkUpsertBookingCustomAttributes**](BookingCustomAttributesApi.md#BulkUpsertBookingCustomAttributes) | **Post** /v2/bookings/custom-attributes/bulk-upsert | BulkUpsertBookingCustomAttributes
[**CreateBookingCustomAttributeDefinition**](BookingCustomAttributesApi.md#CreateBookingCustomAttributeDefinition) | **Post** /v2/bookings/custom-attribute-definitions | CreateBookingCustomAttributeDefinition
[**DeleteBookingCustomAttribute**](BookingCustomAttributesApi.md#DeleteBookingCustomAttribute) | **Delete** /v2/bookings/{booking_id}/custom-attributes/{key} | DeleteBookingCustomAttribute
[**DeleteBookingCustomAttributeDefinition**](BookingCustomAttributesApi.md#DeleteBookingCustomAttributeDefinition) | **Delete** /v2/bookings/custom-attribute-definitions/{key} | DeleteBookingCustomAttributeDefinition
[**ListBookingCustomAttributeDefinitions**](BookingCustomAttributesApi.md#ListBookingCustomAttributeDefinitions) | **Get** /v2/bookings/custom-attribute-definitions | ListBookingCustomAttributeDefinitions
[**ListBookingCustomAttributes**](BookingCustomAttributesApi.md#ListBookingCustomAttributes) | **Get** /v2/bookings/{booking_id}/custom-attributes | ListBookingCustomAttributes
[**RetrieveBookingCustomAttribute**](BookingCustomAttributesApi.md#RetrieveBookingCustomAttribute) | **Get** /v2/bookings/{booking_id}/custom-attributes/{key} | RetrieveBookingCustomAttribute
[**RetrieveBookingCustomAttributeDefinition**](BookingCustomAttributesApi.md#RetrieveBookingCustomAttributeDefinition) | **Get** /v2/bookings/custom-attribute-definitions/{key} | RetrieveBookingCustomAttributeDefinition
[**UpdateBookingCustomAttributeDefinition**](BookingCustomAttributesApi.md#UpdateBookingCustomAttributeDefinition) | **Put** /v2/bookings/custom-attribute-definitions/{key} | UpdateBookingCustomAttributeDefinition
[**UpsertBookingCustomAttribute**](BookingCustomAttributesApi.md#UpsertBookingCustomAttribute) | **Put** /v2/bookings/{booking_id}/custom-attributes/{key} | UpsertBookingCustomAttribute

# **BulkDeleteBookingCustomAttributes**
> BulkDeleteBookingCustomAttributesResponse BulkDeleteBookingCustomAttributes(ctx, body)
BulkDeleteBookingCustomAttributes

Bulk deletes bookings custom attributes.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.  For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus* or *Appointments Premium*.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkDeleteBookingCustomAttributesRequest**](BulkDeleteBookingCustomAttributesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkDeleteBookingCustomAttributesResponse**](BulkDeleteBookingCustomAttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUpsertBookingCustomAttributes**
> BulkUpsertBookingCustomAttributesResponse BulkUpsertBookingCustomAttributes(ctx, body)
BulkUpsertBookingCustomAttributes

Bulk upserts bookings custom attributes.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.  For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus* or *Appointments Premium*.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkUpsertBookingCustomAttributesRequest**](BulkUpsertBookingCustomAttributesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkUpsertBookingCustomAttributesResponse**](BulkUpsertBookingCustomAttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBookingCustomAttributeDefinition**
> CreateBookingCustomAttributeDefinitionResponse CreateBookingCustomAttributeDefinition(ctx, body)
CreateBookingCustomAttributeDefinition

Creates a bookings custom attribute definition.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.  For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus* or *Appointments Premium*.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateBookingCustomAttributeDefinitionRequest**](CreateBookingCustomAttributeDefinitionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateBookingCustomAttributeDefinitionResponse**](CreateBookingCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBookingCustomAttribute**
> DeleteBookingCustomAttributeResponse DeleteBookingCustomAttribute(ctx, bookingId, key)
DeleteBookingCustomAttribute

Deletes a bookings custom attribute.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.  For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus* or *Appointments Premium*.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bookingId** | **string**| The ID of the target [booking](https://developer.squareup.com/reference/square_2024-01-18/objects/Booking). | 
  **key** | **string**| The key of the custom attribute to delete. This key must match the &#x60;key&#x60; of a custom attribute definition in the Square seller account. If the requesting application is not the definition owner, you must use the qualified key. | 

### Return type

[**DeleteBookingCustomAttributeResponse**](DeleteBookingCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBookingCustomAttributeDefinition**
> DeleteBookingCustomAttributeDefinitionResponse DeleteBookingCustomAttributeDefinition(ctx, key)
DeleteBookingCustomAttributeDefinition

Deletes a bookings custom attribute definition.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.  For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus* or *Appointments Premium*.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute definition to delete. | 

### Return type

[**DeleteBookingCustomAttributeDefinitionResponse**](DeleteBookingCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBookingCustomAttributeDefinitions**
> ListBookingCustomAttributeDefinitionsResponse ListBookingCustomAttributeDefinitions(ctx, optional)
ListBookingCustomAttributeDefinitions

Get all bookings custom attribute definitions.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookingCustomAttributesApiListBookingCustomAttributeDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookingCustomAttributesApiListBookingCustomAttributeDefinitionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The maximum number of results to return in a single paged response. This limit is advisory. The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100. The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **cursor** | **optional.String**| The cursor returned in the paged response from the previous call to this endpoint. Provide this cursor to retrieve the next page of results for your original request. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 

### Return type

[**ListBookingCustomAttributeDefinitionsResponse**](ListBookingCustomAttributeDefinitionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBookingCustomAttributes**
> ListBookingCustomAttributesResponse ListBookingCustomAttributes(ctx, bookingId, optional)
ListBookingCustomAttributes

Lists a booking's custom attributes.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bookingId** | **string**| The ID of the target [booking](https://developer.squareup.com/reference/square_2024-01-18/objects/Booking). | 
 **optional** | ***BookingCustomAttributesApiListBookingCustomAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookingCustomAttributesApiListBookingCustomAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The maximum number of results to return in a single paged response. This limit is advisory. The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100. The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **cursor** | **optional.String**| The cursor returned in the paged response from the previous call to this endpoint. Provide this cursor to retrieve the next page of results for your original request. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **withDefinitions** | **optional.Bool**| Indicates whether to return the [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) in the &#x60;definition&#x60; field of each custom attribute. Set this parameter to &#x60;true&#x60; to get the name and description of each custom attribute, information about the data type, or other definition details. The default value is &#x60;false&#x60;. | 

### Return type

[**ListBookingCustomAttributesResponse**](ListBookingCustomAttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveBookingCustomAttribute**
> RetrieveBookingCustomAttributeResponse RetrieveBookingCustomAttribute(ctx, bookingId, key, optional)
RetrieveBookingCustomAttribute

Retrieves a bookings custom attribute.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bookingId** | **string**| The ID of the target [booking](https://developer.squareup.com/reference/square_2024-01-18/objects/Booking). | 
  **key** | **string**| The key of the custom attribute to retrieve. This key must match the &#x60;key&#x60; of a custom attribute definition in the Square seller account. If the requesting application is not the definition owner, you must use the qualified key. | 
 **optional** | ***BookingCustomAttributesApiRetrieveBookingCustomAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookingCustomAttributesApiRetrieveBookingCustomAttributeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withDefinition** | **optional.Bool**| Indicates whether to return the [custom attribute definition](https://developer.squareup.com/reference/square_2024-01-18/objects/CustomAttributeDefinition) in the &#x60;definition&#x60; field of the custom attribute. Set this parameter to &#x60;true&#x60; to get the name and description of the custom attribute, information about the data type, or other definition details. The default value is &#x60;false&#x60;. | 
 **version** | **optional.Int32**| The current version of the custom attribute, which is used for strongly consistent reads to guarantee that you receive the most up-to-date data. When included in the request, Square returns the specified version or a higher version if one exists. If the specified version is higher than the current version, Square returns a &#x60;BAD_REQUEST&#x60; error. | 

### Return type

[**RetrieveBookingCustomAttributeResponse**](RetrieveBookingCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveBookingCustomAttributeDefinition**
> RetrieveBookingCustomAttributeDefinitionResponse RetrieveBookingCustomAttributeDefinition(ctx, key, optional)
RetrieveBookingCustomAttributeDefinition

Retrieves a bookings custom attribute definition.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute definition to retrieve. If the requesting application is not the definition owner, you must use the qualified key. | 
 **optional** | ***BookingCustomAttributesApiRetrieveBookingCustomAttributeDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookingCustomAttributesApiRetrieveBookingCustomAttributeDefinitionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.Int32**| The current version of the custom attribute definition, which is used for strongly consistent reads to guarantee that you receive the most up-to-date data. When included in the request, Square returns the specified version or a higher version if one exists. If the specified version is higher than the current version, Square returns a &#x60;BAD_REQUEST&#x60; error. | 

### Return type

[**RetrieveBookingCustomAttributeDefinitionResponse**](RetrieveBookingCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBookingCustomAttributeDefinition**
> UpdateBookingCustomAttributeDefinitionResponse UpdateBookingCustomAttributeDefinition(ctx, body, key)
UpdateBookingCustomAttributeDefinition

Updates a bookings custom attribute definition.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.  For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus* or *Appointments Premium*.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateBookingCustomAttributeDefinitionRequest**](UpdateBookingCustomAttributeDefinitionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **key** | **string**| The key of the custom attribute definition to update. | 

### Return type

[**UpdateBookingCustomAttributeDefinitionResponse**](UpdateBookingCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpsertBookingCustomAttribute**
> UpsertBookingCustomAttributeResponse UpsertBookingCustomAttribute(ctx, body, bookingId, key)
UpsertBookingCustomAttribute

Upserts a bookings custom attribute.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.  For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus* or *Appointments Premium*.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpsertBookingCustomAttributeRequest**](UpsertBookingCustomAttributeRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **bookingId** | **string**| The ID of the target [booking](https://developer.squareup.com/reference/square_2024-01-18/objects/Booking). | 
  **key** | **string**| The key of the custom attribute to create or update. This key must match the &#x60;key&#x60; of a custom attribute definition in the Square seller account. If the requesting application is not the definition owner, you must use the qualified key. | 

### Return type

[**UpsertBookingCustomAttributeResponse**](UpsertBookingCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

