/*
 * Square Connect API
 *
 * Client library for accessing the Square Connect APIs
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package squareup

// Defines the parameters that can be included in the body of a request to the [Charge](https://developer.squareup.com/reference/square_2024-01-18/transactions-api/charge) endpoint.  Deprecated - recommend using [CreatePayment](https://developer.squareup.com/reference/square_2024-01-18/payments-api/create-payment)
type ChargeRequest struct {
	// A value you specify that uniquely identifies this transaction among transactions you've created.  If you're unsure whether a particular transaction succeeded, you can reattempt it with the same idempotency key without worrying about double-charging the buyer.  See [Idempotency keys](https://developer.squareup.com/docs/working-with-apis/idempotency) for more information.
	IdempotencyKey string `json:"idempotency_key"`
	AmountMoney    *Money `json:"amount_money"`
	// A payment token generated from the [Card.tokenize()](https://developer.squareup.com/reference/sdks/web/payments/objects/Card#Card.tokenize) that represents the card to charge.  The application that provides a payment token to this endpoint must be the _same application_ that generated the payment token with the Web Payments SDK. Otherwise, the nonce is invalid.  Do not provide a value for this field if you provide a value for `customer_card_id`.
	CardNonce string `json:"card_nonce,omitempty"`
	// The ID of the customer card on file to charge. Do not provide a value for this field if you provide a value for `card_nonce`.  If you provide this value, you _must_ also provide a value for `customer_id`.
	CustomerCardId string `json:"customer_card_id,omitempty"`
	// If `true`, the request will only perform an Auth on the provided card. You can then later perform either a Capture (with the [CaptureTransaction](https://developer.squareup.com/reference/square_2024-01-18/transactions-api/capture-transaction) endpoint) or a Void (with the [VoidTransaction](https://developer.squareup.com/reference/square_2024-01-18/transactions-api/void-transaction) endpoint).  Default value: `false`
	DelayCapture bool `json:"delay_capture,omitempty"`
	// An optional ID you can associate with the transaction for your own purposes (such as to associate the transaction with an entity ID in your own database).  This value cannot exceed 40 characters.
	ReferenceId string `json:"reference_id,omitempty"`
	// An optional note to associate with the transaction.  This value cannot exceed 60 characters.
	Note string `json:"note,omitempty"`
	// The ID of the customer to associate this transaction with. This field is required if you provide a value for `customer_card_id`, and optional otherwise.
	CustomerId      string   `json:"customer_id,omitempty"`
	BillingAddress  *Address `json:"billing_address,omitempty"`
	ShippingAddress *Address `json:"shipping_address,omitempty"`
	// The buyer's email address, if available. This value is optional, but this transaction is ineligible for chargeback protection if it is not provided.
	BuyerEmailAddress string `json:"buyer_email_address,omitempty"`
	// The ID of the order to associate with this transaction.  If you provide this value, the `amount_money` value of your request must __exactly match__ the value of the order's `total_money` field.
	OrderId string `json:"order_id,omitempty"`
	// The basic primitive of multi-party transaction. The value is optional. The transaction facilitated by you can be split from here.  If you provide this value, the `amount_money` value in your additional_recipients must not be more than 90% of the `amount_money` value in the charge request. The `location_id` must be the valid location of the app owner merchant.  This field requires the `PAYMENTS_WRITE_ADDITIONAL_RECIPIENTS` OAuth permission.  This field is currently not supported in sandbox.
	AdditionalRecipients []AdditionalRecipient `json:"additional_recipients,omitempty"`
	// A token generated by SqPaymentForm's verifyBuyer() that represents customer's device info and 3ds challenge result.
	VerificationToken string `json:"verification_token,omitempty"`
}
