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

// A single email domain's performance
type EmailDomain struct {
	// The name of the domain (gmail.com, hotmail.com, yahoo.com).
	Domain string `json:"domain,omitempty"`
	// The number of emails sent to that specific domain.
	EmailsSent int32 `json:"emails_sent,omitempty"`
	// The number of bounces at a domain.
	Bounces int32 `json:"bounces,omitempty"`
	// The number of opens for a domain.
	Opens int32 `json:"opens,omitempty"`
	// The number of clicks for a domain.
	Clicks int32 `json:"clicks,omitempty"`
	// The total number of unsubscribes for a domain.
	Unsubs int32 `json:"unsubs,omitempty"`
	// The number of successful deliveries for a domain.
	Delivered int32 `json:"delivered,omitempty"`
	// The percentage of total emails that went to this domain.
	EmailsPct float64 `json:"emails_pct,omitempty"`
	// The percentage of total bounces from this domain.
	BouncesPct float64 `json:"bounces_pct,omitempty"`
	// The percentage of total opens from this domain.
	OpensPct float64 `json:"opens_pct,omitempty"`
	// The percentage of total clicks from this domain.
	ClicksPct float64 `json:"clicks_pct,omitempty"`
	// The percentage of total unsubscribes from this domain.
	UnsubsPct float64 `json:"unsubs_pct,omitempty"`
}
