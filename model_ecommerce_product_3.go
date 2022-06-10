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

// Information about a specific product.
type EcommerceProduct3 struct {
	// A unique identifier for the product.
	Id string `json:"id,omitempty"`
	// The currency code
	CurrencyCode string `json:"currency_code,omitempty"`
	// The title of a product.
	Title string `json:"title,omitempty"`
	// The handle of a product.
	Handle string `json:"handle,omitempty"`
	// The URL for a product.
	Url string `json:"url,omitempty"`
	// The description of a product.
	Description string `json:"description,omitempty"`
	// The type of product.
	Type_ string `json:"type,omitempty"`
	// The vendor for a product.
	Vendor string `json:"vendor,omitempty"`
	// The image URL for a product.
	ImageUrl string `json:"image_url,omitempty"`
	// Returns up to 50 of the product's variants. To retrieve all variants use [Product Variants](https://mailchimp.com/developer/marketing/api/ecommerce-product-variants/).
	Variants []EcommerceProductVariant6 `json:"variants,omitempty"`
	// An array of the product's images.
	Images []EcommerceProductImage5 `json:"images,omitempty"`
	// The date and time the product was published in ISO 8601 format.
	PublishedAtForeign time.Time `json:"published_at_foreign,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
