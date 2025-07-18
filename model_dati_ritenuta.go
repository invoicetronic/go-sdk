/*
Invoicetronic API

The [Invoicetronic API][2] is a RESTful service that allows you to send and receive invoices through the Italian [Servizio di Interscambio (SDI)][1], or Interchange Service. The API is designed to be simple and easy to use, abstracting away SDI complexity while providing complete control over the invoice send/receive process. It provides advanced features as encryption at rest, multi-language pre-flight invoice validation, multiple upload formats, webhooks, event logging, client SDKs, and CLI tools.  For more information, see  [Invoicetronic website][2]  [1]: https://www.fatturapa.gov.it/it/sistemainterscambio/cose-il-sdi/ [2]: https://invoicetronic.com/

API version: 1
Contact: info@invoicetronic.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package invoicetronicsdk

import (
	"encoding/json"
)

// checks if the DatiRitenuta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatiRitenuta{}

// DatiRitenuta struct for DatiRitenuta
type DatiRitenuta struct {
	TipoRitenuta NullableString `json:"tipo_ritenuta,omitempty"`
	ImportoRitenuta *float64 `json:"importo_ritenuta,omitempty"`
	AliquotaRitenuta *float64 `json:"aliquota_ritenuta,omitempty"`
	CausalePagamento NullableString `json:"causale_pagamento,omitempty"`
}

// NewDatiRitenuta instantiates a new DatiRitenuta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatiRitenuta() *DatiRitenuta {
	this := DatiRitenuta{}
	return &this
}

// NewDatiRitenutaWithDefaults instantiates a new DatiRitenuta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatiRitenutaWithDefaults() *DatiRitenuta {
	this := DatiRitenuta{}
	return &this
}

// GetTipoRitenuta returns the TipoRitenuta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiRitenuta) GetTipoRitenuta() string {
	if o == nil || IsNil(o.TipoRitenuta.Get()) {
		var ret string
		return ret
	}
	return *o.TipoRitenuta.Get()
}

// GetTipoRitenutaOk returns a tuple with the TipoRitenuta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiRitenuta) GetTipoRitenutaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TipoRitenuta.Get(), o.TipoRitenuta.IsSet()
}

// HasTipoRitenuta returns a boolean if a field has been set.
func (o *DatiRitenuta) HasTipoRitenuta() bool {
	if o != nil && o.TipoRitenuta.IsSet() {
		return true
	}

	return false
}

// SetTipoRitenuta gets a reference to the given NullableString and assigns it to the TipoRitenuta field.
func (o *DatiRitenuta) SetTipoRitenuta(v string) {
	o.TipoRitenuta.Set(&v)
}
// SetTipoRitenutaNil sets the value for TipoRitenuta to be an explicit nil
func (o *DatiRitenuta) SetTipoRitenutaNil() {
	o.TipoRitenuta.Set(nil)
}

// UnsetTipoRitenuta ensures that no value is present for TipoRitenuta, not even an explicit nil
func (o *DatiRitenuta) UnsetTipoRitenuta() {
	o.TipoRitenuta.Unset()
}

// GetImportoRitenuta returns the ImportoRitenuta field value if set, zero value otherwise.
func (o *DatiRitenuta) GetImportoRitenuta() float64 {
	if o == nil || IsNil(o.ImportoRitenuta) {
		var ret float64
		return ret
	}
	return *o.ImportoRitenuta
}

// GetImportoRitenutaOk returns a tuple with the ImportoRitenuta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatiRitenuta) GetImportoRitenutaOk() (*float64, bool) {
	if o == nil || IsNil(o.ImportoRitenuta) {
		return nil, false
	}
	return o.ImportoRitenuta, true
}

// HasImportoRitenuta returns a boolean if a field has been set.
func (o *DatiRitenuta) HasImportoRitenuta() bool {
	if o != nil && !IsNil(o.ImportoRitenuta) {
		return true
	}

	return false
}

// SetImportoRitenuta gets a reference to the given float64 and assigns it to the ImportoRitenuta field.
func (o *DatiRitenuta) SetImportoRitenuta(v float64) {
	o.ImportoRitenuta = &v
}

// GetAliquotaRitenuta returns the AliquotaRitenuta field value if set, zero value otherwise.
func (o *DatiRitenuta) GetAliquotaRitenuta() float64 {
	if o == nil || IsNil(o.AliquotaRitenuta) {
		var ret float64
		return ret
	}
	return *o.AliquotaRitenuta
}

// GetAliquotaRitenutaOk returns a tuple with the AliquotaRitenuta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatiRitenuta) GetAliquotaRitenutaOk() (*float64, bool) {
	if o == nil || IsNil(o.AliquotaRitenuta) {
		return nil, false
	}
	return o.AliquotaRitenuta, true
}

// HasAliquotaRitenuta returns a boolean if a field has been set.
func (o *DatiRitenuta) HasAliquotaRitenuta() bool {
	if o != nil && !IsNil(o.AliquotaRitenuta) {
		return true
	}

	return false
}

// SetAliquotaRitenuta gets a reference to the given float64 and assigns it to the AliquotaRitenuta field.
func (o *DatiRitenuta) SetAliquotaRitenuta(v float64) {
	o.AliquotaRitenuta = &v
}

// GetCausalePagamento returns the CausalePagamento field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiRitenuta) GetCausalePagamento() string {
	if o == nil || IsNil(o.CausalePagamento.Get()) {
		var ret string
		return ret
	}
	return *o.CausalePagamento.Get()
}

// GetCausalePagamentoOk returns a tuple with the CausalePagamento field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiRitenuta) GetCausalePagamentoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CausalePagamento.Get(), o.CausalePagamento.IsSet()
}

// HasCausalePagamento returns a boolean if a field has been set.
func (o *DatiRitenuta) HasCausalePagamento() bool {
	if o != nil && o.CausalePagamento.IsSet() {
		return true
	}

	return false
}

// SetCausalePagamento gets a reference to the given NullableString and assigns it to the CausalePagamento field.
func (o *DatiRitenuta) SetCausalePagamento(v string) {
	o.CausalePagamento.Set(&v)
}
// SetCausalePagamentoNil sets the value for CausalePagamento to be an explicit nil
func (o *DatiRitenuta) SetCausalePagamentoNil() {
	o.CausalePagamento.Set(nil)
}

// UnsetCausalePagamento ensures that no value is present for CausalePagamento, not even an explicit nil
func (o *DatiRitenuta) UnsetCausalePagamento() {
	o.CausalePagamento.Unset()
}

func (o DatiRitenuta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatiRitenuta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TipoRitenuta.IsSet() {
		toSerialize["tipo_ritenuta"] = o.TipoRitenuta.Get()
	}
	if !IsNil(o.ImportoRitenuta) {
		toSerialize["importo_ritenuta"] = o.ImportoRitenuta
	}
	if !IsNil(o.AliquotaRitenuta) {
		toSerialize["aliquota_ritenuta"] = o.AliquotaRitenuta
	}
	if o.CausalePagamento.IsSet() {
		toSerialize["causale_pagamento"] = o.CausalePagamento.Get()
	}
	return toSerialize, nil
}

type NullableDatiRitenuta struct {
	value *DatiRitenuta
	isSet bool
}

func (v NullableDatiRitenuta) Get() *DatiRitenuta {
	return v.value
}

func (v *NullableDatiRitenuta) Set(val *DatiRitenuta) {
	v.value = val
	v.isSet = true
}

func (v NullableDatiRitenuta) IsSet() bool {
	return v.isSet
}

func (v *NullableDatiRitenuta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatiRitenuta(val *DatiRitenuta) *NullableDatiRitenuta {
	return &NullableDatiRitenuta{value: val, isSet: true}
}

func (v NullableDatiRitenuta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatiRitenuta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


