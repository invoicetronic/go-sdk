# CedentePrestatore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatiAnagrafici** | Pointer to [**DatiAnagraficiCedentePrestatore**](DatiAnagraficiCedentePrestatore.md) |  | [optional] 
**Sede** | Pointer to [**SedeCedentePrestatore**](SedeCedentePrestatore.md) |  | [optional] 
**StabileOrganizzazione** | Pointer to [**StabileOrganizzazione**](StabileOrganizzazione.md) |  | [optional] 
**IscrizioneRea** | Pointer to [**IscrizioneREA**](IscrizioneREA.md) |  | [optional] 
**Contatti** | Pointer to [**Contatti**](Contatti.md) |  | [optional] 
**RiferimentoAmministrazione** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCedentePrestatore

`func NewCedentePrestatore() *CedentePrestatore`

NewCedentePrestatore instantiates a new CedentePrestatore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCedentePrestatoreWithDefaults

`func NewCedentePrestatoreWithDefaults() *CedentePrestatore`

NewCedentePrestatoreWithDefaults instantiates a new CedentePrestatore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatiAnagrafici

`func (o *CedentePrestatore) GetDatiAnagrafici() DatiAnagraficiCedentePrestatore`

GetDatiAnagrafici returns the DatiAnagrafici field if non-nil, zero value otherwise.

### GetDatiAnagraficiOk

`func (o *CedentePrestatore) GetDatiAnagraficiOk() (*DatiAnagraficiCedentePrestatore, bool)`

GetDatiAnagraficiOk returns a tuple with the DatiAnagrafici field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiAnagrafici

`func (o *CedentePrestatore) SetDatiAnagrafici(v DatiAnagraficiCedentePrestatore)`

SetDatiAnagrafici sets DatiAnagrafici field to given value.

### HasDatiAnagrafici

`func (o *CedentePrestatore) HasDatiAnagrafici() bool`

HasDatiAnagrafici returns a boolean if a field has been set.

### GetSede

`func (o *CedentePrestatore) GetSede() SedeCedentePrestatore`

GetSede returns the Sede field if non-nil, zero value otherwise.

### GetSedeOk

`func (o *CedentePrestatore) GetSedeOk() (*SedeCedentePrestatore, bool)`

GetSedeOk returns a tuple with the Sede field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSede

`func (o *CedentePrestatore) SetSede(v SedeCedentePrestatore)`

SetSede sets Sede field to given value.

### HasSede

`func (o *CedentePrestatore) HasSede() bool`

HasSede returns a boolean if a field has been set.

### GetStabileOrganizzazione

`func (o *CedentePrestatore) GetStabileOrganizzazione() StabileOrganizzazione`

GetStabileOrganizzazione returns the StabileOrganizzazione field if non-nil, zero value otherwise.

### GetStabileOrganizzazioneOk

`func (o *CedentePrestatore) GetStabileOrganizzazioneOk() (*StabileOrganizzazione, bool)`

GetStabileOrganizzazioneOk returns a tuple with the StabileOrganizzazione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStabileOrganizzazione

`func (o *CedentePrestatore) SetStabileOrganizzazione(v StabileOrganizzazione)`

SetStabileOrganizzazione sets StabileOrganizzazione field to given value.

### HasStabileOrganizzazione

`func (o *CedentePrestatore) HasStabileOrganizzazione() bool`

HasStabileOrganizzazione returns a boolean if a field has been set.

### GetIscrizioneRea

`func (o *CedentePrestatore) GetIscrizioneRea() IscrizioneREA`

GetIscrizioneRea returns the IscrizioneRea field if non-nil, zero value otherwise.

### GetIscrizioneReaOk

`func (o *CedentePrestatore) GetIscrizioneReaOk() (*IscrizioneREA, bool)`

GetIscrizioneReaOk returns a tuple with the IscrizioneRea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscrizioneRea

`func (o *CedentePrestatore) SetIscrizioneRea(v IscrizioneREA)`

SetIscrizioneRea sets IscrizioneRea field to given value.

### HasIscrizioneRea

`func (o *CedentePrestatore) HasIscrizioneRea() bool`

HasIscrizioneRea returns a boolean if a field has been set.

### GetContatti

`func (o *CedentePrestatore) GetContatti() Contatti`

GetContatti returns the Contatti field if non-nil, zero value otherwise.

### GetContattiOk

`func (o *CedentePrestatore) GetContattiOk() (*Contatti, bool)`

GetContattiOk returns a tuple with the Contatti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContatti

`func (o *CedentePrestatore) SetContatti(v Contatti)`

SetContatti sets Contatti field to given value.

### HasContatti

`func (o *CedentePrestatore) HasContatti() bool`

HasContatti returns a boolean if a field has been set.

### GetRiferimentoAmministrazione

`func (o *CedentePrestatore) GetRiferimentoAmministrazione() string`

GetRiferimentoAmministrazione returns the RiferimentoAmministrazione field if non-nil, zero value otherwise.

### GetRiferimentoAmministrazioneOk

`func (o *CedentePrestatore) GetRiferimentoAmministrazioneOk() (*string, bool)`

GetRiferimentoAmministrazioneOk returns a tuple with the RiferimentoAmministrazione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiferimentoAmministrazione

`func (o *CedentePrestatore) SetRiferimentoAmministrazione(v string)`

SetRiferimentoAmministrazione sets RiferimentoAmministrazione field to given value.

### HasRiferimentoAmministrazione

`func (o *CedentePrestatore) HasRiferimentoAmministrazione() bool`

HasRiferimentoAmministrazione returns a boolean if a field has been set.

### SetRiferimentoAmministrazioneNil

`func (o *CedentePrestatore) SetRiferimentoAmministrazioneNil(b bool)`

 SetRiferimentoAmministrazioneNil sets the value for RiferimentoAmministrazione to be an explicit nil

### UnsetRiferimentoAmministrazione
`func (o *CedentePrestatore) UnsetRiferimentoAmministrazione()`

UnsetRiferimentoAmministrazione ensures that no value is present for RiferimentoAmministrazione, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


