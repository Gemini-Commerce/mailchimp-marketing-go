/*
 * Mailchimp Marketing API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0.76
 * Contact: apihelp@mailchimp.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mcapimarketing
import (
	"time"
)

// Information about a specific order.
type EcommerceOrder3 struct {
	// A unique identifier for the order.
	Id string `json:"id,omitempty"`
	Customer *EcommerceCustomer6 `json:"customer,omitempty"`
	// The unique identifier for the store.
	StoreId string `json:"store_id,omitempty"`
	// A string that uniquely identifies the campaign associated with an order.
	CampaignId string `json:"campaign_id,omitempty"`
	// The URL for the page where the buyer landed when entering the shop.
	LandingSite string `json:"landing_site,omitempty"`
	// The order status. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications).
	FinancialStatus string `json:"financial_status,omitempty"`
	// The fulfillment status for the order. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications).
	FulfillmentStatus string `json:"fulfillment_status,omitempty"`
	// The three-letter ISO 4217 code for the currency that the store accepts.
	CurrencyCode string `json:"currency_code,omitempty"`
	// The order total associated with an order.
	OrderTotal float64 `json:"order_total,omitempty"`
	// The URL for the order.
	OrderUrl string `json:"order_url,omitempty"`
	// The total amount of the discounts to be applied to the price of the order.
	DiscountTotal float64 `json:"discount_total,omitempty"`
	// The tax total associated with an order.
	TaxTotal float64 `json:"tax_total,omitempty"`
	// The shipping total for the order.
	ShippingTotal float64 `json:"shipping_total,omitempty"`
	// The Mailchimp tracking code for the order. Uses the 'mc_tc' parameter in E-Commerce tracking URLs.
	TrackingCode string `json:"tracking_code,omitempty"`
	// The date and time the order was processed in ISO 8601 format.
	ProcessedAtForeign time.Time `json:"processed_at_foreign,omitempty"`
	// The date and time the order was cancelled in ISO 8601 format.
	CancelledAtForeign time.Time `json:"cancelled_at_foreign,omitempty"`
	// The date and time the order was updated in ISO 8601 format.
	UpdatedAtForeign time.Time `json:"updated_at_foreign,omitempty"`
	ShippingAddress *ShippingAddress1 `json:"shipping_address,omitempty"`
	BillingAddress *BillingAddress1 `json:"billing_address,omitempty"`
	// The promo codes applied on the order
	Promos []Orders1Promos `json:"promos,omitempty"`
	// An array of the order's line items.
	Lines []EcommerceOrderLineItem5 `json:"lines,omitempty"`
	Outreach *Outreach1 `json:"outreach,omitempty"`
	// The tracking number associated with the order.
	TrackingNumber string `json:"tracking_number,omitempty"`
	// The tracking carrier associated with the order.
	TrackingCarrier string `json:"tracking_carrier,omitempty"`
	// The tracking URL associated with the order.
	TrackingUrl string `json:"tracking_url,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}