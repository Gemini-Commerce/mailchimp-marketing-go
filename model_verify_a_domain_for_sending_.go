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

// Submit a response to the verification challenge and verify a domain for sending.
type VerifyADomainForSending_ struct {
	// The code that was sent to the email address provided when adding a new domain to verify.
	Code string `json:"code"`
}