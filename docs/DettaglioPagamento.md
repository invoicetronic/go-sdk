# DettaglioPagamento

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Beneficiario** | Pointer to **NullableString** |  | [optional] 
**ModalitaPagamento** | Pointer to **NullableString** |  | [optional] 
**DataRiferimentoTerminiPagamento** | Pointer to **NullableTime** |  | [optional] 
**GiorniTerminiPagamento** | Pointer to **NullableInt32** |  | [optional] 
**DataScadenzaPagamento** | Pointer to **NullableTime** |  | [optional] 
**ImportoPagamento** | Pointer to **float64** |  | [optional] 
**CodUfficioPostale** | Pointer to **NullableString** |  | [optional] 
**CognomeQuietanzante** | Pointer to **NullableString** |  | [optional] 
**NomeQuietanzante** | Pointer to **NullableString** |  | [optional] 
**CfQuietanzante** | Pointer to **NullableString** |  | [optional] 
**TitoloQuietanzante** | Pointer to **NullableString** |  | [optional] 
**IstitutoFinanziario** | Pointer to **NullableString** |  | [optional] 
**Iban** | Pointer to **NullableString** |  | [optional] 
**Abi** | Pointer to **NullableString** |  | [optional] 
**Cab** | Pointer to **NullableString** |  | [optional] 
**Bic** | Pointer to **NullableString** |  | [optional] 
**ScontoPagamentoAnticipato** | Pointer to **NullableFloat64** |  | [optional] 
**DataLimitePagamentoAnticipato** | Pointer to **NullableTime** |  | [optional] 
**PenalitaPagamentiRitardati** | Pointer to **NullableFloat64** |  | [optional] 
**DataDecorrenzaPenale** | Pointer to **NullableTime** |  | [optional] 
**CodicePagamento** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDettaglioPagamento

`func NewDettaglioPagamento() *DettaglioPagamento`

NewDettaglioPagamento instantiates a new DettaglioPagamento object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDettaglioPagamentoWithDefaults

`func NewDettaglioPagamentoWithDefaults() *DettaglioPagamento`

NewDettaglioPagamentoWithDefaults instantiates a new DettaglioPagamento object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeneficiario

`func (o *DettaglioPagamento) GetBeneficiario() string`

GetBeneficiario returns the Beneficiario field if non-nil, zero value otherwise.

### GetBeneficiarioOk

`func (o *DettaglioPagamento) GetBeneficiarioOk() (*string, bool)`

GetBeneficiarioOk returns a tuple with the Beneficiario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiario

`func (o *DettaglioPagamento) SetBeneficiario(v string)`

SetBeneficiario sets Beneficiario field to given value.

### HasBeneficiario

`func (o *DettaglioPagamento) HasBeneficiario() bool`

HasBeneficiario returns a boolean if a field has been set.

### SetBeneficiarioNil

`func (o *DettaglioPagamento) SetBeneficiarioNil(b bool)`

 SetBeneficiarioNil sets the value for Beneficiario to be an explicit nil

### UnsetBeneficiario
`func (o *DettaglioPagamento) UnsetBeneficiario()`

UnsetBeneficiario ensures that no value is present for Beneficiario, not even an explicit nil
### GetModalitaPagamento

`func (o *DettaglioPagamento) GetModalitaPagamento() string`

GetModalitaPagamento returns the ModalitaPagamento field if non-nil, zero value otherwise.

### GetModalitaPagamentoOk

`func (o *DettaglioPagamento) GetModalitaPagamentoOk() (*string, bool)`

GetModalitaPagamentoOk returns a tuple with the ModalitaPagamento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModalitaPagamento

`func (o *DettaglioPagamento) SetModalitaPagamento(v string)`

SetModalitaPagamento sets ModalitaPagamento field to given value.

### HasModalitaPagamento

`func (o *DettaglioPagamento) HasModalitaPagamento() bool`

HasModalitaPagamento returns a boolean if a field has been set.

### SetModalitaPagamentoNil

`func (o *DettaglioPagamento) SetModalitaPagamentoNil(b bool)`

 SetModalitaPagamentoNil sets the value for ModalitaPagamento to be an explicit nil

### UnsetModalitaPagamento
`func (o *DettaglioPagamento) UnsetModalitaPagamento()`

UnsetModalitaPagamento ensures that no value is present for ModalitaPagamento, not even an explicit nil
### GetDataRiferimentoTerminiPagamento

