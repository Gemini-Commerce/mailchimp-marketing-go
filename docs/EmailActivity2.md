# EmailActivity2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **string** | The unique id for the campaign. | [optional] [default to null]
**ListId** | **string** | The unique id for the list. | [optional] [default to null]
**ListIsActive** | **bool** | The status of the list used, namely if it&#x27;s deleted or disabled. | [optional] [default to null]
**EmailId** | **string** | The MD5 hash of the lowercase version of the list member&#x27;s email address. | [optional] [default to null]
**EmailAddress** | **string** | Email address for a subscriber. | [optional] [default to null]
**Activity** | [**[]MemberActivity**](Member Activity.md) | An array of objects, each showing an interaction with the email. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

