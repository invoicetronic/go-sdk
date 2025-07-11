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

// checks if the IscrizioneREA type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IscrizioneREA{}

// IscrizioneREA struct for IscrizioneREA
type IscrizioneREA struct {
	Ufficio NullableString `json:"ufficio,omitempty"`
	NumeroRea NullableString `json:"numero_rea,omitempty"`
	CapitaleSociale NullableFloat64 `json:"capitale_sociale,omitempty"`
	SocioUnico NullableString `json:"socio_unico,omitempty"`
	StatoLiquidazione NullableString `json:"stato_liquidazione,omitempty"`
}

// NewIscrizioneREA instantiates a new IscrizioneREA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIscrizioneREA() *IscrizioneREA {
	this := IscrizioneREA{}
	return &this
}

// NewIscrizioneREAWithDefaults instantiates a new IscrizioneREA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIscrizioneREAWithDefaults() *IscrizioneREA {
	this := IscrizioneREA{}
	return &this
}

// GetUfficio returns the Ufficio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IscrizioneREA) GetUfficio() string {
	if o == nil || IsNil(o.Ufficio.Get()) {
		var ret string
		return ret
	}
	return *o.Ufficio.Get()
}

// GetUfficioOk returns a tuple with the Ufficio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IscrizioneREA) GetUfficioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ufficio.Get(), o.Ufficio.IsSet()
}

// HasUfficio returns a boolean if a field has been set.
func (o *IscrizioneREA) HasUfficio() bool {
	if o != nil && o.Ufficio.IsSet() {
		return true
	}

	return false
}

// SetUfficio gets a reference to the given NullableString and assigns it to the Ufficio field.
func (o *IscrizioneREA) SetUfficio(v string) {
	o.Ufficio.Set(&v)
}
// SetUfficioNil sets the value for Ufficio to be an explicit nil
func (o *IscrizioneREA) SetUfficioNil() {
	o.Ufficio.Set(nil)
}

// UnsetUfficio ensures that no value is present for Ufficio, not even an explicit nil
func (o *IscrizioneREA) UnsetUfficio() {
	o.Ufficio.Unset()
}

// GetNumeroRea returns the NumeroRea field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IscrizioneREA) GetNumeroRea() string {
	if o == nil || IsNil(o.NumeroRea.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroRea.Get()
}

// GetNumeroReaOk returns a tuple with the NumeroRea field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IscrizioneREA) GetNumeroReaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroRea.Get(), o.NumeroRea.IsSet()
}

// HasNumeroRea returns a boolean if a field has been set.
func (o *IscrizioneREA) HasNumeroRea() bool {
	if o != nil && o.NumeroRea.IsSet() {
		return true
	}

	return false
}

// SetNumeroRea gets a reference to the given NullableString and assigns it to the NumeroRea field.
func (o *IscrizioneREA) SetNumeroRea(v string) {
	o.NumeroRea.Set(&v)
}
// SetNumeroReaNil sets the value for NumeroRea to be an explicit nil
func (o *IscrizioneREA) SetNumeroReaNil() {
	o.NumeroRea.Set(nil)
}

// UnsetNumeroRea ensures that no value is present for NumeroRea, not even an explicit nil
func (o *IscrizioneREA) UnsetNumeroRea() {
	o.NumeroRea.Unset()
}

// GetCapitaleSociale returns the CapitaleSociale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IscrizioneREA) GetCapitaleSociale() float64 {
	if o == nil || IsNil(o.CapitaleSociale.Get()) {
		var ret float64
		return ret
	}
	return *o.CapitaleSociale.Get()
}

// GetCapitaleSocialeOk returns a tuple with the CapitaleSociale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IscrizioneREA) GetCapitaleSocialeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CapitaleSociale.Get(), o.CapitaleSociale.IsSet()
}

// HasCapitaleSociale returns a boolean if a field has been set.
func (o *IscrizioneREA) HasCapitaleSociale() bool {
	if o != nil && o.CapitaleSociale.IsSet() {
		return true
	}

	return false
}

// SetCapitaleSociale gets a reference to the given NullableFloat64 and assigns it to the CapitaleSociale field.
func (o *IscrizioneREA) SetCapitaleSociale(v float64) {
	o.CapitaleSociale.Set(&v)
}
// SetCapitaleSocialeNil sets the value for CapitaleSociale to be an explicit nil
func (o *IscrizioneREA) SetCapitaleSocialeNil() {
	o.CapitaleSociale.Set(nil)
}

