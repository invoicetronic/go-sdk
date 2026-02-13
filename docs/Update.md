# Update

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier. Leave it at 0 for new records as it will be set automatically. | [optional] 
**Created** | Pointer to **time.Time** | Creation date. It is set automatically. | [optional] 
**Version** | Pointer to **int32** | Row version, for optimistic concurrency. It is set automatically. | [optional] 
**UserId** | Pointer to **int32** | User id. | [optional] 
**CompanyId** | Pointer to **int32** | Company id. | [optional] 
**SendId** | Pointer to **int32** | Send id. This is the id of the sent invoice to which this update refers to. | [optional] 
**LastUpdate** | Pointer to **time.Time** | Last update from SDI. | [optional] 
**State** | Pointer to **string** | State of the document. These are the possible values, as per the SDI documentation: | [optional] 
**Description** | Pointer to **NullableString** | Description for the state. | [optional] 
**MessageId** | Pointer to **NullableString** | SDI message id. | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) | SDI errors, if any. | [optional] 
**IsRead** | Pointer to **bool** | Whether the item has been read at least once. | [optional] 
**Send** | Pointer to [**SendReduced**](SendReduced.md) |  | [optional] 

## Methods

### NewUpdate

`func NewUpdate() *Update`

NewUpdate instantiates a new Update object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWithDefaults

`func NewUpdateWithDefaults() *Update`

NewUpdateWithDefaults instantiates a new Update object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Update) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Update) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Update) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Update) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *Update) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Update) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Update) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Update) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetVersion

`func (o *Update) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Update) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Update) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Update) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUserId

`func (o *Update) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Update) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Update) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Update) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *Update) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Update) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Update) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *Update) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetSendId

`func (o *Update) GetSendId() int32`

GetSendId returns the SendId field if non-nil, zero value otherwise.

### GetSendIdOk

`func (o *Update) GetSendIdOk() (*int32, bool)`

GetSendIdOk returns a tuple with the SendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendId

`func (o *Update) SetSendId(v int32)`

SetSendId sets SendId field to given value.

### HasSendId

`func (o *Update) HasSendId() bool`

HasSendId returns a boolean if a field has been set.

### GetLastUpdate

`func (o *Update) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *Update) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *Update) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *Update) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetState

`func (o *Update) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Update) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Update) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Update) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDescription

`func (o *Update) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Update) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Update) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Update) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Update) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Update) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMessageId

`func (o *Update) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *Update) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *Update) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *Update) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### SetMessageIdNil

`func (o *Update) SetMessageIdNil(b bool)`

 SetMessageIdNil sets the value for MessageId to be an explicit nil

### UnsetMessageId
`func (o *Update) UnsetMessageId()`

UnsetMessageId ensures that no value is present for MessageId, not even an explicit nil
### GetErrors

`func (o *Update) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Update) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Update) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Update) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *Update) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *Update) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetIsRead

`func (o *Update) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *Update) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *Update) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *Update) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### GetSend

`func (o *Update) GetSend() SendReduced`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *Update) GetSendOk() (*SendReduced, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *Update) SetSend(v SendReduced)`

SetSend sets Send field to given value.

### HasSend

`func (o *Update) HasSend() bool`

HasSend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


