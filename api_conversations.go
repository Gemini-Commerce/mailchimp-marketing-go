/*
 * Mailchimp Marketing API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0.76
 * Contact: apihelp@mailchimp.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type ConversationsApiService service

/* 
ConversationsApiService List conversations
Get a list of conversations for the account. Conversations has been deprecated in favor of Inbox and these endpoints don&#39;t include Inbox data. Past Conversations are still available via this endpoint, but new campaign replies and other Inbox messages aren’t available using this endpoint.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetConversationsOpts - Optional Parameters:
     * @param "Fields" (optional.Interface of []string) -  A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
     * @param "ExcludeFields" (optional.Interface of []string) -  A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
     * @param "Count" (optional.Int32) -  The number of records to return. Default value is 10. Maximum value is 1000
     * @param "Offset" (optional.Int32) -  Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0.
     * @param "HasUnreadMessages" (optional.String) -  Whether the conversation has any unread messages.
     * @param "ListId" (optional.String) -  The unique id for the list.
     * @param "CampaignId" (optional.String) -  The unique id for the campaign.

@return TrackedConversations
*/

type GetConversationsOpts struct { 
	Fields optional.Interface
	ExcludeFields optional.Interface
	Count optional.Int32
	Offset optional.Int32
	HasUnreadMessages optional.String
	ListId optional.String
	CampaignId optional.String
}

func (a *ConversationsApiService) GetConversations(ctx context.Context, localVarOptionals *GetConversationsOpts) (TrackedConversations, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TrackedConversations
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/conversations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeFields.IsSet() {
		localVarQueryParams.Add("exclude_fields", parameterToString(localVarOptionals.ExcludeFields.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HasUnreadMessages.IsSet() {
		localVarQueryParams.Add("has_unread_messages", parameterToString(localVarOptionals.HasUnreadMessages.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ListId.IsSet() {
		localVarQueryParams.Add("list_id", parameterToString(localVarOptionals.ListId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CampaignId.IsSet() {
		localVarQueryParams.Add("campaign_id", parameterToString(localVarOptionals.CampaignId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v TrackedConversations
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v ProblemDetailDocument
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
ConversationsApiService Get conversation
Get details about an individual conversation. Conversations has been deprecated in favor of Inbox and these endpoints don&#39;t include Inbox data. Past Conversations are still available via this endpoint, but new campaign replies and other Inbox messages aren’t available using this endpoint.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param conversationId The unique id for the conversation.
 * @param optional nil or *GetConversationsIdOpts - Optional Parameters:
     * @param "Fields" (optional.Interface of []string) -  A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
     * @param "ExcludeFields" (optional.Interface of []string) -  A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.

@return Conversation
*/

type GetConversationsIdOpts struct { 
	Fields optional.Interface
	ExcludeFields optional.Interface
}

func (a *ConversationsApiService) GetConversationsId(ctx context.Context, conversationId string, localVarOptionals *GetConversationsIdOpts) (Conversation, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue Conversation
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/conversations/{conversation_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"conversation_id"+"}", fmt.Sprintf("%v", conversationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeFields.IsSet() {
		localVarQueryParams.Add("exclude_fields", parameterToString(localVarOptionals.ExcludeFields.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v Conversation
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v ProblemDetailDocument
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
ConversationsApiService List messages
Get messages from a specific conversation. Conversations has been deprecated in favor of Inbox and these endpoints don&#39;t include Inbox data. Past Conversations are still available via this endpoint, but new campaign replies and other Inbox messages aren’t available using this endpoint.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param conversationId The unique id for the conversation.
 * @param optional nil or *GetConversationsIdMessagesOpts - Optional Parameters:
     * @param "Fields" (optional.Interface of []string) -  A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
     * @param "ExcludeFields" (optional.Interface of []string) -  A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
     * @param "IsRead" (optional.String) -  Whether a conversation message has been marked as read.
     * @param "BeforeTimestamp" (optional.Time) -  Restrict the response to messages created before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
     * @param "SinceTimestamp" (optional.Time) -  Restrict the response to messages created after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.

@return CollectionOfConversationMessages
*/

type GetConversationsIdMessagesOpts struct { 
	Fields optional.Interface
	ExcludeFields optional.Interface
	IsRead optional.String
	BeforeTimestamp optional.Time
	SinceTimestamp optional.Time
}

func (a *ConversationsApiService) GetConversationsIdMessages(ctx context.Context, conversationId string, localVarOptionals *GetConversationsIdMessagesOpts) (CollectionOfConversationMessages, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CollectionOfConversationMessages
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/conversations/{conversation_id}/messages"
	localVarPath = strings.Replace(localVarPath, "{"+"conversation_id"+"}", fmt.Sprintf("%v", conversationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeFields.IsSet() {
		localVarQueryParams.Add("exclude_fields", parameterToString(localVarOptionals.ExcludeFields.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.IsRead.IsSet() {
		localVarQueryParams.Add("is_read", parameterToString(localVarOptionals.IsRead.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BeforeTimestamp.IsSet() {
		localVarQueryParams.Add("before_timestamp", parameterToString(localVarOptionals.BeforeTimestamp.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SinceTimestamp.IsSet() {
		localVarQueryParams.Add("since_timestamp", parameterToString(localVarOptionals.SinceTimestamp.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v CollectionOfConversationMessages
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v ProblemDetailDocument
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
ConversationsApiService Get message
Get an individual message in a conversation. Conversations has been deprecated in favor of Inbox and these endpoints don&#39;t include Inbox data. Past Conversations are still available via this endpoint, but new campaign replies and other Inbox messages aren’t available using this endpoint.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param conversationId The unique id for the conversation.
 * @param messageId The unique id for the conversation message.
 * @param optional nil or *GetConversationsIdMessagesIdOpts - Optional Parameters:
     * @param "Fields" (optional.Interface of []string) -  A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
     * @param "ExcludeFields" (optional.Interface of []string) -  A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.

@return ConversationMessage
*/

type GetConversationsIdMessagesIdOpts struct { 
	Fields optional.Interface
	ExcludeFields optional.Interface
}

func (a *ConversationsApiService) GetConversationsIdMessagesId(ctx context.Context, conversationId string, messageId string, localVarOptionals *GetConversationsIdMessagesIdOpts) (ConversationMessage, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ConversationMessage
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/conversations/{conversation_id}/messages/{message_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"conversation_id"+"}", fmt.Sprintf("%v", conversationId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"message_id"+"}", fmt.Sprintf("%v", messageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeFields.IsSet() {
		localVarQueryParams.Add("exclude_fields", parameterToString(localVarOptionals.ExcludeFields.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v ConversationMessage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v ProblemDetailDocument
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}