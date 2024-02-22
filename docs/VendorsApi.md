# {{classname}}

All URIs are relative to *https://connect.squareup.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateVendors**](VendorsApi.md#BulkCreateVendors) | **Post** /v2/vendors/bulk-create | BulkCreateVendors
[**BulkRetrieveVendors**](VendorsApi.md#BulkRetrieveVendors) | **Post** /v2/vendors/bulk-retrieve | BulkRetrieveVendors
[**BulkUpdateVendors**](VendorsApi.md#BulkUpdateVendors) | **Put** /v2/vendors/bulk-update | BulkUpdateVendors
[**CreateVendor**](VendorsApi.md#CreateVendor) | **Post** /v2/vendors/create | CreateVendor
[**RetrieveVendor**](VendorsApi.md#RetrieveVendor) | **Get** /v2/vendors/{vendor_id} | RetrieveVendor
[**SearchVendors**](VendorsApi.md#SearchVendors) | **Post** /v2/vendors/search | SearchVendors
[**UpdateVendor**](VendorsApi.md#UpdateVendor) | **Put** /v2/vendors/{vendor_id} | UpdateVendor

# **BulkCreateVendors**
> BulkCreateVendorsResponse BulkCreateVendors(ctx, body)
BulkCreateVendors

Creates one or more [Vendor](https://developer.squareup.com/reference/square_2024-01-18/objects/Vendor) objects to represent suppliers to a seller.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkCreateVendorsRequest**](BulkCreateVendorsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkCreateVendorsResponse**](BulkCreateVendorsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkRetrieveVendors**
> BulkRetrieveVendorsResponse BulkRetrieveVendors(ctx, body)
BulkRetrieveVendors

Retrieves one or more vendors of specified [Vendor](https://developer.squareup.com/reference/square_2024-01-18/objects/Vendor) IDs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkRetrieveVendorsRequest**](BulkRetrieveVendorsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkRetrieveVendorsResponse**](BulkRetrieveVendorsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUpdateVendors**
> BulkUpdateVendorsResponse BulkUpdateVendors(ctx, body)
BulkUpdateVendors

Updates one or more of existing [Vendor](https://developer.squareup.com/reference/square_2024-01-18/objects/Vendor) objects as suppliers to a seller.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkUpdateVendorsRequest**](BulkUpdateVendorsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkUpdateVendorsResponse**](BulkUpdateVendorsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVendor**
> CreateVendorResponse CreateVendor(ctx, body)
CreateVendor

Creates a single [Vendor](https://developer.squareup.com/reference/square_2024-01-18/objects/Vendor) object to represent a supplier to a seller.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVendorRequest**](CreateVendorRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateVendorResponse**](CreateVendorResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveVendor**
> RetrieveVendorResponse RetrieveVendor(ctx, vendorId)
RetrieveVendor

Retrieves the vendor of a specified [Vendor](https://developer.squareup.com/reference/square_2024-01-18/objects/Vendor) ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorId** | **string**| ID of the [Vendor](https://developer.squareup.com/reference/square_2024-01-18/objects/Vendor) to retrieve. | 

### Return type

[**RetrieveVendorResponse**](RetrieveVendorResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchVendors**
> SearchVendorsResponse SearchVendors(ctx, body)
SearchVendors

Searches for vendors using a filter against supported [Vendor](https://developer.squareup.com/reference/square_2024-01-18/objects/Vendor) properties and a supported sorter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchVendorsRequest**](SearchVendorsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchVendorsResponse**](SearchVendorsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVendor**
> UpdateVendorResponse UpdateVendor(ctx, body, vendorId)
UpdateVendor

Updates an existing [Vendor](https://developer.squareup.com/reference/square_2024-01-18/objects/Vendor) object as a supplier to a seller.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVendorRequest**](UpdateVendorRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **vendorId** | **string**|  | 

### Return type

[**UpdateVendorResponse**](UpdateVendorResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

