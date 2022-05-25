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

// abandonedBrowse automation details. abandonedBrowse is also known as Product Retargeting Email or Retarget Site Visitors on the web.
type AbandonedBrowseAutomation struct {
	// Whether this store supports the abandonedBrowse automation.
	IsSupported bool `json:"is_supported,omitempty"`
	// Unique ID of automation parent campaign.
	Id string `json:"id,omitempty"`
	// Status of the abandonedBrowse automation.
	Status string `json:"status,omitempty"`
}
