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

// Payment include an` itemizations` field that lists the items purchased, along with associated fees, modifiers, and discounts. Each itemization has an `itemization_type` field that indicates which of the following the itemization represents:  <ul> <li>An item variation from the merchant's item library</li> <li>A custom monetary amount</li> <li> An action performed on a Square gift card, such as activating or reloading it. </li> </ul>  *Note**: itemization information included in a `Payment` object reflects details collected **at the time of the payment**. Details such as the name or price of items might have changed since the payment was processed.
type V1PaymentItemization struct {
	// The item's name.
	Name string `json:"name,omitempty"`
	// The quantity of the item purchased. This can be a decimal value.
	Quantity float64 `json:"quantity,omitempty"`
	// The type of purchase that the itemization represents, such as an ITEM or CUSTOM_AMOUNT
	ItemizationType string               `json:"itemization_type,omitempty"`
	ItemDetail      *V1PaymentItemDetail `json:"item_detail,omitempty"`
	// Notes entered by the merchant about the item at the time of payment, if any.
	Notes string `json:"notes,omitempty"`
	// The name of the item variation purchased, if any.
	ItemVariationName   string   `json:"item_variation_name,omitempty"`
	TotalMoney          *V1Money `json:"total_money,omitempty"`
	SingleQuantityMoney *V1Money `json:"single_quantity_money,omitempty"`
	GrossSalesMoney     *V1Money `json:"gross_sales_money,omitempty"`
	DiscountMoney       *V1Money `json:"discount_money,omitempty"`
	NetSalesMoney       *V1Money `json:"net_sales_money,omitempty"`
	// All taxes applied to this itemization.
	Taxes []V1PaymentTax `json:"taxes,omitempty"`
	// All discounts applied to this itemization.
	Discounts []V1PaymentDiscount `json:"discounts,omitempty"`
	// All modifier options applied to this itemization.
	Modifiers []V1PaymentModifier `json:"modifiers,omitempty"`
}
