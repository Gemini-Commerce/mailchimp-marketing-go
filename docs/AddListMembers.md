# AddListMembers

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | Email address for a subscriber. | [optional] [default to null]
**EmailType** | **string** | Type of email this member asked to get (&#x27;html&#x27; or &#x27;text&#x27;). | [optional] [default to null]
**Status** | **string** | Subscriber&#x27;s current status. | [optional] [default to null]
**MergeFields** | [**map[string]interface{}**](interface{}.md) | A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure. | [optional] [default to null]
**Interests** | **map[string]bool** | The key of this object&#x27;s properties is the ID of the interest in question. | [optional] [default to null]
**Language** | **string** | If set/detected, the [subscriber&#x27;s language](https://mailchimp.com/help/view-and-edit-contact-languages/). | [optional] [default to null]
**Vip** | **bool** | [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber. | [optional] [default to null]
**Location** | [***Location**](Location.md) |  | [optional] [default to null]
**IpSignup** | **string** | IP address the subscriber signed up from. | [optional] [default to null]
**TimestampSignup** | [**time.Time**](time.Time.md) | The date and time the subscriber signed up for the list in ISO 8601 format. | [optional] [default to null]
**IpOpt** | **string** | The IP address the subscriber used to confirm their opt-in status. | [optional] [default to null]
**TimestampOpt** | [**time.Time**](time.Time.md) | The date and time the subscribe confirmed their opt-in status in ISO 8601 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

