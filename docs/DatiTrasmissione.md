# DatiTrasmissione

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdTrasmittente** | Pointer to [**IdTrasmittente**](IdTrasmittente.md) |  | [optional] 
**ProgressivoInvio** | Pointer to **NullableString** |  | [optional] 
**FormatoTrasmissione** | Pointer to **NullableString** |  | [optional] 
**CodiceDestinatario** | Pointer to **NullableString** |  | [optional] 
**ContattiTrasmittente** | Pointer to [**ContattiTrasmittente**](ContattiTrasmittente.md) |  | [optional] 
**PecDestinatario** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDatiTrasmissione

`func NewDatiTrasmissione() *DatiTrasmissione`

NewDatiTrasmissione instantiates a new DatiTrasmissione object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatiTrasmissioneWithDefaults

`func NewDatiTrasmissioneWithDefaults() *DatiTrasmissione`

NewDatiTrasmissioneWithDefaults instantiates a new DatiTrasmissione object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdTrasmittente

`func (o *DatiTrasmissione) GetIdTrasmittente() IdTrasmittente`

GetIdTrasmittente returns the IdTrasmittente field if non-nil, zero value otherwise.

### GetIdTrasmittenteOk

`func (o *DatiTrasmissione) GetIdTrasmittenteOk() (*IdTrasmittente, bool)`

GetIdTrasmittenteOk returns a tuple with the IdTrasmittente field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTrasmittente

`func (o *DatiTrasmissione) SetIdTrasmittente(v IdTrasmittente)`

SetIdTrasmittente sets IdTrasmittente field to given value.

### HasIdTrasmittente

`func (o *DatiTrasmissione) HasIdTrasmittente() bool`

HasIdTrasmittente returns a boolean if a field has been set.

### GetProgressivoInvio

`func (o *DatiTrasmissione) GetProgressivoInvio() string`

GetProgressivoInvio returns the ProgressivoInvio field if non-nil, zero value otherwise.

### GetProgressivoInvioOk

`func (o *DatiTrasmissione) GetProgressivoInvioOk() (*string, bool)`

GetProgressivoInvioOk returns a tuple with the ProgressivoInvio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressivoInvio

`func (o *DatiTrasmissione) SetProgressivoInvio(v string)`

SetProgressivoInvio sets ProgressivoInvio field to given value.

### HasProgressivoInvio

`func (o *DatiTrasmissione) HasProgressivoInvio() bool`

HasProgressivoInvio returns a boolean if a field has been set.

### SetProgressivoInvioNil

`func (o *DatiTrasmissione) SetProgressivoInvioNil(b bool)`

 SetProgressivoInvioNil sets the value for ProgressivoInvio to be an explicit nil

### UnsetProgressivoInvio
`func (o *DatiTrasmissione) UnsetProgressivoInvio()`

UnsetProgressivoInvio ensures that no value is present for ProgressivoInvio, not even an explicit nil
### GetFormatoTrasmissione

`func (o *DatiTrasmissione) GetFormatoTrasmissione() string`

GetFormatoTrasmissione returns the FormatoTrasmissione field if non-nil, zero value otherwise.

### GetFormatoTrasmissioneOk

`func (o *DatiTrasmissione) GetFormatoTrasmissioneOk() (*string, bool)`

GetFormatoTrasmissioneOk returns a tuple with the FormatoTrasmissione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatoTrasmissione

`func (o *DatiTrasmissione) SetFormatoTrasmissione(v string)`

SetFormatoTrasmissione sets FormatoTrasmissione field to given value.

### HasFormatoTrasmissione

`func (o *DatiTrasmissione) HasFormatoTrasmissione() bool`

HasFormatoTrasmissione returns a boolean if a field has been set.

### SetFormatoTrasmissioneNil

`func (o *DatiTrasmissione) SetFormatoTrasmissioneNil(b bool)`

 SetFormatoTrasmissioneNil sets the value for FormatoTrasmissione to be an explicit nil

### UnsetFormatoTrasmissione
`func (o *DatiTrasmissione) UnsetFormatoTrasmissione()`

UnsetFormatoTrasmissione ensures that no value is present for FormatoTrasmissione, not even an explicit nil
### GetCodiceDestinatario

`func (o *DatiTrasmissione) GetCodiceDestinatario() string`

GetCodiceDestinatario returns the CodiceDestinatario field if non-nil, zero value otherwise.

### GetCodiceDestinatarioOk

`func (o *DatiTrasmissione) GetCodiceDestinatarioOk() (*string, bool)`

GetCodiceDestinatarioOk returns a tuple with the CodiceDestinatario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodiceDestinatario

`func (o *DatiTrasmissione) SetCodiceDestinatario(v string)`

SetCodiceDestinatario sets CodiceDestinatario field to given value.

### HasCodiceDestinatario

`func (o *DatiTrasmissione) HasCodiceDestinatario() bool`

HasCodiceDestinatario returns a boolean if a field has been set.

### SetCodiceDestinatarioNil

`func (o *DatiTrasmissione) SetCodiceDestinatarioNil(b bool)`

 SetCodiceDestinatarioNil sets the value for CodiceDestinatario to be an explicit nil

### UnsetCodiceDestinatario
`func (o *DatiTrasmissione) UnsetCodiceDestinatario()`

UnsetCodiceDestinatario ensures that no value is present for CodiceDestinatario, not even an explicit nil
### GetContattiTrasmittente

`func (o *DatiTrasmissione) GetContattiTrasmittente() ContattiTrasmittente`

GetContattiTrasmittente returns the ContattiTrasmittente field if non-nil, zero value otherwise.

### GetContattiTrasmittenteOk

`func (o *DatiTrasmissione) GetContattiTrasmittenteOk() (*ContattiTrasmittente, bool)`

GetContattiTrasmittenteOk returns a tuple with the ContattiTrasmittente field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContattiTrasmittente

`func (o *DatiTrasmissione) SetContattiTrasmittente(v ContattiTrasmittente)`

SetContattiTrasmittente sets ContattiTrasmittente field to given value.

### HasContattiTrasmittente

`func (o *DatiTrasmissione) HasContattiTrasmittente() bool`

HasContattiTrasmittente returns a boolean if a field has been set.

### GetPecDestinatario

`func (o *DatiTrasmissione) GetPecDestinatario() string`

GetPecDestinatario returns the PecDestinatario field if non-nil, zero value otherwise.

### GetPecDestinatarioOk

`func (o *DatiTrasmissione) GetPecDestinatarioOk() (*string, bool)`

GetPecDestinatarioOk returns a tuple with the PecDestinatario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPecDestinatario

`func (o *DatiTrasmissione) SetPecDestinatario(v string)`

SetPecDestinatario sets PecDestinatario field to given value.

### HasPecDestinatario

`func (o *DatiTrasmissione) HasPecDestinatario() bool`

HasPecDestinatario returns a boolean if a field has been set.

### SetPecDestinatarioNil

`func (o *DatiTrasmissione) SetPecDestinatarioNil(b bool)`

 SetPecDestinatarioNil sets the value for PecDestinatario to be an explicit nil

### UnsetPecDestinatario
`func (o *DatiTrasmissione) UnsetPecDestinatario()`

UnsetPecDestinatario ensures that no value is present for PecDestinatario, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


