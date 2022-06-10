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

// A summary of Twitter activity for a campaign.
type TwitterStats struct {
	// The number of tweets including a link to the campaign.
	Tweets int32 `json:"tweets,omitempty"`
	// The day and time of the first recorded tweet with a link to the campaign.
	FirstTweet string `json:"first_tweet,omitempty"`
	// The day and time of the last recorded tweet with a link to the campaign.
	LastTweet string `json:"last_tweet,omitempty"`
	// The number of retweets that include a link to the campaign.
	Retweets int32 `json:"retweets,omitempty"`
	// A summary of tweets that include a link to the campaign.
	Statuses []TwitterStatus `json:"statuses,omitempty"`
}
