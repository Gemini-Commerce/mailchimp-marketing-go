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

// A summary of the emails in an Automation workflow.
type AutomationEmails struct {
	// An array of objects, each representing an email in an Automation workflow.
	Emails []AutomationWorkflowEmail1 `json:"emails,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
