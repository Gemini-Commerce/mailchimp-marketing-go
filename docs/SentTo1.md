# SentTo1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailId** | **string** | The MD5 hash of the lowercase version of the list member&#x27;s email address. | [optional] [default to null]
**EmailAddress** | **string** | Email address for a subscriber. | [optional] [default to null]
**MergeFields** | [**map[string]interface{}**](interface{}.md) | A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure. | [optional] [default to null]
**Vip** | **bool** | [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber. | [optional] [default to null]
**Status** | **string** | The status of the member (&#x27;sent&#x27;, &#x27;hard&#x27; for hard bounce, or &#x27;soft&#x27; for soft bounce). | [optional] [default to null]
**OpenCount** | **int32** | The number of times a campaign was opened by this member. | [optional] [default to null]
**LastOpen** | [**time.Time**](time.Time.md) | The date and time of the last open for this member in ISO 8601 format. | [optional] [default to null]
**AbsplitGroup** | **string** | For A/B Split Campaigns, the group the member was apart of (&#x27;a&#x27;, &#x27;b&#x27;, or &#x27;winner&#x27;). | [optional] [default to null]
**GmtOffset** | **int32** | For campaigns sent with timewarp, the time zone group the member is apart of. | [optional] [default to null]
**CampaignId** | **string** | The campaign id. | [optional] [default to null]
**ListId** | **string** | The unique list id. | [optional] [default to null]
**ListIsActive** | **bool** | The status of the list used, namely if it&#x27;s deleted or disabled. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

