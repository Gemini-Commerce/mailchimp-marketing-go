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

// The [conditions of the segment](https://mailchimp.com/help/save-and-manage-segments/). Static and fuzzy segments don't have conditions.
type Conditions2 struct {
	// Match type.
	Match string `json:"match,omitempty"`
	// An array of segment conditions.
	Conditions [][]interface{} `json:"conditions,omitempty"`
}
