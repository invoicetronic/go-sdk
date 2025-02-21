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

// checks if the SedeCedentePrestatore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SedeCedentePrestatore{}

// SedeCedentePrestatore struct for SedeCedentePrestatore
type SedeCedentePrestatore struct {
	Indirizzo NullableString `json:"indirizzo,omitempty"`
	NumeroCivico NullableString `json:"numero_civico,omitempty"`
	Cap NullableString `json:"cap,omitempty"`
	Comune NullableString `json:"comune,omitempty"`
	Provincia NullableString `json:"provincia,omitempty"`
	Nazione NullableString `json:"nazione,omitempty"`
}

// NewSedeCedentePrestatore instantiates a new SedeCedentePrestatore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSedeCedentePrestatore() *SedeCedentePrestatore {
	this := SedeCedentePrestatore{}
	var nazione string = "IT"
	this.Nazione = *NewNullableString(&nazione)
	return &this
}

// NewSedeCedentePrestatoreWithDefaults instantiates a new SedeCedentePrestatore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSedeCedentePrestatoreWithDefaults() *SedeCedentePrestatore {
	this := SedeCedentePrestatore{}
	var nazione string = "IT"
	this.Nazione = *NewNullableString(&nazione)
	return &this
}

// GetIndirizzo returns the Indirizzo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SedeCedentePrestatore) GetIndirizzo() string {
	if o == nil || IsNil(o.Indirizzo.Get()) {
		var ret string
		return ret
	}
	return *o.Indirizzo.Get()
}

// GetIndirizzoOk returns a tuple with the Indirizzo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SedeCedentePrestatore) GetIndirizzoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Indirizzo.Get(), o.Indirizzo.IsSet()
}

// HasIndirizzo returns a boolean if a field has been set.
func (o *SedeCedentePrestatore) HasIndirizzo() bool {
	if o != nil && o.Indirizzo.IsSet() {
		return true
	}

	return false
}

// SetIndirizzo gets a reference to the given NullableString and assigns it to the Indirizzo field.
func (o *SedeCedentePrestatore) SetIndirizzo(v string) {
	o.Indirizzo.Set(&v)
}
// SetIndirizzoNil sets the value for Indirizzo to be an explicit nil
func (o *SedeCedentePrestatore) SetIndirizzoNil() {
	o.Indirizzo.Set(nil)
}

// UnsetIndirizzo ensures that no value is present for Indirizzo, not even an explicit nil
func (o *SedeCedentePrestatore) UnsetIndirizzo() {
	o.Indirizzo.Unset()
}

// GetNumeroCivico returns the NumeroCivico field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SedeCedentePrestatore) GetNumeroCivico() string {
	if o == nil || IsNil(o.NumeroCivico.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroCivico.Get()
}

// GetNumeroCivicoOk returns a tuple with the NumeroCivico field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SedeCedentePrestatore) GetNumeroCivicoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroCivico.Get(), o.NumeroCivico.IsSet()
}

// HasNumeroCivico returns a boolean if a field has been set.
func (o *SedeCedentePrestatore) HasNumeroCivico() bool {
	if o != nil && o.NumeroCivico.IsSet() {
		return true
	}

	return false
}

// SetNumeroCivico gets a reference to the given NullableString and assigns it to the NumeroCivico field.
func (o *SedeCedentePrestatore) SetNumeroCivico(v string) {
	o.NumeroCivico.Set(&v)
}
// SetNumeroCivicoNil sets the value for NumeroCivico to be an explicit nil
func (o *SedeCedentePrestatore) SetNumeroCivicoNil() {
	o.NumeroCivico.Set(nil)
}

// UnsetNumeroCivico ensures that no value is present for NumeroCivico, not even an explicit nil
func (o *SedeCedentePrestatore) UnsetNumeroCivico() {
	o.NumeroCivico.Unset()
}

// GetCap returns the Cap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SedeCedentePrestatore) GetCap() string {
	if o == nil || IsNil(o.Cap.Get()) {
		var ret string
		return ret
	}
	return *o.Cap.Get()
}

// GetCapOk returns a tuple with the Cap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SedeCedentePrestatore) GetCapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cap.Get(), o.Cap.IsSet()
}

// HasCap returns a boolean if a field has been set.
func (o *SedeCedentePrestatore) HasCap() bool {
	if o != nil && o.Cap.IsSet() {
		return true
	}

	return false
}

// SetCap gets a reference to the given NullableString and assigns it to the Cap field.
func (o *SedeCedentePrestatore) SetCap(v string) {
	o.Cap.Set(&v)
}
// SetCapNil sets the value for Cap to be an explicit nil
func (o *SedeCedentePrestatore) SetCapNil() {
	o.Cap.Set(nil)
}

// UnsetCap ensures that no value is present for Cap, not even an explicit nil
func (o *SedeCedentePrestatore) UnsetCap() {
	o.Cap.Unset()
}