`func (o *DettaglioPagamento) GetDataRiferimentoTerminiPagamento() time.Time`

GetDataRiferimentoTerminiPagamento returns the DataRiferimentoTerminiPagamento field if non-nil, zero value otherwise.

### GetDataRiferimentoTerminiPagamentoOk

`func (o *DettaglioPagamento) GetDataRiferimentoTerminiPagamentoOk() (*time.Time, bool)`

GetDataRiferimentoTerminiPagamentoOk returns a tuple with the DataRiferimentoTerminiPagamento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRiferimentoTerminiPagamento

`func (o *DettaglioPagamento) SetDataRiferimentoTerminiPagamento(v time.Time)`

SetDataRiferimentoTerminiPagamento sets DataRiferimentoTerminiPagamento field to given value.

### HasDataRiferimentoTerminiPagamento

`func (o *DettaglioPagamento) HasDataRiferimentoTerminiPagamento() bool`

HasDataRiferimentoTerminiPagamento returns a boolean if a field has been set.

### SetDataRiferimentoTerminiPagamentoNil

`func (o *DettaglioPagamento) SetDataRiferimentoTerminiPagamentoNil(b bool)`

 SetDataRiferimentoTerminiPagamentoNil sets the value for DataRiferimentoTerminiPagamento to be an explicit nil

### UnsetDataRiferimentoTerminiPagamento
`func (o *DettaglioPagamento) UnsetDataRiferimentoTerminiPagamento()`

UnsetDataRiferimentoTerminiPagamento ensures that no value is present for DataRiferimentoTerminiPagamento, not even an explicit nil
### GetGiorniTerminiPagamento

`func (o *DettaglioPagamento) GetGiorniTerminiPagamento() int32`

GetGiorniTerminiPagamento returns the GiorniTerminiPagamento field if non-nil, zero value otherwise.

### GetGiorniTerminiPagamentoOk

`func (o *DettaglioPagamento) GetGiorniTerminiPagamentoOk() (*int32, bool)`

GetGiorniTerminiPagamentoOk returns a tuple with the GiorniTerminiPagamento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiorniTerminiPagamento

`func (o *DettaglioPagamento) SetGiorniTerminiPagamento(v int32)`

SetGiorniTerminiPagamento sets GiorniTerminiPagamento field to given value.

### HasGiorniTerminiPagamento

`func (o *DettaglioPagamento) HasGiorniTerminiPagamento() bool`

HasGiorniTerminiPagamento returns a boolean if a field has been set.

### SetGiorniTerminiPagamentoNil

`func (o *DettaglioPagamento) SetGiorniTerminiPagamentoNil(b bool)`

 SetGiorniTerminiPagamentoNil sets the value for GiorniTerminiPagamento to be an explicit nil

### UnsetGiorniTerminiPagamento
`func (o *DettaglioPagamento) UnsetGiorniTerminiPagamento()`

UnsetGiorniTerminiPagamento ensures that no value is present for GiorniTerminiPagamento, not even an explicit nil
### GetDataScadenzaPagamento

`func (o *DettaglioPagamento) GetDataScadenzaPagamento() time.Time`

GetDataScadenzaPagamento returns the DataScadenzaPagamento field if non-nil, zero value otherwise.

### GetDataScadenzaPagamentoOk

`func (o *DettaglioPagamento) GetDataScadenzaPagamentoOk() (*time.Time, bool)`

GetDataScadenzaPagamentoOk returns a tuple with the DataScadenzaPagamento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataScadenzaPagamento

`func (o *DettaglioPagamento) SetDataScadenzaPagamento(v time.Time)`

SetDataScadenzaPagamento sets DataScadenzaPagamento field to given value.

### HasDataScadenzaPagamento

`func (o *DettaglioPagamento) HasDataScadenzaPagamento() bool`

HasDataScadenzaPagamento returns a boolean if a field has been set.

### SetDataScadenzaPagamentoNil

`func (o *DettaglioPagamento) SetDataScadenzaPagamentoNil(b bool)`

 SetDataScadenzaPagamentoNil sets the value for DataScadenzaPagamento to be an explicit nil

### UnsetDataScadenzaPagamento
`func (o *DettaglioPagamento) UnsetDataScadenzaPagamento()`

UnsetDataScadenzaPagamento ensures that no value is present for DataScadenzaPagamento, not even an explicit nil
### GetImportoPagamento

