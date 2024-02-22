# LoyaltyEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-assigned ID of the loyalty event. | [default to null]
**Type_** | **string** | The type of the loyalty event. | [default to null]
**CreatedAt** | **string** | The timestamp when the event was created, in RFC 3339 format. | [default to null]
**AccumulatePoints** | [***LoyaltyEventAccumulatePoints**](LoyaltyEventAccumulatePoints.md) |  | [optional] [default to null]
**CreateReward** | [***LoyaltyEventCreateReward**](LoyaltyEventCreateReward.md) |  | [optional] [default to null]
**RedeemReward** | [***LoyaltyEventRedeemReward**](LoyaltyEventRedeemReward.md) |  | [optional] [default to null]
**DeleteReward** | [***LoyaltyEventDeleteReward**](LoyaltyEventDeleteReward.md) |  | [optional] [default to null]
**AdjustPoints** | [***LoyaltyEventAdjustPoints**](LoyaltyEventAdjustPoints.md) |  | [optional] [default to null]
**LoyaltyAccountId** | **string** | The ID of the [loyalty account](https://developer.squareup.com/reference/square_2024-01-18/objects/LoyaltyAccount) associated with the event. | [default to null]
**LocationId** | **string** | The ID of the [location](https://developer.squareup.com/reference/square_2024-01-18/objects/Location) where the event occurred. | [optional] [default to null]
**Source** | **string** | Defines whether the event was generated by the Square Point of Sale. | [default to null]
**ExpirePoints** | [***LoyaltyEventExpirePoints**](LoyaltyEventExpirePoints.md) |  | [optional] [default to null]
**OtherEvent** | [***LoyaltyEventOther**](LoyaltyEventOther.md) |  | [optional] [default to null]
**AccumulatePromotionPoints** | [***LoyaltyEventAccumulatePromotionPoints**](LoyaltyEventAccumulatePromotionPoints.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
