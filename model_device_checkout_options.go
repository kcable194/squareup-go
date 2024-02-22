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

type DeviceCheckoutOptions struct {
	// The unique ID of the device intended for this `TerminalCheckout`. A list of `DeviceCode` objects can be retrieved from the /v2/devices/codes endpoint. Match a `DeviceCode.device_id` value with `device_id` to get the associated device code.
	DeviceId string `json:"device_id"`
	// Instructs the device to skip the receipt screen. Defaults to false.
	SkipReceiptScreen bool `json:"skip_receipt_screen,omitempty"`
	// Indicates that signature collection is desired during checkout. Defaults to false.
	CollectSignature bool `json:"collect_signature,omitempty"`
	TipSettings *TipSettings `json:"tip_settings,omitempty"`
	// Show the itemization screen prior to taking a payment. This field is only meaningful when the checkout includes an order ID. Defaults to true.
	ShowItemizedCart bool `json:"show_itemized_cart,omitempty"`
}
