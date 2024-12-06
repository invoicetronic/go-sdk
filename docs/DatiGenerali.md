# DatiGenerali

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatiGeneraliDocumento** | Pointer to [**DatiGeneraliDocumento**](DatiGeneraliDocumento.md) |  | [optional] 
**DatiOrdineAcquisto** | Pointer to [**[]DatiOrdineAcquisto**](DatiOrdineAcquisto.md) |  | [optional] 
**DatiContratto** | Pointer to [**[]DatiContratto**](DatiContratto.md) |  | [optional] 
**DatiConvenzione** | Pointer to [**[]DatiConvenzione**](DatiConvenzione.md) |  | [optional] 
**DatiRicezione** | Pointer to [**[]DatiRicezione**](DatiRicezione.md) |  | [optional] 
**DatiFattureCollegate** | Pointer to [**[]DatiFattureCollegate**](DatiFattureCollegate.md) |  | [optional] 
**DatiSal** | Pointer to [**[]DatiSAL**](DatiSAL.md) |  | [optional] 
**DatiDdt** | Pointer to [**[]DatiDDT**](DatiDDT.md) |  | [optional] 
**DatiTrasporto** | Pointer to [**DatiTrasporto**](DatiTrasporto.md) |  | [optional] 
**FatturaPrincipale** | Pointer to [**FatturaPrincipale**](FatturaPrincipale.md) |  | [optional] 

## Methods

### NewDatiGenerali

`func NewDatiGenerali() *DatiGenerali`

NewDatiGenerali instantiates a new DatiGenerali object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatiGeneraliWithDefaults

`func NewDatiGeneraliWithDefaults() *DatiGenerali`

NewDatiGeneraliWithDefaults instantiates a new DatiGenerali object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatiGeneraliDocumento

`func (o *DatiGenerali) GetDatiGeneraliDocumento() DatiGeneraliDocumento`

GetDatiGeneraliDocumento returns the DatiGeneraliDocumento field if non-nil, zero value otherwise.

### GetDatiGeneraliDocumentoOk

`func (o *DatiGenerali) GetDatiGeneraliDocumentoOk() (*DatiGeneraliDocumento, bool)`

GetDatiGeneraliDocumentoOk returns a tuple with the DatiGeneraliDocumento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiGeneraliDocumento

`func (o *DatiGenerali) SetDatiGeneraliDocumento(v DatiGeneraliDocumento)`

SetDatiGeneraliDocumento sets DatiGeneraliDocumento field to given value.

### HasDatiGeneraliDocumento

`func (o *DatiGenerali) HasDatiGeneraliDocumento() bool`

HasDatiGeneraliDocumento returns a boolean if a field has been set.

### GetDatiOrdineAcquisto

`func (o *DatiGenerali) GetDatiOrdineAcquisto() []DatiOrdineAcquisto`

GetDatiOrdineAcquisto returns the DatiOrdineAcquisto field if non-nil, zero value otherwise.

### GetDatiOrdineAcquistoOk

`func (o *DatiGenerali) GetDatiOrdineAcquistoOk() (*[]DatiOrdineAcquisto, bool)`

GetDatiOrdineAcquistoOk returns a tuple with the DatiOrdineAcquisto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiOrdineAcquisto

`func (o *DatiGenerali) SetDatiOrdineAcquisto(v []DatiOrdineAcquisto)`

SetDatiOrdineAcquisto sets DatiOrdineAcquisto field to given value.

### HasDatiOrdineAcquisto

`func (o *DatiGenerali) HasDatiOrdineAcquisto() bool`

HasDatiOrdineAcquisto returns a boolean if a field has been set.

### SetDatiOrdineAcquistoNil

`func (o *DatiGenerali) SetDatiOrdineAcquistoNil(b bool)`

 SetDatiOrdineAcquistoNil sets the value for DatiOrdineAcquisto to be an explicit nil

### UnsetDatiOrdineAcquisto
`func (o *DatiGenerali) UnsetDatiOrdineAcquisto()`

UnsetDatiOrdineAcquisto ensures that no value is present for DatiOrdineAcquisto, not even an explicit nil
### GetDatiContratto

