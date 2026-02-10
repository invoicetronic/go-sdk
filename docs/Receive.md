# Receive

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
**Identifier** | Pointer to **NullableString** | SDI identifier. This is set by the SDI and is guaranteed to be unique within the SDI system. | [optional] 
**FileName** | Pointer to **NullableString** | Xml file name. | [optional] 
**Format** | Pointer to **NullableString** | SDI format (FPA12, FPR12, FSM10, ...) | [optional] 
**Payload** | **string** | Xml payload. This is the actual xml content, as string. On send, it can be base64 encoded. If it&#39;s not, it will be encoded before sending. It is guaranteed to be encrypted at rest. | 
**LastUpdate** | Pointer to **NullableTime** | Last update from SDI. | [optional] 
**DateSent** | Pointer to **NullableTime** | When the invoice was sent to SDI. | [optional] 
**Documents** | Pointer to [**[]DocumentData**](DocumentData.md) | The invoices included in the payload. This is set by the system, based on the xml content. | [optional] 
**Encoding** | Pointer to **string** | Whether the payload is Base64 encoded or a plain XML (text). | [optional] 
**IsRead** | Pointer to **bool** | Whether the invoice has been read at least once. Set to true only when the invoice is requested with include_payload&#x3D;true. | [optional] 
**MessageId** | Pointer to **NullableString** | SDI message id. | [optional] 

## Methods

### NewReceive

`func NewReceive(payload string, ) *Receive`

NewReceive instantiates a new Receive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiveWithDefaults

`func NewReceiveWithDefaults() *Receive`

NewReceiveWithDefaults instantiates a new Receive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Receive) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Receive) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Receive) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Receive) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *Receive) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Receive) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Receive) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Receive) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetVersion

`func (o *Receive) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Receive) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Receive) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Receive) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUserId

`func (o *Receive) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Receive) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Receive) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Receive) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *Receive) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Receive) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Receive) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *Receive) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCommittente

`func (o *Receive) GetCommittente() string`

GetCommittente returns the Committente field if non-nil, zero value otherwise.

### GetCommittenteOk

`func (o *Receive) GetCommittenteOk() (*string, bool)`

GetCommittenteOk returns a tuple with the Committente field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittente

`func (o *Receive) SetCommittente(v string)`

SetCommittente sets Committente field to given value.

### HasCommittente

`func (o *Receive) HasCommittente() bool`

HasCommittente returns a boolean if a field has been set.

### SetCommittenteNil

`func (o *Receive) SetCommittenteNil(b bool)`

 SetCommittenteNil sets the value for Committente to be an explicit nil

### UnsetCommittente
`func (o *Receive) UnsetCommittente()`

UnsetCommittente ensures that no value is present for Committente, not even an explicit nil
### GetPrestatore

`func (o *Receive) GetPrestatore() string`

GetPrestatore returns the Prestatore field if non-nil, zero value otherwise.

### GetPrestatoreOk

`func (o *Receive) GetPrestatoreOk() (*string, bool)`

GetPrestatoreOk returns a tuple with the Prestatore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestatore

`func (o *Receive) SetPrestatore(v string)`

SetPrestatore sets Prestatore field to given value.

### HasPrestatore

`func (o *Receive) HasPrestatore() bool`

HasPrestatore returns a boolean if a field has been set.

### SetPrestatoreNil

`func (o *Receive) SetPrestatoreNil(b bool)`

 SetPrestatoreNil sets the value for Prestatore to be an explicit nil

### UnsetPrestatore
`func (o *Receive) UnsetPrestatore()`

UnsetPrestatore ensures that no value is present for Prestatore, not even an explicit nil
### GetIdentifier

`func (o *Receive) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Receive) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Receive) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Receive) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *Receive) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *Receive) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetFileName

`func (o *Receive) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Receive) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Receive) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Receive) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *Receive) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *Receive) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFormat

`func (o *Receive) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Receive) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Receive) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *Receive) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *Receive) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *Receive) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetPayload

`func (o *Receive) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Receive) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Receive) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetLastUpdate

`func (o *Receive) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *Receive) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *Receive) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *Receive) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### SetLastUpdateNil

`func (o *Receive) SetLastUpdateNil(b bool)`

 SetLastUpdateNil sets the value for LastUpdate to be an explicit nil

### UnsetLastUpdate
`func (o *Receive) UnsetLastUpdate()`

UnsetLastUpdate ensures that no value is present for LastUpdate, not even an explicit nil
### GetDateSent

`func (o *Receive) GetDateSent() time.Time`

GetDateSent returns the DateSent field if non-nil, zero value otherwise.

### GetDateSentOk

`func (o *Receive) GetDateSentOk() (*time.Time, bool)`

GetDateSentOk returns a tuple with the DateSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSent

`func (o *Receive) SetDateSent(v time.Time)`

SetDateSent sets DateSent field to given value.

### HasDateSent

`func (o *Receive) HasDateSent() bool`

HasDateSent returns a boolean if a field has been set.

### SetDateSentNil

`func (o *Receive) SetDateSentNil(b bool)`

 SetDateSentNil sets the value for DateSent to be an explicit nil

### UnsetDateSent
`func (o *Receive) UnsetDateSent()`

UnsetDateSent ensures that no value is present for DateSent, not even an explicit nil
### GetDocuments

`func (o *Receive) GetDocuments() []DocumentData`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *Receive) GetDocumentsOk() (*[]DocumentData, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *Receive) SetDocuments(v []DocumentData)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *Receive) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### SetDocumentsNil

`func (o *Receive) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *Receive) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil
### GetEncoding

`func (o *Receive) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *Receive) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *Receive) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *Receive) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetIsRead

`func (o *Receive) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *Receive) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *Receive) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *Receive) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### GetMessageId

`func (o *Receive) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *Receive) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *Receive) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *Receive) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### SetMessageIdNil

`func (o *Receive) SetMessageIdNil(b bool)`

 SetMessageIdNil sets the value for MessageId to be an explicit nil

### UnsetMessageId
`func (o *Receive) UnsetMessageId()`

UnsetMessageId ensures that no value is present for MessageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


