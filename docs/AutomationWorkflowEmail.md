# AutomationWorkflowEmail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that uniquely identifies the Automation email. | [optional] [default to null]
**WebId** | **int32** | The ID used in the Mailchimp web application. View this automation in your Mailchimp account at &#x60;https://{dc}.admin.mailchimp.com/campaigns/show/?id&#x3D;{web_id}&#x60;. | [optional] [default to null]
**WorkflowId** | **string** | A string that uniquely identifies an Automation workflow. | [optional] [default to null]
**Position** | **int32** | The position of an Automation email in a workflow. | [optional] [default to null]
**Delay** | [***AutomationDelay1**](Automation Delay_1.md) |  | [optional] [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | The date and time the campaign was created in ISO 8601 format. | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) | The date and time the campaign was started in ISO 8601 format. | [optional] [default to null]
**ArchiveUrl** | **string** | The link to the campaign&#x27;s archive version in ISO 8601 format. | [optional] [default to null]
**Status** | **string** | The current status of the campaign. | [optional] [default to null]
**EmailsSent** | **int32** | The total number of emails sent for this campaign. | [optional] [default to null]
**SendTime** | [**time.Time**](time.Time.md) |  The date and time a campaign was sent in ISO 8601 format | [optional] [default to null]
**ContentType** | **string** | How the campaign&#x27;s content is put together (&#x27;template&#x27;, &#x27;drag_and_drop&#x27;, &#x27;html&#x27;, &#x27;url&#x27;). | [optional] [default to null]
**NeedsBlockRefresh** | **bool** | Determines if the automation email needs its blocks refreshed by opening the web-based campaign editor. | [optional] [default to null]
**HasLogoMergeTag** | **bool** | Determines if the campaign contains the *|BRAND:LOGO|* merge tag. | [optional] [default to null]
**Recipients** | [***List11**](List_11.md) |  | [optional] [default to null]
**Settings** | [***CampaignSettings5**](Campaign Settings_5.md) |  | [optional] [default to null]
**Tracking** | [***CampaignTrackingOptions1**](Campaign Tracking Options_1.md) |  | [optional] [default to null]
**SocialCard** | [***CampaignSocialCard**](Campaign Social Card.md) |  | [optional] [default to null]
**TriggerSettings** | [***AutomationTrigger1**](Automation Trigger_1.md) |  | [optional] [default to null]
**ReportSummary** | [***CampaignReportSummary3**](Campaign Report Summary_3.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

