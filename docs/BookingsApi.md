# {{classname}}

All URIs are relative to *https://connect.squareup.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkRetrieveBookings**](BookingsApi.md#BulkRetrieveBookings) | **Post** /v2/bookings/bulk-retrieve | BulkRetrieveBookings
[**BulkRetrieveTeamMemberBookingProfiles**](BookingsApi.md#BulkRetrieveTeamMemberBookingProfiles) | **Post** /v2/bookings/team-member-booking-profiles/bulk-retrieve | BulkRetrieveTeamMemberBookingProfiles
[**CancelBooking**](BookingsApi.md#CancelBooking) | **Post** /v2/bookings/{booking_id}/cancel | CancelBooking
[**CreateBooking**](BookingsApi.md#CreateBooking) | **Post** /v2/bookings | CreateBooking
[**ListBookings**](BookingsApi.md#ListBookings) | **Get** /v2/bookings | ListBookings
[**ListLocationBookingProfiles**](BookingsApi.md#ListLocationBookingProfiles) | **Get** /v2/bookings/location-booking-profiles | ListLocationBookingProfiles
[**ListTeamMemberBookingProfiles**](BookingsApi.md#ListTeamMemberBookingProfiles) | **Get** /v2/bookings/team-member-booking-profiles | ListTeamMemberBookingProfiles
[**RetrieveBooking**](BookingsApi.md#RetrieveBooking) | **Get** /v2/bookings/{booking_id} | RetrieveBooking
[**RetrieveBusinessBookingProfile**](BookingsApi.md#RetrieveBusinessBookingProfile) | **Get** /v2/bookings/business-booking-profile | RetrieveBusinessBookingProfile
[**RetrieveLocationBookingProfile**](BookingsApi.md#RetrieveLocationBookingProfile) | **Get** /v2/bookings/location-booking-profiles/{location_id} | RetrieveLocationBookingProfile
[**RetrieveTeamMemberBookingProfile**](BookingsApi.md#RetrieveTeamMemberBookingProfile) | **Get** /v2/bookings/team-member-booking-profiles/{team_member_id} | RetrieveTeamMemberBookingProfile
[**SearchAvailability**](BookingsApi.md#SearchAvailability) | **Post** /v2/bookings/availability/search | SearchAvailability
[**UpdateBooking**](BookingsApi.md#UpdateBooking) | **Put** /v2/bookings/{booking_id} | UpdateBooking

# **BulkRetrieveBookings**
> BulkRetrieveBookingsResponse BulkRetrieveBookings(ctx, body)
BulkRetrieveBookings

Bulk-Retrieves a list of bookings by booking IDs.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkRetrieveBookingsRequest**](BulkRetrieveBookingsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkRetrieveBookingsResponse**](BulkRetrieveBookingsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkRetrieveTeamMemberBookingProfiles**
> BulkRetrieveTeamMemberBookingProfilesResponse BulkRetrieveTeamMemberBookingProfiles(ctx, body)
BulkRetrieveTeamMemberBookingProfiles

Retrieves one or more team members' booking profiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkRetrieveTeamMemberBookingProfilesRequest**](BulkRetrieveTeamMemberBookingProfilesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkRetrieveTeamMemberBookingProfilesResponse**](BulkRetrieveTeamMemberBookingProfilesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelBooking**
> CancelBookingResponse CancelBooking(ctx, body, bookingId)
CancelBooking

Cancels an existing booking.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.  For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus* or *Appointments Premium*.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CancelBookingRequest**](CancelBookingRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **bookingId** | **string**| The ID of the [Booking](https://developer.squareup.com/reference/square_2024-01-18/objects/Booking) object representing the to-be-cancelled booking. | 

### Return type

[**CancelBookingResponse**](CancelBookingResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBooking**
> CreateBookingResponse CreateBooking(ctx, body)
CreateBooking

Creates a booking.  The required input must include the following: - `Booking.location_id` - `Booking.start_at` - `Booking.team_member_id` - `Booking.AppointmentSegment.service_variation_id` - `Booking.AppointmentSegment.service_variation_version`  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.  For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus* or *Appointments Premium*.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateBookingRequest**](CreateBookingRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateBookingResponse**](CreateBookingResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBookings**
> ListBookingsResponse ListBookings(ctx, optional)
ListBookings

Retrieve a collection of bookings.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookingsApiListBookingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookingsApiListBookingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The maximum number of results per page to return in a paged response. | 
 **cursor** | **optional.String**| The pagination cursor from the preceding response to return the next page of the results. Do not set this when retrieving the first page of the results. | 
 **customerId** | **optional.String**| The [customer](https://developer.squareup.com/reference/square_2024-01-18/objects/Customer) for whom to retrieve bookings. If this is not set, bookings for all customers are retrieved. | 
 **teamMemberId** | **optional.String**| The team member for whom to retrieve bookings. If this is not set, bookings of all members are retrieved. | 
 **locationId** | **optional.String**| The location for which to retrieve bookings. If this is not set, all locations&#x27; bookings are retrieved. | 
 **startAtMin** | **optional.String**| The RFC 3339 timestamp specifying the earliest of the start time. If this is not set, the current time is used. | 
 **startAtMax** | **optional.String**| The RFC 3339 timestamp specifying the latest of the start time. If this is not set, the time of 31 days after &#x60;start_at_min&#x60; is used. | 

### Return type

[**ListBookingsResponse**](ListBookingsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLocationBookingProfiles**
> ListLocationBookingProfilesResponse ListLocationBookingProfiles(ctx, optional)
ListLocationBookingProfiles

Lists location booking profiles of a seller.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookingsApiListLocationBookingProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookingsApiListLocationBookingProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The maximum number of results to return in a paged response. | 
 **cursor** | **optional.String**| The pagination cursor from the preceding response to return the next page of the results. Do not set this when retrieving the first page of the results. | 

### Return type

[**ListLocationBookingProfilesResponse**](ListLocationBookingProfilesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTeamMemberBookingProfiles**
> ListTeamMemberBookingProfilesResponse ListTeamMemberBookingProfiles(ctx, optional)
ListTeamMemberBookingProfiles

Lists booking profiles for team members.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookingsApiListTeamMemberBookingProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookingsApiListTeamMemberBookingProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookableOnly** | **optional.Bool**| Indicates whether to include only bookable team members in the returned result (&#x60;true&#x60;) or not (&#x60;false&#x60;). | 
 **limit** | **optional.Int32**| The maximum number of results to return in a paged response. | 
 **cursor** | **optional.String**| The pagination cursor from the preceding response to return the next page of the results. Do not set this when retrieving the first page of the results. | 
 **locationId** | **optional.String**| Indicates whether to include only team members enabled at the given location in the returned result. | 

### Return type

[**ListTeamMemberBookingProfilesResponse**](ListTeamMemberBookingProfilesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveBooking**
> RetrieveBookingResponse RetrieveBooking(ctx, bookingId)
RetrieveBooking

Retrieves a booking.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bookingId** | **string**| The ID of the [Booking](https://developer.squareup.com/reference/square_2024-01-18/objects/Booking) object representing the to-be-retrieved booking. | 

### Return type

[**RetrieveBookingResponse**](RetrieveBookingResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveBusinessBookingProfile**
> RetrieveBusinessBookingProfileResponse RetrieveBusinessBookingProfile(ctx, )
RetrieveBusinessBookingProfile

Retrieves a seller's booking profile.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RetrieveBusinessBookingProfileResponse**](RetrieveBusinessBookingProfileResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveLocationBookingProfile**
> RetrieveLocationBookingProfileResponse RetrieveLocationBookingProfile(ctx, locationId)
RetrieveLocationBookingProfile

Retrieves a seller's location booking profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to retrieve the booking profile. | 

### Return type

[**RetrieveLocationBookingProfileResponse**](RetrieveLocationBookingProfileResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveTeamMemberBookingProfile**
> RetrieveTeamMemberBookingProfileResponse RetrieveTeamMemberBookingProfile(ctx, teamMemberId)
RetrieveTeamMemberBookingProfile

Retrieves a team member's booking profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamMemberId** | **string**| The ID of the team member to retrieve. | 

### Return type

[**RetrieveTeamMemberBookingProfileResponse**](RetrieveTeamMemberBookingProfileResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchAvailability**
> SearchAvailabilityResponse SearchAvailability(ctx, body)
SearchAvailability

Searches for availabilities for booking.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_READ` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_READ` and `APPOINTMENTS_READ` for the OAuth scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchAvailabilityRequest**](SearchAvailabilityRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchAvailabilityResponse**](SearchAvailabilityResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBooking**
> UpdateBookingResponse UpdateBooking(ctx, body, bookingId)
UpdateBooking

Updates a booking.  To call this endpoint with buyer-level permissions, set `APPOINTMENTS_WRITE` for the OAuth scope. To call this endpoint with seller-level permissions, set `APPOINTMENTS_ALL_WRITE` and `APPOINTMENTS_WRITE` for the OAuth scope.  For calls to this endpoint with seller-level permissions to succeed, the seller must have subscribed to *Appointments Plus* or *Appointments Premium*.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateBookingRequest**](UpdateBookingRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **bookingId** | **string**| The ID of the [Booking](https://developer.squareup.com/reference/square_2024-01-18/objects/Booking) object representing the to-be-updated booking. | 

### Return type

[**UpdateBookingResponse**](UpdateBookingResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

