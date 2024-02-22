# CatalogSubscriptionPlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the plan. | [default to null]
**Phases** | [**[]SubscriptionPhase**](SubscriptionPhase.md) | A list of SubscriptionPhase containing the [SubscriptionPhase](https://developer.squareup.com/reference/square_2024-01-18/objects/SubscriptionPhase) for this plan. This field it required. Not including this field will throw a REQUIRED_FIELD_MISSING error | [optional] [default to null]
**SubscriptionPlanVariations** | [**[]CatalogObject**](CatalogObject.md) | The list of subscription plan variations available for this product | [optional] [default to null]
**EligibleItemIds** | **[]string** | The list of IDs of &#x60;CatalogItems&#x60; that are eligible for subscription by this SubscriptionPlan&#x27;s variations. | [optional] [default to null]
**EligibleCategoryIds** | **[]string** | The list of IDs of &#x60;CatalogCategory&#x60; that are eligible for subscription by this SubscriptionPlan&#x27;s variations. | [optional] [default to null]
**AllItems** | **bool** | If true, all items in the merchant&#x27;s catalog are subscribable by this SubscriptionPlan. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

