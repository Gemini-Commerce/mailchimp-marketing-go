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

// The tracking options for the Automation.
type AutomationTrackingOptions struct {
	// Whether to [track opens](https://mailchimp.com/help/about-open-tracking/). Defaults to `true`.
	Opens bool `json:"opens,omitempty"`
	// Whether to [track clicks](https://mailchimp.com/help/enable-and-view-click-tracking/) in the HTML version of the Automation. Defaults to `true`.
	HtmlClicks bool `json:"html_clicks,omitempty"`
	// Whether to [track clicks](https://mailchimp.com/help/enable-and-view-click-tracking/) in the plain-text version of the Automation. Defaults to `true`.
	TextClicks bool `json:"text_clicks,omitempty"`
	// Deprecated
	GoalTracking bool `json:"goal_tracking,omitempty"`
	// Whether to enable e-commerce tracking.
	Ecomm360 bool `json:"ecomm360,omitempty"`
	// The custom slug for [Google Analytics](https://mailchimp.com/help/integrate-google-analytics-with-mailchimp/) tracking (max of 50 bytes).
	GoogleAnalytics string `json:"google_analytics,omitempty"`
	// The custom slug for [ClickTale](https://mailchimp.com/help/additional-tracking-options-for-campaigns/) tracking (max of 50 bytes).
	Clicktale string `json:"clicktale,omitempty"`
	Salesforce *SalesforceCrmTracking `json:"salesforce,omitempty"`
	Capsule *CapsuleCrmTracking `json:"capsule,omitempty"`
}
