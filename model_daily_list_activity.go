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

// One day's worth of list activity. Doesn't include Automation activity.
type DailyListActivity struct {
	// The date for the activity summary.
	Day string `json:"day,omitempty"`
	// The total number of emails sent on the date for the activity summary.
	EmailsSent int32 `json:"emails_sent,omitempty"`
	// The number of unique opens.
	UniqueOpens int32 `json:"unique_opens,omitempty"`
	// The number of clicks.
	RecipientClicks int32 `json:"recipient_clicks,omitempty"`
	// The number of hard bounces.
	HardBounce int32 `json:"hard_bounce,omitempty"`
	// The number of soft bounces
	SoftBounce int32 `json:"soft_bounce,omitempty"`
	// The number of subscribes.
	Subs int32 `json:"subs,omitempty"`
	// The number of unsubscribes.
	Unsubs int32 `json:"unsubs,omitempty"`
	// The number of subscribers who may have been added outside of the [double opt-in process](https://mailchimp.com/help/about-double-opt-in/), such as imports or API activity.
	OtherAdds int32 `json:"other_adds,omitempty"`
	// The number of subscribers who may have been removed outside of unsubscribing or reporting an email as spam (for example, deleted subscribers).
	OtherRemoves int32 `json:"other_removes,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
