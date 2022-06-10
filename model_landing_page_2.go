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

// A summary of an individual page's properties.
type LandingPage2 struct {
	// The name of this landing page.
	Name string `json:"name,omitempty"`
	// The title of this landing page seen in the browser's title bar.
	Title string `json:"title,omitempty"`
	// The description of this landing page.
	Description string `json:"description,omitempty"`
	// The ID of the store associated with this landing page.
	StoreId string `json:"store_id,omitempty"`
	// The list's ID associated with this landing page.
	ListId string `json:"list_id,omitempty"`
	Tracking *TrackingSettings `json:"tracking,omitempty"`
}
