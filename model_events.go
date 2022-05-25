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

// A new event for a specific list member
type Events struct {
	// The name for this type of event ('purchased', 'visited', etc). Must be 2-30 characters in length
	Name string `json:"name"`
	// An optional list of properties
	Properties map[string]string `json:"properties,omitempty"`
	// Events created with the is_syncing value set to `true` will not trigger automations.
	IsSyncing bool `json:"is_syncing,omitempty"`
	// The date and time the event occurred in ISO 8601 format.
	OccurredAt time.Time `json:"occurred_at,omitempty"`
}