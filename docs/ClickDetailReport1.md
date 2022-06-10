# ClickDetailReport1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id for the link. | [optional] [default to null]
**Url** | **string** | The URL for the link in the campaign. | [optional] [default to null]
**TotalClicks** | **int32** | The number of total clicks for a link. | [optional] [default to null]
**ClickPercentage** | **float64** | The percentage of total clicks a link generated for a campaign. | [optional] [default to null]
**UniqueClicks** | **int32** | Number of unique clicks for a link. | [optional] [default to null]
**UniqueClickPercentage** | **float64** | The percentage of unique clicks a link generated for a campaign. | [optional] [default to null]
**LastClick** | [**time.Time**](time.Time.md) | The date and time for the last recorded click for a link in ISO 8601 format. | [optional] [default to null]
**AbSplit** | [***AbSplit**](AB Split.md) |  | [optional] [default to null]
**CampaignId** | **string** | The campaign id. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

