# SendReduced

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **NullableString** | SDI identifier. | [optional] 
**Prestatore** | Pointer to **NullableString** | VAT number of the Cedente/Prestatore (vendor). | [optional] 
**MetaData** | Pointer to **map[string]string** | Optional metadata, as json. | [optional] 
**Documents** | Pointer to [**[]DocumentData**](DocumentData.md) | The invoices included in the payload. | [optional] 
**DateSent** | Pointer to **NullableTime** | When the invoice was sent to SDI. | [optional] 

## Methods

### NewSendReduced

`func NewSendReduced() *SendReduced`

NewSendReduced instantiates a new SendReduced object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendReducedWithDefaults

`func NewSendReducedWithDefaults() *SendReduced`

NewSendReducedWithDefaults instantiates a new SendReduced object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *SendReduced) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *SendReduced) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *SendReduced) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *SendReduced) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *SendReduced) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *SendReduced) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetPrestatore

`func (o *SendReduced) GetPrestatore() string`

GetPrestatore returns the Prestatore field if non-nil, zero value otherwise.

### GetPrestatoreOk

`func (o *SendReduced) GetPrestatoreOk() (*string, bool)`

GetPrestatoreOk returns a tuple with the Prestatore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestatore

`func (o *SendReduced) SetPrestatore(v string)`

SetPrestatore sets Prestatore field to given value.

### HasPrestatore

`func (o *SendReduced) HasPrestatore() bool`

HasPrestatore returns a boolean if a field has been set.

### SetPrestatoreNil

`func (o *SendReduced) SetPrestatoreNil(b bool)`

 SetPrestatoreNil sets the value for Prestatore to be an explicit nil

### UnsetPrestatore
`func (o *SendReduced) UnsetPrestatore()`

UnsetPrestatore ensures that no value is present for Prestatore, not even an explicit nil
### GetMetaData

`func (o *SendReduced) GetMetaData() map[string]string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *SendReduced) GetMetaDataOk() (*map[string]string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *SendReduced) SetMetaData(v map[string]string)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *SendReduced) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### SetMetaDataNil

`func (o *SendReduced) SetMetaDataNil(b bool)`

 SetMetaDataNil sets the value for MetaData to be an explicit nil

### UnsetMetaData
`func (o *SendReduced) UnsetMetaData()`

UnsetMetaData ensures that no value is present for MetaData, not even an explicit nil
### GetDocuments

`func (o *SendReduced) GetDocuments() []DocumentData`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *SendReduced) GetDocumentsOk() (*[]DocumentData, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *SendReduced) SetDocuments(v []DocumentData)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *SendReduced) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### SetDocumentsNil

`func (o *SendReduced) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *SendReduced) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil
### GetDateSent

`func (o *SendReduced) GetDateSent() time.Time`

GetDateSent returns the DateSent field if non-nil, zero value otherwise.

### GetDateSentOk

`func (o *SendReduced) GetDateSentOk() (*time.Time, bool)`

GetDateSentOk returns a tuple with the DateSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSent

`func (o *SendReduced) SetDateSent(v time.Time)`

SetDateSent sets DateSent field to given value.

### HasDateSent

`func (o *SendReduced) HasDateSent() bool`

HasDateSent returns a boolean if a field has been set.

### SetDateSentNil

`func (o *SendReduced) SetDateSentNil(b bool)`

 SetDateSentNil sets the value for DateSent to be an explicit nil

### UnsetDateSent
`func (o *SendReduced) UnsetDateSent()`

UnsetDateSent ensures that no value is present for DateSent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


