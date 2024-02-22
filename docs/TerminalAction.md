# TerminalAction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique ID for this &#x60;TerminalAction&#x60;. | [optional] [default to null]
**DeviceId** | **string** | The unique Id of the device intended for this &#x60;TerminalAction&#x60;. The Id can be retrieved from /v2/devices api. | [optional] [default to null]
**DeadlineDuration** | **string** | The duration as an RFC 3339 duration, after which the action will be automatically canceled. TerminalActions that are &#x60;PENDING&#x60; will be automatically &#x60;CANCELED&#x60; and have a cancellation reason of &#x60;TIMED_OUT&#x60;  Default: 5 minutes from creation  Maximum: 5 minutes | [optional] [default to null]
**Status** | **string** | The status of the &#x60;TerminalAction&#x60;. Options: &#x60;PENDING&#x60;, &#x60;IN_PROGRESS&#x60;, &#x60;CANCEL_REQUESTED&#x60;, &#x60;CANCELED&#x60;, &#x60;COMPLETED&#x60; | [optional] [default to null]
**CancelReason** | **string** | The reason why &#x60;TerminalAction&#x60; is canceled. Present if the status is &#x60;CANCELED&#x60;. | [optional] [default to null]
**CreatedAt** | **string** | The time when the &#x60;TerminalAction&#x60; was created as an RFC 3339 timestamp. | [optional] [default to null]
**UpdatedAt** | **string** | The time when the &#x60;TerminalAction&#x60; was last updated as an RFC 3339 timestamp. | [optional] [default to null]
**AppId** | **string** | The ID of the application that created the action. | [optional] [default to null]
**LocationId** | **string** | The location id the action is attached to, if a link can be made. | [optional] [default to null]
**Type_** | **string** | Represents the type of the action. | [optional] [default to null]
**QrCodeOptions** | [***QrCodeOptions**](QrCodeOptions.md) |  | [optional] [default to null]
**SaveCardOptions** | [***SaveCardOptions**](SaveCardOptions.md) |  | [optional] [default to null]
**SignatureOptions** | [***SignatureOptions**](SignatureOptions.md) |  | [optional] [default to null]
**ConfirmationOptions** | [***ConfirmationOptions**](ConfirmationOptions.md) |  | [optional] [default to null]
**ReceiptOptions** | [***ReceiptOptions**](ReceiptOptions.md) |  | [optional] [default to null]
**DataCollectionOptions** | [***DataCollectionOptions**](DataCollectionOptions.md) |  | [optional] [default to null]
**SelectOptions** | [***SelectOptions**](SelectOptions.md) |  | [optional] [default to null]
**DeviceMetadata** | [***DeviceMetadata**](DeviceMetadata.md) |  | [optional] [default to null]
**AwaitNextAction** | **bool** | Indicates the action will be linked to another action and requires a waiting dialog to be displayed instead of returning to the idle screen on completion of the action.  Only supported on SIGNATURE, CONFIRMATION, DATA_COLLECTION, and SELECT types. | [optional] [default to null]
**AwaitNextActionDuration** | **string** | The timeout duration of the waiting dialog as an RFC 3339 duration, after which the waiting dialog will no longer be displayed and the Terminal will return to the idle screen.  Default: 5 minutes from when the waiting dialog is displayed  Maximum: 5 minutes | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