// UnsetCapitaleSociale ensures that no value is present for CapitaleSociale, not even an explicit nil
func (o *IscrizioneREA) UnsetCapitaleSociale() {
	o.CapitaleSociale.Unset()
}

// GetSocioUnico returns the SocioUnico field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IscrizioneREA) GetSocioUnico() string {
	if o == nil || IsNil(o.SocioUnico.Get()) {
		var ret string
		return ret
	}
	return *o.SocioUnico.Get()
}

// GetSocioUnicoOk returns a tuple with the SocioUnico field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IscrizioneREA) GetSocioUnicoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SocioUnico.Get(), o.SocioUnico.IsSet()
}

// HasSocioUnico returns a boolean if a field has been set.
func (o *IscrizioneREA) HasSocioUnico() bool {
	if o != nil && o.SocioUnico.IsSet() {
		return true
	}

	return false
}

// SetSocioUnico gets a reference to the given NullableString and assigns it to the SocioUnico field.
func (o *IscrizioneREA) SetSocioUnico(v string) {
	o.SocioUnico.Set(&v)
}
// SetSocioUnicoNil sets the value for SocioUnico to be an explicit nil
func (o *IscrizioneREA) SetSocioUnicoNil() {
	o.SocioUnico.Set(nil)
}

// UnsetSocioUnico ensures that no value is present for SocioUnico, not even an explicit nil
func (o *IscrizioneREA) UnsetSocioUnico() {
	o.SocioUnico.Unset()
}

// GetStatoLiquidazione returns the StatoLiquidazione field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IscrizioneREA) GetStatoLiquidazione() string {
	if o == nil || IsNil(o.StatoLiquidazione.Get()) {
		var ret string
		return ret
	}
	return *o.StatoLiquidazione.Get()
}

// GetStatoLiquidazioneOk returns a tuple with the StatoLiquidazione field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IscrizioneREA) GetStatoLiquidazioneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatoLiquidazione.Get(), o.StatoLiquidazione.IsSet()
}

// HasStatoLiquidazione returns a boolean if a field has been set.
func (o *IscrizioneREA) HasStatoLiquidazione() bool {
	if o != nil && o.StatoLiquidazione.IsSet() {
		return true
	}

	return false
}

// SetStatoLiquidazione gets a reference to the given NullableString and assigns it to the StatoLiquidazione field.
func (o *IscrizioneREA) SetStatoLiquidazione(v string) {
	o.StatoLiquidazione.Set(&v)
}
// SetStatoLiquidazioneNil sets the value for StatoLiquidazione to be an explicit nil
func (o *IscrizioneREA) SetStatoLiquidazioneNil() {
	o.StatoLiquidazione.Set(nil)
}

// UnsetStatoLiquidazione ensures that no value is present for StatoLiquidazione, not even an explicit nil
func (o *IscrizioneREA) UnsetStatoLiquidazione() {
	o.StatoLiquidazione.Unset()
}

func (o IscrizioneREA) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IscrizioneREA) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ufficio.IsSet() {
		toSerialize["ufficio"] = o.Ufficio.Get()
	}
	if o.NumeroRea.IsSet() {
		toSerialize["numero_rea"] = o.NumeroRea.Get()
	}
	if o.CapitaleSociale.IsSet() {
		toSerialize["capitale_sociale"] = o.CapitaleSociale.Get()
	}
	if o.SocioUnico.IsSet() {
		toSerialize["socio_unico"] = o.SocioUnico.Get()
	}
	if o.StatoLiquidazione.IsSet() {
		toSerialize["stato_liquidazione"] = o.StatoLiquidazione.Get()
	}
	return toSerialize, nil
}

type NullableIscrizioneREA struct {
	value *IscrizioneREA
	isSet bool
}

func (v NullableIscrizioneREA) Get() *IscrizioneREA {
	return v.value
}

func (v *NullableIscrizioneREA) Set(val *IscrizioneREA) {
	v.value = val
	v.isSet = true
}

func (v NullableIscrizioneREA) IsSet() bool {
	return v.isSet
}

func (v *NullableIscrizioneREA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIscrizioneREA(val *IscrizioneREA) *NullableIscrizioneREA {
	return &NullableIscrizioneREA{value: val, isSet: true}
}

func (v NullableIscrizioneREA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIscrizioneREA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


