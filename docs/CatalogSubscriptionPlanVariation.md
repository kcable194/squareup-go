# CatalogSubscriptionPlanVariation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the plan variation. | [default to null]
**Phases** | [**[]SubscriptionPhase**](SubscriptionPhase.md) | A list containing each [SubscriptionPhase](https://developer.squareup.com/reference/square_2024-01-18/objects/SubscriptionPhase) for this plan variation. | [default to null]
**SubscriptionPlanId** | **string** | The id of the subscription plan, if there is one. | [optional] [default to null]
**MonthlyBillingAnchorDate** | **int64** | The day of the month the billing period starts. | [optional] [default to null]
**CanProrate** | **bool** | Whether bills for this plan variation can be split for proration. | [optional] [default to null]
**SuccessorPlanVariationId** | **string** | The ID of a \&quot;successor\&quot; plan variation to this one. If the field is set, and this object is disabled at all locations, it indicates that this variation is deprecated and the object identified by the successor ID be used in its stead. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

