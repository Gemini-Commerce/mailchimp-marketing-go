# AutomationWorkflow2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that identifies the Automation. | [optional] [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | The date and time the Automation was created in ISO 8601 format. | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) | The date and time the Automation was started in ISO 8601 format. | [optional] [default to null]
**Status** | **string** | The current status of the Automation. | [optional] [default to null]
**EmailsSent** | **int32** | The total number of emails sent for the Automation. | [optional] [default to null]
**Recipients** | [***List9**](List_9.md) |  | [optional] [default to null]
**Settings** | [***AutomationCampaignSettings1**](Automation Campaign Settings_1.md) |  | [optional] [default to null]
**Tracking** | [***AutomationTrackingOptions**](Automation Tracking Options.md) |  | [optional] [default to null]
**TriggerSettings** | [***AutomationTrigger1**](Automation Trigger_1.md) |  | [optional] [default to null]
**ReportSummary** | [***CampaignReportSummary2**](Campaign Report Summary_2.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

