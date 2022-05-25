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

type InlineResponse2009AudienceEmailSource struct {
	// Email source name
	Name string `json:"name,omitempty"`
	// Type of the email source
	Type_ string `json:"type,omitempty"`
	// Is the source reference a segment
	IsSegment bool `json:"is_segment,omitempty"`
	// Segment type if this source is tied to a segment
	SegmentType string `json:"segment_type,omitempty"`
	// Associated list name to the source
	ListName string `json:"list_name,omitempty"`
}