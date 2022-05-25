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

// List settings for the outreach
type List10 struct {
	// The unique list id.
	ListId string `json:"list_id,omitempty"`
	// The status of the list used, namely if it's deleted or disabled.
	ListIsActive bool `json:"list_is_active,omitempty"`
	// The name of the list.
	ListName string `json:"list_name,omitempty"`
	// A description of the [segment](https://mailchimp.com/help/save-and-manage-segments/) used for the campaign. Formatted as a string marked up with HTML.
	SegmentText string `json:"segment_text,omitempty"`
	// Count of the recipients on the associated list. Formatted as an integer.
	RecipientCount int32 `json:"recipient_count,omitempty"`
	SegmentOpts *SegmentOptions1 `json:"segment_opts,omitempty"`
}
