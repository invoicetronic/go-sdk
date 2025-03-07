/*
Invoicetronic API

The [Invoicetronic API][2] is a RESTful service that allows you to send and receive invoices through the Italian [Servizio di Interscambio (SDI)][1], or Interchange Service. The API is designed to be simple and easy to use, abstracting away SDI complexity while providing complete control over the invoice send/receive process. It provides advanced features as encryption at rest, multi-language pre-flight invoice validation, multiple upload formats, webhooks, event logging, client SDKs, and CLI tools.  For more information, see  [Invoicetronic website][2]  [1]: https://www.fatturapa.gov.it/it/sistemainterscambio/cose-il-sdi/ [2]: https://invoicetronic.com/

API version: 1
Contact: support@invoicetronic.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package invoicetronicsdk

import (
	"encoding/json"
	"time"
)

// checks if the DettaglioPagamento type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DettaglioPagamento{}

// DettaglioPagamento struct for DettaglioPagamento
type DettaglioPagamento struct {
	Beneficiario NullableString `json:"beneficiario,omitempty"`
	ModalitaPagamento NullableString `json:"modalita_pagamento,omitempty"`
	DataRiferimentoTerminiPagamento NullableTime `json:"data_riferimento_termini_pagamento,omitempty"`
	GiorniTerminiPagamento NullableInt32 `json:"giorni_termini_pagamento,omitempty"`
	DataScadenzaPagamento NullableTime `json:"data_scadenza_pagamento,omitempty"`
	ImportoPagamento *float64 `json:"importo_pagamento,omitempty"`
	CodUfficioPostale NullableString `json:"cod_ufficio_postale,omitempty"`
	CognomeQuietanzante NullableString `json:"cognome_quietanzante,omitempty"`
	NomeQuietanzante NullableString `json:"nome_quietanzante,omitempty"`
	CfQuietanzante NullableString `json:"cf_quietanzante,omitempty"`
	TitoloQuietanzante NullableString `json:"titolo_quietanzante,omitempty"`
	IstitutoFinanziario NullableString `json:"istituto_finanziario,omitempty"`
	Iban NullableString `json:"iban,omitempty"`
	Abi NullableString `json:"abi,omitempty"`
	Cab NullableString `json:"cab,omitempty"`
	Bic NullableString `json:"bic,omitempty"`
	ScontoPagamentoAnticipato NullableFloat64 `json:"sconto_pagamento_anticipato,omitempty"`
	DataLimitePagamentoAnticipato NullableTime `json:"data_limite_pagamento_anticipato,omitempty"`
	PenalitaPagamentiRitardati NullableFloat64 `json:"penalita_pagamenti_ritardati,omitempty"`
	DataDecorrenzaPenale NullableTime `json:"data_decorrenza_penale,omitempty"`
	CodicePagamento NullableString `json:"codice_pagamento,omitempty"`
}

// NewDettaglioPagamento instantiates a new DettaglioPagamento object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDettaglioPagamento() *DettaglioPagamento {
	this := DettaglioPagamento{}
	return &this
}

// NewDettaglioPagamentoWithDefaults instantiates a new DettaglioPagamento object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDettaglioPagamentoWithDefaults() *DettaglioPagamento {
	this := DettaglioPagamento{}
	return &this
}

// GetBeneficiario returns the Beneficiario field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetBeneficiario() string {
	if o == nil || IsNil(o.Beneficiario.Get()) {
		var ret string
		return ret
	}
	return *o.Beneficiario.Get()
}

// GetBeneficiarioOk returns a tuple with the Beneficiario field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetBeneficiarioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Beneficiario.Get(), o.Beneficiario.IsSet()
}

// HasBeneficiario returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasBeneficiario() bool {
	if o != nil && o.Beneficiario.IsSet() {
		return true
	}

	return false
}

// SetBeneficiario gets a reference to the given NullableString and assigns it to the Beneficiario field.
func (o *DettaglioPagamento) SetBeneficiario(v string) {
	o.Beneficiario.Set(&v)
}
// SetBeneficiarioNil sets the value for Beneficiario to be an explicit nil
func (o *DettaglioPagamento) SetBeneficiarioNil() {
	o.Beneficiario.Set(nil)
}

// UnsetBeneficiario ensures that no value is present for Beneficiario, not even an explicit nil
func (o *DettaglioPagamento) UnsetBeneficiario() {
	o.Beneficiario.Unset()
}

// GetModalitaPagamento returns the ModalitaPagamento field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetModalitaPagamento() string {
	if o == nil || IsNil(o.ModalitaPagamento.Get()) {
		var ret string
		return ret
	}
	return *o.ModalitaPagamento.Get()
}

// GetModalitaPagamentoOk returns a tuple with the ModalitaPagamento field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetModalitaPagamentoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModalitaPagamento.Get(), o.ModalitaPagamento.IsSet()
}

// HasModalitaPagamento returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasModalitaPagamento() bool {
	if o != nil && o.ModalitaPagamento.IsSet() {
		return true
	}

	return false
}

// SetModalitaPagamento gets a reference to the given NullableString and assigns it to the ModalitaPagamento field.
func (o *DettaglioPagamento) SetModalitaPagamento(v string) {
	o.ModalitaPagamento.Set(&v)
}
// SetModalitaPagamentoNil sets the value for ModalitaPagamento to be an explicit nil
func (o *DettaglioPagamento) SetModalitaPagamentoNil() {
	o.ModalitaPagamento.Set(nil)
}

// UnsetModalitaPagamento ensures that no value is present for ModalitaPagamento, not even an explicit nil
func (o *DettaglioPagamento) UnsetModalitaPagamento() {
	o.ModalitaPagamento.Unset()
}

// GetDataRiferimentoTerminiPagamento returns the DataRiferimentoTerminiPagamento field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetDataRiferimentoTerminiPagamento() time.Time {
	if o == nil || IsNil(o.DataRiferimentoTerminiPagamento.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DataRiferimentoTerminiPagamento.Get()
}

// GetDataRiferimentoTerminiPagamentoOk returns a tuple with the DataRiferimentoTerminiPagamento field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetDataRiferimentoTerminiPagamentoOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataRiferimentoTerminiPagamento.Get(), o.DataRiferimentoTerminiPagamento.IsSet()
}

// HasDataRiferimentoTerminiPagamento returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasDataRiferimentoTerminiPagamento() bool {
	if o != nil && o.DataRiferimentoTerminiPagamento.IsSet() {
		return true
	}

	return false
}

// SetDataRiferimentoTerminiPagamento gets a reference to the given NullableTime and assigns it to the DataRiferimentoTerminiPagamento field.
func (o *DettaglioPagamento) SetDataRiferimentoTerminiPagamento(v time.Time) {
	o.DataRiferimentoTerminiPagamento.Set(&v)
}
// SetDataRiferimentoTerminiPagamentoNil sets the value for DataRiferimentoTerminiPagamento to be an explicit nil
func (o *DettaglioPagamento) SetDataRiferimentoTerminiPagamentoNil() {
	o.DataRiferimentoTerminiPagamento.Set(nil)
}

// UnsetDataRiferimentoTerminiPagamento ensures that no value is present for DataRiferimentoTerminiPagamento, not even an explicit nil
func (o *DettaglioPagamento) UnsetDataRiferimentoTerminiPagamento() {
	o.DataRiferimentoTerminiPagamento.Unset()
}

// GetGiorniTerminiPagamento returns the GiorniTerminiPagamento field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetGiorniTerminiPagamento() int32 {
	if o == nil || IsNil(o.GiorniTerminiPagamento.Get()) {
		var ret int32
		return ret
	}
	return *o.GiorniTerminiPagamento.Get()
}

// GetGiorniTerminiPagamentoOk returns a tuple with the GiorniTerminiPagamento field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetGiorniTerminiPagamentoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiorniTerminiPagamento.Get(), o.GiorniTerminiPagamento.IsSet()
}

// HasGiorniTerminiPagamento returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasGiorniTerminiPagamento() bool {
	if o != nil && o.GiorniTerminiPagamento.IsSet() {
		return true
	}

	return false
}

// SetGiorniTerminiPagamento gets a reference to the given NullableInt32 and assigns it to the GiorniTerminiPagamento field.
func (o *DettaglioPagamento) SetGiorniTerminiPagamento(v int32) {
	o.GiorniTerminiPagamento.Set(&v)
}
// SetGiorniTerminiPagamentoNil sets the value for GiorniTerminiPagamento to be an explicit nil
func (o *DettaglioPagamento) SetGiorniTerminiPagamentoNil() {
	o.GiorniTerminiPagamento.Set(nil)
}

// UnsetGiorniTerminiPagamento ensures that no value is present for GiorniTerminiPagamento, not even an explicit nil
func (o *DettaglioPagamento) UnsetGiorniTerminiPagamento() {
	o.GiorniTerminiPagamento.Unset()
}

// GetDataScadenzaPagamento returns the DataScadenzaPagamento field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetDataScadenzaPagamento() time.Time {
	if o == nil || IsNil(o.DataScadenzaPagamento.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DataScadenzaPagamento.Get()
}

// GetDataScadenzaPagamentoOk returns a tuple with the DataScadenzaPagamento field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetDataScadenzaPagamentoOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataScadenzaPagamento.Get(), o.DataScadenzaPagamento.IsSet()
}

// HasDataScadenzaPagamento returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasDataScadenzaPagamento() bool {
	if o != nil && o.DataScadenzaPagamento.IsSet() {
		return true
	}

	return false
}

// SetDataScadenzaPagamento gets a reference to the given NullableTime and assigns it to the DataScadenzaPagamento field.
func (o *DettaglioPagamento) SetDataScadenzaPagamento(v time.Time) {
	o.DataScadenzaPagamento.Set(&v)
}
// SetDataScadenzaPagamentoNil sets the value for DataScadenzaPagamento to be an explicit nil
func (o *DettaglioPagamento) SetDataScadenzaPagamentoNil() {
	o.DataScadenzaPagamento.Set(nil)
}

// UnsetDataScadenzaPagamento ensures that no value is present for DataScadenzaPagamento, not even an explicit nil
func (o *DettaglioPagamento) UnsetDataScadenzaPagamento() {
	o.DataScadenzaPagamento.Unset()
}

// GetImportoPagamento returns the ImportoPagamento field value if set, zero value otherwise.
func (o *DettaglioPagamento) GetImportoPagamento() float64 {
	if o == nil || IsNil(o.ImportoPagamento) {
		var ret float64
		return ret
	}
	return *o.ImportoPagamento
}

// GetImportoPagamentoOk returns a tuple with the ImportoPagamento field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DettaglioPagamento) GetImportoPagamentoOk() (*float64, bool) {
	if o == nil || IsNil(o.ImportoPagamento) {
		return nil, false
	}
	return o.ImportoPagamento, true
}

// HasImportoPagamento returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasImportoPagamento() bool {
	if o != nil && !IsNil(o.ImportoPagamento) {
		return true
	}

	return false
}

// SetImportoPagamento gets a reference to the given float64 and assigns it to the ImportoPagamento field.
func (o *DettaglioPagamento) SetImportoPagamento(v float64) {
	o.ImportoPagamento = &v
}

// GetCodUfficioPostale returns the CodUfficioPostale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetCodUfficioPostale() string {
	if o == nil || IsNil(o.CodUfficioPostale.Get()) {
		var ret string
		return ret
	}
	return *o.CodUfficioPostale.Get()
}

// GetCodUfficioPostaleOk returns a tuple with the CodUfficioPostale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetCodUfficioPostaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodUfficioPostale.Get(), o.CodUfficioPostale.IsSet()
}

// HasCodUfficioPostale returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasCodUfficioPostale() bool {
	if o != nil && o.CodUfficioPostale.IsSet() {
		return true
	}

	return false
}

// SetCodUfficioPostale gets a reference to the given NullableString and assigns it to the CodUfficioPostale field.
func (o *DettaglioPagamento) SetCodUfficioPostale(v string) {
	o.CodUfficioPostale.Set(&v)
}
// SetCodUfficioPostaleNil sets the value for CodUfficioPostale to be an explicit nil
func (o *DettaglioPagamento) SetCodUfficioPostaleNil() {
	o.CodUfficioPostale.Set(nil)
}

// UnsetCodUfficioPostale ensures that no value is present for CodUfficioPostale, not even an explicit nil
func (o *DettaglioPagamento) UnsetCodUfficioPostale() {
	o.CodUfficioPostale.Unset()
}

// GetCognomeQuietanzante returns the CognomeQuietanzante field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetCognomeQuietanzante() string {
	if o == nil || IsNil(o.CognomeQuietanzante.Get()) {
		var ret string
		return ret
	}
	return *o.CognomeQuietanzante.Get()
}

// GetCognomeQuietanzanteOk returns a tuple with the CognomeQuietanzante field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetCognomeQuietanzanteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CognomeQuietanzante.Get(), o.CognomeQuietanzante.IsSet()
}

// HasCognomeQuietanzante returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasCognomeQuietanzante() bool {
	if o != nil && o.CognomeQuietanzante.IsSet() {
		return true
	}

	return false
}

// SetCognomeQuietanzante gets a reference to the given NullableString and assigns it to the CognomeQuietanzante field.
func (o *DettaglioPagamento) SetCognomeQuietanzante(v string) {
	o.CognomeQuietanzante.Set(&v)
}
// SetCognomeQuietanzanteNil sets the value for CognomeQuietanzante to be an explicit nil
func (o *DettaglioPagamento) SetCognomeQuietanzanteNil() {
	o.CognomeQuietanzante.Set(nil)
}

// UnsetCognomeQuietanzante ensures that no value is present for CognomeQuietanzante, not even an explicit nil
func (o *DettaglioPagamento) UnsetCognomeQuietanzante() {
	o.CognomeQuietanzante.Unset()
}

// GetNomeQuietanzante returns the NomeQuietanzante field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetNomeQuietanzante() string {
	if o == nil || IsNil(o.NomeQuietanzante.Get()) {
		var ret string
		return ret
	}
	return *o.NomeQuietanzante.Get()
}

// GetNomeQuietanzanteOk returns a tuple with the NomeQuietanzante field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetNomeQuietanzanteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NomeQuietanzante.Get(), o.NomeQuietanzante.IsSet()
}

// HasNomeQuietanzante returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasNomeQuietanzante() bool {
	if o != nil && o.NomeQuietanzante.IsSet() {
		return true
	}

	return false
}

// SetNomeQuietanzante gets a reference to the given NullableString and assigns it to the NomeQuietanzante field.
func (o *DettaglioPagamento) SetNomeQuietanzante(v string) {
	o.NomeQuietanzante.Set(&v)
}
// SetNomeQuietanzanteNil sets the value for NomeQuietanzante to be an explicit nil
func (o *DettaglioPagamento) SetNomeQuietanzanteNil() {
	o.NomeQuietanzante.Set(nil)
}

// UnsetNomeQuietanzante ensures that no value is present for NomeQuietanzante, not even an explicit nil
func (o *DettaglioPagamento) UnsetNomeQuietanzante() {
	o.NomeQuietanzante.Unset()
}

// GetCfQuietanzante returns the CfQuietanzante field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetCfQuietanzante() string {
	if o == nil || IsNil(o.CfQuietanzante.Get()) {
		var ret string
		return ret
	}
	return *o.CfQuietanzante.Get()
}

// GetCfQuietanzanteOk returns a tuple with the CfQuietanzante field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetCfQuietanzanteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CfQuietanzante.Get(), o.CfQuietanzante.IsSet()
}

// HasCfQuietanzante returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasCfQuietanzante() bool {
	if o != nil && o.CfQuietanzante.IsSet() {
		return true
	}

	return false
}

// SetCfQuietanzante gets a reference to the given NullableString and assigns it to the CfQuietanzante field.
func (o *DettaglioPagamento) SetCfQuietanzante(v string) {
	o.CfQuietanzante.Set(&v)
}
// SetCfQuietanzanteNil sets the value for CfQuietanzante to be an explicit nil
func (o *DettaglioPagamento) SetCfQuietanzanteNil() {
	o.CfQuietanzante.Set(nil)
}

// UnsetCfQuietanzante ensures that no value is present for CfQuietanzante, not even an explicit nil
func (o *DettaglioPagamento) UnsetCfQuietanzante() {
	o.CfQuietanzante.Unset()
}

// GetTitoloQuietanzante returns the TitoloQuietanzante field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetTitoloQuietanzante() string {
	if o == nil || IsNil(o.TitoloQuietanzante.Get()) {
		var ret string
		return ret
	}
	return *o.TitoloQuietanzante.Get()
}

// GetTitoloQuietanzanteOk returns a tuple with the TitoloQuietanzante field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetTitoloQuietanzanteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TitoloQuietanzante.Get(), o.TitoloQuietanzante.IsSet()
}

// HasTitoloQuietanzante returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasTitoloQuietanzante() bool {
	if o != nil && o.TitoloQuietanzante.IsSet() {
		return true
	}

	return false
}

// SetTitoloQuietanzante gets a reference to the given NullableString and assigns it to the TitoloQuietanzante field.
func (o *DettaglioPagamento) SetTitoloQuietanzante(v string) {
	o.TitoloQuietanzante.Set(&v)
}
// SetTitoloQuietanzanteNil sets the value for TitoloQuietanzante to be an explicit nil
func (o *DettaglioPagamento) SetTitoloQuietanzanteNil() {
	o.TitoloQuietanzante.Set(nil)
}

// UnsetTitoloQuietanzante ensures that no value is present for TitoloQuietanzante, not even an explicit nil
func (o *DettaglioPagamento) UnsetTitoloQuietanzante() {
	o.TitoloQuietanzante.Unset()
}

// GetIstitutoFinanziario returns the IstitutoFinanziario field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetIstitutoFinanziario() string {
	if o == nil || IsNil(o.IstitutoFinanziario.Get()) {
		var ret string
		return ret
	}
	return *o.IstitutoFinanziario.Get()
}

// GetIstitutoFinanziarioOk returns a tuple with the IstitutoFinanziario field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetIstitutoFinanziarioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IstitutoFinanziario.Get(), o.IstitutoFinanziario.IsSet()
}

// HasIstitutoFinanziario returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasIstitutoFinanziario() bool {
	if o != nil && o.IstitutoFinanziario.IsSet() {
		return true
	}

	return false
}

// SetIstitutoFinanziario gets a reference to the given NullableString and assigns it to the IstitutoFinanziario field.
func (o *DettaglioPagamento) SetIstitutoFinanziario(v string) {
	o.IstitutoFinanziario.Set(&v)
}
// SetIstitutoFinanziarioNil sets the value for IstitutoFinanziario to be an explicit nil
func (o *DettaglioPagamento) SetIstitutoFinanziarioNil() {
	o.IstitutoFinanziario.Set(nil)
}

// UnsetIstitutoFinanziario ensures that no value is present for IstitutoFinanziario, not even an explicit nil
func (o *DettaglioPagamento) UnsetIstitutoFinanziario() {
	o.IstitutoFinanziario.Unset()
}

// GetIban returns the Iban field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetIban() string {
	if o == nil || IsNil(o.Iban.Get()) {
		var ret string
		return ret
	}
	return *o.Iban.Get()
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetIbanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iban.Get(), o.Iban.IsSet()
}

// HasIban returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasIban() bool {
	if o != nil && o.Iban.IsSet() {
		return true
	}

	return false
}

// SetIban gets a reference to the given NullableString and assigns it to the Iban field.
func (o *DettaglioPagamento) SetIban(v string) {
	o.Iban.Set(&v)
}
// SetIbanNil sets the value for Iban to be an explicit nil
func (o *DettaglioPagamento) SetIbanNil() {
	o.Iban.Set(nil)
}

// UnsetIban ensures that no value is present for Iban, not even an explicit nil
func (o *DettaglioPagamento) UnsetIban() {
	o.Iban.Unset()
}

// GetAbi returns the Abi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetAbi() string {
	if o == nil || IsNil(o.Abi.Get()) {
		var ret string
		return ret
	}
	return *o.Abi.Get()
}

// GetAbiOk returns a tuple with the Abi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetAbiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Abi.Get(), o.Abi.IsSet()
}

// HasAbi returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasAbi() bool {
	if o != nil && o.Abi.IsSet() {
		return true
	}

	return false
}

// SetAbi gets a reference to the given NullableString and assigns it to the Abi field.
func (o *DettaglioPagamento) SetAbi(v string) {
	o.Abi.Set(&v)
}
// SetAbiNil sets the value for Abi to be an explicit nil
func (o *DettaglioPagamento) SetAbiNil() {
	o.Abi.Set(nil)
}

// UnsetAbi ensures that no value is present for Abi, not even an explicit nil
func (o *DettaglioPagamento) UnsetAbi() {
	o.Abi.Unset()
}

// GetCab returns the Cab field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetCab() string {
	if o == nil || IsNil(o.Cab.Get()) {
		var ret string
		return ret
	}
	return *o.Cab.Get()
}

// GetCabOk returns a tuple with the Cab field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetCabOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cab.Get(), o.Cab.IsSet()
}

// HasCab returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasCab() bool {
	if o != nil && o.Cab.IsSet() {
		return true
	}

	return false
}

// SetCab gets a reference to the given NullableString and assigns it to the Cab field.
func (o *DettaglioPagamento) SetCab(v string) {
	o.Cab.Set(&v)
}
// SetCabNil sets the value for Cab to be an explicit nil
func (o *DettaglioPagamento) SetCabNil() {
	o.Cab.Set(nil)
}

// UnsetCab ensures that no value is present for Cab, not even an explicit nil
func (o *DettaglioPagamento) UnsetCab() {
	o.Cab.Unset()
}

// GetBic returns the Bic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetBic() string {
	if o == nil || IsNil(o.Bic.Get()) {
		var ret string
		return ret
	}
	return *o.Bic.Get()
}

// GetBicOk returns a tuple with the Bic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetBicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bic.Get(), o.Bic.IsSet()
}

// HasBic returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasBic() bool {
	if o != nil && o.Bic.IsSet() {
		return true
	}

	return false
}

// SetBic gets a reference to the given NullableString and assigns it to the Bic field.
func (o *DettaglioPagamento) SetBic(v string) {
	o.Bic.Set(&v)
}
// SetBicNil sets the value for Bic to be an explicit nil
func (o *DettaglioPagamento) SetBicNil() {
	o.Bic.Set(nil)
}

// UnsetBic ensures that no value is present for Bic, not even an explicit nil
func (o *DettaglioPagamento) UnsetBic() {
	o.Bic.Unset()
}

// GetScontoPagamentoAnticipato returns the ScontoPagamentoAnticipato field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetScontoPagamentoAnticipato() float64 {
	if o == nil || IsNil(o.ScontoPagamentoAnticipato.Get()) {
		var ret float64
		return ret
	}
	return *o.ScontoPagamentoAnticipato.Get()
}

// GetScontoPagamentoAnticipatoOk returns a tuple with the ScontoPagamentoAnticipato field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetScontoPagamentoAnticipatoOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScontoPagamentoAnticipato.Get(), o.ScontoPagamentoAnticipato.IsSet()
}

// HasScontoPagamentoAnticipato returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasScontoPagamentoAnticipato() bool {
	if o != nil && o.ScontoPagamentoAnticipato.IsSet() {
		return true
	}

	return false
}

// SetScontoPagamentoAnticipato gets a reference to the given NullableFloat64 and assigns it to the ScontoPagamentoAnticipato field.
func (o *DettaglioPagamento) SetScontoPagamentoAnticipato(v float64) {
	o.ScontoPagamentoAnticipato.Set(&v)
}
// SetScontoPagamentoAnticipatoNil sets the value for ScontoPagamentoAnticipato to be an explicit nil
func (o *DettaglioPagamento) SetScontoPagamentoAnticipatoNil() {
	o.ScontoPagamentoAnticipato.Set(nil)
}

// UnsetScontoPagamentoAnticipato ensures that no value is present for ScontoPagamentoAnticipato, not even an explicit nil
func (o *DettaglioPagamento) UnsetScontoPagamentoAnticipato() {
	o.ScontoPagamentoAnticipato.Unset()
}

// GetDataLimitePagamentoAnticipato returns the DataLimitePagamentoAnticipato field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetDataLimitePagamentoAnticipato() time.Time {
	if o == nil || IsNil(o.DataLimitePagamentoAnticipato.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DataLimitePagamentoAnticipato.Get()
}

// GetDataLimitePagamentoAnticipatoOk returns a tuple with the DataLimitePagamentoAnticipato field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetDataLimitePagamentoAnticipatoOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataLimitePagamentoAnticipato.Get(), o.DataLimitePagamentoAnticipato.IsSet()
}

// HasDataLimitePagamentoAnticipato returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasDataLimitePagamentoAnticipato() bool {
	if o != nil && o.DataLimitePagamentoAnticipato.IsSet() {
		return true
	}

	return false
}

// SetDataLimitePagamentoAnticipato gets a reference to the given NullableTime and assigns it to the DataLimitePagamentoAnticipato field.
func (o *DettaglioPagamento) SetDataLimitePagamentoAnticipato(v time.Time) {
	o.DataLimitePagamentoAnticipato.Set(&v)
}
// SetDataLimitePagamentoAnticipatoNil sets the value for DataLimitePagamentoAnticipato to be an explicit nil
func (o *DettaglioPagamento) SetDataLimitePagamentoAnticipatoNil() {
	o.DataLimitePagamentoAnticipato.Set(nil)
}

// UnsetDataLimitePagamentoAnticipato ensures that no value is present for DataLimitePagamentoAnticipato, not even an explicit nil
func (o *DettaglioPagamento) UnsetDataLimitePagamentoAnticipato() {
	o.DataLimitePagamentoAnticipato.Unset()
}

// GetPenalitaPagamentiRitardati returns the PenalitaPagamentiRitardati field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetPenalitaPagamentiRitardati() float64 {
	if o == nil || IsNil(o.PenalitaPagamentiRitardati.Get()) {
		var ret float64
		return ret
	}
	return *o.PenalitaPagamentiRitardati.Get()
}

// GetPenalitaPagamentiRitardatiOk returns a tuple with the PenalitaPagamentiRitardati field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetPenalitaPagamentiRitardatiOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PenalitaPagamentiRitardati.Get(), o.PenalitaPagamentiRitardati.IsSet()
}

// HasPenalitaPagamentiRitardati returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasPenalitaPagamentiRitardati() bool {
	if o != nil && o.PenalitaPagamentiRitardati.IsSet() {
		return true
	}

	return false
}

// SetPenalitaPagamentiRitardati gets a reference to the given NullableFloat64 and assigns it to the PenalitaPagamentiRitardati field.
func (o *DettaglioPagamento) SetPenalitaPagamentiRitardati(v float64) {
	o.PenalitaPagamentiRitardati.Set(&v)
}
// SetPenalitaPagamentiRitardatiNil sets the value for PenalitaPagamentiRitardati to be an explicit nil
func (o *DettaglioPagamento) SetPenalitaPagamentiRitardatiNil() {
	o.PenalitaPagamentiRitardati.Set(nil)
}

// UnsetPenalitaPagamentiRitardati ensures that no value is present for PenalitaPagamentiRitardati, not even an explicit nil
func (o *DettaglioPagamento) UnsetPenalitaPagamentiRitardati() {
	o.PenalitaPagamentiRitardati.Unset()
}

// GetDataDecorrenzaPenale returns the DataDecorrenzaPenale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetDataDecorrenzaPenale() time.Time {
	if o == nil || IsNil(o.DataDecorrenzaPenale.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DataDecorrenzaPenale.Get()
}

// GetDataDecorrenzaPenaleOk returns a tuple with the DataDecorrenzaPenale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetDataDecorrenzaPenaleOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataDecorrenzaPenale.Get(), o.DataDecorrenzaPenale.IsSet()
}

// HasDataDecorrenzaPenale returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasDataDecorrenzaPenale() bool {
	if o != nil && o.DataDecorrenzaPenale.IsSet() {
		return true
	}

	return false
}

// SetDataDecorrenzaPenale gets a reference to the given NullableTime and assigns it to the DataDecorrenzaPenale field.
func (o *DettaglioPagamento) SetDataDecorrenzaPenale(v time.Time) {
	o.DataDecorrenzaPenale.Set(&v)
}
// SetDataDecorrenzaPenaleNil sets the value for DataDecorrenzaPenale to be an explicit nil
func (o *DettaglioPagamento) SetDataDecorrenzaPenaleNil() {
	o.DataDecorrenzaPenale.Set(nil)
}

// UnsetDataDecorrenzaPenale ensures that no value is present for DataDecorrenzaPenale, not even an explicit nil
func (o *DettaglioPagamento) UnsetDataDecorrenzaPenale() {
	o.DataDecorrenzaPenale.Unset()
}

// GetCodicePagamento returns the CodicePagamento field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioPagamento) GetCodicePagamento() string {
	if o == nil || IsNil(o.CodicePagamento.Get()) {
		var ret string
		return ret
	}
	return *o.CodicePagamento.Get()
}

// GetCodicePagamentoOk returns a tuple with the CodicePagamento field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioPagamento) GetCodicePagamentoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodicePagamento.Get(), o.CodicePagamento.IsSet()
}

// HasCodicePagamento returns a boolean if a field has been set.
func (o *DettaglioPagamento) HasCodicePagamento() bool {
	if o != nil && o.CodicePagamento.IsSet() {
		return true
	}

	return false
}

// SetCodicePagamento gets a reference to the given NullableString and assigns it to the CodicePagamento field.
func (o *DettaglioPagamento) SetCodicePagamento(v string) {
	o.CodicePagamento.Set(&v)
}
// SetCodicePagamentoNil sets the value for CodicePagamento to be an explicit nil
func (o *DettaglioPagamento) SetCodicePagamentoNil() {
	o.CodicePagamento.Set(nil)
}

// UnsetCodicePagamento ensures that no value is present for CodicePagamento, not even an explicit nil
func (o *DettaglioPagamento) UnsetCodicePagamento() {
	o.CodicePagamento.Unset()
}

func (o DettaglioPagamento) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DettaglioPagamento) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Beneficiario.IsSet() {
		toSerialize["beneficiario"] = o.Beneficiario.Get()
	}
	if o.ModalitaPagamento.IsSet() {
		toSerialize["modalita_pagamento"] = o.ModalitaPagamento.Get()
	}
	if o.DataRiferimentoTerminiPagamento.IsSet() {
		toSerialize["data_riferimento_termini_pagamento"] = o.DataRiferimentoTerminiPagamento.Get()
	}
	if o.GiorniTerminiPagamento.IsSet() {
		toSerialize["giorni_termini_pagamento"] = o.GiorniTerminiPagamento.Get()
	}
	if o.DataScadenzaPagamento.IsSet() {
		toSerialize["data_scadenza_pagamento"] = o.DataScadenzaPagamento.Get()
	}
	if !IsNil(o.ImportoPagamento) {
		toSerialize["importo_pagamento"] = o.ImportoPagamento
	}
	if o.CodUfficioPostale.IsSet() {
		toSerialize["cod_ufficio_postale"] = o.CodUfficioPostale.Get()
	}
	if o.CognomeQuietanzante.IsSet() {
		toSerialize["cognome_quietanzante"] = o.CognomeQuietanzante.Get()
	}
	if o.NomeQuietanzante.IsSet() {
		toSerialize["nome_quietanzante"] = o.NomeQuietanzante.Get()
	}
	if o.CfQuietanzante.IsSet() {
		toSerialize["cf_quietanzante"] = o.CfQuietanzante.Get()
	}
	if o.TitoloQuietanzante.IsSet() {
		toSerialize["titolo_quietanzante"] = o.TitoloQuietanzante.Get()
	}
	if o.IstitutoFinanziario.IsSet() {
		toSerialize["istituto_finanziario"] = o.IstitutoFinanziario.Get()
	}
	if o.Iban.IsSet() {
		toSerialize["iban"] = o.Iban.Get()
	}
	if o.Abi.IsSet() {
		toSerialize["abi"] = o.Abi.Get()
	}
	if o.Cab.IsSet() {
		toSerialize["cab"] = o.Cab.Get()
	}
	if o.Bic.IsSet() {
		toSerialize["bic"] = o.Bic.Get()
	}
	if o.ScontoPagamentoAnticipato.IsSet() {
		toSerialize["sconto_pagamento_anticipato"] = o.ScontoPagamentoAnticipato.Get()
	}
	if o.DataLimitePagamentoAnticipato.IsSet() {
		toSerialize["data_limite_pagamento_anticipato"] = o.DataLimitePagamentoAnticipato.Get()
	}
	if o.PenalitaPagamentiRitardati.IsSet() {
		toSerialize["penalita_pagamenti_ritardati"] = o.PenalitaPagamentiRitardati.Get()
	}
	if o.DataDecorrenzaPenale.IsSet() {
		toSerialize["data_decorrenza_penale"] = o.DataDecorrenzaPenale.Get()
	}
	if o.CodicePagamento.IsSet() {
		toSerialize["codice_pagamento"] = o.CodicePagamento.Get()
	}
	return toSerialize, nil
}

type NullableDettaglioPagamento struct {
	value *DettaglioPagamento
	isSet bool
}

func (v NullableDettaglioPagamento) Get() *DettaglioPagamento {
	return v.value
}

func (v *NullableDettaglioPagamento) Set(val *DettaglioPagamento) {
	v.value = val
	v.isSet = true
}

func (v NullableDettaglioPagamento) IsSet() bool {
	return v.isSet
}

func (v *NullableDettaglioPagamento) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDettaglioPagamento(val *DettaglioPagamento) *NullableDettaglioPagamento {
	return &NullableDettaglioPagamento{value: val, isSet: true}
}

func (v NullableDettaglioPagamento) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDettaglioPagamento) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


