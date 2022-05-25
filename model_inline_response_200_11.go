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

type InlineResponse20011 struct {
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
	Audience *InlineResponse2009Audience `json:"audience,omitempty"`
	AudienceActivity *InlineResponse20011AudienceActivity `json:"audience_activity,omitempty"`
	Budget *InlineResponse2009Budget `json:"budget,omitempty"`
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	Channel *InlineResponse2009Channel `json:"channel,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	EmailSourceName string `json:"email_source_name,omitempty"`
	EndTime time.Time `json:"end_time,omitempty"`
	HasSegment bool `json:"has_segment,omitempty"`
	// Unique ID of an Outreach
	Id string `json:"id,omitempty"`
	// Title or name of an Outreach
	Name string `json:"name,omitempty"`
	NeedsAttention bool `json:"needs_attention,omitempty"`
	PausedAt time.Time `json:"paused_at,omitempty"`
	PublishedTime time.Time `json:"published_time,omitempty"`
	Recipients *List10 `json:"recipients,omitempty"`
	ReportSummary *InlineResponse20011ReportSummary `json:"report_summary,omitempty"`
	// Outreach report availability
	ShowReport bool `json:"show_report,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	// Campaign, Ad, or Page status
	Status string `json:"status,omitempty"`
	// The URL of the thumbnail for this outreach
	Thumbnail string `json:"thumbnail,omitempty"`
	// Supported Campaign, Ad, Page type
	Type_ string `json:"type,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	WasCanceledByFacebook bool `json:"was_canceled_by_facebook,omitempty"`
	// Web ID
	WebId int32 `json:"web_id,omitempty"`
}
