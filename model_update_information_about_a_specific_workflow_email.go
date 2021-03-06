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

// Update information about an individual Automation workflow email.
type UpdateInformationAboutASpecificWorkflowEmail struct {
	Settings *CampaignSettings `json:"settings,omitempty"`
	Delay *AutomationDelay `json:"delay,omitempty"`
}
