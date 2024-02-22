# Booking

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique ID of this object representing a booking. | [optional] [default to null]
**Version** | **int32** | The revision number for the booking used for optimistic concurrency. | [optional] [default to null]
**Status** | **string** | The status of the booking, describing where the booking stands with respect to the booking state machine. | [optional] [default to null]
**CreatedAt** | **string** | The RFC 3339 timestamp specifying the creation time of this booking. | [optional] [default to null]
**UpdatedAt** | **string** | The RFC 3339 timestamp specifying the most recent update time of this booking. | [optional] [default to null]
**StartAt** | **string** | The RFC 3339 timestamp specifying the starting time of this booking. | [optional] [default to null]
**LocationId** | **string** | The ID of the [Location](https://developer.squareup.com/reference/square_2024-01-18/objects/Location) object representing the location where the booked service is provided. Once set when the booking is created, its value cannot be changed. | [optional] [default to null]
**CustomerId** | **string** | The ID of the [Customer](https://developer.squareup.com/reference/square_2024-01-18/objects/Customer) object representing the customer receiving the booked service. | [optional] [default to null]
**CustomerNote** | **string** | The free-text field for the customer to supply notes about the booking. For example, the note can be preferences that cannot be expressed by supported attributes of a relevant [CatalogObject](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogObject) instance. | [optional] [default to null]
**SellerNote** | **string** | The free-text field for the seller to supply notes about the booking. For example, the note can be preferences that cannot be expressed by supported attributes of a specific [CatalogObject](https://developer.squareup.com/reference/square_2024-01-18/objects/CatalogObject) instance. This field should not be visible to customers. | [optional] [default to null]
**AppointmentSegments** | [**[]AppointmentSegment**](AppointmentSegment.md) | A list of appointment segments for this booking. | [optional] [default to null]
**TransitionTimeMinutes** | **int32** | Additional time at the end of a booking. Applications should not make this field visible to customers of a seller. | [optional] [default to null]
**AllDay** | **bool** | Whether the booking is of a full business day. | [optional] [default to null]
**LocationType** | **string** | The type of location where the booking is held. Access to this field requires seller-level permissions. | [optional] [default to null]
**CreatorDetails** | [***BookingCreatorDetails**](BookingCreatorDetails.md) |  | [optional] [default to null]
**Source** | **string** | The source of the booking. Access to this field requires seller-level permissions. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