`func (o *DatiGenerali) GetDatiContratto() []DatiContratto`

GetDatiContratto returns the DatiContratto field if non-nil, zero value otherwise.

### GetDatiContrattoOk

`func (o *DatiGenerali) GetDatiContrattoOk() (*[]DatiContratto, bool)`

GetDatiContrattoOk returns a tuple with the DatiContratto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiContratto

`func (o *DatiGenerali) SetDatiContratto(v []DatiContratto)`

SetDatiContratto sets DatiContratto field to given value.

### HasDatiContratto

`func (o *DatiGenerali) HasDatiContratto() bool`

HasDatiContratto returns a boolean if a field has been set.

### SetDatiContrattoNil

`func (o *DatiGenerali) SetDatiContrattoNil(b bool)`

 SetDatiContrattoNil sets the value for DatiContratto to be an explicit nil

### UnsetDatiContratto
`func (o *DatiGenerali) UnsetDatiContratto()`

UnsetDatiContratto ensures that no value is present for DatiContratto, not even an explicit nil
### GetDatiConvenzione

`func (o *DatiGenerali) GetDatiConvenzione() []DatiConvenzione`

GetDatiConvenzione returns the DatiConvenzione field if non-nil, zero value otherwise.

### GetDatiConvenzioneOk

`func (o *DatiGenerali) GetDatiConvenzioneOk() (*[]DatiConvenzione, bool)`

GetDatiConvenzioneOk returns a tuple with the DatiConvenzione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiConvenzione

`func (o *DatiGenerali) SetDatiConvenzione(v []DatiConvenzione)`

SetDatiConvenzione sets DatiConvenzione field to given value.

### HasDatiConvenzione

`func (o *DatiGenerali) HasDatiConvenzione() bool`

HasDatiConvenzione returns a boolean if a field has been set.

### SetDatiConvenzioneNil

`func (o *DatiGenerali) SetDatiConvenzioneNil(b bool)`

 SetDatiConvenzioneNil sets the value for DatiConvenzione to be an explicit nil

### UnsetDatiConvenzione
`func (o *DatiGenerali) UnsetDatiConvenzione()`

UnsetDatiConvenzione ensures that no value is present for DatiConvenzione, not even an explicit nil
### GetDatiRicezione

`func (o *DatiGenerali) GetDatiRicezione() []DatiRicezione`

GetDatiRicezione returns the DatiRicezione field if non-nil, zero value otherwise.

### GetDatiRicezioneOk

`func (o *DatiGenerali) GetDatiRicezioneOk() (*[]DatiRicezione, bool)`

GetDatiRicezioneOk returns a tuple with the DatiRicezione field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiRicezione

`func (o *DatiGenerali) SetDatiRicezione(v []DatiRicezione)`

SetDatiRicezione sets DatiRicezione field to given value.

### HasDatiRicezione

`func (o *DatiGenerali) HasDatiRicezione() bool`

HasDatiRicezione returns a boolean if a field has been set.

### SetDatiRicezioneNil

`func (o *DatiGenerali) SetDatiRicezioneNil(b bool)`

 SetDatiRicezioneNil sets the value for DatiRicezione to be an explicit nil

### UnsetDatiRicezione
`func (o *DatiGenerali) UnsetDatiRicezione()`

UnsetDatiRicezione ensures that no value is present for DatiRicezione, not even an explicit nil
### GetDatiFattureCollegate

`func (o *DatiGenerali) GetDatiFattureCollegate() []DatiFattureCollegate`

GetDatiFattureCollegate returns the DatiFattureCollegate field if non-nil, zero value otherwise.

### GetDatiFattureCollegateOk

`func (o *DatiGenerali) GetDatiFattureCollegateOk() (*[]DatiFattureCollegate, bool)`

GetDatiFattureCollegateOk returns a tuple with the DatiFattureCollegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiFattureCollegate

`func (o *DatiGenerali) SetDatiFattureCollegate(v []DatiFattureCollegate)`

SetDatiFattureCollegate sets DatiFattureCollegate field to given value.

### HasDatiFattureCollegate

`func (o *DatiGenerali) HasDatiFattureCollegate() bool`

