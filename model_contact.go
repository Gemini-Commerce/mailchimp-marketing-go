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

// Information about the contact.
type Contact struct {
	// The MD5 hash of the lowercase version of the list member's email address.
	EmailId string `json:"email_id,omitempty"`
	// The ID of this contact.
	ContactId string `json:"contact_id,omitempty"`
	// The contact's current status.
	Status string `json:"status,omitempty"`
	// The contact's email address.
	Email string `json:"email,omitempty"`
	// The contact's full name.
	FullName string `json:"full_name,omitempty"`
	// Indicates whether a contact consents to 1:1 messaging.
	ConsentsToOneToOneMessaging bool `json:"consents_to_one_to_one_messaging,omitempty"`
	// URL for the contact's avatar or profile image.
	AvatarUrl string `json:"avatar_url,omitempty"`
}