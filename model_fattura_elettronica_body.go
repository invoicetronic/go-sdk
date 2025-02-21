/*
Italian eInvoice API v1

The [Italian eInvoice API][2] is a RESTful API that allows you to send and receive invoices through the Italian [Servizio di Interscambio (SDI)][1], or Interchange Service. The API is designed by Invoicetronic to be simple and easy to use, abstracting away SDI complexity while providing complete control over the invoice send/receive process. The API also provides advanced features as encryption at rest, invoice validation, multiple upload formats, webhooks, event logging, client SDKs for commonly used languages, and CLI tools.  For more information, see  [Invoicetronic website][2]  [1]: https://www.fatturapa.gov.it/it/sistemainterscambio/cose-il-sdi/ [2]: https://invoicetronic.com/

API version: 1
Contact: support@invoicetronic.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package invoicesdk

import (
	"encoding/json"
)

// checks if the FatturaElettronicaBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FatturaElettronicaBody{}

// FatturaElettronicaBody struct for FatturaElettronicaBody
type FatturaElettronicaBody struct {
	DatiGenerali *DatiGenerali `json:"dati_generali,omitempty"`
	DatiBeniServizi *DatiBeniServizi `json:"dati_beni_servizi,omitempty"`
	DatiVeicoli *DatiVeicoli `json:"dati_veicoli,omitempty"`
	DatiPagamento []DatiPagamento `json:"dati_pagamento,omitempty"`
	Allegati []Allegati `json:"allegati,omitempty"`
}

// NewFatturaElettronicaBody instantiates a new FatturaElettronicaBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFatturaElettronicaBody() *FatturaElettronicaBody {
	this := FatturaElettronicaBody{}
	return &this
}

// NewFatturaElettronicaBodyWithDefaults instantiates a new FatturaElettronicaBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFatturaElettronicaBodyWithDefaults() *FatturaElettronicaBody {
	this := FatturaElettronicaBody{}
	return &this
}

// GetDatiGenerali returns the DatiGenerali field value if set, zero value otherwise.
func (o *FatturaElettronicaBody) GetDatiGenerali() DatiGenerali {
	if o == nil || IsNil(o.DatiGenerali) {
		var ret DatiGenerali
		return ret
	}
	return *o.DatiGenerali
}

// GetDatiGeneraliOk returns a tuple with the DatiGenerali field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FatturaElettronicaBody) GetDatiGeneraliOk() (*DatiGenerali, bool) {
	if o == nil || IsNil(o.DatiGenerali) {
		return nil, false
	}
	return o.DatiGenerali, true
}

// HasDatiGenerali returns a boolean if a field has been set.
func (o *FatturaElettronicaBody) HasDatiGenerali() bool {
	if o != nil && !IsNil(o.DatiGenerali) {
		return true
	}

	return false
}

// SetDatiGenerali gets a reference to the given DatiGenerali and assigns it to the DatiGenerali field.
func (o *FatturaElettronicaBody) SetDatiGenerali(v DatiGenerali) {
	o.DatiGenerali = &v
}

// GetDatiBeniServizi returns the DatiBeniServizi field value if set, zero value otherwise.
func (o *FatturaElettronicaBody) GetDatiBeniServizi() DatiBeniServizi {
	if o == nil || IsNil(o.DatiBeniServizi) {
		var ret DatiBeniServizi
		return ret
	}
	return *o.DatiBeniServizi
}

// GetDatiBeniServiziOk returns a tuple with the DatiBeniServizi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FatturaElettronicaBody) GetDatiBeniServiziOk() (*DatiBeniServizi, bool) {
	if o == nil || IsNil(o.DatiBeniServizi) {
		return nil, false
	}
	return o.DatiBeniServizi, true
}

// HasDatiBeniServizi returns a boolean if a field has been set.
func (o *FatturaElettronicaBody) HasDatiBeniServizi() bool {
	if o != nil && !IsNil(o.DatiBeniServizi) {
		return true
	}

	return false
}

// SetDatiBeniServizi gets a reference to the given DatiBeniServizi and assigns it to the DatiBeniServizi field.
func (o *FatturaElettronicaBody) SetDatiBeniServizi(v DatiBeniServizi) {
	o.DatiBeniServizi = &v
}

// GetDatiVeicoli returns the DatiVeicoli field value if set, zero value otherwise.
func (o *FatturaElettronicaBody) GetDatiVeicoli() DatiVeicoli {
	if o == nil || IsNil(o.DatiVeicoli) {
		var ret DatiVeicoli
		return ret
	}
	return *o.DatiVeicoli
}

// GetDatiVeicoliOk returns a tuple with the DatiVeicoli field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FatturaElettronicaBody) GetDatiVeicoliOk() (*DatiVeicoli, bool) {
	if o == nil || IsNil(o.DatiVeicoli) {
		return nil, false
	}
	return o.DatiVeicoli, true
}

// HasDatiVeicoli returns a boolean if a field has been set.
func (o *FatturaElettronicaBody) HasDatiVeicoli() bool {
	if o != nil && !IsNil(o.DatiVeicoli) {
		return true
	}

	return false
}

// SetDatiVeicoli gets a reference to the given DatiVeicoli and assigns it to the DatiVeicoli field.
func (o *FatturaElettronicaBody) SetDatiVeicoli(v DatiVeicoli) {
	o.DatiVeicoli = &v
}

// GetDatiPagamento returns the DatiPagamento field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FatturaElettronicaBody) GetDatiPagamento() []DatiPagamento {
	if o == nil {
		var ret []DatiPagamento
		return ret
	}
	return o.DatiPagamento
}

// GetDatiPagamentoOk returns a tuple with the DatiPagamento field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FatturaElettronicaBody) GetDatiPagamentoOk() ([]DatiPagamento, bool) {
	if o == nil || IsNil(o.DatiPagamento) {
		return nil, false
	}
	return o.DatiPagamento, true
}

// HasDatiPagamento returns a boolean if a field has been set.
func (o *FatturaElettronicaBody) HasDatiPagamento() bool {
	if o != nil && !IsNil(o.DatiPagamento) {
		return true
	}

	return false
}

// SetDatiPagamento gets a reference to the given []DatiPagamento and assigns it to the DatiPagamento field.
func (o *FatturaElettronicaBody) SetDatiPagamento(v []DatiPagamento) {
	o.DatiPagamento = v
}

// GetAllegati returns the Allegati field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FatturaElettronicaBody) GetAllegati() []Allegati {
	if o == nil {
		var ret []Allegati
		return ret
	}
	return o.Allegati
}

// GetAllegatiOk returns a tuple with the Allegati field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FatturaElettronicaBody) GetAllegatiOk() ([]Allegati, bool) {
	if o == nil || IsNil(o.Allegati) {
		return nil, false
	}
	return o.Allegati, true
}

// HasAllegati returns a boolean if a field has been set.
func (o *FatturaElettronicaBody) HasAllegati() bool {
	if o != nil && !IsNil(o.Allegati) {
		return true
	}

	return false
}

// SetAllegati gets a reference to the given []Allegati and assigns it to the Allegati field.
func (o *FatturaElettronicaBody) SetAllegati(v []Allegati) {
	o.Allegati = v
}

func (o FatturaElettronicaBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FatturaElettronicaBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatiGenerali) {
		toSerialize["dati_generali"] = o.DatiGenerali
	}
	if !IsNil(o.DatiBeniServizi) {
		toSerialize["dati_beni_servizi"] = o.DatiBeniServizi
	}
	if !IsNil(o.DatiVeicoli) {
		toSerialize["dati_veicoli"] = o.DatiVeicoli
	}
	if o.DatiPagamento != nil {
		toSerialize["dati_pagamento"] = o.DatiPagamento
	}
	if o.Allegati != nil {
		toSerialize["allegati"] = o.Allegati
	}
	return toSerialize, nil
}

type NullableFatturaElettronicaBody struct {
	value *FatturaElettronicaBody
	isSet bool
}

func (v NullableFatturaElettronicaBody) Get() *FatturaElettronicaBody {
	return v.value
}

func (v *NullableFatturaElettronicaBody) Set(val *FatturaElettronicaBody) {
	v.value = val
	v.isSet = true
}

func (v NullableFatturaElettronicaBody) IsSet() bool {
	return v.isSet
}

func (v *NullableFatturaElettronicaBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFatturaElettronicaBody(val *FatturaElettronicaBody) *NullableFatturaElettronicaBody {
	return &NullableFatturaElettronicaBody{value: val, isSet: true}
}

func (v NullableFatturaElettronicaBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFatturaElettronicaBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


