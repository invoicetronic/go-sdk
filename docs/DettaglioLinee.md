# DettaglioLinee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumeroLinea** | Pointer to **int32** |  | [optional] 
**TipoCessionePrestazione** | Pointer to **NullableString** |  | [optional] 
**CodiceArticolo** | Pointer to [**[]CodiceArticolo**](CodiceArticolo.md) |  | [optional] 
**Descrizione** | Pointer to **NullableString** |  | [optional] 
**Quantita** | Pointer to **NullableFloat64** |  | [optional] 
**UnitaMisura** | Pointer to **NullableString** |  | [optional] 
**DataInizioPeriodo** | Pointer to **NullableTime** |  | [optional] 
**DataFinePeriodo** | Pointer to **NullableTime** |  | [optional] 
**PrezzoUnitario** | Pointer to **float64** |  | [optional] 
**ScontoMaggiorazione** | Pointer to [**[]ScontoMaggiorazione**](ScontoMaggiorazione.md) |  | [optional] 
**PrezzoTotale** | Pointer to **float64** |  | [optional] 
**AliquotaIva** | Pointer to **float64** |  | [optional] 
**Ritenuta** | Pointer to **NullableString** |  | [optional] 
**Natura** | Pointer to **NullableString** |  | [optional] 
**RiferimentoAmministrazione** | Pointer to **NullableString** |  | [optional] 
**AltriDatiGestionali** | Pointer to [**[]AltriDatiGestionali**](AltriDatiGestionali.md) |  | [optional] 

## Methods

### NewDettaglioLinee

`func NewDettaglioLinee() *DettaglioLinee`

NewDettaglioLinee instantiates a new DettaglioLinee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDettaglioLineeWithDefaults

`func NewDettaglioLineeWithDefaults() *DettaglioLinee`

NewDettaglioLineeWithDefaults instantiates a new DettaglioLinee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumeroLinea

`func (o *DettaglioLinee) GetNumeroLinea() int32`

GetNumeroLinea returns the NumeroLinea field if non-nil, zero value otherwise.

### GetNumeroLineaOk

`func (o *DettaglioLinee) GetNumeroLineaOk() (*int32, bool)`

GetNumeroLineaOk returns a tuple with the NumeroLinea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroLinea

`func (o *DettaglioLinee) SetNumeroLinea(v int32)`

SetNumeroLinea sets NumeroLinea field to given value.

### HasNumeroLinea

`func (o *DettaglioLinee) HasNumeroLinea() bool`

HasNumeroLinea returns a boolean if a field has been set.

### GetTipoCessionePrestazione

`func (o *DettaglioLinee) GetTipoCessionePrestazione() string`

GetTipoCessionePrestazione returns the TipoCessionePrestazione field if non-nil, zero value otherwise.

### GetTipoCessionePrestazioneOk

`func (o *DettaglioLinee) GetTipoCessionePrestazioneOk() (*string, bool)`

GetTipoCessionePrestazioneOk returns a tuple with the TipoCessionePrestazione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoCessionePrestazione

`func (o *DettaglioLinee) SetTipoCessionePrestazione(v string)`

SetTipoCessionePrestazione sets TipoCessionePrestazione field to given value.

### HasTipoCessionePrestazione

`func (o *DettaglioLinee) HasTipoCessionePrestazione() bool`

HasTipoCessionePrestazione returns a boolean if a field has been set.

### SetTipoCessionePrestazioneNil

`func (o *DettaglioLinee) SetTipoCessionePrestazioneNil(b bool)`

 SetTipoCessionePrestazioneNil sets the value for TipoCessionePrestazione to be an explicit nil

### UnsetTipoCessionePrestazione
`func (o *DettaglioLinee) UnsetTipoCessionePrestazione()`

UnsetTipoCessionePrestazione ensures that no value is present for TipoCessionePrestazione, not even an explicit nil
### GetCodiceArticolo

`func (o *DettaglioLinee) GetCodiceArticolo() []CodiceArticolo`

GetCodiceArticolo returns the CodiceArticolo field if non-nil, zero value otherwise.

### GetCodiceArticoloOk

`func (o *DettaglioLinee) GetCodiceArticoloOk() (*[]CodiceArticolo, bool)`

GetCodiceArticoloOk returns a tuple with the CodiceArticolo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodiceArticolo

`func (o *DettaglioLinee) SetCodiceArticolo(v []CodiceArticolo)`

SetCodiceArticolo sets CodiceArticolo field to given value.

### HasCodiceArticolo

`func (o *DettaglioLinee) HasCodiceArticolo() bool`

HasCodiceArticolo returns a boolean if a field has been set.

