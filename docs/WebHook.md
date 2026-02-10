# WebHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier. Leave it at 0 for new records as it will be set automatically. | [optional] 
**Created** | Pointer to **time.Time** | Creation date. It is set automatically. | [optional] 
**Version** | Pointer to **int32** | Row version, for optimistic concurrency. It is set automatically. | [optional] 
**UserId** | Pointer to **int32** | User id. | [optional] 
**CompanyId** | Pointer to **NullableInt32** | Company id. | [optional] 
**Url** | **string** | The url of your application&#39;s endpoint that will receive a POST request when the webhook is fired. | 
**Enabled** | Pointer to **bool** | Whether the webhook is enabled. On creation, this is set to &#x60;true&#x60;. | [optional] 
**Secret** | Pointer to **NullableString** | The secret used to generate webhook signatures, only returned on creation. You should store this value securely and validate it on every call, to ensure that the caller is InvoicetronicApi. | [optional] 
**Description** | Pointer to **NullableString** | An optional description. | [optional] 
**Events** | Pointer to **[]string** | List of events that trigger the webhook.  See Invoicetronic.SupportedEvents.Available for a list of valid event names. | [optional] 

## Methods

### NewWebHook

`func NewWebHook(url string, ) *WebHook`

NewWebHook instantiates a new WebHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookWithDefaults

`func NewWebHookWithDefaults() *WebHook`

NewWebHookWithDefaults instantiates a new WebHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebHook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebHook) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebHook) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WebHook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *WebHook) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebHook) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebHook) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WebHook) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetVersion

`func (o *WebHook) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WebHook) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WebHook) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WebHook) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUserId

`func (o *WebHook) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebHook) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebHook) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WebHook) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *WebHook) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *WebHook) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *WebHook) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *WebHook) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *WebHook) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *WebHook) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetUrl

`func (o *WebHook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebHook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebHook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEnabled

`func (o *WebHook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebHook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebHook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebHook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSecret

`func (o *WebHook) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebHook) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebHook) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *WebHook) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *WebHook) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *WebHook) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetDescription

`func (o *WebHook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebHook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebHook) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebHook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebHook) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebHook) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEvents

`func (o *WebHook) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebHook) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebHook) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WebHook) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *WebHook) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *WebHook) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


