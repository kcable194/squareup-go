# LoyaltyProgramRewardDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** | Indicates the scope of the reward tier. DEPRECATED at version 2020-12-16. You can find this information in the &#x60;product_set_data&#x60; field of the &#x60;PRODUCT_SET&#x60; catalog object referenced by the pricing rule. For &#x60;ORDER&#x60; scopes, &#x60;all_products&#x60; is true. For &#x60;ITEM_VARIATION&#x60; or &#x60;CATEGORY&#x60; scopes, &#x60;product_ids_any&#x60; is a list of catalog object IDs of the given type. | [default to null]
**DiscountType** | **string** | The type of discount the reward tier offers. DEPRECATED at version 2020-12-16. You can find this information in the &#x60;discount_data.discount_type&#x60; field of the &#x60;DISCOUNT&#x60; catalog object referenced by the pricing rule. | [default to null]
**PercentageDiscount** | **string** | The fixed percentage of the discount. Present if &#x60;discount_type&#x60; is &#x60;FIXED_PERCENTAGE&#x60;. For example, a 7.25% off discount will be represented as \&quot;7.25\&quot;. DEPRECATED at version 2020-12-16. You can find this information in the &#x60;discount_data.percentage&#x60; field of the &#x60;DISCOUNT&#x60; catalog object referenced by the pricing rule. | [optional] [default to null]
**CatalogObjectIds** | **[]string** | The list of catalog objects to which this reward can be applied. They are either all item-variation ids or category ids, depending on the &#x60;type&#x60; field. DEPRECATED at version 2020-12-16. You can find this information in the &#x60;product_set_data.product_ids_any&#x60; field of the &#x60;PRODUCT_SET&#x60; catalog object referenced by the pricing rule. | [optional] [default to null]
**FixedDiscountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**MaxDiscountMoney** | [***Money**](Money.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

