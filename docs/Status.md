# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationLeft** | Pointer to **int32** | Operations (invoices and validations) left. | [optional] 
**SignatureLeft** | Pointer to **int32** | Signatures left. | [optional] 
**HasActiveSeat** | Pointer to **bool** | Whether the current API key has an active Desk seat. | [optional] 
**IsSubKey** | Pointer to **bool** | Whether the current API key is a sub-key (restricted key). | [optional] 

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationLeft

`func (o *Status) GetOperationLeft() int32`

GetOperationLeft returns the OperationLeft field if non-nil, zero value otherwise.

### GetOperationLeftOk

`func (o *Status) GetOperationLeftOk() (*int32, bool)`

GetOperationLeftOk returns a tuple with the OperationLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationLeft

`func (o *Status) SetOperationLeft(v int32)`

SetOperationLeft sets OperationLeft field to given value.

### HasOperationLeft

`func (o *Status) HasOperationLeft() bool`

HasOperationLeft returns a boolean if a field has been set.

### GetSignatureLeft

`func (o *Status) GetSignatureLeft() int32`

GetSignatureLeft returns the SignatureLeft field if non-nil, zero value otherwise.

### GetSignatureLeftOk

`func (o *Status) GetSignatureLeftOk() (*int32, bool)`

GetSignatureLeftOk returns a tuple with the SignatureLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureLeft

`func (o *Status) SetSignatureLeft(v int32)`

SetSignatureLeft sets SignatureLeft field to given value.

### HasSignatureLeft

`func (o *Status) HasSignatureLeft() bool`

HasSignatureLeft returns a boolean if a field has been set.

### GetHasActiveSeat

`func (o *Status) GetHasActiveSeat() bool`

GetHasActiveSeat returns the HasActiveSeat field if non-nil, zero value otherwise.

### GetHasActiveSeatOk

`func (o *Status) GetHasActiveSeatOk() (*bool, bool)`

GetHasActiveSeatOk returns a tuple with the HasActiveSeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasActiveSeat

`func (o *Status) SetHasActiveSeat(v bool)`

SetHasActiveSeat sets HasActiveSeat field to given value.

### HasHasActiveSeat

`func (o *Status) HasHasActiveSeat() bool`

HasHasActiveSeat returns a boolean if a field has been set.

### GetIsSubKey

`func (o *Status) GetIsSubKey() bool`

GetIsSubKey returns the IsSubKey field if non-nil, zero value otherwise.

### GetIsSubKeyOk

`func (o *Status) GetIsSubKeyOk() (*bool, bool)`

GetIsSubKeyOk returns a tuple with the IsSubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubKey

`func (o *Status) SetIsSubKey(v bool)`

SetIsSubKey sets IsSubKey field to given value.

### HasIsSubKey

`func (o *Status) HasIsSubKey() bool`

HasIsSubKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