### SetCodiceArticoloNil

`func (o *DettaglioLinee) SetCodiceArticoloNil(b bool)`

 SetCodiceArticoloNil sets the value for CodiceArticolo to be an explicit nil

### UnsetCodiceArticolo
`func (o *DettaglioLinee) UnsetCodiceArticolo()`

UnsetCodiceArticolo ensures that no value is present for CodiceArticolo, not even an explicit nil
### GetDescrizione

`func (o *DettaglioLinee) GetDescrizione() string`

GetDescrizione returns the Descrizione field if non-nil, zero value otherwise.

### GetDescrizioneOk

`func (o *DettaglioLinee) GetDescrizioneOk() (*string, bool)`

GetDescrizioneOk returns a tuple with the Descrizione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescrizione

`func (o *DettaglioLinee) SetDescrizione(v string)`

SetDescrizione sets Descrizione field to given value.

### HasDescrizione

`func (o *DettaglioLinee) HasDescrizione() bool`

HasDescrizione returns a boolean if a field has been set.

### SetDescrizioneNil

`func (o *DettaglioLinee) SetDescrizioneNil(b bool)`

 SetDescrizioneNil sets the value for Descrizione to be an explicit nil

### UnsetDescrizione
`func (o *DettaglioLinee) UnsetDescrizione()`

UnsetDescrizione ensures that no value is present for Descrizione, not even an explicit nil
### GetQuantita

`func (o *DettaglioLinee) GetQuantita() float64`

GetQuantita returns the Quantita field if non-nil, zero value otherwise.

### GetQuantitaOk

`func (o *DettaglioLinee) GetQuantitaOk() (*float64, bool)`

GetQuantitaOk returns a tuple with the Quantita field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantita

`func (o *DettaglioLinee) SetQuantita(v float64)`

SetQuantita sets Quantita field to given value.

### HasQuantita

`func (o *DettaglioLinee) HasQuantita() bool`

HasQuantita returns a boolean if a field has been set.

### SetQuantitaNil

`func (o *DettaglioLinee) SetQuantitaNil(b bool)`

 SetQuantitaNil sets the value for Quantita to be an explicit nil

### UnsetQuantita
`func (o *DettaglioLinee) UnsetQuantita()`

UnsetQuantita ensures that no value is present for Quantita, not even an explicit nil
### GetUnitaMisura

`func (o *DettaglioLinee) GetUnitaMisura() string`

GetUnitaMisura returns the UnitaMisura field if non-nil, zero value otherwise.

### GetUnitaMisuraOk

`func (o *DettaglioLinee) GetUnitaMisuraOk() (*string, bool)`

GetUnitaMisuraOk returns a tuple with the UnitaMisura field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitaMisura

`func (o *DettaglioLinee) SetUnitaMisura(v string)`

SetUnitaMisura sets UnitaMisura field to given value.

### HasUnitaMisura

`func (o *DettaglioLinee) HasUnitaMisura() bool`

HasUnitaMisura returns a boolean if a field has been set.

### SetUnitaMisuraNil

`func (o *DettaglioLinee) SetUnitaMisuraNil(b bool)`

 SetUnitaMisuraNil sets the value for UnitaMisura to be an explicit nil

### UnsetUnitaMisura
`func (o *DettaglioLinee) UnsetUnitaMisura()`

UnsetUnitaMisura ensures that no value is present for UnitaMisura, not even an explicit nil
### GetDataInizioPeriodo

`func (o *DettaglioLinee) GetDataInizioPeriodo() time.Time`

GetDataInizioPeriodo returns the DataInizioPeriodo field if non-nil, zero value otherwise.

### GetDataInizioPeriodoOk

`func (o *DettaglioLinee) GetDataInizioPeriodoOk() (*time.Time, bool)`

GetDataInizioPeriodoOk returns a tuple with the DataInizioPeriodo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInizioPeriodo

`func (o *DettaglioLinee) SetDataInizioPeriodo(v time.Time)`

SetDataInizioPeriodo sets DataInizioPeriodo field to given value.

### HasDataInizioPeriodo

`func (o *DettaglioLinee) HasDataInizioPeriodo() bool`

HasDataInizioPeriodo returns a boolean if a field has been set.

### SetDataInizioPeriodoNil

`func (o *DettaglioLinee) SetDataInizioPeriodoNil(b bool)`

 SetDataInizioPeriodoNil sets the value for DataInizioPeriodo to be an explicit nil

### UnsetDataInizioPeriodo
`func (o *DettaglioLinee) UnsetDataInizioPeriodo()`

UnsetDataInizioPeriodo ensures that no value is present for DataInizioPeriodo, not even an explicit nil
### GetDataFinePeriodo