`func (o *DettaglioPagamento) GetImportoPagamento() float64`

GetImportoPagamento returns the ImportoPagamento field if non-nil, zero value otherwise.

### GetImportoPagamentoOk

`func (o *DettaglioPagamento) GetImportoPagamentoOk() (*float64, bool)`

GetImportoPagamentoOk returns a tuple with the ImportoPagamento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportoPagamento

`func (o *DettaglioPagamento) SetImportoPagamento(v float64)`

SetImportoPagamento sets ImportoPagamento field to given value.

### HasImportoPagamento

`func (o *DettaglioPagamento) HasImportoPagamento() bool`

HasImportoPagamento returns a boolean if a field has been set.

### GetCodUfficioPostale

`func (o *DettaglioPagamento) GetCodUfficioPostale() string`

GetCodUfficioPostale returns the CodUfficioPostale field if non-nil, zero value otherwise.

### GetCodUfficioPostaleOk

`func (o *DettaglioPagamento) GetCodUfficioPostaleOk() (*string, bool)`

GetCodUfficioPostaleOk returns a tuple with the CodUfficioPostale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodUfficioPostale

`func (o *DettaglioPagamento) SetCodUfficioPostale(v string)`

SetCodUfficioPostale sets CodUfficioPostale field to given value.

### HasCodUfficioPostale

`func (o *DettaglioPagamento) HasCodUfficioPostale() bool`

HasCodUfficioPostale returns a boolean if a field has been set.

### SetCodUfficioPostaleNil

`func (o *DettaglioPagamento) SetCodUfficioPostaleNil(b bool)`

 SetCodUfficioPostaleNil sets the value for CodUfficioPostale to be an explicit nil

### UnsetCodUfficioPostale
`func (o *DettaglioPagamento) UnsetCodUfficioPostale()`

UnsetCodUfficioPostale ensures that no value is present for CodUfficioPostale, not even an explicit nil
### GetCognomeQuietanzante

`func (o *DettaglioPagamento) GetCognomeQuietanzante() string`

GetCognomeQuietanzante returns the CognomeQuietanzante field if non-nil, zero value otherwise.

### GetCognomeQuietanzanteOk

`func (o *DettaglioPagamento) GetCognomeQuietanzanteOk() (*string, bool)`

GetCognomeQuietanzanteOk returns a tuple with the CognomeQuietanzante field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCognomeQuietanzante

`func (o *DettaglioPagamento) SetCognomeQuietanzante(v string)`

SetCognomeQuietanzante sets CognomeQuietanzante field to given value.

### HasCognomeQuietanzante

`func (o *DettaglioPagamento) HasCognomeQuietanzante() bool`

HasCognomeQuietanzante returns a boolean if a field has been set.

### SetCognomeQuietanzanteNil

`func (o *DettaglioPagamento) SetCognomeQuietanzanteNil(b bool)`

 SetCognomeQuietanzanteNil sets the value for CognomeQuietanzante to be an explicit nil

### UnsetCognomeQuietanzante
`func (o *DettaglioPagamento) UnsetCognomeQuietanzante()`

UnsetCognomeQuietanzante ensures that no value is present for CognomeQuietanzante, not even an explicit nil
### GetNomeQuietanzante

`func (o *DettaglioPagamento) GetNomeQuietanzante() string`

GetNomeQuietanzante returns the NomeQuietanzante field if non-nil, zero value otherwise.

### GetNomeQuietanzanteOk

`func (o *DettaglioPagamento) GetNomeQuietanzanteOk() (*string, bool)`

GetNomeQuietanzanteOk returns a tuple with the NomeQuietanzante field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomeQuietanzante

`func (o *DettaglioPagamento) SetNomeQuietanzante(v string)`

SetNomeQuietanzante sets NomeQuietanzante field to given value.

### HasNomeQuietanzante

`func (o *DettaglioPagamento) HasNomeQuietanzante() bool`

HasNomeQuietanzante returns a boolean if a field has been set.

### SetNomeQuietanzanteNil

`func (o *DettaglioPagamento) SetNomeQuietanzanteNil(b bool)`

 SetNomeQuietanzanteNil sets the value for NomeQuietanzante to be an explicit nil

### UnsetNomeQuietanzante
`func (o *DettaglioPagamento) UnsetNomeQuietanzante()`

UnsetNomeQuietanzante ensures that no value is present for NomeQuietanzante, not even an explicit nil
### GetCfQuietanzante

