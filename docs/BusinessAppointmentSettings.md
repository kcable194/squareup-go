# BusinessAppointmentSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationTypes** | **[]string** | Types of the location allowed for bookings. | [optional] [default to null]
**AlignmentTime** | **string** | The time unit of the service duration for bookings. | [optional] [default to null]
**MinBookingLeadTimeSeconds** | **int32** | The minimum lead time in seconds before a service can be booked. A booking must be created at least this amount of time before its starting time. | [optional] [default to null]
**MaxBookingLeadTimeSeconds** | **int32** | The maximum lead time in seconds before a service can be booked. A booking must be created at most this amount of time before its starting time. | [optional] [default to null]
**AnyTeamMemberBookingEnabled** | **bool** | Indicates whether a customer can choose from all available time slots and have a staff member assigned automatically (&#x60;true&#x60;) or not (&#x60;false&#x60;). | [optional] [default to null]
**MultipleServiceBookingEnabled** | **bool** | Indicates whether a customer can book multiple services in a single online booking. | [optional] [default to null]
**MaxAppointmentsPerDayLimitType** | **string** | Indicates whether the daily appointment limit applies to team members or to business locations. | [optional] [default to null]
**MaxAppointmentsPerDayLimit** | **int32** | The maximum number of daily appointments per team member or per location. | [optional] [default to null]
**CancellationWindowSeconds** | **int32** | The cut-off time in seconds for allowing clients to cancel or reschedule an appointment. | [optional] [default to null]
**CancellationFeeMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CancellationPolicy** | **string** | The cancellation policy adopted by the seller. | [optional] [default to null]
**CancellationPolicyText** | **string** | The free-form text of the seller&#x27;s cancellation policy. | [optional] [default to null]
**SkipBookingFlowStaffSelection** | **bool** | Indicates whether customers has an assigned staff member (&#x60;true&#x60;) or can select s staff member of their choice (&#x60;false&#x60;). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

