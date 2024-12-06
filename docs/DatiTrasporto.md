# DatiTrasporto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatiAnagraficiVettore** | Pointer to [**DatiAnagraficiVettore**](DatiAnagraficiVettore.md) |  | [optional] 
**MezzoTrasporto** | Pointer to **NullableString** |  | [optional] 
**CausaleTrasporto** | Pointer to **NullableString** |  | [optional] 
**NumeroColli** | Pointer to **NullableInt32** |  | [optional] 
**Descrizione** | Pointer to **NullableString** |  | [optional] 
**UnitaMisuraPeso** | Pointer to **NullableString** |  | [optional] 
**PesoLordo** | Pointer to **NullableFloat64** |  | [optional] 
**PesoNetto** | Pointer to **NullableFloat64** |  | [optional] 
**DataOraRitiro** | Pointer to **NullableTime** |  | [optional] 
**DataInizioTrasporto** | Pointer to **NullableTime** |  | [optional] 
**TipoResa** | Pointer to **NullableString** |  | [optional] 
**IndirizzoResa** | Pointer to [**IndirizzoResa**](IndirizzoResa.md) |  | [optional] 
**DataOraConsegna** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewDatiTrasporto

`func NewDatiTrasporto() *DatiTrasporto`

NewDatiTrasporto instantiates a new DatiTrasporto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatiTrasportoWithDefaults

`func NewDatiTrasportoWithDefaults() *DatiTrasporto`

NewDatiTrasportoWithDefaults instantiates a new DatiTrasporto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatiAnagraficiVettore

`func (o *DatiTrasporto) GetDatiAnagraficiVettore() DatiAnagraficiVettore`

GetDatiAnagraficiVettore returns the DatiAnagraficiVettore field if non-nil, zero value otherwise.

### GetDatiAnagraficiVettoreOk

`func (o *DatiTrasporto) GetDatiAnagraficiVettoreOk() (*DatiAnagraficiVettore, bool)`

GetDatiAnagraficiVettoreOk returns a tuple with the DatiAnagraficiVettore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiAnagraficiVettore

`func (o *DatiTrasporto) SetDatiAnagraficiVettore(v DatiAnagraficiVettore)`

SetDatiAnagraficiVettore sets DatiAnagraficiVettore field to given value.

### HasDatiAnagraficiVettore

`func (o *DatiTrasporto) HasDatiAnagraficiVettore() bool`

HasDatiAnagraficiVettore returns a boolean if a field has been set.

### GetMezzoTrasporto

`func (o *DatiTrasporto) GetMezzoTrasporto() string`

GetMezzoTrasporto returns the MezzoTrasporto field if non-nil, zero value otherwise.

### GetMezzoTrasportoOk

`func (o *DatiTrasporto) GetMezzoTrasportoOk() (*string, bool)`

GetMezzoTrasportoOk returns a tuple with the MezzoTrasporto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMezzoTrasporto

`func (o *DatiTrasporto) SetMezzoTrasporto(v string)`

SetMezzoTrasporto sets MezzoTrasporto field to given value.

### HasMezzoTrasporto

`func (o *DatiTrasporto) HasMezzoTrasporto() bool`

HasMezzoTrasporto returns a boolean if a field has been set.

### SetMezzoTrasportoNil

`func (o *DatiTrasporto) SetMezzoTrasportoNil(b bool)`

 SetMezzoTrasportoNil sets the value for MezzoTrasporto to be an explicit nil

### UnsetMezzoTrasporto
`func (o *DatiTrasporto) UnsetMezzoTrasporto()`

UnsetMezzoTrasporto ensures that no value is present for MezzoTrasporto, not even an explicit nil
### GetCausaleTrasporto

`func (o *DatiTrasporto) GetCausaleTrasporto() string`

GetCausaleTrasporto returns the CausaleTrasporto field if non-nil, zero value otherwise.

