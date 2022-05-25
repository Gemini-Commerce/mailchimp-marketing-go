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

// The events that can trigger the webhook and whether they are enabled.
type Events1 struct {
	// Whether the webhook is triggered when a list subscriber is added.
	Subscribe bool `json:"subscribe,omitempty"`
	// Whether the webhook is triggered when a list member unsubscribes.
	Unsubscribe bool `json:"unsubscribe,omitempty"`
	// Whether the webhook is triggered when a subscriber's profile is updated.
	Profile bool `json:"profile,omitempty"`
	// Whether the webhook is triggered when a subscriber's email address is cleaned from the list.
	Cleaned bool `json:"cleaned,omitempty"`
	// Whether the webhook is triggered when a subscriber's email address is changed.
	Upemail bool `json:"upemail,omitempty"`
	// Whether the webhook is triggered when a campaign is sent or cancelled.
	Campaign bool `json:"campaign,omitempty"`
}
