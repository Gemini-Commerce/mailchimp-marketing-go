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

type AllOfinlineResponse20010FacebookAdsItems struct {
	EmailSourceName string `json:"email_source_name,omitempty"`
	PausedAt time.Time `json:"paused_at,omitempty"`
	EndTime time.Time `json:"end_time,omitempty"`
	NeedsAttention bool `json:"needs_attention,omitempty"`
	WasCanceledByFacebook bool `json:"was_canceled_by_facebook,omitempty"`
	// Check if this ad is connected to a facebook page
	IsConnected bool `json:"is_connected,omitempty"`
	// Check if this ad has audience setup
	HasAudience bool `json:"has_audience,omitempty"`
	// Check if this ad has content
	HasContent bool `json:"has_content,omitempty"`
	// Channel settings
	Channel *interface{} `json:"channel,omitempty"`
	// Check if this ad is connected to a facebook page
	Feedback *interface{} `json:"feedback,omitempty"`
	// Connected Site
	Site *interface{} `json:"site,omitempty"`
	// Audience settings
	Audience *interface{} `json:"audience,omitempty"`
	Budget *interface{} `json:"budget,omitempty"`
	Content *interface{} `json:"content,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []interface{} `json:"_links,omitempty"`
}
