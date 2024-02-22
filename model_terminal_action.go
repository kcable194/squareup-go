/*
 * Square Connect API
 *
 * Client library for accessing the Square Connect APIs
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Represents an action processed by the Square Terminal.
type TerminalAction struct {
	// A unique ID for this `TerminalAction`.
	Id string `json:"id,omitempty"`
	// The unique Id of the device intended for this `TerminalAction`. The Id can be retrieved from /v2/devices api.
	DeviceId string `json:"device_id,omitempty"`
	// The duration as an RFC 3339 duration, after which the action will be automatically canceled. TerminalActions that are `PENDING` will be automatically `CANCELED` and have a cancellation reason of `TIMED_OUT`  Default: 5 minutes from creation  Maximum: 5 minutes
	DeadlineDuration string `json:"deadline_duration,omitempty"`
	// The status of the `TerminalAction`. Options: `PENDING`, `IN_PROGRESS`, `CANCEL_REQUESTED`, `CANCELED`, `COMPLETED`
	Status string `json:"status,omitempty"`
	// The reason why `TerminalAction` is canceled. Present if the status is `CANCELED`.
	CancelReason string `json:"cancel_reason,omitempty"`
	// The time when the `TerminalAction` was created as an RFC 3339 timestamp.
	CreatedAt string `json:"created_at,omitempty"`
	// The time when the `TerminalAction` was last updated as an RFC 3339 timestamp.
	UpdatedAt string `json:"updated_at,omitempty"`
	// The ID of the application that created the action.
	AppId string `json:"app_id,omitempty"`
	// The location id the action is attached to, if a link can be made.
	LocationId string `json:"location_id,omitempty"`
	// Represents the type of the action.
	Type_ string `json:"type,omitempty"`
	QrCodeOptions *QrCodeOptions `json:"qr_code_options,omitempty"`
	SaveCardOptions *SaveCardOptions `json:"save_card_options,omitempty"`
	SignatureOptions *SignatureOptions `json:"signature_options,omitempty"`
	ConfirmationOptions *ConfirmationOptions `json:"confirmation_options,omitempty"`
	ReceiptOptions *ReceiptOptions `json:"receipt_options,omitempty"`
	DataCollectionOptions *DataCollectionOptions `json:"data_collection_options,omitempty"`
	SelectOptions *SelectOptions `json:"select_options,omitempty"`
	DeviceMetadata *DeviceMetadata `json:"device_metadata,omitempty"`
	// Indicates the action will be linked to another action and requires a waiting dialog to be displayed instead of returning to the idle screen on completion of the action.  Only supported on SIGNATURE, CONFIRMATION, DATA_COLLECTION, and SELECT types.
	AwaitNextAction bool `json:"await_next_action,omitempty"`
	// The timeout duration of the waiting dialog as an RFC 3339 duration, after which the waiting dialog will no longer be displayed and the Terminal will return to the idle screen.  Default: 5 minutes from when the waiting dialog is displayed  Maximum: 5 minutes
	AwaitNextActionDuration string `json:"await_next_action_duration,omitempty"`
}
