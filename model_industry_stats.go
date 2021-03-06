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

// The average campaign statistics for your industry.
type IndustryStats struct {
	// The type of business industry associated with your account. For example: retail, education, etc.
	Type_ string `json:"type,omitempty"`
	// The industry open rate.
	OpenRate float64 `json:"open_rate,omitempty"`
	// The industry click rate.
	ClickRate float64 `json:"click_rate,omitempty"`
	// The industry bounce rate.
	BounceRate float64 `json:"bounce_rate,omitempty"`
	// The industry unopened rate.
	UnopenRate float64 `json:"unopen_rate,omitempty"`
	// The industry unsubscribe rate.
	UnsubRate float64 `json:"unsub_rate,omitempty"`
	// The industry abuse rate.
	AbuseRate float64 `json:"abuse_rate,omitempty"`
}
