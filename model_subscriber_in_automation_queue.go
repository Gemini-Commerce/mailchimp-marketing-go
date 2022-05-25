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

// Information about subscribers in an Automation email queue.
type SubscriberInAutomationQueue struct {
	// The MD5 hash of the lowercase version of the list member's email address.
	Id string `json:"id,omitempty"`
	// A string that uniquely identifies an Automation workflow.
	WorkflowId string `json:"workflow_id,omitempty"`
	// A string that uniquely identifies an email in an Automation workflow.
	EmailId string `json:"email_id,omitempty"`
	// A string that uniquely identifies a list.
	ListId string `json:"list_id,omitempty"`
	// The list member's email address.
	EmailAddress string `json:"email_address"`
	// The date and time of the next send for the workflow email in ISO 8601 format.
	NextSend time.Time `json:"next_send,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
