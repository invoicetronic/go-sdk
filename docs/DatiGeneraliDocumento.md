# DatiGeneraliDocumento

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TipoDocumento** | Pointer to **NullableString** |  | [optional] 
**Divisa** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **time.Time** |  | [optional] 
**Numero** | Pointer to **NullableString** |  | [optional] 
**DatiRitenuta** | Pointer to [**[]DatiRitenuta**](DatiRitenuta.md) |  | [optional] 
**DatiBollo** | Pointer to [**DatiBollo**](DatiBollo.md) |  | [optional] 
**DatiCassaPrevidenziale** | Pointer to [**[]DatiCassaPrevidenziale**](DatiCassaPrevidenziale.md) |  | [optional] 
**ScontoMaggiorazione** | Pointer to [**[]ScontoMaggiorazione**](ScontoMaggiorazione.md) |  | [optional] 
**ImportoTotaleDocumento** | Pointer to **NullableFloat64** |  | [optional] 
**Arrotondamento** | Pointer to **NullableFloat64** |  | [optional] 
**Causale** | Pointer to **[]string** |  | [optional] 
**Art73** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDatiGeneraliDocumento

`func NewDatiGeneraliDocumento() *DatiGeneraliDocumento`

NewDatiGeneraliDocumento instantiates a new DatiGeneraliDocumento object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatiGeneraliDocumentoWithDefaults

`func NewDatiGeneraliDocumentoWithDefaults() *DatiGeneraliDocumento`

NewDatiGeneraliDocumentoWithDefaults instantiates a new DatiGeneraliDocumento object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTipoDocumento

`func (o *DatiGeneraliDocumento) GetTipoDocumento() string`

GetTipoDocumento returns the TipoDocumento field if non-nil, zero value otherwise.

### GetTipoDocumentoOk

`func (o *DatiGeneraliDocumento) GetTipoDocumentoOk() (*string, bool)`

GetTipoDocumentoOk returns a tuple with the TipoDocumento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoDocumento

`func (o *DatiGeneraliDocumento) SetTipoDocumento(v string)`

SetTipoDocumento sets TipoDocumento field to given value.

### HasTipoDocumento

`func (o *DatiGeneraliDocumento) HasTipoDocumento() bool`

HasTipoDocumento returns a boolean if a field has been set.

### SetTipoDocumentoNil

`func (o *DatiGeneraliDocumento) SetTipoDocumentoNil(b bool)`

 SetTipoDocumentoNil sets the value for TipoDocumento to be an explicit nil

### UnsetTipoDocumento
`func (o *DatiGeneraliDocumento) UnsetTipoDocumento()`

UnsetTipoDocumento ensures that no value is present for TipoDocumento, not even an explicit nil
### GetDivisa

`func (o *DatiGeneraliDocumento) GetDivisa() string`

GetDivisa returns the Divisa field if non-nil, zero value otherwise.

### GetDivisaOk

`func (o *DatiGeneraliDocumento) GetDivisaOk() (*string, bool)`

GetDivisaOk returns a tuple with the Divisa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisa

`func (o *DatiGeneraliDocumento) SetDivisa(v string)`

SetDivisa sets Divisa field to given value.

### HasDivisa

`func (o *DatiGeneraliDocumento) HasDivisa() bool`

HasDivisa returns a boolean if a field has been set.

### SetDivisaNil

`func (o *DatiGeneraliDocumento) SetDivisaNil(b bool)`

 SetDivisaNil sets the value for Divisa to be an explicit nil

### UnsetDivisa
`func (o *DatiGeneraliDocumento) UnsetDivisa()`

UnsetDivisa ensures that no value is present for Divisa, not even an explicit nil
### GetData

`func (o *DatiGeneraliDocumento) GetData() time.Time`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatiGeneraliDocumento) GetDataOk() (*time.Time, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatiGeneraliDocumento) SetData(v time.Time)`

SetData sets Data field to given value.

### HasData

`func (o *DatiGeneraliDocumento) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNumero

`func (o *DatiGeneraliDocumento) GetNumero() string`

GetNumero returns the Numero field if non-nil, zero value otherwise.

### GetNumeroOk

`func (o *DatiGeneraliDocumento) GetNumeroOk() (*string, bool)`

GetNumeroOk returns a tuple with the Numero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumero

`func (o *DatiGeneraliDocumento) SetNumero(v string)`

SetNumero sets Numero field to given value.

### HasNumero

`func (o *DatiGeneraliDocumento) HasNumero() bool`

HasNumero returns a boolean if a field has been set.

### SetNumeroNil

`func (o *DatiGeneraliDocumento) SetNumeroNil(b bool)`

 SetNumeroNil sets the value for Numero to be an explicit nil

