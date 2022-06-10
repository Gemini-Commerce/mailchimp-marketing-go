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
import (
	"time"
)

// A specific feedback message from a specific campaign.
type CampaignFeedback3 struct {
	// The individual id for the feedback item.
	FeedbackId int32 `json:"feedback_id,omitempty"`
	// If a reply, the id of the parent feedback item.
	ParentId int32 `json:"parent_id,omitempty"`
	// The block id for the editable block that the feedback addresses.
	BlockId int32 `json:"block_id,omitempty"`
	// The content of the feedback.
	Message string `json:"message"`
	// The status of feedback.
	IsComplete bool `json:"is_complete,omitempty"`
	// The login name of the user who created the feedback.
	CreatedBy string `json:"created_by,omitempty"`
	// The date and time the feedback item was created in ISO 8601 format.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The date and time the feedback was last updated in ISO 8601 format.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The source of the feedback.
	Source string `json:"source,omitempty"`
	// The unique id for the campaign.
	CampaignId string `json:"campaign_id,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
