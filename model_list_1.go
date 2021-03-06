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

// List settings for the campaign.
type List1 struct {
	// The unique list id.
	ListId string `json:"list_id"`
	SegmentOpts *SegmentOptions `json:"segment_opts,omitempty"`
}
