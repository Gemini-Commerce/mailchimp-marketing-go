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

// [Default values for campaigns](https://mailchimp.com/help/edit-your-emails-subject-preview-text-from-name-or-from-email-address/) created for this list.
type CampaignDefaults1 struct {
	// The default from name for campaigns sent to this list.
	FromName string `json:"from_name"`
	// The default from email for campaigns sent to this list.
	FromEmail string `json:"from_email"`
	// The default subject line for campaigns sent to this list.
	Subject string `json:"subject"`
	// The default language for this lists's forms.
	Language string `json:"language"`
}