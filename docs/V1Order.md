# V1Order

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Any errors that occurred during the request. | [optional] [default to null]
**Id** | **string** | The order&#x27;s unique identifier. | [optional] [default to null]
**BuyerEmail** | **string** | The email address of the order&#x27;s buyer. | [optional] [default to null]
**RecipientName** | **string** | The name of the order&#x27;s buyer. | [optional] [default to null]
**RecipientPhoneNumber** | **string** | The phone number to use for the order&#x27;s delivery. | [optional] [default to null]
**State** | **string** | Whether the tax is an ADDITIVE tax or an INCLUSIVE tax. | [optional] [default to null]
**ShippingAddress** | [***Address**](Address.md) |  | [optional] [default to null]
**SubtotalMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**TotalShippingMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**TotalTaxMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**TotalPriceMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**TotalDiscountMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**CreatedAt** | **string** | The time when the order was created, in ISO 8601 format. | [optional] [default to null]
**UpdatedAt** | **string** | The time when the order was last modified, in ISO 8601 format. | [optional] [default to null]
**ExpiresAt** | **string** | The time when the order expires if no action is taken, in ISO 8601 format. | [optional] [default to null]
**PaymentId** | **string** | The unique identifier of the payment associated with the order. | [optional] [default to null]
**BuyerNote** | **string** | A note provided by the buyer when the order was created, if any. | [optional] [default to null]
**CompletedNote** | **string** | A note provided by the merchant when the order&#x27;s state was set to COMPLETED, if any | [optional] [default to null]
**RefundedNote** | **string** | A note provided by the merchant when the order&#x27;s state was set to REFUNDED, if any. | [optional] [default to null]
**CanceledNote** | **string** | A note provided by the merchant when the order&#x27;s state was set to CANCELED, if any. | [optional] [default to null]
**Tender** | [***V1Tender**](V1Tender.md) |  | [optional] [default to null]
**OrderHistory** | [**[]V1OrderHistoryEntry**](V1OrderHistoryEntry.md) | The history of actions associated with the order. | [optional] [default to null]
**PromoCode** | **string** | The promo code provided by the buyer, if any. | [optional] [default to null]
**BtcReceiveAddress** | **string** | For Bitcoin transactions, the address that the buyer sent Bitcoin to. | [optional] [default to null]
**BtcPriceSatoshi** | **float64** | For Bitcoin transactions, the price of the buyer&#x27;s order in satoshi (100 million satoshi equals 1 BTC). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

