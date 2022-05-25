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

// The settings specific to A/B test campaigns.
type AbTestOptions2 struct {
	// ID for the winning combination.
	WinningCombinationId string `json:"winning_combination_id,omitempty"`
	// ID of the campaign that was sent to the remaining recipients based on the winning combination.
	WinningCampaignId string `json:"winning_campaign_id,omitempty"`
	// The combination that performs the best. This may be determined automatically by click rate, open rate, or total revenue -- or you may choose manually based on the reporting data you find the most valuable. For Multivariate Campaigns testing send_time, winner_criteria is ignored. For Multivariate Campaigns with 'manual' as the winner_criteria, the winner must be chosen in the Mailchimp web application.
	WinnerCriteria string `json:"winner_criteria"`
	// The number of minutes to wait before choosing the winning campaign. The value of wait_time must be greater than 0 and in whole hours, specified in minutes.
	WaitTime int32 `json:"wait_time,omitempty"`
	// The percentage of recipients to send the test combinations to, must be a value between 10 and 100.
	TestSize int32 `json:"test_size,omitempty"`
	// The possible subject lines to test. If no subject lines are provided, settings.subject_line will be used.
	SubjectLines []string `json:"subject_lines,omitempty"`
	// The possible send times to test. The times provided should be in the format YYYY-MM-DD HH:MM:SS. If send_times are provided to test, the test_size will be set to 100% and winner_criteria will be ignored.
	SendTimes []time.Time `json:"send_times,omitempty"`
	// The possible from names. The number of from_names provided must match the number of reply_to_addresses. If no from_names are provided, settings.from_name will be used.
	FromNames []string `json:"from_names,omitempty"`
	// The possible reply-to addresses. The number of reply_to_addresses provided must match the number of from_names. If no reply_to_addresses are provided, settings.reply_to will be used.
	ReplyToAddresses []string `json:"reply_to_addresses,omitempty"`
	// Descriptions of possible email contents. To set campaign contents, make a PUT request to /campaigns/{campaign_id}/content with the field 'variate_contents'.
	Contents []string `json:"contents,omitempty"`
	// Combinations of possible variables used to build emails.
	Combinations []AbTestOptionsCombinations `json:"combinations,omitempty"`
}
