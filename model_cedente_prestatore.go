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

// checks if the CedentePrestatore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CedentePrestatore{}

// CedentePrestatore struct for CedentePrestatore
type CedentePrestatore struct {
	DatiAnagrafici *DatiAnagraficiCedentePrestatore `json:"dati_anagrafici,omitempty"`
	Sede *SedeCedentePrestatore `json:"sede,omitempty"`
	StabileOrganizzazione *StabileOrganizzazione `json:"stabile_organizzazione,omitempty"`
	IscrizioneRea *IscrizioneREA `json:"iscrizione_rea,omitempty"`
	Contatti *Contatti `json:"contatti,omitempty"`
	RiferimentoAmministrazione NullableString `json:"riferimento_amministrazione,omitempty"`
}

// NewCedentePrestatore instantiates a new CedentePrestatore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCedentePrestatore() *CedentePrestatore {
	this := CedentePrestatore{}
	return &this
}

// NewCedentePrestatoreWithDefaults instantiates a new CedentePrestatore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCedentePrestatoreWithDefaults() *CedentePrestatore {
	this := CedentePrestatore{}
	return &this
}

// GetDatiAnagrafici returns the DatiAnagrafici field value if set, zero value otherwise.
func (o *CedentePrestatore) GetDatiAnagrafici() DatiAnagraficiCedentePrestatore {
	if o == nil || IsNil(o.DatiAnagrafici) {
		var ret DatiAnagraficiCedentePrestatore
		return ret
	}
	return *o.DatiAnagrafici
}

// GetDatiAnagraficiOk returns a tuple with the DatiAnagrafici field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CedentePrestatore) GetDatiAnagraficiOk() (*DatiAnagraficiCedentePrestatore, bool) {
	if o == nil || IsNil(o.DatiAnagrafici) {
		return nil, false
	}
	return o.DatiAnagrafici, true
}

// HasDatiAnagrafici returns a boolean if a field has been set.
func (o *CedentePrestatore) HasDatiAnagrafici() bool {
	if o != nil && !IsNil(o.DatiAnagrafici) {
		return true
	}

	return false
}

// SetDatiAnagrafici gets a reference to the given DatiAnagraficiCedentePrestatore and assigns it to the DatiAnagrafici field.
func (o *CedentePrestatore) SetDatiAnagrafici(v DatiAnagraficiCedentePrestatore) {
	o.DatiAnagrafici = &v
}

// GetSede returns the Sede field value if set, zero value otherwise.
func (o *CedentePrestatore) GetSede() SedeCedentePrestatore {
	if o == nil || IsNil(o.Sede) {
		var ret SedeCedentePrestatore
		return ret
	}
	return *o.Sede
}

// GetSedeOk returns a tuple with the Sede field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CedentePrestatore) GetSedeOk() (*SedeCedentePrestatore, bool) {
	if o == nil || IsNil(o.Sede) {
		return nil, false
	}
	return o.Sede, true
}

// HasSede returns a boolean if a field has been set.
func (o *CedentePrestatore) HasSede() bool {
	if o != nil && !IsNil(o.Sede) {
		return true
	}

	return false
}

// SetSede gets a reference to the given SedeCedentePrestatore and assigns it to the Sede field.
func (o *CedentePrestatore) SetSede(v SedeCedentePrestatore) {
	o.Sede = &v
}

// GetStabileOrganizzazione returns the StabileOrganizzazione field value if set, zero value otherwise.
func (o *CedentePrestatore) GetStabileOrganizzazione() StabileOrganizzazione {
	if o == nil || IsNil(o.StabileOrganizzazione) {
		var ret StabileOrganizzazione
		return ret
	}
	return *o.StabileOrganizzazione
}

// GetStabileOrganizzazioneOk returns a tuple with the StabileOrganizzazione field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CedentePrestatore) GetStabileOrganizzazioneOk() (*StabileOrganizzazione, bool) {
	if o == nil || IsNil(o.StabileOrganizzazione) {
		return nil, false
	}
	return o.StabileOrganizzazione, true
}

// HasStabileOrganizzazione returns a boolean if a field has been set.
func (o *CedentePrestatore) HasStabileOrganizzazione() bool {
	if o != nil && !IsNil(o.StabileOrganizzazione) {
		return true
	}

	return false
}

// SetStabileOrganizzazione gets a reference to the given StabileOrganizzazione and assigns it to the StabileOrganizzazione field.
func (o *CedentePrestatore) SetStabileOrganizzazione(v StabileOrganizzazione) {
	o.StabileOrganizzazione = &v
}

