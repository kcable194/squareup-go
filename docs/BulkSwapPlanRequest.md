# BulkSwapPlanRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPlanVariationId** | **string** | The ID of the new subscription plan variation.  This field is required. | [default to null]
**OldPlanVariationId** | **string** | The ID of the plan variation whose subscriptions should be swapped. Active subscriptions using this plan variation will be subscribed to the new plan variation on their next billing day. | [default to null]
**LocationId** | **string** | The ID of the location to associate with the swapped subscriptions. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

