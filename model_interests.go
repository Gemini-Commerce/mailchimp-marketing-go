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

// A list of this category's interests
type Interests struct {
	// An array of this category's interests
	Interests []Interest `json:"interests,omitempty"`
	// The unique list id that the interests belong to.
	ListId string `json:"list_id,omitempty"`
	// The id for the interest category.
	CategoryId string `json:"category_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}