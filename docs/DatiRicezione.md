# DatiRicezione

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiferimentoNumeroLinea** | Pointer to **[]int32** |  | [optional] 
**IdDocumento** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **NullableTime** |  | [optional] 
**NumItem** | Pointer to **NullableString** |  | [optional] 
**CodiceCommessaConvenzione** | Pointer to **NullableString** |  | [optional] 
**CodiceCup** | Pointer to **NullableString** |  | [optional] 
**CodiceCig** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDatiRicezione

`func NewDatiRicezione() *DatiRicezione`

NewDatiRicezione instantiates a new DatiRicezione object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatiRicezioneWithDefaults

`func NewDatiRicezioneWithDefaults() *DatiRicezione`

NewDatiRicezioneWithDefaults instantiates a new DatiRicezione object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiferimentoNumeroLinea

`func (o *DatiRicezione) GetRiferimentoNumeroLinea() []int32`

GetRiferimentoNumeroLinea returns the RiferimentoNumeroLinea field if non-nil, zero value otherwise.

### GetRiferimentoNumeroLineaOk

`func (o *DatiRicezione) GetRiferimentoNumeroLineaOk() (*[]int32, bool)`

GetRiferimentoNumeroLineaOk returns a tuple with the RiferimentoNumeroLinea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiferimentoNumeroLinea

`func (o *DatiRicezione) SetRiferimentoNumeroLinea(v []int32)`

SetRiferimentoNumeroLinea sets RiferimentoNumeroLinea field to given value.

### HasRiferimentoNumeroLinea

`func (o *DatiRicezione) HasRiferimentoNumeroLinea() bool`

HasRiferimentoNumeroLinea returns a boolean if a field has been set.

### SetRiferimentoNumeroLineaNil

`func (o *DatiRicezione) SetRiferimentoNumeroLineaNil(b bool)`

 SetRiferimentoNumeroLineaNil sets the value for RiferimentoNumeroLinea to be an explicit nil

### UnsetRiferimentoNumeroLinea
`func (o *DatiRicezione) UnsetRiferimentoNumeroLinea()`

UnsetRiferimentoNumeroLinea ensures that no value is present for RiferimentoNumeroLinea, not even an explicit nil
### GetIdDocumento

`func (o *DatiRicezione) GetIdDocumento() string`

GetIdDocumento returns the IdDocumento field if non-nil, zero value otherwise.

### GetIdDocumentoOk

`func (o *DatiRicezione) GetIdDocumentoOk() (*string, bool)`

GetIdDocumentoOk returns a tuple with the IdDocumento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdDocumento

`func (o *DatiRicezione) SetIdDocumento(v string)`

SetIdDocumento sets IdDocumento field to given value.

### HasIdDocumento

`func (o *DatiRicezione) HasIdDocumento() bool`

HasIdDocumento returns a boolean if a field has been set.

### SetIdDocumentoNil

`func (o *DatiRicezione) SetIdDocumentoNil(b bool)`

 SetIdDocumentoNil sets the value for IdDocumento to be an explicit nil

### UnsetIdDocumento
`func (o *DatiRicezione) UnsetIdDocumento()`

UnsetIdDocumento ensures that no value is present for IdDocumento, not even an explicit nil
### GetData

`func (o *DatiRicezione) GetData() time.Time`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DatiRicezione) GetDataOk() (*time.Time, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DatiRicezione) SetData(v time.Time)`

SetData sets Data field to given value.

### HasData

`func (o *DatiRicezione) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *DatiRicezione) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *DatiRicezione) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetNumItem

`func (o *DatiRicezione) GetNumItem() string`

GetNumItem returns the NumItem field if non-nil, zero value otherwise.

### GetNumItemOk

`func (o *DatiRicezione) GetNumItemOk() (*string, bool)`

GetNumItemOk returns a tuple with the NumItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumItem

`func (o *DatiRicezione) SetNumItem(v string)`

