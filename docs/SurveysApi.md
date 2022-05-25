# \SurveysApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostListsIdSurveysIdActionsPublish**](SurveysApi.md#PostListsIdSurveysIdActionsPublish) | **Post** /lists/{list_id}/surveys/{survey_id}/actions/publish | Publish a Survey
[**PostListsIdSurveysIdActionsUnpublish**](SurveysApi.md#PostListsIdSurveysIdActionsUnpublish) | **Post** /lists/{list_id}/surveys/{survey_id}/actions/unpublish | Unpublish a Survey


# **PostListsIdSurveysIdActionsPublish**
> PostListsIdSurveysIdActionsPublish(ctx, listId, surveyId)
Publish a Survey

Publish a survey that is in draft, unpublished, or has been previously published and edited.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **surveyId** | **string**| The ID of the survey. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsIdSurveysIdActionsUnpublish**
> PostListsIdSurveysIdActionsUnpublish(ctx, listId, surveyId)
Unpublish a Survey

Unpublish a survey that has been published.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **surveyId** | **string**| The ID of the survey. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