`func (o *DettaglioPagamento) GetCfQuietanzante() string`

GetCfQuietanzante returns the CfQuietanzante field if non-nil, zero value otherwise.

### GetCfQuietanzanteOk

`func (o *DettaglioPagamento) GetCfQuietanzanteOk() (*string, bool)`

GetCfQuietanzanteOk returns a tuple with the CfQuietanzante field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfQuietanzante

`func (o *DettaglioPagamento) SetCfQuietanzante(v string)`

SetCfQuietanzante sets CfQuietanzante field to given value.

### HasCfQuietanzante

`func (o *DettaglioPagamento) HasCfQuietanzante() bool`

HasCfQuietanzante returns a boolean if a field has been set.

### SetCfQuietanzanteNil

`func (o *DettaglioPagamento) SetCfQuietanzanteNil(b bool)`

 SetCfQuietanzanteNil sets the value for CfQuietanzante to be an explicit nil

### UnsetCfQuietanzante
`func (o *DettaglioPagamento) UnsetCfQuietanzante()`

UnsetCfQuietanzante ensures that no value is present for CfQuietanzante, not even an explicit nil
### GetTitoloQuietanzante

`func (o *DettaglioPagamento) GetTitoloQuietanzante() string`

GetTitoloQuietanzante returns the TitoloQuietanzante field if non-nil, zero value otherwise.

### GetTitoloQuietanzanteOk

`func (o *DettaglioPagamento) GetTitoloQuietanzanteOk() (*string, bool)`

GetTitoloQuietanzanteOk returns a tuple with the TitoloQuietanzante field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitoloQuietanzante

`func (o *DettaglioPagamento) SetTitoloQuietanzante(v string)`

SetTitoloQuietanzante sets TitoloQuietanzante field to given value.

### HasTitoloQuietanzante

`func (o *DettaglioPagamento) HasTitoloQuietanzante() bool`

HasTitoloQuietanzante returns a boolean if a field has been set.

### SetTitoloQuietanzanteNil

`func (o *DettaglioPagamento) SetTitoloQuietanzanteNil(b bool)`

 SetTitoloQuietanzanteNil sets the value for TitoloQuietanzante to be an explicit nil

### UnsetTitoloQuietanzante
`func (o *DettaglioPagamento) UnsetTitoloQuietanzante()`

UnsetTitoloQuietanzante ensures that no value is present for TitoloQuietanzante, not even an explicit nil
### GetIstitutoFinanziario

`func (o *DettaglioPagamento) GetIstitutoFinanziario() string`

GetIstitutoFinanziario returns the IstitutoFinanziario field if non-nil, zero value otherwise.

### GetIstitutoFinanziarioOk

`func (o *DettaglioPagamento) GetIstitutoFinanziarioOk() (*string, bool)`

GetIstitutoFinanziarioOk returns a tuple with the IstitutoFinanziario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIstitutoFinanziario

`func (o *DettaglioPagamento) SetIstitutoFinanziario(v string)`

SetIstitutoFinanziario sets IstitutoFinanziario field to given value.

### HasIstitutoFinanziario

`func (o *DettaglioPagamento) HasIstitutoFinanziario() bool`

HasIstitutoFinanziario returns a boolean if a field has been set.

### SetIstitutoFinanziarioNil

`func (o *DettaglioPagamento) SetIstitutoFinanziarioNil(b bool)`

 SetIstitutoFinanziarioNil sets the value for IstitutoFinanziario to be an explicit nil

### UnsetIstitutoFinanziario
`func (o *DettaglioPagamento) UnsetIstitutoFinanziario()`

UnsetIstitutoFinanziario ensures that no value is present for IstitutoFinanziario, not even an explicit nil
### GetIban

