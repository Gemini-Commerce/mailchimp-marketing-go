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

// An individual message in a conversation. Conversation tracking is a feature available to paid accounts that lets you view replies to your campaigns in your Mailchimp account.
type ConversationMessage struct {
	// A string that uniquely identifies this message
	Id string `json:"id,omitempty"`
	// A string that identifies this message's conversation
	ConversationId string `json:"conversation_id,omitempty"`
	// The list's web ID
	ListId int32 `json:"list_id,omitempty"`
	// A label representing the sender of this message
	FromLabel string `json:"from_label,omitempty"`
	// A label representing the email of the sender of this message
	FromEmail string `json:"from_email,omitempty"`
	// The subject of this message
	Subject string `json:"subject,omitempty"`
	// The plain-text content of the message
	Message string `json:"message,omitempty"`
	// Whether this message has been marked as read
	Read bool `json:"read,omitempty"`
	// The date and time the message was either sent or received in ISO 8601 format.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}