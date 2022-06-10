# EcommerceOrder3

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the order. | [optional] [default to null]
**Customer** | [***EcommerceCustomer6**](Ecommerce Customer_6.md) |  | [optional] [default to null]
**StoreId** | **string** | The unique identifier for the store. | [optional] [default to null]
**CampaignId** | **string** | A string that uniquely identifies the campaign associated with an order. | [optional] [default to null]
**LandingSite** | **string** | The URL for the page where the buyer landed when entering the shop. | [optional] [default to null]
**FinancialStatus** | **string** | The order status. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications). | [optional] [default to null]
**FulfillmentStatus** | **string** | The fulfillment status for the order. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications). | [optional] [default to null]
**CurrencyCode** | **string** | The three-letter ISO 4217 code for the currency that the store accepts. | [optional] [default to null]
**OrderTotal** | **float64** | The order total associated with an order. | [optional] [default to null]
**OrderUrl** | **string** | The URL for the order. | [optional] [default to null]
**DiscountTotal** | **float64** | The total amount of the discounts to be applied to the price of the order. | [optional] [default to null]
**TaxTotal** | **float64** | The tax total associated with an order. | [optional] [default to null]
**ShippingTotal** | **float64** | The shipping total for the order. | [optional] [default to null]
**TrackingCode** | **string** | The Mailchimp tracking code for the order. Uses the &#x27;mc_tc&#x27; parameter in E-Commerce tracking URLs. | [optional] [default to null]
**ProcessedAtForeign** | [**time.Time**](time.Time.md) | The date and time the order was processed in ISO 8601 format. | [optional] [default to null]
**CancelledAtForeign** | [**time.Time**](time.Time.md) | The date and time the order was cancelled in ISO 8601 format. | [optional] [default to null]
**UpdatedAtForeign** | [**time.Time**](time.Time.md) | The date and time the order was updated in ISO 8601 format. | [optional] [default to null]
**ShippingAddress** | [***ShippingAddress1**](Shipping Address_1.md) |  | [optional] [default to null]
**BillingAddress** | [***BillingAddress1**](Billing Address_1.md) |  | [optional] [default to null]
**Promos** | [**[]Orders1Promos**](Orders_1_promos.md) | The promo codes applied on the order | [optional] [default to null]
**Lines** | [**[]EcommerceOrderLineItem5**](Ecommerce Order Line Item_5.md) | An array of the order&#x27;s line items. | [optional] [default to null]
**Outreach** | [***Outreach1**](Outreach_1.md) |  | [optional] [default to null]
**TrackingNumber** | **string** | The tracking number associated with the order. | [optional] [default to null]
**TrackingCarrier** | **string** | The tracking carrier associated with the order. | [optional] [default to null]
**TrackingUrl** | **string** | The tracking URL associated with the order. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

