/*
 * Mailchimp Marketing API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0.76
 * Contact: apihelp@mailchimp.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Information about a specific cart.
type EcommerceCart1 struct {
	// A unique identifier for the cart.
	Id string `json:"id"`
	Customer *EcommerceCustomer1 `json:"customer"`
	// A string that uniquely identifies the campaign for a cart.
	CampaignId string `json:"campaign_id,omitempty"`
	// The URL for the cart. This parameter is required for [Abandoned Cart](https://mailchimp.com/help/create-an-abandoned-cart-email/) automations.
	CheckoutUrl string `json:"checkout_url,omitempty"`
	// The three-letter ISO 4217 code for the currency that the cart uses.
	CurrencyCode string `json:"currency_code"`
	// The order total for the cart.
	OrderTotal float32 `json:"order_total"`
	// The total tax for the cart.
	TaxTotal float32 `json:"tax_total,omitempty"`
	// An array of the cart's line items.
	Lines []EcommerceCartLineItem1 `json:"lines"`
}