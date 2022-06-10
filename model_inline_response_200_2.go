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

// An account export.
type InlineResponse2002 struct {
	// The ID for the export.
	ExportId int32 `json:"export_id,omitempty"`
	// Start time for the export.
	Started time.Time `json:"started,omitempty"`
	// If finished, the finish time for the export.
	Finished time.Time `json:"finished,omitempty"`
	// The size of the uncompressed export in bytes.
	SizeInBytes int32 `json:"size_in_bytes,omitempty"`
	// If the export is finished, the download URL for an export. URLs are only valid for 90 days after the export completes.
	DownloadUrl string `json:"download_url,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
