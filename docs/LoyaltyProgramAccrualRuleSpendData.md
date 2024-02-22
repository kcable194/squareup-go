# LoyaltyProgramAccrualRuleSpendData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountMoney** | [***Money**](Money.md) |  | [default to null]
**ExcludedCategoryIds** | **[]string** | The IDs of any &#x60;CATEGORY&#x60; catalog objects that are excluded from points accrual.  You can use the [BatchRetrieveCatalogObjects](https://developer.squareup.com/reference/square_2024-01-18/catalog-api/batch-retrieve-catalog-objects) endpoint to retrieve information about the excluded categories. | [optional] [default to null]
**ExcludedItemVariationIds** | **[]string** | The IDs of any &#x60;ITEM_VARIATION&#x60; catalog objects that are excluded from points accrual.  You can use the [BatchRetrieveCatalogObjects](https://developer.squareup.com/reference/square_2024-01-18/catalog-api/batch-retrieve-catalog-objects) endpoint to retrieve information about the excluded item variations. | [optional] [default to null]
**TaxMode** | **string** | Indicates how taxes should be treated when calculating the purchase amount used for points accrual. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

