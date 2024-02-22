# Tender

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The tender&#x27;s unique ID. It is the associated payment ID. | [optional] [default to null]
**LocationId** | **string** | The ID of the transaction&#x27;s associated location. | [optional] [default to null]
**TransactionId** | **string** | The ID of the tender&#x27;s associated transaction. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp for when the tender was created, in RFC 3339 format. | [optional] [default to null]
**Note** | **string** | An optional note associated with the tender at the time of payment. | [optional] [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TipMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**ProcessingFeeMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CustomerId** | **string** | If the tender is associated with a customer or represents a customer&#x27;s card on file, this is the ID of the associated customer. | [optional] [default to null]
**Type_** | **string** | The type of tender, such as &#x60;CARD&#x60; or &#x60;CASH&#x60;. | [default to null]
**CardDetails** | [***TenderCardDetails**](TenderCardDetails.md) |  | [optional] [default to null]
**CashDetails** | [***TenderCashDetails**](TenderCashDetails.md) |  | [optional] [default to null]
**BankAccountDetails** | [***TenderBankAccountDetails**](TenderBankAccountDetails.md) |  | [optional] [default to null]
**BuyNowPayLaterDetails** | [***TenderBuyNowPayLaterDetails**](TenderBuyNowPayLaterDetails.md) |  | [optional] [default to null]
**SquareAccountDetails** | [***TenderSquareAccountDetails**](TenderSquareAccountDetails.md) |  | [optional] [default to null]
**AdditionalRecipients** | [**[]AdditionalRecipient**](AdditionalRecipient.md) | Additional recipients (other than the merchant) receiving a portion of this tender. For example, fees assessed on the purchase by a third party integration. | [optional] [default to null]
**PaymentId** | **string** | The ID of the [Payment](https://developer.squareup.com/reference/square_2024-01-18/objects/Payment) that corresponds to this tender. This value is only present for payments created with the v2 Payments API. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

