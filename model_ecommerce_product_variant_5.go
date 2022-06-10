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

// Information about a specific product variant.
type EcommerceProductVariant5 struct {
	// The title of a product variant.
	Title string `json:"title,omitempty"`
	// The URL for a product variant.
	Url string `json:"url,omitempty"`
	// The stock keeping unit (SKU) of a product variant.
	Sku string `json:"sku,omitempty"`
	// The price of a product variant.
	Price float64 `json:"price,omitempty"`
	// The inventory quantity of a product variant.
	InventoryQuantity int32 `json:"inventory_quantity,omitempty"`
	// The image URL for a product variant.
	ImageUrl string `json:"image_url,omitempty"`
	// The backorders of a product variant.
	Backorders string `json:"backorders,omitempty"`
	// The visibility of a product variant.
	Visibility string `json:"visibility,omitempty"`
}
