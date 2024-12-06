# DatiAnagraficiCedentePrestatore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdFiscaleIva** | Pointer to [**IdFiscaleIVA**](IdFiscaleIVA.md) |  | [optional] 
**CodiceFiscale** | Pointer to **NullableString** |  | [optional] 
**Anagrafica** | Pointer to [**Anagrafica**](Anagrafica.md) |  | [optional] 
**AlboProfessionale** | Pointer to **NullableString** |  | [optional] 
**ProvinciaAlbo** | Pointer to **NullableString** |  | [optional] 
**NumeroIscrizioneAlbo** | Pointer to **NullableString** |  | [optional] 
**DataIscrizioneAlbo** | Pointer to **NullableTime** |  | [optional] 
**RegimeFiscale** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDatiAnagraficiCedentePrestatore

`func NewDatiAnagraficiCedentePrestatore() *DatiAnagraficiCedentePrestatore`

NewDatiAnagraficiCedentePrestatore instantiates a new DatiAnagraficiCedentePrestatore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatiAnagraficiCedentePrestatoreWithDefaults

`func NewDatiAnagraficiCedentePrestatoreWithDefaults() *DatiAnagraficiCedentePrestatore`

NewDatiAnagraficiCedentePrestatoreWithDefaults instantiates a new DatiAnagraficiCedentePrestatore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdFiscaleIva

`func (o *DatiAnagraficiCedentePrestatore) GetIdFiscaleIva() IdFiscaleIVA`

GetIdFiscaleIva returns the IdFiscaleIva field if non-nil, zero value otherwise.

### GetIdFiscaleIvaOk

`func (o *DatiAnagraficiCedentePrestatore) GetIdFiscaleIvaOk() (*IdFiscaleIVA, bool)`

GetIdFiscaleIvaOk returns a tuple with the IdFiscaleIva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdFiscaleIva

`func (o *DatiAnagraficiCedentePrestatore) SetIdFiscaleIva(v IdFiscaleIVA)`

SetIdFiscaleIva sets IdFiscaleIva field to given value.

### HasIdFiscaleIva

`func (o *DatiAnagraficiCedentePrestatore) HasIdFiscaleIva() bool`

HasIdFiscaleIva returns a boolean if a field has been set.

### GetCodiceFiscale

`func (o *DatiAnagraficiCedentePrestatore) GetCodiceFiscale() string`

GetCodiceFiscale returns the CodiceFiscale field if non-nil, zero value otherwise.

### GetCodiceFiscaleOk

`func (o *DatiAnagraficiCedentePrestatore) GetCodiceFiscaleOk() (*string, bool)`

GetCodiceFiscaleOk returns a tuple with the CodiceFiscale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodiceFiscale

`func (o *DatiAnagraficiCedentePrestatore) SetCodiceFiscale(v string)`

SetCodiceFiscale sets CodiceFiscale field to given value.

### HasCodiceFiscale

`func (o *DatiAnagraficiCedentePrestatore) HasCodiceFiscale() bool`

HasCodiceFiscale returns a boolean if a field has been set.

### SetCodiceFiscaleNil

`func (o *DatiAnagraficiCedentePrestatore) SetCodiceFiscaleNil(b bool)`

 SetCodiceFiscaleNil sets the value for CodiceFiscale to be an explicit nil

### UnsetCodiceFiscale
`func (o *DatiAnagraficiCedentePrestatore) UnsetCodiceFiscale()`

UnsetCodiceFiscale ensures that no value is present for CodiceFiscale, not even an explicit nil
### GetAnagrafica

`func (o *DatiAnagraficiCedentePrestatore) GetAnagrafica() Anagrafica`

GetAnagrafica returns the Anagrafica field if non-nil, zero value otherwise.

### GetAnagraficaOk

`func (o *DatiAnagraficiCedentePrestatore) GetAnagraficaOk() (*Anagrafica, bool)`

GetAnagraficaOk returns a tuple with the Anagrafica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnagrafica

`func (o *DatiAnagraficiCedentePrestatore) SetAnagrafica(v Anagrafica)`

SetAnagrafica sets Anagrafica field to given value.

### HasAnagrafica

`func (o *DatiAnagraficiCedentePrestatore) HasAnagrafica() bool`

HasAnagrafica returns a boolean if a field has been set.

### GetAlboProfessionale

`func (o *DatiAnagraficiCedentePrestatore) GetAlboProfessionale() string`

GetAlboProfessionale returns the AlboProfessionale field if non-nil, zero value otherwise.

### GetAlboProfessionaleOk

`func (o *DatiAnagraficiCedentePrestatore) GetAlboProfessionaleOk() (*string, bool)`

GetAlboProfessionaleOk returns a tuple with the AlboProfessionale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlboProfessionale

`func (o *DatiAnagraficiCedentePrestatore) SetAlboProfessionale(v string)`

SetAlboProfessionale sets AlboProfessionale field to given value.

### HasAlboProfessionale

`func (o *DatiAnagraficiCedentePrestatore) HasAlboProfessionale() bool`

HasAlboProfessionale returns a boolean if a field has been set.

### SetAlboProfessionaleNil

`func (o *DatiAnagraficiCedentePrestatore) SetAlboProfessionaleNil(b bool)`

 SetAlboProfessionaleNil sets the value for AlboProfessionale to be an explicit nil

### UnsetAlboProfessionale
`func (o *DatiAnagraficiCedentePrestatore) UnsetAlboProfessionale()`

UnsetAlboProfessionale ensures that no value is present for AlboProfessionale, not even an explicit nil
### GetProvinciaAlbo

