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

// A collection of landing pages.
type InlineResponse2008 struct {
	// The landing pages on the account
	LandingPages []LandingPage3 `json:"landing_pages,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
