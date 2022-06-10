# Unsubscribes2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailId** | **string** | The MD5 hash of the lowercase version of the list member&#x27;s email address. | [optional] [default to null]
**EmailAddress** | **string** | Email address for a subscriber. | [optional] [default to null]
**MergeFields** | [**map[string]interface{}**](interface{}.md) | A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure. | [optional] [default to null]
**Vip** | **bool** | [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber. | [optional] [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The date and time the member opted-out in ISO 8601 format. | [optional] [default to null]
**Reason** | **string** | If available, the reason listed by the member for unsubscribing. | [optional] [default to null]
**CampaignId** | **string** | The campaign id. | [optional] [default to null]
**ListId** | **string** | The list id. | [optional] [default to null]
**ListIsActive** | **bool** | The status of the list used, namely if it&#x27;s deleted or disabled. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