`func (o *DettaglioLinee) GetDataFinePeriodo() time.Time`

GetDataFinePeriodo returns the DataFinePeriodo field if non-nil, zero value otherwise.

### GetDataFinePeriodoOk

`func (o *DettaglioLinee) GetDataFinePeriodoOk() (*time.Time, bool)`

GetDataFinePeriodoOk returns a tuple with the DataFinePeriodo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFinePeriodo

`func (o *DettaglioLinee) SetDataFinePeriodo(v time.Time)`

SetDataFinePeriodo sets DataFinePeriodo field to given value.

### HasDataFinePeriodo

`func (o *DettaglioLinee) HasDataFinePeriodo() bool`

HasDataFinePeriodo returns a boolean if a field has been set.

### SetDataFinePeriodoNil

`func (o *DettaglioLinee) SetDataFinePeriodoNil(b bool)`

 SetDataFinePeriodoNil sets the value for DataFinePeriodo to be an explicit nil

### UnsetDataFinePeriodo
`func (o *DettaglioLinee) UnsetDataFinePeriodo()`

UnsetDataFinePeriodo ensures that no value is present for DataFinePeriodo, not even an explicit nil
### GetPrezzoUnitario

`func (o *DettaglioLinee) GetPrezzoUnitario() float64`

GetPrezzoUnitario returns the PrezzoUnitario field if non-nil, zero value otherwise.

### GetPrezzoUnitarioOk

`func (o *DettaglioLinee) GetPrezzoUnitarioOk() (*float64, bool)`

GetPrezzoUnitarioOk returns a tuple with the PrezzoUnitario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrezzoUnitario

`func (o *DettaglioLinee) SetPrezzoUnitario(v float64)`

SetPrezzoUnitario sets PrezzoUnitario field to given value.

### HasPrezzoUnitario

`func (o *DettaglioLinee) HasPrezzoUnitario() bool`

HasPrezzoUnitario returns a boolean if a field has been set.

### GetScontoMaggiorazione

`func (o *DettaglioLinee) GetScontoMaggiorazione() []ScontoMaggiorazione`

GetScontoMaggiorazione returns the ScontoMaggiorazione field if non-nil, zero value otherwise.

### GetScontoMaggiorazioneOk

`func (o *DettaglioLinee) GetScontoMaggiorazioneOk() (*[]ScontoMaggiorazione, bool)`

GetScontoMaggiorazioneOk returns a tuple with the ScontoMaggiorazione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScontoMaggiorazione

`func (o *DettaglioLinee) SetScontoMaggiorazione(v []ScontoMaggiorazione)`

SetScontoMaggiorazione sets ScontoMaggiorazione field to given value.

### HasScontoMaggiorazione

`func (o *DettaglioLinee) HasScontoMaggiorazione() bool`

HasScontoMaggiorazione returns a boolean if a field has been set.

### SetScontoMaggiorazioneNil

`func (o *DettaglioLinee) SetScontoMaggiorazioneNil(b bool)`

 SetScontoMaggiorazioneNil sets the value for ScontoMaggiorazione to be an explicit nil

### UnsetScontoMaggiorazione
`func (o *DettaglioLinee) UnsetScontoMaggiorazione()`

UnsetScontoMaggiorazione ensures that no value is present for ScontoMaggiorazione, not even an explicit nil
### GetPrezzoTotale

`func (o *DettaglioLinee) GetPrezzoTotale() float64`

GetPrezzoTotale returns the PrezzoTotale field if non-nil, zero value otherwise.

### GetPrezzoTotaleOk

`func (o *DettaglioLinee) GetPrezzoTotaleOk() (*float64, bool)`

GetPrezzoTotaleOk returns a tuple with the PrezzoTotale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrezzoTotale

`func (o *DettaglioLinee) SetPrezzoTotale(v float64)`

SetPrezzoTotale sets PrezzoTotale field to given value.

### HasPrezzoTotale

`func (o *DettaglioLinee) HasPrezzoTotale() bool`

HasPrezzoTotale returns a boolean if a field has been set.

### GetAliquotaIva

`func (o *DettaglioLinee) GetAliquotaIva() float64`

GetAliquotaIva returns the AliquotaIva field if non-nil, zero value otherwise.

### GetAliquotaIvaOk

`func (o *DettaglioLinee) GetAliquotaIvaOk() (*float64, bool)`

GetAliquotaIvaOk returns a tuple with the AliquotaIva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliquotaIva

`func (o *DettaglioLinee) SetAliquotaIva(v float64)`

SetAliquotaIva sets AliquotaIva field to given value.

### HasAliquotaIva

