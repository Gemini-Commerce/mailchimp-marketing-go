# AbTestingOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SplitTest** | **string** | The type of AB split to run. | [optional] [default to null]
**PickWinner** | **string** | How we should evaluate a winner. Based on &#x27;opens&#x27;, &#x27;clicks&#x27;, or &#x27;manual&#x27;. | [optional] [default to null]
**WaitUnits** | **string** | How unit of time for measuring the winner (&#x27;hours&#x27; or &#x27;days&#x27;). This cannot be changed after a campaign is sent. | [optional] [default to null]
**WaitTime** | **int32** | The amount of time to wait before picking a winner. This cannot be changed after a campaign is sent. | [optional] [default to null]
**SplitSize** | **int32** | The size of the split groups. Campaigns split based on &#x27;schedule&#x27; are forced to have a 50/50 split. Valid split integers are between 1-50. | [optional] [default to null]
**FromNameA** | **string** | For campaigns split on &#x27;From Name&#x27;, the name for Group A. | [optional] [default to null]
**FromNameB** | **string** | For campaigns split on &#x27;From Name&#x27;, the name for Group B. | [optional] [default to null]
**ReplyEmailA** | **string** | For campaigns split on &#x27;From Name&#x27;, the reply-to address for Group A. | [optional] [default to null]
**ReplyEmailB** | **string** | For campaigns split on &#x27;From Name&#x27;, the reply-to address for Group B. | [optional] [default to null]
**SubjectA** | **string** | For campaigns split on &#x27;Subject Line&#x27;, the subject line for Group A. | [optional] [default to null]
**SubjectB** | **string** | For campaigns split on &#x27;Subject Line&#x27;, the subject line for Group B. | [optional] [default to null]
**SendTimeA** | [**time.Time**](time.Time.md) | The send time for Group A. | [optional] [default to null]
**SendTimeB** | [**time.Time**](time.Time.md) | The send time for Group B. | [optional] [default to null]
**SendTimeWinner** | **string** | The send time for the winning version. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

