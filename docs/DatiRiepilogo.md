# DatiRiepilogo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliquotaIva** | Pointer to **float64** |  | [optional] 
**Natura** | Pointer to **NullableString** |  | [optional] 
**SpeseAccessorie** | Pointer to **NullableFloat64** |  | [optional] 
**Arrotondamento** | Pointer to **NullableFloat64** |  | [optional] 
**ImponibileImporto** | Pointer to **float64** |  | [optional] 
**Imposta** | Pointer to **float64** |  | [optional] 
**EsigibilitaIva** | Pointer to **NullableString** |  | [optional] 
**RiferimentoNormativo** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDatiRiepilogo

`func NewDatiRiepilogo() *DatiRiepilogo`

NewDatiRiepilogo instantiates a new DatiRiepilogo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatiRiepilogoWithDefaults

`func NewDatiRiepilogoWithDefaults() *DatiRiepilogo`

NewDatiRiepilogoWithDefaults instantiates a new DatiRiepilogo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliquotaIva

`func (o *DatiRiepilogo) GetAliquotaIva() float64`

GetAliquotaIva returns the AliquotaIva field if non-nil, zero value otherwise.

### GetAliquotaIvaOk

`func (o *DatiRiepilogo) GetAliquotaIvaOk() (*float64, bool)`

GetAliquotaIvaOk returns a tuple with the AliquotaIva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliquotaIva

`func (o *DatiRiepilogo) SetAliquotaIva(v float64)`

SetAliquotaIva sets AliquotaIva field to given value.

### HasAliquotaIva

`func (o *DatiRiepilogo) HasAliquotaIva() bool`

HasAliquotaIva returns a boolean if a field has been set.

### GetNatura

`func (o *DatiRiepilogo) GetNatura() string`

GetNatura returns the Natura field if non-nil, zero value otherwise.

### GetNaturaOk

`func (o *DatiRiepilogo) GetNaturaOk() (*string, bool)`

GetNaturaOk returns a tuple with the Natura field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatura

`func (o *DatiRiepilogo) SetNatura(v string)`

SetNatura sets Natura field to given value.

### HasNatura

`func (o *DatiRiepilogo) HasNatura() bool`

HasNatura returns a boolean if a field has been set.

### SetNaturaNil

`func (o *DatiRiepilogo) SetNaturaNil(b bool)`

 SetNaturaNil sets the value for Natura to be an explicit nil

### UnsetNatura
`func (o *DatiRiepilogo) UnsetNatura()`

UnsetNatura ensures that no value is present for Natura, not even an explicit nil
### GetSpeseAccessorie

`func (o *DatiRiepilogo) GetSpeseAccessorie() float64`

GetSpeseAccessorie returns the SpeseAccessorie field if non-nil, zero value otherwise.

### GetSpeseAccessorieOk

`func (o *DatiRiepilogo) GetSpeseAccessorieOk() (*float64, bool)`

GetSpeseAccessorieOk returns a tuple with the SpeseAccessorie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeseAccessorie

`func (o *DatiRiepilogo) SetSpeseAccessorie(v float64)`

SetSpeseAccessorie sets SpeseAccessorie field to given value.

### HasSpeseAccessorie

`func (o *DatiRiepilogo) HasSpeseAccessorie() bool`

HasSpeseAccessorie returns a boolean if a field has been set.

### SetSpeseAccessorieNil

`func (o *DatiRiepilogo) SetSpeseAccessorieNil(b bool)`

 SetSpeseAccessorieNil sets the value for SpeseAccessorie to be an explicit nil

### UnsetSpeseAccessorie
`func (o *DatiRiepilogo) UnsetSpeseAccessorie()`

UnsetSpeseAccessorie ensures that no value is present for SpeseAccessorie, not even an explicit nil
### GetArrotondamento

`func (o *DatiRiepilogo) GetArrotondamento() float64`

GetArrotondamento returns the Arrotondamento field if non-nil, zero value otherwise.

### GetArrotondamentoOk

`func (o *DatiRiepilogo) GetArrotondamentoOk() (*float64, bool)`

GetArrotondamentoOk returns a tuple with the Arrotondamento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrotondamento

`func (o *DatiRiepilogo) SetArrotondamento(v float64)`

SetArrotondamento sets Arrotondamento field to given value.

