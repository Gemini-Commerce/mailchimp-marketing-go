# EcommerceCustomer4

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the customer. Limited to 50 characters. | [default to null]
**EmailAddress** | **string** | The customer&#x27;s email address. | [default to null]
**OptInStatus** | **bool** | The customer&#x27;s opt-in status. This value will never overwrite the opt-in status of a pre-existing Mailchimp list member, but will apply to list members that are added through the e-commerce API endpoints. Customers who don&#x27;t opt in to your Mailchimp list [will be added as &#x60;Transactional&#x60; members](https://mailchimp.com/developer/marketing/docs/e-commerce/#customers). | [default to null]
**Company** | **string** | The customer&#x27;s company. | [optional] [default to null]
**FirstName** | **string** | The customer&#x27;s first name. | [optional] [default to null]
**LastName** | **string** | The customer&#x27;s last name. | [optional] [default to null]
**Address** | [***Address1**](Address_1.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

