# SubscriptionEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the subscription event. | [default to null]
**SubscriptionEventType** | **string** | Type of the subscription event. | [default to null]
**EffectiveDate** | **string** | The &#x60;YYYY-MM-DD&#x60;-formatted date (for example, 2013-01-15) when the subscription event occurred. | [default to null]
**MonthlyBillingAnchorDate** | **int32** | The day-of-the-month the billing anchor date was changed to, if applicable. | [optional] [default to null]
**Info** | [***Info**](Info.md) |  | [optional] [default to null]
**Phases** | [**[]Phase**](Phase.md) | A list of Phases, to pass phase-specific information used in the swap. | [optional] [default to null]
**PlanVariationId** | **string** | The ID of the subscription plan variation associated with the subscription. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