// GetIscrizioneRea returns the IscrizioneRea field value if set, zero value otherwise.
func (o *CedentePrestatore) GetIscrizioneRea() IscrizioneREA {
	if o == nil || IsNil(o.IscrizioneRea) {
		var ret IscrizioneREA
		return ret
	}
	return *o.IscrizioneRea
}

// GetIscrizioneReaOk returns a tuple with the IscrizioneRea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CedentePrestatore) GetIscrizioneReaOk() (*IscrizioneREA, bool) {
	if o == nil || IsNil(o.IscrizioneRea) {
		return nil, false
	}
	return o.IscrizioneRea, true
}

// HasIscrizioneRea returns a boolean if a field has been set.
func (o *CedentePrestatore) HasIscrizioneRea() bool {
	if o != nil && !IsNil(o.IscrizioneRea) {
		return true
	}

	return false
}

// SetIscrizioneRea gets a reference to the given IscrizioneREA and assigns it to the IscrizioneRea field.
func (o *CedentePrestatore) SetIscrizioneRea(v IscrizioneREA) {
	o.IscrizioneRea = &v
}

// GetContatti returns the Contatti field value if set, zero value otherwise.
func (o *CedentePrestatore) GetContatti() Contatti {
	if o == nil || IsNil(o.Contatti) {
		var ret Contatti
		return ret
	}
	return *o.Contatti
}

// GetContattiOk returns a tuple with the Contatti field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CedentePrestatore) GetContattiOk() (*Contatti, bool) {
	if o == nil || IsNil(o.Contatti) {
		return nil, false
	}
	return o.Contatti, true
}

// HasContatti returns a boolean if a field has been set.
func (o *CedentePrestatore) HasContatti() bool {
	if o != nil && !IsNil(o.Contatti) {
		return true
	}

	return false
}

// SetContatti gets a reference to the given Contatti and assigns it to the Contatti field.
func (o *CedentePrestatore) SetContatti(v Contatti) {
	o.Contatti = &v
}

// GetRiferimentoAmministrazione returns the RiferimentoAmministrazione field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CedentePrestatore) GetRiferimentoAmministrazione() string {
	if o == nil || IsNil(o.RiferimentoAmministrazione.Get()) {
		var ret string
		return ret
	}
	return *o.RiferimentoAmministrazione.Get()
}

// GetRiferimentoAmministrazioneOk returns a tuple with the RiferimentoAmministrazione field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CedentePrestatore) GetRiferimentoAmministrazioneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RiferimentoAmministrazione.Get(), o.RiferimentoAmministrazione.IsSet()
}

// HasRiferimentoAmministrazione returns a boolean if a field has been set.
func (o *CedentePrestatore) HasRiferimentoAmministrazione() bool {
	if o != nil && o.RiferimentoAmministrazione.IsSet() {
		return true
	}

	return false
}

// SetRiferimentoAmministrazione gets a reference to the given NullableString and assigns it to the RiferimentoAmministrazione field.
func (o *CedentePrestatore) SetRiferimentoAmministrazione(v string) {
	o.RiferimentoAmministrazione.Set(&v)
}
// SetRiferimentoAmministrazioneNil sets the value for RiferimentoAmministrazione to be an explicit nil
func (o *CedentePrestatore) SetRiferimentoAmministrazioneNil() {
	o.RiferimentoAmministrazione.Set(nil)
}

// UnsetRiferimentoAmministrazione ensures that no value is present for RiferimentoAmministrazione, not even an explicit nil
func (o *CedentePrestatore) UnsetRiferimentoAmministrazione() {
	o.RiferimentoAmministrazione.Unset()
}

func (o CedentePrestatore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CedentePrestatore) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IscrizioneRea) {
		toSerialize["iscrizione_rea"] = o.IscrizioneRea
	}
	if !IsNil(o.Contatti) {
		toSerialize["contatti"] = o.Contatti
	}
	if o.RiferimentoAmministrazione.IsSet() {
		toSerialize["riferimento_amministrazione"] = o.RiferimentoAmministrazione.Get()
	}
	return toSerialize, nil
}

type NullableCedentePrestatore struct {
	value *CedentePrestatore
	isSet bool
}

func (v NullableCedentePrestatore) Get() *CedentePrestatore {
	return v.value
}

func (v *NullableCedentePrestatore) Set(val *CedentePrestatore) {
	v.value = val
	v.isSet = true
}

func (v NullableCedentePrestatore) IsSet() bool {
	return v.isSet
}

func (v *NullableCedentePrestatore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCedentePrestatore(val *CedentePrestatore) *NullableCedentePrestatore {
	return &NullableCedentePrestatore{value: val, isSet: true}
}

func (v NullableCedentePrestatore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCedentePrestatore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


