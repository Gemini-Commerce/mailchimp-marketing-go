# SubscriberList

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the list. | [default to null]
**Contact** | [***ListContact**](List Contact.md) |  | [default to null]
**PermissionReminder** | **string** | The [permission reminder](https://mailchimp.com/help/edit-the-permission-reminder/) for the list. | [default to null]
**UseArchiveBar** | **bool** | Whether campaigns for this list use the [Archive Bar](https://mailchimp.com/help/about-email-campaign-archives-and-pages/) in archives by default. | [optional] [default to false]
**CampaignDefaults** | [***CampaignDefaults**](Campaign Defaults.md) |  | [default to null]
**NotifyOnSubscribe** | **string** | The email address to send [subscribe notifications](https://mailchimp.com/help/change-subscribe-and-unsubscribe-notifications/) to. | [optional] [default to false]
**NotifyOnUnsubscribe** | **string** | The email address to send [unsubscribe notifications](https://mailchimp.com/help/change-subscribe-and-unsubscribe-notifications/) to. | [optional] [default to false]
**EmailTypeOption** | **bool** | Whether the list supports [multiple formats for emails](https://mailchimp.com/help/change-audience-name-defaults/). When set to &#x60;true&#x60;, subscribers can choose whether they want to receive HTML or plain-text emails. When set to &#x60;false&#x60;, subscribers will receive HTML emails, with a plain-text alternative backup. | [default to null]
**DoubleOptin** | **bool** | Whether or not to require the subscriber to confirm subscription via email. | [optional] [default to false]
**MarketingPermissions** | **bool** | Whether or not the list has marketing permissions (eg. GDPR) enabled. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