### UnsetNumero
`func (o *DatiGeneraliDocumento) UnsetNumero()`

UnsetNumero ensures that no value is present for Numero, not even an explicit nil
### GetDatiRitenuta

`func (o *DatiGeneraliDocumento) GetDatiRitenuta() []DatiRitenuta`

GetDatiRitenuta returns the DatiRitenuta field if non-nil, zero value otherwise.

### GetDatiRitenutaOk

`func (o *DatiGeneraliDocumento) GetDatiRitenutaOk() (*[]DatiRitenuta, bool)`

GetDatiRitenutaOk returns a tuple with the DatiRitenuta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiRitenuta

`func (o *DatiGeneraliDocumento) SetDatiRitenuta(v []DatiRitenuta)`

SetDatiRitenuta sets DatiRitenuta field to given value.

### HasDatiRitenuta

`func (o *DatiGeneraliDocumento) HasDatiRitenuta() bool`

HasDatiRitenuta returns a boolean if a field has been set.

### SetDatiRitenutaNil

`func (o *DatiGeneraliDocumento) SetDatiRitenutaNil(b bool)`

 SetDatiRitenutaNil sets the value for DatiRitenuta to be an explicit nil

### UnsetDatiRitenuta
`func (o *DatiGeneraliDocumento) UnsetDatiRitenuta()`

UnsetDatiRitenuta ensures that no value is present for DatiRitenuta, not even an explicit nil
### GetDatiBollo

`func (o *DatiGeneraliDocumento) GetDatiBollo() DatiBollo`

GetDatiBollo returns the DatiBollo field if non-nil, zero value otherwise.

### GetDatiBolloOk

`func (o *DatiGeneraliDocumento) GetDatiBolloOk() (*DatiBollo, bool)`

GetDatiBolloOk returns a tuple with the DatiBollo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiBollo

`func (o *DatiGeneraliDocumento) SetDatiBollo(v DatiBollo)`

SetDatiBollo sets DatiBollo field to given value.

### HasDatiBollo

`func (o *DatiGeneraliDocumento) HasDatiBollo() bool`

HasDatiBollo returns a boolean if a field has been set.

### GetDatiCassaPrevidenziale

`func (o *DatiGeneraliDocumento) GetDatiCassaPrevidenziale() []DatiCassaPrevidenziale`

GetDatiCassaPrevidenziale returns the DatiCassaPrevidenziale field if non-nil, zero value otherwise.

### GetDatiCassaPrevidenzialeOk

`func (o *DatiGeneraliDocumento) GetDatiCassaPrevidenzialeOk() (*[]DatiCassaPrevidenziale, bool)`

GetDatiCassaPrevidenzialeOk returns a tuple with the DatiCassaPrevidenziale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiCassaPrevidenziale

`func (o *DatiGeneraliDocumento) SetDatiCassaPrevidenziale(v []DatiCassaPrevidenziale)`

SetDatiCassaPrevidenziale sets DatiCassaPrevidenziale field to given value.

### HasDatiCassaPrevidenziale

`func (o *DatiGeneraliDocumento) HasDatiCassaPrevidenziale() bool`

HasDatiCassaPrevidenziale returns a boolean if a field has been set.

### SetDatiCassaPrevidenzialeNil

`func (o *DatiGeneraliDocumento) SetDatiCassaPrevidenzialeNil(b bool)`

 SetDatiCassaPrevidenzialeNil sets the value for DatiCassaPrevidenziale to be an explicit nil

### UnsetDatiCassaPrevidenziale
`func (o *DatiGeneraliDocumento) UnsetDatiCassaPrevidenziale()`

UnsetDatiCassaPrevidenziale ensures that no value is present for DatiCassaPrevidenziale, not even an explicit nil
### GetScontoMaggiorazione

`func (o *DatiGeneraliDocumento) GetScontoMaggiorazione() []ScontoMaggiorazione`

GetScontoMaggiorazione returns the ScontoMaggiorazione field if non-nil, zero value otherwise.

### GetScontoMaggiorazioneOk

`func (o *DatiGeneraliDocumento) GetScontoMaggiorazioneOk() (*[]ScontoMaggiorazione, bool)`

GetScontoMaggiorazioneOk returns a tuple with the ScontoMaggiorazione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScontoMaggiorazione

`func (o *DatiGeneraliDocumento) SetScontoMaggiorazione(v []ScontoMaggiorazione)`

SetScontoMaggiorazione sets ScontoMaggiorazione field to given value.

### HasScontoMaggiorazione

`func (o *DatiGeneraliDocumento) HasScontoMaggiorazione() bool`

HasScontoMaggiorazione returns a boolean if a field has been set.

### SetScontoMaggiorazioneNil

