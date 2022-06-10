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

// This object represents a link from the resource where it is found to another resource or action that may be performed.
type ResourceLink struct {
	// As with an HTML 'rel' attribute, this describes the type of link.
	Rel string `json:"rel,omitempty"`
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href string `json:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method string `json:"method,omitempty"`
	// For GETs, this is a URL representing the schema that the response should conform to.
	TargetSchema string `json:"targetSchema,omitempty"`
	// For HTTP methods that can receive bodies (POST and PUT), this is a URL representing the schema that the body should conform to.
	Schema string `json:"schema,omitempty"`
}