### GetCausaleTrasportoOk

`func (o *DatiTrasporto) GetCausaleTrasportoOk() (*string, bool)`

GetCausaleTrasportoOk returns a tuple with the CausaleTrasporto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCausaleTrasporto

`func (o *DatiTrasporto) SetCausaleTrasporto(v string)`

SetCausaleTrasporto sets CausaleTrasporto field to given value.

### HasCausaleTrasporto

`func (o *DatiTrasporto) HasCausaleTrasporto() bool`

HasCausaleTrasporto returns a boolean if a field has been set.

### SetCausaleTrasportoNil

`func (o *DatiTrasporto) SetCausaleTrasportoNil(b bool)`

 SetCausaleTrasportoNil sets the value for CausaleTrasporto to be an explicit nil

### UnsetCausaleTrasporto
`func (o *DatiTrasporto) UnsetCausaleTrasporto()`

UnsetCausaleTrasporto ensures that no value is present for CausaleTrasporto, not even an explicit nil
### GetNumeroColli

`func (o *DatiTrasporto) GetNumeroColli() int32`

GetNumeroColli returns the NumeroColli field if non-nil, zero value otherwise.

### GetNumeroColliOk

`func (o *DatiTrasporto) GetNumeroColliOk() (*int32, bool)`

GetNumeroColliOk returns a tuple with the NumeroColli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroColli

`func (o *DatiTrasporto) SetNumeroColli(v int32)`

SetNumeroColli sets NumeroColli field to given value.

### HasNumeroColli

`func (o *DatiTrasporto) HasNumeroColli() bool`

HasNumeroColli returns a boolean if a field has been set.

### SetNumeroColliNil

`func (o *DatiTrasporto) SetNumeroColliNil(b bool)`

 SetNumeroColliNil sets the value for NumeroColli to be an explicit nil

### UnsetNumeroColli
`func (o *DatiTrasporto) UnsetNumeroColli()`

UnsetNumeroColli ensures that no value is present for NumeroColli, not even an explicit nil
### GetDescrizione

`func (o *DatiTrasporto) GetDescrizione() string`

GetDescrizione returns the Descrizione field if non-nil, zero value otherwise.

### GetDescrizioneOk

`func (o *DatiTrasporto) GetDescrizioneOk() (*string, bool)`

GetDescrizioneOk returns a tuple with the Descrizione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescrizione

`func (o *DatiTrasporto) SetDescrizione(v string)`

SetDescrizione sets Descrizione field to given value.

### HasDescrizione

`func (o *DatiTrasporto) HasDescrizione() bool`

HasDescrizione returns a boolean if a field has been set.

### SetDescrizioneNil

`func (o *DatiTrasporto) SetDescrizioneNil(b bool)`

 SetDescrizioneNil sets the value for Descrizione to be an explicit nil

### UnsetDescrizione
`func (o *DatiTrasporto) UnsetDescrizione()`

UnsetDescrizione ensures that no value is present for Descrizione, not even an explicit nil
### GetUnitaMisuraPeso

`func (o *DatiTrasporto) GetUnitaMisuraPeso() string`

GetUnitaMisuraPeso returns the UnitaMisuraPeso field if non-nil, zero value otherwise.

### GetUnitaMisuraPesoOk

`func (o *DatiTrasporto) GetUnitaMisuraPesoOk() (*string, bool)`

GetUnitaMisuraPesoOk returns a tuple with the UnitaMisuraPeso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitaMisuraPeso

`func (o *DatiTrasporto) SetUnitaMisuraPeso(v string)`

SetUnitaMisuraPeso sets UnitaMisuraPeso field to given value.

### HasUnitaMisuraPeso

`func (o *DatiTrasporto) HasUnitaMisuraPeso() bool`

HasUnitaMisuraPeso returns a boolean if a field has been set.

### SetUnitaMisuraPesoNil

