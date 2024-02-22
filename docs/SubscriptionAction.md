# SubscriptionAction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of an action scoped to a subscription. | [optional] [default to null]
**Type_** | **string** | The type of the action. | [optional] [default to null]
**EffectiveDate** | **string** | The &#x60;YYYY-MM-DD&#x60;-formatted date when the action occurs on the subscription. | [optional] [default to null]
**MonthlyBillingAnchorDate** | **int32** | The new billing anchor day value, for a &#x60;CHANGE_BILLING_ANCHOR_DATE&#x60; action. | [optional] [default to null]
**Phases** | [**[]Phase**](Phase.md) | A list of Phases, to pass phase-specific information used in the swap. | [optional] [default to null]
**NewPlanVariationId** | **string** | The target subscription plan variation that a subscription switches to, for a &#x60;SWAP_PLAN&#x60; action. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

