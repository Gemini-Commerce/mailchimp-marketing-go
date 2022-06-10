# ListMembers2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The MD5 hash of the lowercase version of the list member&#x27;s email address. | [optional] [default to null]
**EmailAddress** | **string** | Email address for a subscriber. | [optional] [default to null]
**UniqueEmailId** | **string** | An identifier for the address across all of Mailchimp. | [optional] [default to null]
**ContactId** | **string** | As Mailchimp evolves beyond email, you may eventually have contacts without email addresses. While the &#x60;id&#x60; is the MD5 hash of their email address, this &#x60;contact_id&#x60; is agnostic of contact’s inclusion of an email address. | [optional] [default to null]
**FullName** | **string** | The contact&#x27;s full name. | [optional] [default to null]
**WebId** | **int32** | The ID used in the Mailchimp web application. View this member in your Mailchimp account at &#x60;https://{dc}.admin.mailchimp.com/lists/members/view?id&#x3D;{web_id}&#x60;. | [optional] [default to null]
**EmailType** | **string** | Type of email this member asked to get (&#x27;html&#x27; or &#x27;text&#x27;). | [optional] [default to null]
**Status** | **string** | Subscriber&#x27;s current status. | [optional] [default to null]
**UnsubscribeReason** | **string** | A subscriber&#x27;s reason for unsubscribing. | [optional] [default to null]
**ConsentsToOneToOneMessaging** | **bool** | Indicates whether a contact consents to 1:1 messaging. | [optional] [default to null]
**MergeFields** | [**map[string]interface{}**](interface{}.md) | A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure. | [optional] [default to null]
**Interests** | **map[string]bool** | The key of this object&#x27;s properties is the ID of the interest in question. | [optional] [default to null]
**Stats** | [***SubscriberStats1**](Subscriber Stats_1.md) |  | [optional] [default to null]
**IpSignup** | **string** | IP address the subscriber signed up from. | [optional] [default to null]
**TimestampSignup** | [**time.Time**](time.Time.md) | The date and time the subscriber signed up for the list in ISO 8601 format. | [optional] [default to null]
**IpOpt** | **string** | The IP address the subscriber used to confirm their opt-in status. | [optional] [default to null]
**TimestampOpt** | [**time.Time**](time.Time.md) | The date and time the subscribe confirmed their opt-in status in ISO 8601 format. | [optional] [default to null]
**MemberRating** | **int32** | Star rating for this member, between 1 and 5. | [optional] [default to null]
**LastChanged** | [**time.Time**](time.Time.md) | The date and time the member&#x27;s info was last changed in ISO 8601 format. | [optional] [default to null]
**Language** | **string** | If set/detected, the [subscriber&#x27;s language](https://mailchimp.com/help/view-and-edit-contact-languages/). | [optional] [default to null]
**Vip** | **bool** | [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber. | [optional] [default to null]
**EmailClient** | **string** | The list member&#x27;s email client. | [optional] [default to null]
**Location** | [***Location3**](Location_3.md) |  | [optional] [default to null]
**MarketingPermissions** | [**[]MarketingPermission1**](Marketing Permission_1.md) | The marketing permissions for the subscriber. | [optional] [default to null]
**LastNote** | [***Notes**](Notes.md) |  | [optional] [default to null]
**Source** | **string** | The source from which the subscriber was added to this list. | [optional] [default to null]
**TagsCount** | **int32** | The number of tags applied to this member. | [optional] [default to null]
**Tags** | [**[]ListMembers1Tags**](List Members_1_tags.md) | Returns up to 50 tags applied to this member. To retrieve all tags see [Member Tags](https://mailchimp.com/developer/marketing/api/list-member-tags/). | [optional] [default to null]
**ListId** | **string** | The list id. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

