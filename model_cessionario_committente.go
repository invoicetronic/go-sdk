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

// checks if the CessionarioCommittente type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CessionarioCommittente{}

// CessionarioCommittente struct for CessionarioCommittente
type CessionarioCommittente struct {
	DatiAnagrafici *DatiAnagraficiCessionarioCommittente `json:"dati_anagrafici,omitempty"`
	Sede *SedeCessionarioCommittente `json:"sede,omitempty"`
	StabileOrganizzazione *StabileOrganizzazione `json:"stabile_organizzazione,omitempty"`
	RappresentanteFiscale *RappresentanteFiscaleCessionarioCommittente `json:"rappresentante_fiscale,omitempty"`
}

// NewCessionarioCommittente instantiates a new CessionarioCommittente object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCessionarioCommittente() *CessionarioCommittente {
	this := CessionarioCommittente{}
	return &this
}

// NewCessionarioCommittenteWithDefaults instantiates a new CessionarioCommittente object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCessionarioCommittenteWithDefaults() *CessionarioCommittente {
	this := CessionarioCommittente{}
	return &this
}

// GetDatiAnagrafici returns the DatiAnagrafici field value if set, zero value otherwise.
func (o *CessionarioCommittente) GetDatiAnagrafici() DatiAnagraficiCessionarioCommittente {
	if o == nil || IsNil(o.DatiAnagrafici) {
		var ret DatiAnagraficiCessionarioCommittente
		return ret
	}
	return *o.DatiAnagrafici
}

// GetDatiAnagraficiOk returns a tuple with the DatiAnagrafici field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CessionarioCommittente) GetDatiAnagraficiOk() (*DatiAnagraficiCessionarioCommittente, bool) {
	if o == nil || IsNil(o.DatiAnagrafici) {
		return nil, false
	}
	return o.DatiAnagrafici, true
}

// HasDatiAnagrafici returns a boolean if a field has been set.
func (o *CessionarioCommittente) HasDatiAnagrafici() bool {
	if o != nil && !IsNil(o.DatiAnagrafici) {
		return true
	}

	return false
}

// SetDatiAnagrafici gets a reference to the given DatiAnagraficiCessionarioCommittente and assigns it to the DatiAnagrafici field.
func (o *CessionarioCommittente) SetDatiAnagrafici(v DatiAnagraficiCessionarioCommittente) {
	o.DatiAnagrafici = &v
}

// GetSede returns the Sede field value if set, zero value otherwise.
func (o *CessionarioCommittente) GetSede() SedeCessionarioCommittente {
	if o == nil || IsNil(o.Sede) {
		var ret SedeCessionarioCommittente
		return ret
	}
	return *o.Sede
}

// GetSedeOk returns a tuple with the Sede field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CessionarioCommittente) GetSedeOk() (*SedeCessionarioCommittente, bool) {
	if o == nil || IsNil(o.Sede) {
		return nil, false
	}
	return o.Sede, true
}

// HasSede returns a boolean if a field has been set.
func (o *CessionarioCommittente) HasSede() bool {
	if o != nil && !IsNil(o.Sede) {
		return true
	}

	return false
}

// SetSede gets a reference to the given SedeCessionarioCommittente and assigns it to the Sede field.
func (o *CessionarioCommittente) SetSede(v SedeCessionarioCommittente) {
	o.Sede = &v
}

// GetStabileOrganizzazione returns the StabileOrganizzazione field value if set, zero value otherwise.
func (o *CessionarioCommittente) GetStabileOrganizzazione() StabileOrganizzazione {
	if o == nil || IsNil(o.StabileOrganizzazione) {
		var ret StabileOrganizzazione
		return ret
	}
	return *o.StabileOrganizzazione
}

// GetStabileOrganizzazioneOk returns a tuple with the StabileOrganizzazione field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CessionarioCommittente) GetStabileOrganizzazioneOk() (*StabileOrganizzazione, bool) {
	if o == nil || IsNil(o.StabileOrganizzazione) {
		return nil, false
	}
	return o.StabileOrganizzazione, true
}

// HasStabileOrganizzazione returns a boolean if a field has been set.
func (o *CessionarioCommittente) HasStabileOrganizzazione() bool {
	if o != nil && !IsNil(o.StabileOrganizzazione) {
		return true
	}

	return false
}

// SetStabileOrganizzazione gets a reference to the given StabileOrganizzazione and assigns it to the StabileOrganizzazione field.
func (o *CessionarioCommittente) SetStabileOrganizzazione(v StabileOrganizzazione) {
	o.StabileOrganizzazione = &v
}

// GetRappresentanteFiscale returns the RappresentanteFiscale field value if set, zero value otherwise.
func (o *CessionarioCommittente) GetRappresentanteFiscale() RappresentanteFiscaleCessionarioCommittente {
	if o == nil || IsNil(o.RappresentanteFiscale) {
		var ret RappresentanteFiscaleCessionarioCommittente
		return ret
	}
	return *o.RappresentanteFiscale
}

// GetRappresentanteFiscaleOk returns a tuple with the RappresentanteFiscale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CessionarioCommittente) GetRappresentanteFiscaleOk() (*RappresentanteFiscaleCessionarioCommittente, bool) {
	if o == nil || IsNil(o.RappresentanteFiscale) {
		return nil, false
	}
	return o.RappresentanteFiscale, true
}

// HasRappresentanteFiscale returns a boolean if a field has been set.
func (o *CessionarioCommittente) HasRappresentanteFiscale() bool {
	if o != nil && !IsNil(o.RappresentanteFiscale) {
		return true
	}

	return false
}

// SetRappresentanteFiscale gets a reference to the given RappresentanteFiscaleCessionarioCommittente and assigns it to the RappresentanteFiscale field.
func (o *CessionarioCommittente) SetRappresentanteFiscale(v RappresentanteFiscaleCessionarioCommittente) {
	o.RappresentanteFiscale = &v
}

func (o CessionarioCommittente) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CessionarioCommittente) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatiAnagrafici) {
		toSerialize["dati_anagrafici"] = o.DatiAnagrafici
	}
	if !IsNil(o.Sede) {
		toSerialize["sede"] = o.Sede
	}
	if !IsNil(o.StabileOrganizzazione) {
		toSerialize["stabile_organizzazione"] = o.StabileOrganizzazione
	}
	if !IsNil(o.RappresentanteFiscale) {
		toSerialize["rappresentante_fiscale"] = o.RappresentanteFiscale
	}
	return toSerialize, nil
}

type NullableCessionarioCommittente struct {
	value *CessionarioCommittente
	isSet bool
}

func (v NullableCessionarioCommittente) Get() *CessionarioCommittente {
	return v.value
}

func (v *NullableCessionarioCommittente) Set(val *CessionarioCommittente) {
	v.value = val
	v.isSet = true
}

func (v NullableCessionarioCommittente) IsSet() bool {
	return v.isSet
}

func (v *NullableCessionarioCommittente) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCessionarioCommittente(val *CessionarioCommittente) *NullableCessionarioCommittente {
	return &NullableCessionarioCommittente{value: val, isSet: true}
}

func (v NullableCessionarioCommittente) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCessionarioCommittente) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


