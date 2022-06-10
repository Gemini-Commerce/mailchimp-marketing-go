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

// A list of members who have unsubscribed from a specific campaign.
type Unsubscribes struct {
	// An array of objects, each representing a member who unsubscribed from a campaign.
	Unsubscribes []Unsubscribes2 `json:"unsubscribes,omitempty"`
	// The campaign id.
	CampaignId string `json:"campaign_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
