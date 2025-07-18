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

// checks if the ScontoMaggiorazione type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScontoMaggiorazione{}

// ScontoMaggiorazione struct for ScontoMaggiorazione
type ScontoMaggiorazione struct {
	Tipo NullableString `json:"tipo,omitempty"`
	Percentuale NullableFloat64 `json:"percentuale,omitempty"`
	Importo NullableFloat64 `json:"importo,omitempty"`
}

// NewScontoMaggiorazione instantiates a new ScontoMaggiorazione object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScontoMaggiorazione() *ScontoMaggiorazione {
	this := ScontoMaggiorazione{}
	return &this
}

// NewScontoMaggiorazioneWithDefaults instantiates a new ScontoMaggiorazione object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScontoMaggiorazioneWithDefaults() *ScontoMaggiorazione {
	this := ScontoMaggiorazione{}
	return &this
}

// GetTipo returns the Tipo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScontoMaggiorazione) GetTipo() string {
	if o == nil || IsNil(o.Tipo.Get()) {
		var ret string
		return ret
	}
	return *o.Tipo.Get()
}

// GetTipoOk returns a tuple with the Tipo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScontoMaggiorazione) GetTipoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tipo.Get(), o.Tipo.IsSet()
}

// HasTipo returns a boolean if a field has been set.
func (o *ScontoMaggiorazione) HasTipo() bool {
	if o != nil && o.Tipo.IsSet() {
		return true
	}

	return false
}

// SetTipo gets a reference to the given NullableString and assigns it to the Tipo field.
func (o *ScontoMaggiorazione) SetTipo(v string) {
	o.Tipo.Set(&v)
}
// SetTipoNil sets the value for Tipo to be an explicit nil
func (o *ScontoMaggiorazione) SetTipoNil() {
	o.Tipo.Set(nil)
}

// UnsetTipo ensures that no value is present for Tipo, not even an explicit nil
func (o *ScontoMaggiorazione) UnsetTipo() {
	o.Tipo.Unset()
}

// GetPercentuale returns the Percentuale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScontoMaggiorazione) GetPercentuale() float64 {
	if o == nil || IsNil(o.Percentuale.Get()) {
		var ret float64
		return ret
	}
	return *o.Percentuale.Get()
}

// GetPercentualeOk returns a tuple with the Percentuale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScontoMaggiorazione) GetPercentualeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Percentuale.Get(), o.Percentuale.IsSet()
}

// HasPercentuale returns a boolean if a field has been set.
func (o *ScontoMaggiorazione) HasPercentuale() bool {
	if o != nil && o.Percentuale.IsSet() {
		return true
	}

	return false
}

// SetPercentuale gets a reference to the given NullableFloat64 and assigns it to the Percentuale field.
func (o *ScontoMaggiorazione) SetPercentuale(v float64) {
	o.Percentuale.Set(&v)
}
// SetPercentualeNil sets the value for Percentuale to be an explicit nil
func (o *ScontoMaggiorazione) SetPercentualeNil() {
	o.Percentuale.Set(nil)
}

// UnsetPercentuale ensures that no value is present for Percentuale, not even an explicit nil
func (o *ScontoMaggiorazione) UnsetPercentuale() {
	o.Percentuale.Unset()
}

// GetImporto returns the Importo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScontoMaggiorazione) GetImporto() float64 {
	if o == nil || IsNil(o.Importo.Get()) {
		var ret float64
		return ret
	}
	return *o.Importo.Get()
}

// GetImportoOk returns a tuple with the Importo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScontoMaggiorazione) GetImportoOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Importo.Get(), o.Importo.IsSet()
}

// HasImporto returns a boolean if a field has been set.
func (o *ScontoMaggiorazione) HasImporto() bool {
	if o != nil && o.Importo.IsSet() {
		return true
	}

	return false
}

// SetImporto gets a reference to the given NullableFloat64 and assigns it to the Importo field.
func (o *ScontoMaggiorazione) SetImporto(v float64) {
	o.Importo.Set(&v)
}
// SetImportoNil sets the value for Importo to be an explicit nil
func (o *ScontoMaggiorazione) SetImportoNil() {
	o.Importo.Set(nil)
}

// UnsetImporto ensures that no value is present for Importo, not even an explicit nil
func (o *ScontoMaggiorazione) UnsetImporto() {
	o.Importo.Unset()
}

func (o ScontoMaggiorazione) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScontoMaggiorazione) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Tipo.IsSet() {
		toSerialize["tipo"] = o.Tipo.Get()
	}
	if o.Percentuale.IsSet() {
		toSerialize["percentuale"] = o.Percentuale.Get()
	}
	if o.Importo.IsSet() {
		toSerialize["importo"] = o.Importo.Get()
	}
	return toSerialize, nil
}

type NullableScontoMaggiorazione struct {
	value *ScontoMaggiorazione
	isSet bool
}

func (v NullableScontoMaggiorazione) Get() *ScontoMaggiorazione {
	return v.value
}

func (v *NullableScontoMaggiorazione) Set(val *ScontoMaggiorazione) {
	v.value = val
	v.isSet = true
}

func (v NullableScontoMaggiorazione) IsSet() bool {
	return v.isSet
}

func (v *NullableScontoMaggiorazione) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScontoMaggiorazione(val *ScontoMaggiorazione) *NullableScontoMaggiorazione {
	return &NullableScontoMaggiorazione{value: val, isSet: true}
}

func (v NullableScontoMaggiorazione) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScontoMaggiorazione) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


