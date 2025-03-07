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
)

// checks if the DatiAnagrafici type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatiAnagrafici{}

// DatiAnagrafici struct for DatiAnagrafici
type DatiAnagrafici struct {
	IdFiscaleIva *IdFiscaleIVA `json:"id_fiscale_iva,omitempty"`
	CodiceFiscale NullableString `json:"codice_fiscale,omitempty"`
	Anagrafica *Anagrafica `json:"anagrafica,omitempty"`
}

// NewDatiAnagrafici instantiates a new DatiAnagrafici object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatiAnagrafici() *DatiAnagrafici {
	this := DatiAnagrafici{}
	return &this
}

// NewDatiAnagraficiWithDefaults instantiates a new DatiAnagrafici object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatiAnagraficiWithDefaults() *DatiAnagrafici {
	this := DatiAnagrafici{}
	return &this
}

// GetIdFiscaleIva returns the IdFiscaleIva field value if set, zero value otherwise.
func (o *DatiAnagrafici) GetIdFiscaleIva() IdFiscaleIVA {
	if o == nil || IsNil(o.IdFiscaleIva) {
		var ret IdFiscaleIVA
		return ret
	}
	return *o.IdFiscaleIva
}

// GetIdFiscaleIvaOk returns a tuple with the IdFiscaleIva field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatiAnagrafici) GetIdFiscaleIvaOk() (*IdFiscaleIVA, bool) {
	if o == nil || IsNil(o.IdFiscaleIva) {
		return nil, false
	}
	return o.IdFiscaleIva, true
}

// HasIdFiscaleIva returns a boolean if a field has been set.
func (o *DatiAnagrafici) HasIdFiscaleIva() bool {
	if o != nil && !IsNil(o.IdFiscaleIva) {
		return true
	}

	return false
}

// SetIdFiscaleIva gets a reference to the given IdFiscaleIVA and assigns it to the IdFiscaleIva field.
func (o *DatiAnagrafici) SetIdFiscaleIva(v IdFiscaleIVA) {
	o.IdFiscaleIva = &v
}

// GetCodiceFiscale returns the CodiceFiscale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiAnagrafici) GetCodiceFiscale() string {
	if o == nil || IsNil(o.CodiceFiscale.Get()) {
		var ret string
		return ret
	}
	return *o.CodiceFiscale.Get()
}

// GetCodiceFiscaleOk returns a tuple with the CodiceFiscale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiAnagrafici) GetCodiceFiscaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodiceFiscale.Get(), o.CodiceFiscale.IsSet()
}

// HasCodiceFiscale returns a boolean if a field has been set.
func (o *DatiAnagrafici) HasCodiceFiscale() bool {
	if o != nil && o.CodiceFiscale.IsSet() {
		return true
	}

	return false
}

// SetCodiceFiscale gets a reference to the given NullableString and assigns it to the CodiceFiscale field.
func (o *DatiAnagrafici) SetCodiceFiscale(v string) {
	o.CodiceFiscale.Set(&v)
}
// SetCodiceFiscaleNil sets the value for CodiceFiscale to be an explicit nil
func (o *DatiAnagrafici) SetCodiceFiscaleNil() {
	o.CodiceFiscale.Set(nil)
}

// UnsetCodiceFiscale ensures that no value is present for CodiceFiscale, not even an explicit nil
func (o *DatiAnagrafici) UnsetCodiceFiscale() {
	o.CodiceFiscale.Unset()
}

// GetAnagrafica returns the Anagrafica field value if set, zero value otherwise.
func (o *DatiAnagrafici) GetAnagrafica() Anagrafica {
	if o == nil || IsNil(o.Anagrafica) {
		var ret Anagrafica
		return ret
	}
	return *o.Anagrafica
}

// GetAnagraficaOk returns a tuple with the Anagrafica field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatiAnagrafici) GetAnagraficaOk() (*Anagrafica, bool) {
	if o == nil || IsNil(o.Anagrafica) {
		return nil, false
	}
	return o.Anagrafica, true
}

// HasAnagrafica returns a boolean if a field has been set.
func (o *DatiAnagrafici) HasAnagrafica() bool {
	if o != nil && !IsNil(o.Anagrafica) {
		return true
	}

	return false
}

// SetAnagrafica gets a reference to the given Anagrafica and assigns it to the Anagrafica field.
func (o *DatiAnagrafici) SetAnagrafica(v Anagrafica) {
	o.Anagrafica = &v
}

func (o DatiAnagrafici) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatiAnagrafici) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdFiscaleIva) {
		toSerialize["id_fiscale_iva"] = o.IdFiscaleIva
	}
	if o.CodiceFiscale.IsSet() {
		toSerialize["codice_fiscale"] = o.CodiceFiscale.Get()
	}
	if !IsNil(o.Anagrafica) {
		toSerialize["anagrafica"] = o.Anagrafica
	}
	return toSerialize, nil
}

type NullableDatiAnagrafici struct {
	value *DatiAnagrafici
	isSet bool
}

func (v NullableDatiAnagrafici) Get() *DatiAnagrafici {
	return v.value
}

func (v *NullableDatiAnagrafici) Set(val *DatiAnagrafici) {
	v.value = val
	v.isSet = true
}

func (v NullableDatiAnagrafici) IsSet() bool {
	return v.isSet
}

func (v *NullableDatiAnagrafici) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatiAnagrafici(val *DatiAnagrafici) *NullableDatiAnagrafici {
	return &NullableDatiAnagrafici{value: val, isSet: true}
}

func (v NullableDatiAnagrafici) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatiAnagrafici) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


