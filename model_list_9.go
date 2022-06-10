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

// List settings for the Automation.
type List9 struct {
	// The unique list id.
	ListId string `json:"list_id,omitempty"`
	// The status of the list used, namely if it's deleted or disabled.
	ListIsActive bool `json:"list_is_active,omitempty"`
	// List Name.
	ListName string `json:"list_name,omitempty"`
	SegmentOpts *SegmentOptions2 `json:"segment_opts,omitempty"`
	// The id of the store.
	StoreId string `json:"store_id,omitempty"`
}
