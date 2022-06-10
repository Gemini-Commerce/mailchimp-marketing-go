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

// The schedule for sending the RSS Campaign.
type SendingSchedule1 struct {
	// The hour to send the campaign in local time. Acceptable hours are 0-23. For example, '4' would be 4am in [your account's default time zone](https://mailchimp.com/help/set-account-details/).
	Hour int32 `json:"hour,omitempty"`
	DailySend *DailySendingDays `json:"daily_send,omitempty"`
	// The day of the week to send a weekly RSS Campaign.
	WeeklySendDay string `json:"weekly_send_day,omitempty"`
	// The day of the month to send a monthly RSS Campaign. Acceptable days are 0-31, where '0' is always the last day of a month. Months with fewer than the selected number of days will not have an RSS campaign sent out that day. For example, RSS Campaigns set to send on the 30th will not go out in February.
	MonthlySendDate float64 `json:"monthly_send_date,omitempty"`
}