`func (o *DettaglioLinee) HasAliquotaIva() bool`

HasAliquotaIva returns a boolean if a field has been set.

### GetRitenuta

`func (o *DettaglioLinee) GetRitenuta() string`

GetRitenuta returns the Ritenuta field if non-nil, zero value otherwise.

### GetRitenutaOk

`func (o *DettaglioLinee) GetRitenutaOk() (*string, bool)`

GetRitenutaOk returns a tuple with the Ritenuta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRitenuta

`func (o *DettaglioLinee) SetRitenuta(v string)`

SetRitenuta sets Ritenuta field to given value.

### HasRitenuta

`func (o *DettaglioLinee) HasRitenuta() bool`

HasRitenuta returns a boolean if a field has been set.

### SetRitenutaNil

`func (o *DettaglioLinee) SetRitenutaNil(b bool)`

 SetRitenutaNil sets the value for Ritenuta to be an explicit nil

### UnsetRitenuta
`func (o *DettaglioLinee) UnsetRitenuta()`

UnsetRitenuta ensures that no value is present for Ritenuta, not even an explicit nil
### GetNatura

`func (o *DettaglioLinee) GetNatura() string`

GetNatura returns the Natura field if non-nil, zero value otherwise.

### GetNaturaOk

`func (o *DettaglioLinee) GetNaturaOk() (*string, bool)`

GetNaturaOk returns a tuple with the Natura field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatura

`func (o *DettaglioLinee) SetNatura(v string)`

SetNatura sets Natura field to given value.

### HasNatura

`func (o *DettaglioLinee) HasNatura() bool`

HasNatura returns a boolean if a field has been set.

### SetNaturaNil

`func (o *DettaglioLinee) SetNaturaNil(b bool)`

 SetNaturaNil sets the value for Natura to be an explicit nil

### UnsetNatura
`func (o *DettaglioLinee) UnsetNatura()`

UnsetNatura ensures that no value is present for Natura, not even an explicit nil
### GetRiferimentoAmministrazione

`func (o *DettaglioLinee) GetRiferimentoAmministrazione() string`

GetRiferimentoAmministrazione returns the RiferimentoAmministrazione field if non-nil, zero value otherwise.

### GetRiferimentoAmministrazioneOk

`func (o *DettaglioLinee) GetRiferimentoAmministrazioneOk() (*string, bool)`

GetRiferimentoAmministrazioneOk returns a tuple with the RiferimentoAmministrazione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiferimentoAmministrazione

`func (o *DettaglioLinee) SetRiferimentoAmministrazione(v string)`

SetRiferimentoAmministrazione sets RiferimentoAmministrazione field to given value.

### HasRiferimentoAmministrazione

`func (o *DettaglioLinee) HasRiferimentoAmministrazione() bool`

HasRiferimentoAmministrazione returns a boolean if a field has been set.

### SetRiferimentoAmministrazioneNil

`func (o *DettaglioLinee) SetRiferimentoAmministrazioneNil(b bool)`

 SetRiferimentoAmministrazioneNil sets the value for RiferimentoAmministrazione to be an explicit nil

### UnsetRiferimentoAmministrazione
`func (o *DettaglioLinee) UnsetRiferimentoAmministrazione()`

UnsetRiferimentoAmministrazione ensures that no value is present for RiferimentoAmministrazione, not even an explicit nil
### GetAltriDatiGestionali

`func (o *DettaglioLinee) GetAltriDatiGestionali() []AltriDatiGestionali`

GetAltriDatiGestionali returns the AltriDatiGestionali field if non-nil, zero value otherwise.

### GetAltriDatiGestionaliOk

`func (o *DettaglioLinee) GetAltriDatiGestionaliOk() (*[]AltriDatiGestionali, bool)`

GetAltriDatiGestionaliOk returns a tuple with the AltriDatiGestionali field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltriDatiGestionali

`func (o *DettaglioLinee) SetAltriDatiGestionali(v []AltriDatiGestionali)`

SetAltriDatiGestionali sets AltriDatiGestionali field to given value.

### HasAltriDatiGestionali

`func (o *DettaglioLinee) HasAltriDatiGestionali() bool`

HasAltriDatiGestionali returns a boolean if a field has been set.

### SetAltriDatiGestionaliNil

`func (o *DettaglioLinee) SetAltriDatiGestionaliNil(b bool)`

 SetAltriDatiGestionaliNil sets the value for AltriDatiGestionali to be an explicit nil

### UnsetAltriDatiGestionali
`func (o *DettaglioLinee) UnsetAltriDatiGestionali()`

UnsetAltriDatiGestionali ensures that no value is present for AltriDatiGestionali, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


