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

// Deprecated
type SalesforceCrmTracking struct {
	// Create a campaign in a connected Salesforce account.
	Campaign bool `json:"campaign,omitempty"`
	// Update contact notes for a campaign based on subscriber email addresses.
	Notes bool `json:"notes,omitempty"`
}
