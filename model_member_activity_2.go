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

// Member activity events.
type MemberActivity2 struct {
	// The type of action recorded for the subscriber.
	Action string `json:"action,omitempty"`
	// The date and time recorded for the action.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// For clicks, the URL the subscriber clicked on.
	Url string `json:"url,omitempty"`
	// The type of campaign that was sent.
	Type_ string `json:"type,omitempty"`
	// The web-based ID for the campaign.
	CampaignId string `json:"campaign_id,omitempty"`
	// If set, the campaign's title.
	Title string `json:"title,omitempty"`
	// The ID of the parent campaign.
	ParentCampaign string `json:"parent_campaign,omitempty"`
}