`func (o *DatiTrasporto) SetUnitaMisuraPesoNil(b bool)`

 SetUnitaMisuraPesoNil sets the value for UnitaMisuraPeso to be an explicit nil

### UnsetUnitaMisuraPeso
`func (o *DatiTrasporto) UnsetUnitaMisuraPeso()`

UnsetUnitaMisuraPeso ensures that no value is present for UnitaMisuraPeso, not even an explicit nil
### GetPesoLordo

`func (o *DatiTrasporto) GetPesoLordo() float64`

GetPesoLordo returns the PesoLordo field if non-nil, zero value otherwise.

### GetPesoLordoOk

`func (o *DatiTrasporto) GetPesoLordoOk() (*float64, bool)`

GetPesoLordoOk returns a tuple with the PesoLordo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPesoLordo

`func (o *DatiTrasporto) SetPesoLordo(v float64)`

SetPesoLordo sets PesoLordo field to given value.

### HasPesoLordo

`func (o *DatiTrasporto) HasPesoLordo() bool`

HasPesoLordo returns a boolean if a field has been set.

### SetPesoLordoNil

`func (o *DatiTrasporto) SetPesoLordoNil(b bool)`

 SetPesoLordoNil sets the value for PesoLordo to be an explicit nil

### UnsetPesoLordo
`func (o *DatiTrasporto) UnsetPesoLordo()`

UnsetPesoLordo ensures that no value is present for PesoLordo, not even an explicit nil
### GetPesoNetto

`func (o *DatiTrasporto) GetPesoNetto() float64`

GetPesoNetto returns the PesoNetto field if non-nil, zero value otherwise.

### GetPesoNettoOk

`func (o *DatiTrasporto) GetPesoNettoOk() (*float64, bool)`

GetPesoNettoOk returns a tuple with the PesoNetto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPesoNetto

`func (o *DatiTrasporto) SetPesoNetto(v float64)`

SetPesoNetto sets PesoNetto field to given value.

### HasPesoNetto

`func (o *DatiTrasporto) HasPesoNetto() bool`

HasPesoNetto returns a boolean if a field has been set.

### SetPesoNettoNil

`func (o *DatiTrasporto) SetPesoNettoNil(b bool)`

 SetPesoNettoNil sets the value for PesoNetto to be an explicit nil

### UnsetPesoNetto
`func (o *DatiTrasporto) UnsetPesoNetto()`

UnsetPesoNetto ensures that no value is present for PesoNetto, not even an explicit nil
### GetDataOraRitiro

`func (o *DatiTrasporto) GetDataOraRitiro() time.Time`

GetDataOraRitiro returns the DataOraRitiro field if non-nil, zero value otherwise.

### GetDataOraRitiroOk

`func (o *DatiTrasporto) GetDataOraRitiroOk() (*time.Time, bool)`

GetDataOraRitiroOk returns a tuple with the DataOraRitiro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataOraRitiro

`func (o *DatiTrasporto) SetDataOraRitiro(v time.Time)`

SetDataOraRitiro sets DataOraRitiro field to given value.

### HasDataOraRitiro

`func (o *DatiTrasporto) HasDataOraRitiro() bool`

HasDataOraRitiro returns a boolean if a field has been set.

### SetDataOraRitiroNil

`func (o *DatiTrasporto) SetDataOraRitiroNil(b bool)`

 SetDataOraRitiroNil sets the value for DataOraRitiro to be an explicit nil

### UnsetDataOraRitiro
`func (o *DatiTrasporto) UnsetDataOraRitiro()`

UnsetDataOraRitiro ensures that no value is present for DataOraRitiro, not even an explicit nil
### GetDataInizioTrasporto

`func (o *DatiTrasporto) GetDataInizioTrasporto() time.Time`

GetDataInizioTrasporto returns the DataInizioTrasporto field if non-nil, zero value otherwise.

### GetDataInizioTrasportoOk

