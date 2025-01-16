# WebHookHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier. Leave it at 0 for new records as it will be set automatically. | [optional] 
**Created** | Pointer to **time.Time** | Creation date. It is set automatically. | [optional] 
**Version** | Pointer to **int32** | Row version, for optimistic concurrency. It is set automatically. | [optional] 
**WebHookId** | Pointer to **int32** | Webhook id. | [optional] 
**UserId** | Pointer to **int32** | User id. | [optional] 
**Event** | Pointer to **NullableString** | Event name. | [optional] 
**StatusCode** | Pointer to **int32** | Status code. | [optional] 
**DateTime** | Pointer to **time.Time** | Date and time of the request. | [optional] 
**Success** | Pointer to **bool** | Wether the request was successful. | [optional] [readonly] 

## Methods

### NewWebHookHistory

`func NewWebHookHistory() *WebHookHistory`

NewWebHookHistory instantiates a new WebHookHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookHistoryWithDefaults

`func NewWebHookHistoryWithDefaults() *WebHookHistory`

NewWebHookHistoryWithDefaults instantiates a new WebHookHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebHookHistory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebHookHistory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebHookHistory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WebHookHistory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *WebHookHistory) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebHookHistory) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebHookHistory) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WebHookHistory) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetVersion

`func (o *WebHookHistory) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WebHookHistory) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WebHookHistory) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WebHookHistory) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWebHookId

`func (o *WebHookHistory) GetWebHookId() int32`

GetWebHookId returns the WebHookId field if non-nil, zero value otherwise.

### GetWebHookIdOk

`func (o *WebHookHistory) GetWebHookIdOk() (*int32, bool)`

GetWebHookIdOk returns a tuple with the WebHookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookId

`func (o *WebHookHistory) SetWebHookId(v int32)`

SetWebHookId sets WebHookId field to given value.

### HasWebHookId

`func (o *WebHookHistory) HasWebHookId() bool`

HasWebHookId returns a boolean if a field has been set.

### GetUserId

`func (o *WebHookHistory) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebHookHistory) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebHookHistory) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WebHookHistory) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEvent

`func (o *WebHookHistory) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebHookHistory) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebHookHistory) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WebHookHistory) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *WebHookHistory) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *WebHookHistory) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetStatusCode

`func (o *WebHookHistory) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *WebHookHistory) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *WebHookHistory) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *WebHookHistory) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetDateTime

`func (o *WebHookHistory) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *WebHookHistory) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *WebHookHistory) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *WebHookHistory) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetSuccess

`func (o *WebHookHistory) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *WebHookHistory) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *WebHookHistory) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *WebHookHistory) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


