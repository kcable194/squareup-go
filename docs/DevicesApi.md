# {{classname}}

All URIs are relative to *https://connect.squareup.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceCode**](DevicesApi.md#CreateDeviceCode) | **Post** /v2/devices/codes | CreateDeviceCode
[**GetDevice**](DevicesApi.md#GetDevice) | **Get** /v2/devices/{device_id} | GetDevice
[**GetDeviceCode**](DevicesApi.md#GetDeviceCode) | **Get** /v2/devices/codes/{id} | GetDeviceCode
[**ListDeviceCodes**](DevicesApi.md#ListDeviceCodes) | **Get** /v2/devices/codes | ListDeviceCodes
[**ListDevices**](DevicesApi.md#ListDevices) | **Get** /v2/devices | ListDevices

# **CreateDeviceCode**
> CreateDeviceCodeResponse CreateDeviceCode(ctx, body)
CreateDeviceCode

Creates a DeviceCode that can be used to login to a Square Terminal device to enter the connected terminal mode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateDeviceCodeRequest**](CreateDeviceCodeRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateDeviceCodeResponse**](CreateDeviceCodeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevice**
> GetDeviceResponse GetDevice(ctx, deviceId)
GetDevice

Retrieves Device with the associated `device_id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **string**| The unique ID for the desired &#x60;Device&#x60;. | 

### Return type

[**GetDeviceResponse**](GetDeviceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeviceCode**
> GetDeviceCodeResponse GetDeviceCode(ctx, id)
GetDeviceCode

Retrieves DeviceCode with the associated ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The unique identifier for the device code. | 

### Return type

[**GetDeviceCodeResponse**](GetDeviceCodeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeviceCodes**
> ListDeviceCodesResponse ListDeviceCodes(ctx, optional)
ListDeviceCodes

Lists all DeviceCodes associated with the merchant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesApiListDeviceCodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiListDeviceCodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query.  See [Paginating results](https://developer.squareup.com/docs/working-with-apis/pagination) for more information. | 
 **locationId** | **optional.String**| If specified, only returns DeviceCodes of the specified location. Returns DeviceCodes of all locations if empty. | 
 **productType** | **optional.String**| If specified, only returns DeviceCodes targeting the specified product type. Returns DeviceCodes of all product types if empty. | 
 **status** | **optional.String**| If specified, returns DeviceCodes with the specified statuses. Returns DeviceCodes of status &#x60;PAIRED&#x60; and &#x60;UNPAIRED&#x60; if empty. | 

### Return type

[**ListDeviceCodesResponse**](ListDeviceCodesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDevices**
> ListDevicesResponse ListDevices(ctx, optional)
ListDevices

List devices associated with the merchant. Currently, only Terminal API devices are supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevicesApiListDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevicesApiListDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. See [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination) for more information. | 
 **sortOrder** | **optional.String**| The order in which results are listed. - &#x60;ASC&#x60; - Oldest to newest. - &#x60;DESC&#x60; - Newest to oldest (default). | 
 **limit** | **optional.Int32**| The number of results to return in a single page. | 
 **locationId** | **optional.String**| If present, only returns devices at the target location. | 

### Return type

[**ListDevicesResponse**](ListDevicesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