`func (o *DatiTrasporto) GetDataInizioTrasportoOk() (*time.Time, bool)`

GetDataInizioTrasportoOk returns a tuple with the DataInizioTrasporto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInizioTrasporto

`func (o *DatiTrasporto) SetDataInizioTrasporto(v time.Time)`

SetDataInizioTrasporto sets DataInizioTrasporto field to given value.

### HasDataInizioTrasporto

`func (o *DatiTrasporto) HasDataInizioTrasporto() bool`

HasDataInizioTrasporto returns a boolean if a field has been set.

### SetDataInizioTrasportoNil

`func (o *DatiTrasporto) SetDataInizioTrasportoNil(b bool)`

 SetDataInizioTrasportoNil sets the value for DataInizioTrasporto to be an explicit nil

### UnsetDataInizioTrasporto
`func (o *DatiTrasporto) UnsetDataInizioTrasporto()`

UnsetDataInizioTrasporto ensures that no value is present for DataInizioTrasporto, not even an explicit nil
### GetTipoResa

`func (o *DatiTrasporto) GetTipoResa() string`

GetTipoResa returns the TipoResa field if non-nil, zero value otherwise.

### GetTipoResaOk

`func (o *DatiTrasporto) GetTipoResaOk() (*string, bool)`

GetTipoResaOk returns a tuple with the TipoResa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoResa

`func (o *DatiTrasporto) SetTipoResa(v string)`

SetTipoResa sets TipoResa field to given value.

### HasTipoResa

`func (o *DatiTrasporto) HasTipoResa() bool`

HasTipoResa returns a boolean if a field has been set.

### SetTipoResaNil

`func (o *DatiTrasporto) SetTipoResaNil(b bool)`

 SetTipoResaNil sets the value for TipoResa to be an explicit nil

### UnsetTipoResa
`func (o *DatiTrasporto) UnsetTipoResa()`

UnsetTipoResa ensures that no value is present for TipoResa, not even an explicit nil
### GetIndirizzoResa

`func (o *DatiTrasporto) GetIndirizzoResa() IndirizzoResa`

GetIndirizzoResa returns the IndirizzoResa field if non-nil, zero value otherwise.

### GetIndirizzoResaOk

`func (o *DatiTrasporto) GetIndirizzoResaOk() (*IndirizzoResa, bool)`

GetIndirizzoResaOk returns a tuple with the IndirizzoResa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirizzoResa

`func (o *DatiTrasporto) SetIndirizzoResa(v IndirizzoResa)`

SetIndirizzoResa sets IndirizzoResa field to given value.

### HasIndirizzoResa

`func (o *DatiTrasporto) HasIndirizzoResa() bool`

HasIndirizzoResa returns a boolean if a field has been set.

### GetDataOraConsegna

`func (o *DatiTrasporto) GetDataOraConsegna() time.Time`

GetDataOraConsegna returns the DataOraConsegna field if non-nil, zero value otherwise.

### GetDataOraConsegnaOk

`func (o *DatiTrasporto) GetDataOraConsegnaOk() (*time.Time, bool)`

GetDataOraConsegnaOk returns a tuple with the DataOraConsegna field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataOraConsegna

`func (o *DatiTrasporto) SetDataOraConsegna(v time.Time)`

SetDataOraConsegna sets DataOraConsegna field to given value.

### HasDataOraConsegna

`func (o *DatiTrasporto) HasDataOraConsegna() bool`

HasDataOraConsegna returns a boolean if a field has been set.

### SetDataOraConsegnaNil

`func (o *DatiTrasporto) SetDataOraConsegnaNil(b bool)`

 SetDataOraConsegnaNil sets the value for DataOraConsegna to be an explicit nil

### UnsetDataOraConsegna
`func (o *DatiTrasporto) UnsetDataOraConsegna()`

UnsetDataOraConsegna ensures that no value is present for DataOraConsegna, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


