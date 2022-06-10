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

type InlineResponse20011ReportSummary struct {
	Opens int32 `json:"opens,omitempty"`
	UniqueOpens int32 `json:"unique_opens,omitempty"`
	OpenRate float64 `json:"open_rate,omitempty"`
	Clicks int32 `json:"clicks,omitempty"`
	SubscriberClicks int32 `json:"subscriber_clicks,omitempty"`
	ClickRate float64 `json:"click_rate,omitempty"`
	Visits int32 `json:"visits,omitempty"`
	UniqueVisits int32 `json:"unique_visits,omitempty"`
	ConversionRate float64 `json:"conversion_rate,omitempty"`
	Subscribes int32 `json:"subscribes,omitempty"`
	Ecommerce *InlineResponse20011ReportSummaryEcommerce `json:"ecommerce,omitempty"`
	Impressions float64 `json:"impressions,omitempty"`
	Reach int32 `json:"reach,omitempty"`
	Engagements int32 `json:"engagements,omitempty"`
	TotalSent int32 `json:"total_sent,omitempty"`
}
