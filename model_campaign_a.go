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

// Stats for Campaign A.
type CampaignA struct {
	// Bounces for Campaign A.
	Bounces int32 `json:"bounces,omitempty"`
	// Abuse reports for Campaign A.
	AbuseReports int32 `json:"abuse_reports,omitempty"`
	// Unsubscribes for Campaign A.
	Unsubs int32 `json:"unsubs,omitempty"`
	// Recipient Clicks for Campaign A.
	RecipientClicks int32 `json:"recipient_clicks,omitempty"`
	// Forwards for Campaign A.
	Forwards int32 `json:"forwards,omitempty"`
	// Opens from forwards for Campaign A.
	ForwardsOpens int32 `json:"forwards_opens,omitempty"`
	// Opens for Campaign A.
	Opens int32 `json:"opens,omitempty"`
	// The last open for Campaign A.
	LastOpen string `json:"last_open,omitempty"`
	// Unique opens for Campaign A.
	UniqueOpens int32 `json:"unique_opens,omitempty"`
}