`func (o *DatiGeneraliDocumento) SetScontoMaggiorazioneNil(b bool)`

 SetScontoMaggiorazioneNil sets the value for ScontoMaggiorazione to be an explicit nil

### UnsetScontoMaggiorazione
`func (o *DatiGeneraliDocumento) UnsetScontoMaggiorazione()`

UnsetScontoMaggiorazione ensures that no value is present for ScontoMaggiorazione, not even an explicit nil
### GetImportoTotaleDocumento

`func (o *DatiGeneraliDocumento) GetImportoTotaleDocumento() float64`

GetImportoTotaleDocumento returns the ImportoTotaleDocumento field if non-nil, zero value otherwise.

### GetImportoTotaleDocumentoOk

`func (o *DatiGeneraliDocumento) GetImportoTotaleDocumentoOk() (*float64, bool)`

GetImportoTotaleDocumentoOk returns a tuple with the ImportoTotaleDocumento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportoTotaleDocumento

`func (o *DatiGeneraliDocumento) SetImportoTotaleDocumento(v float64)`

SetImportoTotaleDocumento sets ImportoTotaleDocumento field to given value.

### HasImportoTotaleDocumento

`func (o *DatiGeneraliDocumento) HasImportoTotaleDocumento() bool`

HasImportoTotaleDocumento returns a boolean if a field has been set.

### SetImportoTotaleDocumentoNil

`func (o *DatiGeneraliDocumento) SetImportoTotaleDocumentoNil(b bool)`

 SetImportoTotaleDocumentoNil sets the value for ImportoTotaleDocumento to be an explicit nil

### UnsetImportoTotaleDocumento
`func (o *DatiGeneraliDocumento) UnsetImportoTotaleDocumento()`

UnsetImportoTotaleDocumento ensures that no value is present for ImportoTotaleDocumento, not even an explicit nil
### GetArrotondamento

`func (o *DatiGeneraliDocumento) GetArrotondamento() float64`

GetArrotondamento returns the Arrotondamento field if non-nil, zero value otherwise.

### GetArrotondamentoOk

`func (o *DatiGeneraliDocumento) GetArrotondamentoOk() (*float64, bool)`

GetArrotondamentoOk returns a tuple with the Arrotondamento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrotondamento

`func (o *DatiGeneraliDocumento) SetArrotondamento(v float64)`

SetArrotondamento sets Arrotondamento field to given value.

### HasArrotondamento

`func (o *DatiGeneraliDocumento) HasArrotondamento() bool`

HasArrotondamento returns a boolean if a field has been set.

### SetArrotondamentoNil

`func (o *DatiGeneraliDocumento) SetArrotondamentoNil(b bool)`

 SetArrotondamentoNil sets the value for Arrotondamento to be an explicit nil

### UnsetArrotondamento
`func (o *DatiGeneraliDocumento) UnsetArrotondamento()`

UnsetArrotondamento ensures that no value is present for Arrotondamento, not even an explicit nil
### GetCausale

`func (o *DatiGeneraliDocumento) GetCausale() []string`

GetCausale returns the Causale field if non-nil, zero value otherwise.

### GetCausaleOk

`func (o *DatiGeneraliDocumento) GetCausaleOk() (*[]string, bool)`

GetCausaleOk returns a tuple with the Causale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCausale

`func (o *DatiGeneraliDocumento) SetCausale(v []string)`

SetCausale sets Causale field to given value.

### HasCausale

`func (o *DatiGeneraliDocumento) HasCausale() bool`

HasCausale returns a boolean if a field has been set.

### SetCausaleNil

`func (o *DatiGeneraliDocumento) SetCausaleNil(b bool)`

 SetCausaleNil sets the value for Causale to be an explicit nil

### UnsetCausale
`func (o *DatiGeneraliDocumento) UnsetCausale()`

UnsetCausale ensures that no value is present for Causale, not even an explicit nil
### GetArt73

`func (o *DatiGeneraliDocumento) GetArt73() string`

GetArt73 returns the Art73 field if non-nil, zero value otherwise.

### GetArt73Ok

`func (o *DatiGeneraliDocumento) GetArt73Ok() (*string, bool)`

GetArt73Ok returns a tuple with the Art73 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArt73

`func (o *DatiGeneraliDocumento) SetArt73(v string)`

SetArt73 sets Art73 field to given value.

### HasArt73

`func (o *DatiGeneraliDocumento) HasArt73() bool`

HasArt73 returns a boolean if a field has been set.

### SetArt73Nil

`func (o *DatiGeneraliDocumento) SetArt73Nil(b bool)`

 SetArt73Nil sets the value for Art73 to be an explicit nil

### UnsetArt73
`func (o *DatiGeneraliDocumento) UnsetArt73()`

UnsetArt73 ensures that no value is present for Art73, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