// GetComune returns the Comune field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SedeCedentePrestatore) GetComune() string {
	if o == nil || IsNil(o.Comune.Get()) {
		var ret string
		return ret
	}
	return *o.Comune.Get()
}

// GetComuneOk returns a tuple with the Comune field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SedeCedentePrestatore) GetComuneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comune.Get(), o.Comune.IsSet()
}

// HasComune returns a boolean if a field has been set.
func (o *SedeCedentePrestatore) HasComune() bool {
	if o != nil && o.Comune.IsSet() {
		return true
	}

	return false
}

// SetComune gets a reference to the given NullableString and assigns it to the Comune field.
func (o *SedeCedentePrestatore) SetComune(v string) {
	o.Comune.Set(&v)
}
// SetComuneNil sets the value for Comune to be an explicit nil
func (o *SedeCedentePrestatore) SetComuneNil() {
	o.Comune.Set(nil)
}

// UnsetComune ensures that no value is present for Comune, not even an explicit nil
func (o *SedeCedentePrestatore) UnsetComune() {
	o.Comune.Unset()
}

// GetProvincia returns the Provincia field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SedeCedentePrestatore) GetProvincia() string {
	if o == nil || IsNil(o.Provincia.Get()) {
		var ret string
		return ret
	}
	return *o.Provincia.Get()
}

// GetProvinciaOk returns a tuple with the Provincia field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SedeCedentePrestatore) GetProvinciaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provincia.Get(), o.Provincia.IsSet()
}

// HasProvincia returns a boolean if a field has been set.
func (o *SedeCedentePrestatore) HasProvincia() bool {
	if o != nil && o.Provincia.IsSet() {
		return true
	}

	return false
}

// SetProvincia gets a reference to the given NullableString and assigns it to the Provincia field.
func (o *SedeCedentePrestatore) SetProvincia(v string) {
	o.Provincia.Set(&v)
}
// SetProvinciaNil sets the value for Provincia to be an explicit nil
func (o *SedeCedentePrestatore) SetProvinciaNil() {
	o.Provincia.Set(nil)
}

// UnsetProvincia ensures that no value is present for Provincia, not even an explicit nil
func (o *SedeCedentePrestatore) UnsetProvincia() {
	o.Provincia.Unset()
}

// GetNazione returns the Nazione field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SedeCedentePrestatore) GetNazione() string {
	if o == nil || IsNil(o.Nazione.Get()) {
		var ret string
		return ret
	}
	return *o.Nazione.Get()
}

// GetNazioneOk returns a tuple with the Nazione field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SedeCedentePrestatore) GetNazioneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nazione.Get(), o.Nazione.IsSet()
}

// HasNazione returns a boolean if a field has been set.
func (o *SedeCedentePrestatore) HasNazione() bool {
	if o != nil && o.Nazione.IsSet() {
		return true
	}

	return false
}

// SetNazione gets a reference to the given NullableString and assigns it to the Nazione field.
func (o *SedeCedentePrestatore) SetNazione(v string) {
	o.Nazione.Set(&v)
}
// SetNazioneNil sets the value for Nazione to be an explicit nil
func (o *SedeCedentePrestatore) SetNazioneNil() {
	o.Nazione.Set(nil)
}

// UnsetNazione ensures that no value is present for Nazione, not even an explicit nil
func (o *SedeCedentePrestatore) UnsetNazione() {
	o.Nazione.Unset()
}

func (o SedeCedentePrestatore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SedeCedentePrestatore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Indirizzo.IsSet() {
		toSerialize["indirizzo"] = o.Indirizzo.Get()
	}
	if o.NumeroCivico.IsSet() {
		toSerialize["numero_civico"] = o.NumeroCivico.Get()
	}
	if o.Cap.IsSet() {
		toSerialize["cap"] = o.Cap.Get()
	}
	if o.Comune.IsSet() {
		toSerialize["comune"] = o.Comune.Get()
	}
	if o.Provincia.IsSet() {
		toSerialize["provincia"] = o.Provincia.Get()
	}
	if o.Nazione.IsSet() {
		toSerialize["nazione"] = o.Nazione.Get()
	}
	return toSerialize, nil
}

type NullableSedeCedentePrestatore struct {
	value *SedeCedentePrestatore
	isSet bool
}

func (v NullableSedeCedentePrestatore) Get() *SedeCedentePrestatore {
	return v.value
}

func (v *NullableSedeCedentePrestatore) Set(val *SedeCedentePrestatore) {
	v.value = val
	v.isSet = true
}

func (v NullableSedeCedentePrestatore) IsSet() bool {
	return v.isSet
}

func (v *NullableSedeCedentePrestatore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSedeCedentePrestatore(val *SedeCedentePrestatore) *NullableSedeCedentePrestatore {
	return &NullableSedeCedentePrestatore{value: val, isSet: true}
}

func (v NullableSedeCedentePrestatore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSedeCedentePrestatore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


