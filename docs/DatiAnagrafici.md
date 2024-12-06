# DatiAnagrafici

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdFiscaleIva** | Pointer to [**IdFiscaleIVA**](IdFiscaleIVA.md) |  | [optional] 
**CodiceFiscale** | Pointer to **NullableString** |  | [optional] 
**Anagrafica** | Pointer to [**Anagrafica**](Anagrafica.md) |  | [optional] 

## Methods

### NewDatiAnagrafici

`func NewDatiAnagrafici() *DatiAnagrafici`

NewDatiAnagrafici instantiates a new DatiAnagrafici object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatiAnagraficiWithDefaults

`func NewDatiAnagraficiWithDefaults() *DatiAnagrafici`

NewDatiAnagraficiWithDefaults instantiates a new DatiAnagrafici object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdFiscaleIva

`func (o *DatiAnagrafici) GetIdFiscaleIva() IdFiscaleIVA`

GetIdFiscaleIva returns the IdFiscaleIva field if non-nil, zero value otherwise.

### GetIdFiscaleIvaOk

`func (o *DatiAnagrafici) GetIdFiscaleIvaOk() (*IdFiscaleIVA, bool)`

GetIdFiscaleIvaOk returns a tuple with the IdFiscaleIva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdFiscaleIva

`func (o *DatiAnagrafici) SetIdFiscaleIva(v IdFiscaleIVA)`

SetIdFiscaleIva sets IdFiscaleIva field to given value.

### HasIdFiscaleIva

`func (o *DatiAnagrafici) HasIdFiscaleIva() bool`

HasIdFiscaleIva returns a boolean if a field has been set.

### GetCodiceFiscale

`func (o *DatiAnagrafici) GetCodiceFiscale() string`

GetCodiceFiscale returns the CodiceFiscale field if non-nil, zero value otherwise.

### GetCodiceFiscaleOk

`func (o *DatiAnagrafici) GetCodiceFiscaleOk() (*string, bool)`

GetCodiceFiscaleOk returns a tuple with the CodiceFiscale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodiceFiscale

`func (o *DatiAnagrafici) SetCodiceFiscale(v string)`

SetCodiceFiscale sets CodiceFiscale field to given value.

### HasCodiceFiscale

`func (o *DatiAnagrafici) HasCodiceFiscale() bool`

HasCodiceFiscale returns a boolean if a field has been set.

### SetCodiceFiscaleNil

`func (o *DatiAnagrafici) SetCodiceFiscaleNil(b bool)`

 SetCodiceFiscaleNil sets the value for CodiceFiscale to be an explicit nil

### UnsetCodiceFiscale
`func (o *DatiAnagrafici) UnsetCodiceFiscale()`

UnsetCodiceFiscale ensures that no value is present for CodiceFiscale, not even an explicit nil
### GetAnagrafica

`func (o *DatiAnagrafici) GetAnagrafica() Anagrafica`

GetAnagrafica returns the Anagrafica field if non-nil, zero value otherwise.

### GetAnagraficaOk

`func (o *DatiAnagrafici) GetAnagraficaOk() (*Anagrafica, bool)`

GetAnagraficaOk returns a tuple with the Anagrafica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnagrafica

`func (o *DatiAnagrafici) SetAnagrafica(v Anagrafica)`

SetAnagrafica sets Anagrafica field to given value.

### HasAnagrafica

`func (o *DatiAnagrafici) HasAnagrafica() bool`

HasAnagrafica returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


