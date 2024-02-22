# LoyaltyEventRedeemReward

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoyaltyProgramId** | **string** | The ID of the [loyalty program](https://developer.squareup.com/reference/square_2024-01-18/objects/LoyaltyProgram). | [default to null]
**RewardId** | **string** | The ID of the redeemed [loyalty reward](https://developer.squareup.com/reference/square_2024-01-18/objects/LoyaltyReward). This field is returned only if the event source is &#x60;LOYALTY_API&#x60;. | [optional] [default to null]
**OrderId** | **string** | The ID of the [order](https://developer.squareup.com/reference/square_2024-01-18/objects/Order) that redeemed the reward. This field is returned only if the Orders API is used to process orders. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

