# IndirizzoResa

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

### NewIndirizzoResa

`func NewIndirizzoResa() *IndirizzoResa`

NewIndirizzoResa instantiates a new IndirizzoResa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndirizzoResaWithDefaults

`func NewIndirizzoResaWithDefaults() *IndirizzoResa`

NewIndirizzoResaWithDefaults instantiates a new IndirizzoResa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndirizzo

`func (o *IndirizzoResa) GetIndirizzo() string`

GetIndirizzo returns the Indirizzo field if non-nil, zero value otherwise.

### GetIndirizzoOk

`func (o *IndirizzoResa) GetIndirizzoOk() (*string, bool)`

GetIndirizzoOk returns a tuple with the Indirizzo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirizzo

`func (o *IndirizzoResa) SetIndirizzo(v string)`

SetIndirizzo sets Indirizzo field to given value.

### HasIndirizzo

`func (o *IndirizzoResa) HasIndirizzo() bool`

HasIndirizzo returns a boolean if a field has been set.

### SetIndirizzoNil

`func (o *IndirizzoResa) SetIndirizzoNil(b bool)`

 SetIndirizzoNil sets the value for Indirizzo to be an explicit nil

### UnsetIndirizzo
`func (o *IndirizzoResa) UnsetIndirizzo()`

UnsetIndirizzo ensures that no value is present for Indirizzo, not even an explicit nil
### GetNumeroCivico

`func (o *IndirizzoResa) GetNumeroCivico() string`

GetNumeroCivico returns the NumeroCivico field if non-nil, zero value otherwise.

### GetNumeroCivicoOk

`func (o *IndirizzoResa) GetNumeroCivicoOk() (*string, bool)`

GetNumeroCivicoOk returns a tuple with the NumeroCivico field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroCivico

`func (o *IndirizzoResa) SetNumeroCivico(v string)`

SetNumeroCivico sets NumeroCivico field to given value.

### HasNumeroCivico

`func (o *IndirizzoResa) HasNumeroCivico() bool`

HasNumeroCivico returns a boolean if a field has been set.

### SetNumeroCivicoNil

`func (o *IndirizzoResa) SetNumeroCivicoNil(b bool)`

 SetNumeroCivicoNil sets the value for NumeroCivico to be an explicit nil

### UnsetNumeroCivico
`func (o *IndirizzoResa) UnsetNumeroCivico()`

UnsetNumeroCivico ensures that no value is present for NumeroCivico, not even an explicit nil
### GetCap

`func (o *IndirizzoResa) GetCap() string`

GetCap returns the Cap field if non-nil, zero value otherwise.

### GetCapOk

`func (o *IndirizzoResa) GetCapOk() (*string, bool)`

GetCapOk returns a tuple with the Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCap

`func (o *IndirizzoResa) SetCap(v string)`

SetCap sets Cap field to given value.

### HasCap

`func (o *IndirizzoResa) HasCap() bool`

HasCap returns a boolean if a field has been set.

### SetCapNil

`func (o *IndirizzoResa) SetCapNil(b bool)`

 SetCapNil sets the value for Cap to be an explicit nil

### UnsetCap
`func (o *IndirizzoResa) UnsetCap()`

UnsetCap ensures that no value is present for Cap, not even an explicit nil
### GetComune

`func (o *IndirizzoResa) GetComune() string`

GetComune returns the Comune field if non-nil, zero value otherwise.

### GetComuneOk

`func (o *IndirizzoResa) GetComuneOk() (*string, bool)`

GetComuneOk returns a tuple with the Comune field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComune

`func (o *IndirizzoResa) SetComune(v string)`

SetComune sets Comune field to given value.

### HasComune

`func (o *IndirizzoResa) HasComune() bool`

HasComune returns a boolean if a field has been set.

### SetComuneNil

`func (o *IndirizzoResa) SetComuneNil(b bool)`

 SetComuneNil sets the value for Comune to be an explicit nil

### UnsetComune
`func (o *IndirizzoResa) UnsetComune()`

UnsetComune ensures that no value is present for Comune, not even an explicit nil
### GetProvincia

`func (o *IndirizzoResa) GetProvincia() string`

GetProvincia returns the Provincia field if non-nil, zero value otherwise.

### GetProvinciaOk

`func (o *IndirizzoResa) GetProvinciaOk() (*string, bool)`

GetProvinciaOk returns a tuple with the Provincia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvincia

`func (o *IndirizzoResa) SetProvincia(v string)`

SetProvincia sets Provincia field to given value.

### HasProvincia

`func (o *IndirizzoResa) HasProvincia() bool`

HasProvincia returns a boolean if a field has been set.

### SetProvinciaNil

`func (o *IndirizzoResa) SetProvinciaNil(b bool)`

 SetProvinciaNil sets the value for Provincia to be an explicit nil

### UnsetProvincia
`func (o *IndirizzoResa) UnsetProvincia()`

UnsetProvincia ensures that no value is present for Provincia, not even an explicit nil
### GetNazione

`func (o *IndirizzoResa) GetNazione() string`

GetNazione returns the Nazione field if non-nil, zero value otherwise.

### GetNazioneOk

`func (o *IndirizzoResa) GetNazioneOk() (*string, bool)`

GetNazioneOk returns a tuple with the Nazione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNazione

`func (o *IndirizzoResa) SetNazione(v string)`

SetNazione sets Nazione field to given value.

### HasNazione

`func (o *IndirizzoResa) HasNazione() bool`

HasNazione returns a boolean if a field has been set.

### SetNazioneNil

`func (o *IndirizzoResa) SetNazioneNil(b bool)`

 SetNazioneNil sets the value for Nazione to be an explicit nil

### UnsetNazione
`func (o *IndirizzoResa) UnsetNazione()`

UnsetNazione ensures that no value is present for Nazione, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


