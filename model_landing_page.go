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

// A summary of an individual landing page's settings and content.
type LandingPage struct {
	// A string that uniquely identifies this landing page.
	Id string `json:"id,omitempty"`
	// The name of this landing page.
	Name string `json:"name,omitempty"`
	// The title of this landing page seen in the browser's title bar.
	Title string `json:"title,omitempty"`
	// The description of this landing page.
	Description string `json:"description,omitempty"`
	// The template_id of this landing page.
	TemplateId int32 `json:"template_id,omitempty"`
	// The status of this landing page.
	Status string `json:"status,omitempty"`
	// The list's ID associated with this landing page.
	ListId string `json:"list_id,omitempty"`
	// The ID of the store associated with this landing page.
	StoreId string `json:"store_id,omitempty"`
	// The ID used in the Mailchimp web application.
	WebId int32 `json:"web_id,omitempty"`
	// Created by mobile or web
	CreatedBySource string `json:"created_by_source,omitempty"`
	// The url of the published landing page.
	Url string `json:"url,omitempty"`
	// The time this landing page was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The time this landing page was published.
	PublishedAt time.Time `json:"published_at,omitempty"`
	// The time this landing page was unpublished.
	UnpublishedAt time.Time `json:"unpublished_at,omitempty"`
	// The time this landing page was updated at.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Tracking *TrackingSettings `json:"tracking,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}