### HasArrotondamento

`func (o *DatiRiepilogo) HasArrotondamento() bool`

HasArrotondamento returns a boolean if a field has been set.

### SetArrotondamentoNil

`func (o *DatiRiepilogo) SetArrotondamentoNil(b bool)`

 SetArrotondamentoNil sets the value for Arrotondamento to be an explicit nil

### UnsetArrotondamento
`func (o *DatiRiepilogo) UnsetArrotondamento()`

UnsetArrotondamento ensures that no value is present for Arrotondamento, not even an explicit nil
### GetImponibileImporto

`func (o *DatiRiepilogo) GetImponibileImporto() float64`

GetImponibileImporto returns the ImponibileImporto field if non-nil, zero value otherwise.

### GetImponibileImportoOk

`func (o *DatiRiepilogo) GetImponibileImportoOk() (*float64, bool)`

GetImponibileImportoOk returns a tuple with the ImponibileImporto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImponibileImporto

`func (o *DatiRiepilogo) SetImponibileImporto(v float64)`

SetImponibileImporto sets ImponibileImporto field to given value.

### HasImponibileImporto

`func (o *DatiRiepilogo) HasImponibileImporto() bool`

HasImponibileImporto returns a boolean if a field has been set.

### GetImposta

`func (o *DatiRiepilogo) GetImposta() float64`

GetImposta returns the Imposta field if non-nil, zero value otherwise.

### GetImpostaOk

`func (o *DatiRiepilogo) GetImpostaOk() (*float64, bool)`

GetImpostaOk returns a tuple with the Imposta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImposta

`func (o *DatiRiepilogo) SetImposta(v float64)`

SetImposta sets Imposta field to given value.

### HasImposta

`func (o *DatiRiepilogo) HasImposta() bool`

HasImposta returns a boolean if a field has been set.

### GetEsigibilitaIva

`func (o *DatiRiepilogo) GetEsigibilitaIva() string`

GetEsigibilitaIva returns the EsigibilitaIva field if non-nil, zero value otherwise.

### GetEsigibilitaIvaOk

`func (o *DatiRiepilogo) GetEsigibilitaIvaOk() (*string, bool)`

GetEsigibilitaIvaOk returns a tuple with the EsigibilitaIva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsigibilitaIva

`func (o *DatiRiepilogo) SetEsigibilitaIva(v string)`

SetEsigibilitaIva sets EsigibilitaIva field to given value.

### HasEsigibilitaIva

`func (o *DatiRiepilogo) HasEsigibilitaIva() bool`

HasEsigibilitaIva returns a boolean if a field has been set.

### SetEsigibilitaIvaNil

`func (o *DatiRiepilogo) SetEsigibilitaIvaNil(b bool)`

 SetEsigibilitaIvaNil sets the value for EsigibilitaIva to be an explicit nil

### UnsetEsigibilitaIva
`func (o *DatiRiepilogo) UnsetEsigibilitaIva()`

UnsetEsigibilitaIva ensures that no value is present for EsigibilitaIva, not even an explicit nil
### GetRiferimentoNormativo

`func (o *DatiRiepilogo) GetRiferimentoNormativo() string`

GetRiferimentoNormativo returns the RiferimentoNormativo field if non-nil, zero value otherwise.

### GetRiferimentoNormativoOk

`func (o *DatiRiepilogo) GetRiferimentoNormativoOk() (*string, bool)`

GetRiferimentoNormativoOk returns a tuple with the RiferimentoNormativo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiferimentoNormativo

`func (o *DatiRiepilogo) SetRiferimentoNormativo(v string)`

SetRiferimentoNormativo sets RiferimentoNormativo field to given value.

### HasRiferimentoNormativo

`func (o *DatiRiepilogo) HasRiferimentoNormativo() bool`

HasRiferimentoNormativo returns a boolean if a field has been set.

### SetRiferimentoNormativoNil

`func (o *DatiRiepilogo) SetRiferimentoNormativoNil(b bool)`

 SetRiferimentoNormativoNil sets the value for RiferimentoNormativo to be an explicit nil

### UnsetRiferimentoNormativo
`func (o *DatiRiepilogo) UnsetRiferimentoNormativo()`

UnsetRiferimentoNormativo ensures that no value is present for RiferimentoNormativo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


