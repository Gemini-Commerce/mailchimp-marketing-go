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

// The clicks and visits data from the last five weeks.
type WeeklyClicksAndVisitsData struct {
	// The total number of clicks in a week.
	Clicks []WeeklyClicksAndVisitsDataClicks `json:"clicks,omitempty"`
	// The total number of visits in a week.
	Visits []WeeklyClicksAndVisitsDataVisits `json:"visits,omitempty"`
	UniqueVisits []DailyClicksAndVisitsDataUniqueVisits `json:"unique_visits,omitempty"`
}