`func (o *DettaglioPagamento) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *DettaglioPagamento) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *DettaglioPagamento) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *DettaglioPagamento) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *DettaglioPagamento) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *DettaglioPagamento) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetAbi

`func (o *DettaglioPagamento) GetAbi() string`

GetAbi returns the Abi field if non-nil, zero value otherwise.

### GetAbiOk

`func (o *DettaglioPagamento) GetAbiOk() (*string, bool)`

GetAbiOk returns a tuple with the Abi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbi

`func (o *DettaglioPagamento) SetAbi(v string)`

SetAbi sets Abi field to given value.

### HasAbi

`func (o *DettaglioPagamento) HasAbi() bool`

HasAbi returns a boolean if a field has been set.

### SetAbiNil

`func (o *DettaglioPagamento) SetAbiNil(b bool)`

 SetAbiNil sets the value for Abi to be an explicit nil

### UnsetAbi
`func (o *DettaglioPagamento) UnsetAbi()`

UnsetAbi ensures that no value is present for Abi, not even an explicit nil
### GetCab

`func (o *DettaglioPagamento) GetCab() string`

GetCab returns the Cab field if non-nil, zero value otherwise.

### GetCabOk

`func (o *DettaglioPagamento) GetCabOk() (*string, bool)`

GetCabOk returns a tuple with the Cab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCab

`func (o *DettaglioPagamento) SetCab(v string)`

SetCab sets Cab field to given value.

### HasCab

`func (o *DettaglioPagamento) HasCab() bool`

HasCab returns a boolean if a field has been set.

### SetCabNil

`func (o *DettaglioPagamento) SetCabNil(b bool)`

 SetCabNil sets the value for Cab to be an explicit nil

### UnsetCab
`func (o *DettaglioPagamento) UnsetCab()`

UnsetCab ensures that no value is present for Cab, not even an explicit nil
### GetBic

`func (o *DettaglioPagamento) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *DettaglioPagamento) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *DettaglioPagamento) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *DettaglioPagamento) HasBic() bool`

HasBic returns a boolean if a field has been set.

### SetBicNil

`func (o *DettaglioPagamento) SetBicNil(b bool)`

 SetBicNil sets the value for Bic to be an explicit nil

### UnsetBic
`func (o *DettaglioPagamento) UnsetBic()`

UnsetBic ensures that no value is present for Bic, not even an explicit nil
### GetScontoPagamentoAnticipato

`func (o *DettaglioPagamento) GetScontoPagamentoAnticipato() float64`

GetScontoPagamentoAnticipato returns the ScontoPagamentoAnticipato field if non-nil, zero value otherwise.

### GetScontoPagamentoAnticipatoOk

`func (o *DettaglioPagamento) GetScontoPagamentoAnticipatoOk() (*float64, bool)`

GetScontoPagamentoAnticipatoOk returns a tuple with the ScontoPagamentoAnticipato field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScontoPagamentoAnticipato

`func (o *DettaglioPagamento) SetScontoPagamentoAnticipato(v float64)`

SetScontoPagamentoAnticipato sets ScontoPagamentoAnticipato field to given value.

### HasScontoPagamentoAnticipato

`func (o *DettaglioPagamento) HasScontoPagamentoAnticipato() bool`

HasScontoPagamentoAnticipato returns a boolean if a field has been set.

### SetScontoPagamentoAnticipatoNil

`func (o *DettaglioPagamento) SetScontoPagamentoAnticipatoNil(b bool)`

 SetScontoPagamentoAnticipatoNil sets the value for ScontoPagamentoAnticipato to be an explicit nil

### UnsetScontoPagamentoAnticipato
`func (o *DettaglioPagamento) UnsetScontoPagamentoAnticipato()`

UnsetScontoPagamentoAnticipato ensures that no value is present for ScontoPagamentoAnticipato, not even an explicit nil
### GetDataLimitePagamentoAnticipato

`func (o *DettaglioPagamento) GetDataLimitePagamentoAnticipato() time.Time`

GetDataLimitePagamentoAnticipato returns the DataLimitePagamentoAnticipato field if non-nil, zero value otherwise.

### GetDataLimitePagamentoAnticipatoOk

`func (o *DettaglioPagamento) GetDataLimitePagamentoAnticipatoOk() (*time.Time, bool)`

GetDataLimitePagamentoAnticipatoOk returns a tuple with the DataLimitePagamentoAnticipato field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLimitePagamentoAnticipato

`func (o *DettaglioPagamento) SetDataLimitePagamentoAnticipato(v time.Time)`

SetDataLimitePagamentoAnticipato sets DataLimitePagamentoAnticipato field to given value.

### HasDataLimitePagamentoAnticipato

`func (o *DettaglioPagamento) HasDataLimitePagamentoAnticipato() bool`

HasDataLimitePagamentoAnticipato returns a boolean if a field has been set.

### SetDataLimitePagamentoAnticipatoNil

`func (o *DettaglioPagamento) SetDataLimitePagamentoAnticipatoNil(b bool)`

 SetDataLimitePagamentoAnticipatoNil sets the value for DataLimitePagamentoAnticipato to be an explicit nil

### UnsetDataLimitePagamentoAnticipato
`func (o *DettaglioPagamento) UnsetDataLimitePagamentoAnticipato()`

UnsetDataLimitePagamentoAnticipato ensures that no value is present for DataLimitePagamentoAnticipato, not even an explicit nil
### GetPenalitaPagamentiRitardati

`func (o *DettaglioPagamento) GetPenalitaPagamentiRitardati() float64`

GetPenalitaPagamentiRitardati returns the PenalitaPagamentiRitardati field if non-nil, zero value otherwise.

### GetPenalitaPagamentiRitardatiOk

`func (o *DettaglioPagamento) GetPenalitaPagamentiRitardatiOk() (*float64, bool)`

GetPenalitaPagamentiRitardatiOk returns a tuple with the PenalitaPagamentiRitardati field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenalitaPagamentiRitardati

`func (o *DettaglioPagamento) SetPenalitaPagamentiRitardati(v float64)`

SetPenalitaPagamentiRitardati sets PenalitaPagamentiRitardati field to given value.

### HasPenalitaPagamentiRitardati

`func (o *DettaglioPagamento) HasPenalitaPagamentiRitardati() bool`

HasPenalitaPagamentiRitardati returns a boolean if a field has been set.

### SetPenalitaPagamentiRitardatiNil

`func (o *DettaglioPagamento) SetPenalitaPagamentiRitardatiNil(b bool)`

 SetPenalitaPagamentiRitardatiNil sets the value for PenalitaPagamentiRitardati to be an explicit nil

### UnsetPenalitaPagamentiRitardati
`func (o *DettaglioPagamento) UnsetPenalitaPagamentiRitardati()`

UnsetPenalitaPagamentiRitardati ensures that no value is present for PenalitaPagamentiRitardati, not even an explicit nil
### GetDataDecorrenzaPenale

`func (o *DettaglioPagamento) GetDataDecorrenzaPenale() time.Time`

GetDataDecorrenzaPenale returns the DataDecorrenzaPenale field if non-nil, zero value otherwise.

### GetDataDecorrenzaPenaleOk

`func (o *DettaglioPagamento) GetDataDecorrenzaPenaleOk() (*time.Time, bool)`

GetDataDecorrenzaPenaleOk returns a tuple with the DataDecorrenzaPenale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDecorrenzaPenale

`func (o *DettaglioPagamento) SetDataDecorrenzaPenale(v time.Time)`

SetDataDecorrenzaPenale sets DataDecorrenzaPenale field to given value.

### HasDataDecorrenzaPenale

`func (o *DettaglioPagamento) HasDataDecorrenzaPenale() bool`

HasDataDecorrenzaPenale returns a boolean if a field has been set.

### SetDataDecorrenzaPenaleNil

`func (o *DettaglioPagamento) SetDataDecorrenzaPenaleNil(b bool)`

 SetDataDecorrenzaPenaleNil sets the value for DataDecorrenzaPenale to be an explicit nil

### UnsetDataDecorrenzaPenale
`func (o *DettaglioPagamento) UnsetDataDecorrenzaPenale()`

UnsetDataDecorrenzaPenale ensures that no value is present for DataDecorrenzaPenale, not even an explicit nil
### GetCodicePagamento

`func (o *DettaglioPagamento) GetCodicePagamento() string`

GetCodicePagamento returns the CodicePagamento field if non-nil, zero value otherwise.

### GetCodicePagamentoOk

`func (o *DettaglioPagamento) GetCodicePagamentoOk() (*string, bool)`

GetCodicePagamentoOk returns a tuple with the CodicePagamento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodicePagamento

`func (o *DettaglioPagamento) SetCodicePagamento(v string)`

SetCodicePagamento sets CodicePagamento field to given value.

### HasCodicePagamento

`func (o *DettaglioPagamento) HasCodicePagamento() bool`

HasCodicePagamento returns a boolean if a field has been set.

### SetCodicePagamentoNil

`func (o *DettaglioPagamento) SetCodicePagamentoNil(b bool)`

 SetCodicePagamentoNil sets the value for CodicePagamento to be an explicit nil

### UnsetCodicePagamento
`func (o *DettaglioPagamento) UnsetCodicePagamento()`

UnsetCodicePagamento ensures that no value is present for CodicePagamento, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


