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

// A subscriber who clicked a specific URL in a specific campaign.
type ClickDetailMember struct {
	// The MD5 hash of the lowercase version of the list member's email address.
	EmailId string `json:"email_id,omitempty"`
	// Email address for a subscriber.
	EmailAddress string `json:"email_address,omitempty"`
	// A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.
	MergeFields map[string]interface{} `json:"merge_fields,omitempty"`
	// [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.
	Vip bool `json:"vip,omitempty"`
	// The total number of times the subscriber clicked on the link.
	Clicks int32 `json:"clicks,omitempty"`
	// The campaign id.
	CampaignId string `json:"campaign_id,omitempty"`
	// The id for the tracked URL in the campaign.
	UrlId string `json:"url_id,omitempty"`
	// The list id.
	ListId string `json:"list_id,omitempty"`
	// The status of the list used, namely if it's deleted or disabled.
	ListIsActive bool `json:"list_is_active,omitempty"`
	// The status of the member, namely if they are subscribed, unsubscribed, deleted, non-subscribed, transactional, pending, or need reconfirmation.
	ContactStatus string `json:"contact_status,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
