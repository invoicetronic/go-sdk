# Send

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier. Leave it at 0 for new records as it will be set automatically. | [optional] 
**Created** | Pointer to **time.Time** | Creation date. It is set automatically. | [optional] 
**Version** | Pointer to **int32** | Row version, for optimistic concurrency. It is set automatically. | [optional] 
**UserId** | Pointer to **int32** | User id. | [optional] 
**CompanyId** | Pointer to **int32** | Company id. On send, this is the sender and must be set in advance. On receive, it will be  automatically set based on the recipient&#39;s VAT number. If a matching company is not found, the invoice will be rejected until the company is created. | [optional] 
**Committente** | Pointer to **NullableString** | VAT number of the Cessionario/Committente (customer). This is automatically set based on the recipient&#39;s VAT number. | [optional] 
**Prestatore** | Pointer to **NullableString** | VAT number of the Cedente/Prestatore (vendor). This is automatically set based on the sender&#39;s VAT number. | [optional] 
**Identifier** | Pointer to **NullableString** | SDI identifier. This is set by the SDI and is guaranted to be unique within the SDI system. | [optional] 
**FileName** | Pointer to **NullableString** | Xml file name. | [optional] 
**Format** | Pointer to **NullableString** | SDI format (FPA12, FPR12, FSM10, ...) | [optional] 
**Payload** | Pointer to **NullableString** | Xml payloaad. This is the actual xml content, as string. On send, it can be base64 encoded. If it&#39;s not, it will be encoded before sending. It is guaranteed to be cyphered at rest. | [optional] 
**LastUpdate** | Pointer to **NullableTime** | Last update from SDI. | [optional] 
**DateSent** | Pointer to **NullableTime** | When the invoice was sent to SDI. | [optional] 
**Documents** | Pointer to [**[]DocumentData**](DocumentData.md) | The invoices included in the payload. This is set by the system, based on the xml content. | [optional] 
**MetaData** | Pointer to **map[string]string** | Optional metadata, as json. | [optional] 

## Methods

### NewSend

`func NewSend() *Send`

NewSend instantiates a new Send object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendWithDefaults

`func NewSendWithDefaults() *Send`

NewSendWithDefaults instantiates a new Send object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Send) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Send) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Send) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Send) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *Send) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Send) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Send) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Send) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetVersion

`func (o *Send) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Send) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Send) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Send) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUserId

`func (o *Send) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Send) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Send) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Send) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *Send) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Send) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Send) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *Send) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCommittente

`func (o *Send) GetCommittente() string`

GetCommittente returns the Committente field if non-nil, zero value otherwise.

### GetCommittenteOk

`func (o *Send) GetCommittenteOk() (*string, bool)`

GetCommittenteOk returns a tuple with the Committente field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittente

`func (o *Send) SetCommittente(v string)`

SetCommittente sets Committente field to given value.

### HasCommittente

`func (o *Send) HasCommittente() bool`

HasCommittente returns a boolean if a field has been set.

### SetCommittenteNil

`func (o *Send) SetCommittenteNil(b bool)`

 SetCommittenteNil sets the value for Committente to be an explicit nil

### UnsetCommittente
`func (o *Send) UnsetCommittente()`

UnsetCommittente ensures that no value is present for Committente, not even an explicit nil
### GetPrestatore

`func (o *Send) GetPrestatore() string`

GetPrestatore returns the Prestatore field if non-nil, zero value otherwise.

### GetPrestatoreOk

`func (o *Send) GetPrestatoreOk() (*string, bool)`

GetPrestatoreOk returns a tuple with the Prestatore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestatore

`func (o *Send) SetPrestatore(v string)`

SetPrestatore sets Prestatore field to given value.

### HasPrestatore

`func (o *Send) HasPrestatore() bool`

HasPrestatore returns a boolean if a field has been set.

### SetPrestatoreNil

`func (o *Send) SetPrestatoreNil(b bool)`

 SetPrestatoreNil sets the value for Prestatore to be an explicit nil

### UnsetPrestatore
`func (o *Send) UnsetPrestatore()`

UnsetPrestatore ensures that no value is present for Prestatore, not even an explicit nil
### GetIdentifier

`func (o *Send) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Send) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Send) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Send) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *Send) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *Send) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetFileName

`func (o *Send) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Send) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Send) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Send) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *Send) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *Send) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFormat

`func (o *Send) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Send) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Send) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *Send) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *Send) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *Send) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetPayload

`func (o *Send) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Send) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Send) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Send) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *Send) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *Send) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetLastUpdate

`func (o *Send) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *Send) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *Send) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *Send) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### SetLastUpdateNil

`func (o *Send) SetLastUpdateNil(b bool)`

 SetLastUpdateNil sets the value for LastUpdate to be an explicit nil

### UnsetLastUpdate
`func (o *Send) UnsetLastUpdate()`

UnsetLastUpdate ensures that no value is present for LastUpdate, not even an explicit nil
### GetDateSent

`func (o *Send) GetDateSent() time.Time`

GetDateSent returns the DateSent field if non-nil, zero value otherwise.

### GetDateSentOk

`func (o *Send) GetDateSentOk() (*time.Time, bool)`

GetDateSentOk returns a tuple with the DateSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSent

`func (o *Send) SetDateSent(v time.Time)`

SetDateSent sets DateSent field to given value.

### HasDateSent

`func (o *Send) HasDateSent() bool`

HasDateSent returns a boolean if a field has been set.

### SetDateSentNil

`func (o *Send) SetDateSentNil(b bool)`

 SetDateSentNil sets the value for DateSent to be an explicit nil

### UnsetDateSent
`func (o *Send) UnsetDateSent()`

UnsetDateSent ensures that no value is present for DateSent, not even an explicit nil
### GetDocuments

`func (o *Send) GetDocuments() []DocumentData`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *Send) GetDocumentsOk() (*[]DocumentData, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *Send) SetDocuments(v []DocumentData)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *Send) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### SetDocumentsNil

`func (o *Send) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *Send) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil
### GetMetaData

`func (o *Send) GetMetaData() map[string]string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *Send) GetMetaDataOk() (*map[string]string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *Send) SetMetaData(v map[string]string)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *Send) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### SetMetaDataNil

`func (o *Send) SetMetaDataNil(b bool)`

 SetMetaDataNil sets the value for MetaData to be an explicit nil

### UnsetMetaData
`func (o *Send) UnsetMetaData()`

UnsetMetaData ensures that no value is present for MetaData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


