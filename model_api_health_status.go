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

// API health status.
type ApiHealthStatus struct {
	// This will return a constant string value if the request is successful. Ex. \"Everything's Chimpy!\"
	HealthStatus string `json:"health_status,omitempty"`
}