HasDatiFattureCollegate returns a boolean if a field has been set.

### SetDatiFattureCollegateNil

`func (o *DatiGenerali) SetDatiFattureCollegateNil(b bool)`

 SetDatiFattureCollegateNil sets the value for DatiFattureCollegate to be an explicit nil

### UnsetDatiFattureCollegate
`func (o *DatiGenerali) UnsetDatiFattureCollegate()`

UnsetDatiFattureCollegate ensures that no value is present for DatiFattureCollegate, not even an explicit nil
### GetDatiSal

`func (o *DatiGenerali) GetDatiSal() []DatiSAL`

GetDatiSal returns the DatiSal field if non-nil, zero value otherwise.

### GetDatiSalOk

`func (o *DatiGenerali) GetDatiSalOk() (*[]DatiSAL, bool)`

GetDatiSalOk returns a tuple with the DatiSal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiSal

`func (o *DatiGenerali) SetDatiSal(v []DatiSAL)`

SetDatiSal sets DatiSal field to given value.

### HasDatiSal

`func (o *DatiGenerali) HasDatiSal() bool`

HasDatiSal returns a boolean if a field has been set.

### SetDatiSalNil

`func (o *DatiGenerali) SetDatiSalNil(b bool)`

 SetDatiSalNil sets the value for DatiSal to be an explicit nil

### UnsetDatiSal
`func (o *DatiGenerali) UnsetDatiSal()`

UnsetDatiSal ensures that no value is present for DatiSal, not even an explicit nil
### GetDatiDdt

`func (o *DatiGenerali) GetDatiDdt() []DatiDDT`

GetDatiDdt returns the DatiDdt field if non-nil, zero value otherwise.

### GetDatiDdtOk

`func (o *DatiGenerali) GetDatiDdtOk() (*[]DatiDDT, bool)`

GetDatiDdtOk returns a tuple with the DatiDdt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiDdt

`func (o *DatiGenerali) SetDatiDdt(v []DatiDDT)`

SetDatiDdt sets DatiDdt field to given value.

### HasDatiDdt

`func (o *DatiGenerali) HasDatiDdt() bool`

HasDatiDdt returns a boolean if a field has been set.

### SetDatiDdtNil

`func (o *DatiGenerali) SetDatiDdtNil(b bool)`

 SetDatiDdtNil sets the value for DatiDdt to be an explicit nil

### UnsetDatiDdt
`func (o *DatiGenerali) UnsetDatiDdt()`

UnsetDatiDdt ensures that no value is present for DatiDdt, not even an explicit nil
### GetDatiTrasporto

`func (o *DatiGenerali) GetDatiTrasporto() DatiTrasporto`

GetDatiTrasporto returns the DatiTrasporto field if non-nil, zero value otherwise.

### GetDatiTrasportoOk

`func (o *DatiGenerali) GetDatiTrasportoOk() (*DatiTrasporto, bool)`

GetDatiTrasportoOk returns a tuple with the DatiTrasporto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatiTrasporto

`func (o *DatiGenerali) SetDatiTrasporto(v DatiTrasporto)`

SetDatiTrasporto sets DatiTrasporto field to given value.

### HasDatiTrasporto

`func (o *DatiGenerali) HasDatiTrasporto() bool`

HasDatiTrasporto returns a boolean if a field has been set.

### GetFatturaPrincipale

`func (o *DatiGenerali) GetFatturaPrincipale() FatturaPrincipale`

GetFatturaPrincipale returns the FatturaPrincipale field if non-nil, zero value otherwise.

### GetFatturaPrincipaleOk

`func (o *DatiGenerali) GetFatturaPrincipaleOk() (*FatturaPrincipale, bool)`

GetFatturaPrincipaleOk returns a tuple with the FatturaPrincipale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFatturaPrincipale

`func (o *DatiGenerali) SetFatturaPrincipale(v FatturaPrincipale)`

SetFatturaPrincipale sets FatturaPrincipale field to given value.

### HasFatturaPrincipale

`func (o *DatiGenerali) HasFatturaPrincipale() bool`

HasFatturaPrincipale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


