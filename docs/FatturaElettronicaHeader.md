# FatturaElettronicaHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatiTrasmissione** | Pointer to [**DatiTrasmissione**](DatiTrasmissione.md) |  | [optional] 
**CedentePrestatore** | Pointer to [**CedentePrestatore**](CedentePrestatore.md) |  | [optional] 
**RappresentanteFiscale** | Pointer to [**RappresentanteFiscale**](RappresentanteFiscale.md) |  | [optional] 
**CessionarioCommittente** | Pointer to [**CessionarioCommittente**](CessionarioCommittente.md) |  | [optional] 
**TerzoIntermediarioOSoggettoEmittente** | Pointer to [**TerzoIntermediarioOSoggettoEmittente**](TerzoIntermediarioOSoggettoEmittente.md) |  | [optional] 
**SoggettoEmittente** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFatturaElettronicaHeader

`func NewFatturaElettronicaHeader() *FatturaElettronicaHeader`

NewFatturaElettronicaHeader instantiates a new FatturaElettronicaHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFatturaElettronicaHeaderWithDefaults

`func NewFatturaElettronicaHeaderWithDefaults() *FatturaElettronicaHeader`

NewFatturaElettronicaHeaderWithDefaults instantiates a new FatturaElettronicaHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatiTrasmissione

`func (o *FatturaElettronicaHeader) GetDatiTrasmissione() DatiTrasmissione`

GetDatiTrasmissione returns the DatiTrasmissione field if non-nil, zero value otherwise.

### GetDatiTrasmissioneOk

`func (o *FatturaElettronicaHeader) GetDatiTrasmissioneOk() (*DatiTrasmissione, bool)`

GetDatiTrasmissioneOk returns a tuple with the DatiTrasmissione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiTrasmissione

`func (o *FatturaElettronicaHeader) SetDatiTrasmissione(v DatiTrasmissione)`

SetDatiTrasmissione sets DatiTrasmissione field to given value.

### HasDatiTrasmissione

`func (o *FatturaElettronicaHeader) HasDatiTrasmissione() bool`

HasDatiTrasmissione returns a boolean if a field has been set.

### GetCedentePrestatore

`func (o *FatturaElettronicaHeader) GetCedentePrestatore() CedentePrestatore`

GetCedentePrestatore returns the CedentePrestatore field if non-nil, zero value otherwise.

### GetCedentePrestatoreOk

`func (o *FatturaElettronicaHeader) GetCedentePrestatoreOk() (*CedentePrestatore, bool)`

GetCedentePrestatoreOk returns a tuple with the CedentePrestatore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCedentePrestatore

`func (o *FatturaElettronicaHeader) SetCedentePrestatore(v CedentePrestatore)`

SetCedentePrestatore sets CedentePrestatore field to given value.

### HasCedentePrestatore

`func (o *FatturaElettronicaHeader) HasCedentePrestatore() bool`

HasCedentePrestatore returns a boolean if a field has been set.

### GetRappresentanteFiscale

`func (o *FatturaElettronicaHeader) GetRappresentanteFiscale() RappresentanteFiscale`

GetRappresentanteFiscale returns the RappresentanteFiscale field if non-nil, zero value otherwise.

### GetRappresentanteFiscaleOk

`func (o *FatturaElettronicaHeader) GetRappresentanteFiscaleOk() (*RappresentanteFiscale, bool)`

GetRappresentanteFiscaleOk returns a tuple with the RappresentanteFiscale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRappresentanteFiscale

`func (o *FatturaElettronicaHeader) SetRappresentanteFiscale(v RappresentanteFiscale)`

SetRappresentanteFiscale sets RappresentanteFiscale field to given value.

### HasRappresentanteFiscale

`func (o *FatturaElettronicaHeader) HasRappresentanteFiscale() bool`

HasRappresentanteFiscale returns a boolean if a field has been set.

### GetCessionarioCommittente

`func (o *FatturaElettronicaHeader) GetCessionarioCommittente() CessionarioCommittente`

GetCessionarioCommittente returns the CessionarioCommittente field if non-nil, zero value otherwise.

### GetCessionarioCommittenteOk

`func (o *FatturaElettronicaHeader) GetCessionarioCommittenteOk() (*CessionarioCommittente, bool)`

GetCessionarioCommittenteOk returns a tuple with the CessionarioCommittente field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCessionarioCommittente

`func (o *FatturaElettronicaHeader) SetCessionarioCommittente(v CessionarioCommittente)`

SetCessionarioCommittente sets CessionarioCommittente field to given value.

### HasCessionarioCommittente

`func (o *FatturaElettronicaHeader) HasCessionarioCommittente() bool`

HasCessionarioCommittente returns a boolean if a field has been set.

### GetTerzoIntermediarioOSoggettoEmittente

`func (o *FatturaElettronicaHeader) GetTerzoIntermediarioOSoggettoEmittente() TerzoIntermediarioOSoggettoEmittente`

GetTerzoIntermediarioOSoggettoEmittente returns the TerzoIntermediarioOSoggettoEmittente field if non-nil, zero value otherwise.

### GetTerzoIntermediarioOSoggettoEmittenteOk

`func (o *FatturaElettronicaHeader) GetTerzoIntermediarioOSoggettoEmittenteOk() (*TerzoIntermediarioOSoggettoEmittente, bool)`

GetTerzoIntermediarioOSoggettoEmittenteOk returns a tuple with the TerzoIntermediarioOSoggettoEmittente field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerzoIntermediarioOSoggettoEmittente

`func (o *FatturaElettronicaHeader) SetTerzoIntermediarioOSoggettoEmittente(v TerzoIntermediarioOSoggettoEmittente)`

SetTerzoIntermediarioOSoggettoEmittente sets TerzoIntermediarioOSoggettoEmittente field to given value.

### HasTerzoIntermediarioOSoggettoEmittente

`func (o *FatturaElettronicaHeader) HasTerzoIntermediarioOSoggettoEmittente() bool`

HasTerzoIntermediarioOSoggettoEmittente returns a boolean if a field has been set.

### GetSoggettoEmittente

`func (o *FatturaElettronicaHeader) GetSoggettoEmittente() string`

GetSoggettoEmittente returns the SoggettoEmittente field if non-nil, zero value otherwise.

### GetSoggettoEmittenteOk

`func (o *FatturaElettronicaHeader) GetSoggettoEmittenteOk() (*string, bool)`

GetSoggettoEmittenteOk returns a tuple with the SoggettoEmittente field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoggettoEmittente

`func (o *FatturaElettronicaHeader) SetSoggettoEmittente(v string)`

SetSoggettoEmittente sets SoggettoEmittente field to given value.

### HasSoggettoEmittente

`func (o *FatturaElettronicaHeader) HasSoggettoEmittente() bool`

HasSoggettoEmittente returns a boolean if a field has been set.

### SetSoggettoEmittenteNil

`func (o *FatturaElettronicaHeader) SetSoggettoEmittenteNil(b bool)`

 SetSoggettoEmittenteNil sets the value for SoggettoEmittente to be an explicit nil

### UnsetSoggettoEmittente
`func (o *FatturaElettronicaHeader) UnsetSoggettoEmittente()`

UnsetSoggettoEmittente ensures that no value is present for SoggettoEmittente, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


