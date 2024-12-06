# FatturaOrdinaria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SistemaEmittente** | Pointer to **NullableString** |  | [optional] 
**FatturaElettronicaHeader** | Pointer to [**FatturaElettronicaHeader**](FatturaElettronicaHeader.md) |  | [optional] 
**FatturaElettronicaBody** | Pointer to [**[]FatturaElettronicaBody**](FatturaElettronicaBody.md) |  | [optional] 

## Methods

### NewFatturaOrdinaria

`func NewFatturaOrdinaria() *FatturaOrdinaria`

NewFatturaOrdinaria instantiates a new FatturaOrdinaria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFatturaOrdinariaWithDefaults

`func NewFatturaOrdinariaWithDefaults() *FatturaOrdinaria`

NewFatturaOrdinariaWithDefaults instantiates a new FatturaOrdinaria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSistemaEmittente

`func (o *FatturaOrdinaria) GetSistemaEmittente() string`

GetSistemaEmittente returns the SistemaEmittente field if non-nil, zero value otherwise.

### GetSistemaEmittenteOk

`func (o *FatturaOrdinaria) GetSistemaEmittenteOk() (*string, bool)`

GetSistemaEmittenteOk returns a tuple with the SistemaEmittente field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSistemaEmittente

`func (o *FatturaOrdinaria) SetSistemaEmittente(v string)`

SetSistemaEmittente sets SistemaEmittente field to given value.

### HasSistemaEmittente

`func (o *FatturaOrdinaria) HasSistemaEmittente() bool`

HasSistemaEmittente returns a boolean if a field has been set.

### SetSistemaEmittenteNil

`func (o *FatturaOrdinaria) SetSistemaEmittenteNil(b bool)`

 SetSistemaEmittenteNil sets the value for SistemaEmittente to be an explicit nil

### UnsetSistemaEmittente
`func (o *FatturaOrdinaria) UnsetSistemaEmittente()`

UnsetSistemaEmittente ensures that no value is present for SistemaEmittente, not even an explicit nil
### GetFatturaElettronicaHeader

`func (o *FatturaOrdinaria) GetFatturaElettronicaHeader() FatturaElettronicaHeader`

GetFatturaElettronicaHeader returns the FatturaElettronicaHeader field if non-nil, zero value otherwise.

### GetFatturaElettronicaHeaderOk

`func (o *FatturaOrdinaria) GetFatturaElettronicaHeaderOk() (*FatturaElettronicaHeader, bool)`

GetFatturaElettronicaHeaderOk returns a tuple with the FatturaElettronicaHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFatturaElettronicaHeader

`func (o *FatturaOrdinaria) SetFatturaElettronicaHeader(v FatturaElettronicaHeader)`

SetFatturaElettronicaHeader sets FatturaElettronicaHeader field to given value.

### HasFatturaElettronicaHeader

`func (o *FatturaOrdinaria) HasFatturaElettronicaHeader() bool`

HasFatturaElettronicaHeader returns a boolean if a field has been set.

### GetFatturaElettronicaBody

`func (o *FatturaOrdinaria) GetFatturaElettronicaBody() []FatturaElettronicaBody`

GetFatturaElettronicaBody returns the FatturaElettronicaBody field if non-nil, zero value otherwise.

### GetFatturaElettronicaBodyOk

`func (o *FatturaOrdinaria) GetFatturaElettronicaBodyOk() (*[]FatturaElettronicaBody, bool)`

GetFatturaElettronicaBodyOk returns a tuple with the FatturaElettronicaBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFatturaElettronicaBody

`func (o *FatturaOrdinaria) SetFatturaElettronicaBody(v []FatturaElettronicaBody)`

SetFatturaElettronicaBody sets FatturaElettronicaBody field to given value.

### HasFatturaElettronicaBody

`func (o *FatturaOrdinaria) HasFatturaElettronicaBody() bool`

HasFatturaElettronicaBody returns a boolean if a field has been set.

### SetFatturaElettronicaBodyNil

`func (o *FatturaOrdinaria) SetFatturaElettronicaBodyNil(b bool)`

 SetFatturaElettronicaBodyNil sets the value for FatturaElettronicaBody to be an explicit nil

### UnsetFatturaElettronicaBody
`func (o *FatturaOrdinaria) UnsetFatturaElettronicaBody()`

UnsetFatturaElettronicaBody ensures that no value is present for FatturaElettronicaBody, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


