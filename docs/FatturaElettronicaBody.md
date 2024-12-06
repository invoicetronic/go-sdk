# FatturaElettronicaBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatiGenerali** | Pointer to [**DatiGenerali**](DatiGenerali.md) |  | [optional] 
**DatiBeniServizi** | Pointer to [**DatiBeniServizi**](DatiBeniServizi.md) |  | [optional] 
**DatiVeicoli** | Pointer to [**DatiVeicoli**](DatiVeicoli.md) |  | [optional] 
**DatiPagamento** | Pointer to [**[]DatiPagamento**](DatiPagamento.md) |  | [optional] 
**Allegati** | Pointer to [**[]Allegati**](Allegati.md) |  | [optional] 

## Methods

### NewFatturaElettronicaBody

`func NewFatturaElettronicaBody() *FatturaElettronicaBody`

NewFatturaElettronicaBody instantiates a new FatturaElettronicaBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFatturaElettronicaBodyWithDefaults

`func NewFatturaElettronicaBodyWithDefaults() *FatturaElettronicaBody`

NewFatturaElettronicaBodyWithDefaults instantiates a new FatturaElettronicaBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatiGenerali

`func (o *FatturaElettronicaBody) GetDatiGenerali() DatiGenerali`

GetDatiGenerali returns the DatiGenerali field if non-nil, zero value otherwise.

### GetDatiGeneraliOk

`func (o *FatturaElettronicaBody) GetDatiGeneraliOk() (*DatiGenerali, bool)`

GetDatiGeneraliOk returns a tuple with the DatiGenerali field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiGenerali

`func (o *FatturaElettronicaBody) SetDatiGenerali(v DatiGenerali)`

SetDatiGenerali sets DatiGenerali field to given value.

### HasDatiGenerali

`func (o *FatturaElettronicaBody) HasDatiGenerali() bool`

HasDatiGenerali returns a boolean if a field has been set.

### GetDatiBeniServizi

`func (o *FatturaElettronicaBody) GetDatiBeniServizi() DatiBeniServizi`

GetDatiBeniServizi returns the DatiBeniServizi field if non-nil, zero value otherwise.

### GetDatiBeniServiziOk

`func (o *FatturaElettronicaBody) GetDatiBeniServiziOk() (*DatiBeniServizi, bool)`

GetDatiBeniServiziOk returns a tuple with the DatiBeniServizi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiBeniServizi

`func (o *FatturaElettronicaBody) SetDatiBeniServizi(v DatiBeniServizi)`

SetDatiBeniServizi sets DatiBeniServizi field to given value.

### HasDatiBeniServizi

`func (o *FatturaElettronicaBody) HasDatiBeniServizi() bool`

HasDatiBeniServizi returns a boolean if a field has been set.

### GetDatiVeicoli

`func (o *FatturaElettronicaBody) GetDatiVeicoli() DatiVeicoli`

GetDatiVeicoli returns the DatiVeicoli field if non-nil, zero value otherwise.

### GetDatiVeicoliOk

`func (o *FatturaElettronicaBody) GetDatiVeicoliOk() (*DatiVeicoli, bool)`

GetDatiVeicoliOk returns a tuple with the DatiVeicoli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiVeicoli

`func (o *FatturaElettronicaBody) SetDatiVeicoli(v DatiVeicoli)`

SetDatiVeicoli sets DatiVeicoli field to given value.

### HasDatiVeicoli

`func (o *FatturaElettronicaBody) HasDatiVeicoli() bool`

HasDatiVeicoli returns a boolean if a field has been set.

### GetDatiPagamento

`func (o *FatturaElettronicaBody) GetDatiPagamento() []DatiPagamento`

GetDatiPagamento returns the DatiPagamento field if non-nil, zero value otherwise.

### GetDatiPagamentoOk

`func (o *FatturaElettronicaBody) GetDatiPagamentoOk() (*[]DatiPagamento, bool)`

GetDatiPagamentoOk returns a tuple with the DatiPagamento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiPagamento

`func (o *FatturaElettronicaBody) SetDatiPagamento(v []DatiPagamento)`

SetDatiPagamento sets DatiPagamento field to given value.

### HasDatiPagamento

`func (o *FatturaElettronicaBody) HasDatiPagamento() bool`

HasDatiPagamento returns a boolean if a field has been set.

### SetDatiPagamentoNil

`func (o *FatturaElettronicaBody) SetDatiPagamentoNil(b bool)`

 SetDatiPagamentoNil sets the value for DatiPagamento to be an explicit nil

### UnsetDatiPagamento
`func (o *FatturaElettronicaBody) UnsetDatiPagamento()`

UnsetDatiPagamento ensures that no value is present for DatiPagamento, not even an explicit nil
### GetAllegati

`func (o *FatturaElettronicaBody) GetAllegati() []Allegati`

GetAllegati returns the Allegati field if non-nil, zero value otherwise.

### GetAllegatiOk

`func (o *FatturaElettronicaBody) GetAllegatiOk() (*[]Allegati, bool)`

GetAllegatiOk returns a tuple with the Allegati field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllegati

`func (o *FatturaElettronicaBody) SetAllegati(v []Allegati)`

SetAllegati sets Allegati field to given value.

### HasAllegati

`func (o *FatturaElettronicaBody) HasAllegati() bool`

HasAllegati returns a boolean if a field has been set.

### SetAllegatiNil

`func (o *FatturaElettronicaBody) SetAllegatiNil(b bool)`

 SetAllegatiNil sets the value for Allegati to be an explicit nil

### UnsetAllegati
`func (o *FatturaElettronicaBody) UnsetAllegati()`

UnsetAllegati ensures that no value is present for Allegati, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


