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

// Trigger settings for the Automation.
type AutomationTrigger struct {
	// The type of Automation workflow. Currently only supports 'abandonedCart'.
	WorkflowType string `json:"workflow_type"`
}
