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

// The last 50 member events for a list.
type MemberActivityEvents struct {
	// An array of objects, each representing a member event.
	Activity []MemberActivity `json:"activity,omitempty"`
	// The MD5 hash of the lowercase version of the list member's email address.
	EmailId string `json:"email_id,omitempty"`
	// As Mailchimp evolves beyond email, you may eventually have contacts without email addresses. While the `email_id` is the MD5 hash of their email address, this `contact_id` is agnostic of contact’s inclusion of an email address.
	ContactId string `json:"contact_id,omitempty"`
	// The list id.
	ListId string `json:"list_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
