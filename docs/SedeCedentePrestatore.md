# SedeCedentePrestatore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indirizzo** | Pointer to **NullableString** |  | [optional] 
**NumeroCivico** | Pointer to **NullableString** |  | [optional] 
**Cap** | Pointer to **NullableString** |  | [optional] 
**Comune** | Pointer to **NullableString** |  | [optional] 
**Provincia** | Pointer to **NullableString** |  | [optional] 
**Nazione** | Pointer to **NullableString** |  | [optional] [default to "IT"]

## Methods

### NewSedeCedentePrestatore

`func NewSedeCedentePrestatore() *SedeCedentePrestatore`

NewSedeCedentePrestatore instantiates a new SedeCedentePrestatore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSedeCedentePrestatoreWithDefaults

`func NewSedeCedentePrestatoreWithDefaults() *SedeCedentePrestatore`

NewSedeCedentePrestatoreWithDefaults instantiates a new SedeCedentePrestatore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndirizzo

`func (o *SedeCedentePrestatore) GetIndirizzo() string`

GetIndirizzo returns the Indirizzo field if non-nil, zero value otherwise.

### GetIndirizzoOk

`func (o *SedeCedentePrestatore) GetIndirizzoOk() (*string, bool)`

GetIndirizzoOk returns a tuple with the Indirizzo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirizzo

`func (o *SedeCedentePrestatore) SetIndirizzo(v string)`

SetIndirizzo sets Indirizzo field to given value.

### HasIndirizzo

`func (o *SedeCedentePrestatore) HasIndirizzo() bool`

HasIndirizzo returns a boolean if a field has been set.

### SetIndirizzoNil

`func (o *SedeCedentePrestatore) SetIndirizzoNil(b bool)`

 SetIndirizzoNil sets the value for Indirizzo to be an explicit nil

### UnsetIndirizzo
`func (o *SedeCedentePrestatore) UnsetIndirizzo()`

UnsetIndirizzo ensures that no value is present for Indirizzo, not even an explicit nil
### GetNumeroCivico

`func (o *SedeCedentePrestatore) GetNumeroCivico() string`

GetNumeroCivico returns the NumeroCivico field if non-nil, zero value otherwise.

### GetNumeroCivicoOk

`func (o *SedeCedentePrestatore) GetNumeroCivicoOk() (*string, bool)`

GetNumeroCivicoOk returns a tuple with the NumeroCivico field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroCivico

`func (o *SedeCedentePrestatore) SetNumeroCivico(v string)`

SetNumeroCivico sets NumeroCivico field to given value.

### HasNumeroCivico

`func (o *SedeCedentePrestatore) HasNumeroCivico() bool`

HasNumeroCivico returns a boolean if a field has been set.

### SetNumeroCivicoNil

`func (o *SedeCedentePrestatore) SetNumeroCivicoNil(b bool)`

 SetNumeroCivicoNil sets the value for NumeroCivico to be an explicit nil

### UnsetNumeroCivico
`func (o *SedeCedentePrestatore) UnsetNumeroCivico()`

UnsetNumeroCivico ensures that no value is present for NumeroCivico, not even an explicit nil
### GetCap

`func (o *SedeCedentePrestatore) GetCap() string`

GetCap returns the Cap field if non-nil, zero value otherwise.

### GetCapOk

`func (o *SedeCedentePrestatore) GetCapOk() (*string, bool)`

GetCapOk returns a tuple with the Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCap

`func (o *SedeCedentePrestatore) SetCap(v string)`

SetCap sets Cap field to given value.

### HasCap

`func (o *SedeCedentePrestatore) HasCap() bool`

HasCap returns a boolean if a field has been set.

### SetCapNil

`func (o *SedeCedentePrestatore) SetCapNil(b bool)`

 SetCapNil sets the value for Cap to be an explicit nil

### UnsetCap
`func (o *SedeCedentePrestatore) UnsetCap()`

UnsetCap ensures that no value is present for Cap, not even an explicit nil
### GetComune

`func (o *SedeCedentePrestatore) GetComune() string`

GetComune returns the Comune field if non-nil, zero value otherwise.

### GetComuneOk

`func (o *SedeCedentePrestatore) GetComuneOk() (*string, bool)`

GetComuneOk returns a tuple with the Comune field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComune

`func (o *SedeCedentePrestatore) SetComune(v string)`

SetComune sets Comune field to given value.

### HasComune

`func (o *SedeCedentePrestatore) HasComune() bool`

HasComune returns a boolean if a field has been set.

### SetComuneNil

`func (o *SedeCedentePrestatore) SetComuneNil(b bool)`

 SetComuneNil sets the value for Comune to be an explicit nil

### UnsetComune
`func (o *SedeCedentePrestatore) UnsetComune()`

UnsetComune ensures that no value is present for Comune, not even an explicit nil
### GetProvincia

`func (o *SedeCedentePrestatore) GetProvincia() string`

GetProvincia returns the Provincia field if non-nil, zero value otherwise.

### GetProvinciaOk

`func (o *SedeCedentePrestatore) GetProvinciaOk() (*string, bool)`

GetProvinciaOk returns a tuple with the Provincia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvincia

`func (o *SedeCedentePrestatore) SetProvincia(v string)`

SetProvincia sets Provincia field to given value.

### HasProvincia

`func (o *SedeCedentePrestatore) HasProvincia() bool`

HasProvincia returns a boolean if a field has been set.

### SetProvinciaNil

`func (o *SedeCedentePrestatore) SetProvinciaNil(b bool)`

 SetProvinciaNil sets the value for Provincia to be an explicit nil

### UnsetProvincia
`func (o *SedeCedentePrestatore) UnsetProvincia()`

UnsetProvincia ensures that no value is present for Provincia, not even an explicit nil
### GetNazione

`func (o *SedeCedentePrestatore) GetNazione() string`

GetNazione returns the Nazione field if non-nil, zero value otherwise.

### GetNazioneOk

`func (o *SedeCedentePrestatore) GetNazioneOk() (*string, bool)`

GetNazioneOk returns a tuple with the Nazione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNazione

`func (o *SedeCedentePrestatore) SetNazione(v string)`

SetNazione sets Nazione field to given value.

### HasNazione

`func (o *SedeCedentePrestatore) HasNazione() bool`

HasNazione returns a boolean if a field has been set.

### SetNazioneNil

`func (o *SedeCedentePrestatore) SetNazioneNil(b bool)`

 SetNazioneNil sets the value for Nazione to be an explicit nil

### UnsetNazione
`func (o *SedeCedentePrestatore) UnsetNazione()`

UnsetNazione ensures that no value is present for Nazione, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


