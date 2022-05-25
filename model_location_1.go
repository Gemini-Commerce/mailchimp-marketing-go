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

// Subscriber location information.
type Location1 struct {
	// The location latitude.
	Latitude float32 `json:"latitude,omitempty"`
	// The location longitude.
	Longitude float32 `json:"longitude,omitempty"`
	// The time difference in hours from GMT.
	Gmtoff int32 `json:"gmtoff,omitempty"`
	// The offset for timezones where daylight saving time is observed.
	Dstoff int32 `json:"dstoff,omitempty"`
	// The unique code for the location country.
	CountryCode string `json:"country_code,omitempty"`
	// The timezone for the location.
	Timezone string `json:"timezone,omitempty"`
}