`func (o *DatiAnagraficiCedentePrestatore) GetProvinciaAlbo() string`

GetProvinciaAlbo returns the ProvinciaAlbo field if non-nil, zero value otherwise.

### GetProvinciaAlboOk

`func (o *DatiAnagraficiCedentePrestatore) GetProvinciaAlboOk() (*string, bool)`

GetProvinciaAlboOk returns a tuple with the ProvinciaAlbo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvinciaAlbo

`func (o *DatiAnagraficiCedentePrestatore) SetProvinciaAlbo(v string)`

SetProvinciaAlbo sets ProvinciaAlbo field to given value.

### HasProvinciaAlbo

`func (o *DatiAnagraficiCedentePrestatore) HasProvinciaAlbo() bool`

HasProvinciaAlbo returns a boolean if a field has been set.

### SetProvinciaAlboNil

`func (o *DatiAnagraficiCedentePrestatore) SetProvinciaAlboNil(b bool)`

 SetProvinciaAlboNil sets the value for ProvinciaAlbo to be an explicit nil

### UnsetProvinciaAlbo
`func (o *DatiAnagraficiCedentePrestatore) UnsetProvinciaAlbo()`

UnsetProvinciaAlbo ensures that no value is present for ProvinciaAlbo, not even an explicit nil
### GetNumeroIscrizioneAlbo

`func (o *DatiAnagraficiCedentePrestatore) GetNumeroIscrizioneAlbo() string`

GetNumeroIscrizioneAlbo returns the NumeroIscrizioneAlbo field if non-nil, zero value otherwise.

### GetNumeroIscrizioneAlboOk

`func (o *DatiAnagraficiCedentePrestatore) GetNumeroIscrizioneAlboOk() (*string, bool)`

GetNumeroIscrizioneAlboOk returns a tuple with the NumeroIscrizioneAlbo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroIscrizioneAlbo

`func (o *DatiAnagraficiCedentePrestatore) SetNumeroIscrizioneAlbo(v string)`

SetNumeroIscrizioneAlbo sets NumeroIscrizioneAlbo field to given value.

### HasNumeroIscrizioneAlbo

`func (o *DatiAnagraficiCedentePrestatore) HasNumeroIscrizioneAlbo() bool`

HasNumeroIscrizioneAlbo returns a boolean if a field has been set.

### SetNumeroIscrizioneAlboNil

`func (o *DatiAnagraficiCedentePrestatore) SetNumeroIscrizioneAlboNil(b bool)`

 SetNumeroIscrizioneAlboNil sets the value for NumeroIscrizioneAlbo to be an explicit nil

### UnsetNumeroIscrizioneAlbo
`func (o *DatiAnagraficiCedentePrestatore) UnsetNumeroIscrizioneAlbo()`

UnsetNumeroIscrizioneAlbo ensures that no value is present for NumeroIscrizioneAlbo, not even an explicit nil
### GetDataIscrizioneAlbo

`func (o *DatiAnagraficiCedentePrestatore) GetDataIscrizioneAlbo() time.Time`

GetDataIscrizioneAlbo returns the DataIscrizioneAlbo field if non-nil, zero value otherwise.

### GetDataIscrizioneAlboOk

`func (o *DatiAnagraficiCedentePrestatore) GetDataIscrizioneAlboOk() (*time.Time, bool)`

GetDataIscrizioneAlboOk returns a tuple with the DataIscrizioneAlbo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIscrizioneAlbo

`func (o *DatiAnagraficiCedentePrestatore) SetDataIscrizioneAlbo(v time.Time)`

SetDataIscrizioneAlbo sets DataIscrizioneAlbo field to given value.

### HasDataIscrizioneAlbo

`func (o *DatiAnagraficiCedentePrestatore) HasDataIscrizioneAlbo() bool`

HasDataIscrizioneAlbo returns a boolean if a field has been set.

### SetDataIscrizioneAlboNil

`func (o *DatiAnagraficiCedentePrestatore) SetDataIscrizioneAlboNil(b bool)`

 SetDataIscrizioneAlboNil sets the value for DataIscrizioneAlbo to be an explicit nil

### UnsetDataIscrizioneAlbo
`func (o *DatiAnagraficiCedentePrestatore) UnsetDataIscrizioneAlbo()`

UnsetDataIscrizioneAlbo ensures that no value is present for DataIscrizioneAlbo, not even an explicit nil
### GetRegimeFiscale

`func (o *DatiAnagraficiCedentePrestatore) GetRegimeFiscale() string`

GetRegimeFiscale returns the RegimeFiscale field if non-nil, zero value otherwise.

### GetRegimeFiscaleOk

`func (o *DatiAnagraficiCedentePrestatore) GetRegimeFiscaleOk() (*string, bool)`

GetRegimeFiscaleOk returns a tuple with the RegimeFiscale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegimeFiscale

`func (o *DatiAnagraficiCedentePrestatore) SetRegimeFiscale(v string)`

SetRegimeFiscale sets RegimeFiscale field to given value.

### HasRegimeFiscale

`func (o *DatiAnagraficiCedentePrestatore) HasRegimeFiscale() bool`

HasRegimeFiscale returns a boolean if a field has been set.

### SetRegimeFiscaleNil

`func (o *DatiAnagraficiCedentePrestatore) SetRegimeFiscaleNil(b bool)`

 SetRegimeFiscaleNil sets the value for RegimeFiscale to be an explicit nil

### UnsetRegimeFiscale
`func (o *DatiAnagraficiCedentePrestatore) UnsetRegimeFiscale()`

UnsetRegimeFiscale ensures that no value is present for RegimeFiscale, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