SetNumItem sets NumItem field to given value.

### HasNumItem

`func (o *DatiRicezione) HasNumItem() bool`

HasNumItem returns a boolean if a field has been set.

### SetNumItemNil

`func (o *DatiRicezione) SetNumItemNil(b bool)`

 SetNumItemNil sets the value for NumItem to be an explicit nil

### UnsetNumItem
`func (o *DatiRicezione) UnsetNumItem()`

UnsetNumItem ensures that no value is present for NumItem, not even an explicit nil
### GetCodiceCommessaConvenzione

`func (o *DatiRicezione) GetCodiceCommessaConvenzione() string`

GetCodiceCommessaConvenzione returns the CodiceCommessaConvenzione field if non-nil, zero value otherwise.

### GetCodiceCommessaConvenzioneOk

`func (o *DatiRicezione) GetCodiceCommessaConvenzioneOk() (*string, bool)`

GetCodiceCommessaConvenzioneOk returns a tuple with the CodiceCommessaConvenzione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodiceCommessaConvenzione

`func (o *DatiRicezione) SetCodiceCommessaConvenzione(v string)`

SetCodiceCommessaConvenzione sets CodiceCommessaConvenzione field to given value.

### HasCodiceCommessaConvenzione

`func (o *DatiRicezione) HasCodiceCommessaConvenzione() bool`

HasCodiceCommessaConvenzione returns a boolean if a field has been set.

### SetCodiceCommessaConvenzioneNil

`func (o *DatiRicezione) SetCodiceCommessaConvenzioneNil(b bool)`

 SetCodiceCommessaConvenzioneNil sets the value for CodiceCommessaConvenzione to be an explicit nil

### UnsetCodiceCommessaConvenzione
`func (o *DatiRicezione) UnsetCodiceCommessaConvenzione()`

UnsetCodiceCommessaConvenzione ensures that no value is present for CodiceCommessaConvenzione, not even an explicit nil
### GetCodiceCup

`func (o *DatiRicezione) GetCodiceCup() string`

GetCodiceCup returns the CodiceCup field if non-nil, zero value otherwise.

### GetCodiceCupOk

`func (o *DatiRicezione) GetCodiceCupOk() (*string, bool)`

GetCodiceCupOk returns a tuple with the CodiceCup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodiceCup

`func (o *DatiRicezione) SetCodiceCup(v string)`

SetCodiceCup sets CodiceCup field to given value.

### HasCodiceCup

`func (o *DatiRicezione) HasCodiceCup() bool`

HasCodiceCup returns a boolean if a field has been set.

### SetCodiceCupNil

`func (o *DatiRicezione) SetCodiceCupNil(b bool)`

 SetCodiceCupNil sets the value for CodiceCup to be an explicit nil

### UnsetCodiceCup
`func (o *DatiRicezione) UnsetCodiceCup()`

UnsetCodiceCup ensures that no value is present for CodiceCup, not even an explicit nil
### GetCodiceCig

`func (o *DatiRicezione) GetCodiceCig() string`

GetCodiceCig returns the CodiceCig field if non-nil, zero value otherwise.

### GetCodiceCigOk

`func (o *DatiRicezione) GetCodiceCigOk() (*string, bool)`

GetCodiceCigOk returns a tuple with the CodiceCig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodiceCig

`func (o *DatiRicezione) SetCodiceCig(v string)`

SetCodiceCig sets CodiceCig field to given value.

### HasCodiceCig

`func (o *DatiRicezione) HasCodiceCig() bool`

HasCodiceCig returns a boolean if a field has been set.

### SetCodiceCigNil

`func (o *DatiRicezione) SetCodiceCigNil(b bool)`

 SetCodiceCigNil sets the value for CodiceCig to be an explicit nil

### UnsetCodiceCig
`func (o *DatiRicezione) UnsetCodiceCig()`

UnsetCodiceCig ensures that no value is present for CodiceCig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


