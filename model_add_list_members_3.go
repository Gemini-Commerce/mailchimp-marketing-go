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

import (
	"time"
)

// Individuals who are currently or have been previously subscribed to this list, including members who have bounced or unsubscribed.
type AddListMembers3 struct {
	// Email address for a subscriber.
	EmailAddress string `json:"email_address,omitempty"`
	// Type of email this member asked to get ('html' or 'text').
	EmailType string `json:"email_type,omitempty"`
	// Subscriber's current status.
	Status string `json:"status,omitempty"`
	// A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.
	MergeFields map[string]interface{} `json:"merge_fields,omitempty"`
	// The key of this object's properties is the ID of the interest in question.
	Interests map[string]bool `json:"interests,omitempty"`
	// If set/detected, the [subscriber's language](https://mailchimp.com/help/view-and-edit-contact-languages/).
	Language string `json:"language,omitempty"`
	// [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.
	Vip bool `json:"vip,omitempty"`
	Location *Location `json:"location,omitempty"`
	// The marketing permissions for the subscriber.
	MarketingPermissions []MarketingPermission1 `json:"marketing_permissions,omitempty"`
	// IP address the subscriber signed up from.
	IpSignup string `json:"ip_signup,omitempty"`
	// The date and time the subscriber signed up for the list in ISO 8601 format.
	TimestampSignup time.Time `json:"timestamp_signup,omitempty"`
	// The IP address the subscriber used to confirm their opt-in status.
	IpOpt string `json:"ip_opt,omitempty"`
	// The date and time the subscriber confirmed their opt-in status in ISO 8601 format.
	TimestampOpt time.Time `json:"timestamp_opt,omitempty"`
}
