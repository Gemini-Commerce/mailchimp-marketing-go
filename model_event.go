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

import (
	"time"
)

// A specific event for a contact.
type Event struct {
	// The date and time the event occurred in ISO 8601 format.
	OccurredAt time.Time `json:"occurred_at,omitempty"`
	// The name for this type of event ('purchased', 'visited', etc). Must be 2-30 characters in length
	Name string `json:"name,omitempty"`
	// An optional list of properties
	Properties map[string]string `json:"properties,omitempty"`
}
