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

// Assign subscribers to interests to group them together. Interests are referred to as 'group names' in the Mailchimp application.
type Interest3 struct {
	// The id for the interest category.
	CategoryId string `json:"category_id,omitempty"`
	// The ID for the list that this interest belongs to.
	ListId string `json:"list_id,omitempty"`
	// The ID for the interest.
	Id string `json:"id,omitempty"`
	// The name of the interest. This can be shown publicly on a subscription form.
	Name string `json:"name,omitempty"`
	// The number of subscribers associated with this interest.
	SubscriberCount string `json:"subscriber_count,omitempty"`
	// The display order for interests.
	DisplayOrder int32 `json:"display_order,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
