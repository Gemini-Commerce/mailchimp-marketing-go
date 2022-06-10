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

// Information about a specific order line.
type EcommerceOrderLineItem1 struct {
	// A unique identifier for the product associated with the order line item.
	ProductId string `json:"product_id,omitempty"`
	// A unique identifier for the product variant associated with the order line item.
	ProductVariantId string `json:"product_variant_id,omitempty"`
	// The quantity of an order line item.
	Quantity int32 `json:"quantity,omitempty"`
	// The price of an order line item.
	Price float64 `json:"price,omitempty"`
	// The total discount amount applied to this line item.
	Discount float64 `json:"discount,omitempty"`
}
