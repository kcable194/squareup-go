# CashDrawerShift

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The shift unique ID. | [optional] [default to null]
**State** | **string** | The shift current state. | [optional] [default to null]
**OpenedAt** | **string** | The time when the shift began, in ISO 8601 format. | [optional] [default to null]
**EndedAt** | **string** | The time when the shift ended, in ISO 8601 format. | [optional] [default to null]
**ClosedAt** | **string** | The time when the shift was closed, in ISO 8601 format. | [optional] [default to null]
**Description** | **string** | The free-form text description of a cash drawer by an employee. | [optional] [default to null]
**OpenedCashMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CashPaymentMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CashRefundsMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CashPaidInMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CashPaidOutMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**ExpectedCashMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**ClosedCashMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Device** | [***CashDrawerDevice**](CashDrawerDevice.md) |  | [optional] [default to null]
**CreatedAt** | **string** | The shift start time in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The shift updated at time in RFC 3339 format. | [optional] [default to null]
**LocationId** | **string** | The ID of the location the cash drawer shift belongs to. | [optional] [default to null]
**TeamMemberIds** | **[]string** | The IDs of all team members that were logged into Square Point of Sale at any point while the cash drawer shift was open. | [optional] [default to null]
**OpeningTeamMemberId** | **string** | The ID of the team member that started the cash drawer shift. | [optional] [default to null]
**EndingTeamMemberId** | **string** | The ID of the team member that ended the cash drawer shift. | [optional] [default to null]
**ClosingTeamMemberId** | **string** | The ID of the team member that closed the cash drawer shift by auditing the cash drawer contents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

