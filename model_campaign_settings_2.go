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

// The settings for your campaign, including subject, from name, reply-to address, and more.
type CampaignSettings2 struct {
	// The subject line for the campaign.
	SubjectLine string `json:"subject_line"`
	// The preview text for the campaign.
	PreviewText string `json:"preview_text,omitempty"`
	// The title of the campaign.
	Title string `json:"title,omitempty"`
	// The 'from' name on the campaign (not an email address).
	FromName string `json:"from_name"`
	// The reply-to email address for the campaign.
	ReplyTo string `json:"reply_to"`
	// Use Mailchimp Conversation feature to manage out-of-office replies.
	UseConversation bool `json:"use_conversation,omitempty"`
	// The campaign's custom 'To' name. Typically the first name [audience field](https://mailchimp.com/help/getting-started-with-merge-tags/).
	ToName string `json:"to_name,omitempty"`
	// If the campaign is listed in a folder, the id for that folder.
	FolderId string `json:"folder_id,omitempty"`
	// Whether Mailchimp [authenticated](https://mailchimp.com/help/about-email-authentication/) the campaign. Defaults to `true`.
	Authenticate bool `json:"authenticate,omitempty"`
	// Automatically append Mailchimp's [default footer](https://mailchimp.com/help/about-campaign-footers/) to the campaign.
	AutoFooter bool `json:"auto_footer,omitempty"`
	// Automatically inline the CSS included with the campaign content.
	InlineCss bool `json:"inline_css,omitempty"`
	// Automatically tweet a link to the [campaign archive](https://mailchimp.com/help/about-email-campaign-archives-and-pages/) page when the campaign is sent.
	AutoTweet bool `json:"auto_tweet,omitempty"`
	// An array of [Facebook](https://mailchimp.com/help/connect-or-disconnect-the-facebook-integration/) page ids to auto-post to.
	AutoFbPost []string `json:"auto_fb_post,omitempty"`
	// Allows Facebook comments on the campaign (also force-enables the Campaign Archive toolbar). Defaults to `true`.
	FbComments bool `json:"fb_comments,omitempty"`
	// The id of the template to use.
	TemplateId int32 `json:"template_id,omitempty"`
}
