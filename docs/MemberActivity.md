# MemberActivity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | One of the following actions: &#x27;open&#x27;, &#x27;click&#x27;, or &#x27;bounce&#x27; | [optional] [default to null]
**Type_** | **string** | If the action is a &#x27;bounce&#x27;, the type of bounce received: &#x27;hard&#x27;, &#x27;soft&#x27;. | [optional] [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The date and time recorded for the action in ISO 8601 format. | [optional] [default to null]
**Url** | **string** | If the action is a &#x27;click&#x27;, the URL on which the member clicked. | [optional] [default to null]
**Ip** | **string** | The IP address recorded for the action. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

