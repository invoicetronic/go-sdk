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
	"time"
)

// checks if the DatiAnagraficiCedentePrestatore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatiAnagraficiCedentePrestatore{}

// DatiAnagraficiCedentePrestatore struct for DatiAnagraficiCedentePrestatore
type DatiAnagraficiCedentePrestatore struct {
	IdFiscaleIva *IdFiscaleIVA `json:"id_fiscale_iva,omitempty"`
	CodiceFiscale NullableString `json:"codice_fiscale,omitempty"`
	Anagrafica *Anagrafica `json:"anagrafica,omitempty"`
	AlboProfessionale NullableString `json:"albo_professionale,omitempty"`
	ProvinciaAlbo NullableString `json:"provincia_albo,omitempty"`
	NumeroIscrizioneAlbo NullableString `json:"numero_iscrizione_albo,omitempty"`
	DataIscrizioneAlbo NullableTime `json:"data_iscrizione_albo,omitempty"`
	RegimeFiscale NullableString `json:"regime_fiscale,omitempty"`
}

// NewDatiAnagraficiCedentePrestatore instantiates a new DatiAnagraficiCedentePrestatore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatiAnagraficiCedentePrestatore() *DatiAnagraficiCedentePrestatore {
	this := DatiAnagraficiCedentePrestatore{}
	return &this
}

// NewDatiAnagraficiCedentePrestatoreWithDefaults instantiates a new DatiAnagraficiCedentePrestatore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatiAnagraficiCedentePrestatoreWithDefaults() *DatiAnagraficiCedentePrestatore {
	this := DatiAnagraficiCedentePrestatore{}
	return &this
}

// GetIdFiscaleIva returns the IdFiscaleIva field value if set, zero value otherwise.
func (o *DatiAnagraficiCedentePrestatore) GetIdFiscaleIva() IdFiscaleIVA {
	if o == nil || IsNil(o.IdFiscaleIva) {
		var ret IdFiscaleIVA
		return ret
	}
	return *o.IdFiscaleIva
}

// GetIdFiscaleIvaOk returns a tuple with the IdFiscaleIva field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatiAnagraficiCedentePrestatore) GetIdFiscaleIvaOk() (*IdFiscaleIVA, bool) {
	if o == nil || IsNil(o.IdFiscaleIva) {
		return nil, false
	}
	return o.IdFiscaleIva, true
}

// HasIdFiscaleIva returns a boolean if a field has been set.
func (o *DatiAnagraficiCedentePrestatore) HasIdFiscaleIva() bool {
	if o != nil && !IsNil(o.IdFiscaleIva) {
		return true
	}

	return false
}

// SetIdFiscaleIva gets a reference to the given IdFiscaleIVA and assigns it to the IdFiscaleIva field.
func (o *DatiAnagraficiCedentePrestatore) SetIdFiscaleIva(v IdFiscaleIVA) {
	o.IdFiscaleIva = &v
}

// GetCodiceFiscale returns the CodiceFiscale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiAnagraficiCedentePrestatore) GetCodiceFiscale() string {
	if o == nil || IsNil(o.CodiceFiscale.Get()) {
		var ret string
		return ret
	}
	return *o.CodiceFiscale.Get()
}

// GetCodiceFiscaleOk returns a tuple with the CodiceFiscale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiAnagraficiCedentePrestatore) GetCodiceFiscaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodiceFiscale.Get(), o.CodiceFiscale.IsSet()
}

// HasCodiceFiscale returns a boolean if a field has been set.
func (o *DatiAnagraficiCedentePrestatore) HasCodiceFiscale() bool {
	if o != nil && o.CodiceFiscale.IsSet() {
		return true
	}

	return false
}

// SetCodiceFiscale gets a reference to the given NullableString and assigns it to the CodiceFiscale field.
func (o *DatiAnagraficiCedentePrestatore) SetCodiceFiscale(v string) {
	o.CodiceFiscale.Set(&v)
}
// SetCodiceFiscaleNil sets the value for CodiceFiscale to be an explicit nil
func (o *DatiAnagraficiCedentePrestatore) SetCodiceFiscaleNil() {
	o.CodiceFiscale.Set(nil)
}

// UnsetCodiceFiscale ensures that no value is present for CodiceFiscale, not even an explicit nil
func (o *DatiAnagraficiCedentePrestatore) UnsetCodiceFiscale() {
	o.CodiceFiscale.Unset()
}

// GetAnagrafica returns the Anagrafica field value if set, zero value otherwise.
func (o *DatiAnagraficiCedentePrestatore) GetAnagrafica() Anagrafica {
	if o == nil || IsNil(o.Anagrafica) {
		var ret Anagrafica
		return ret
	}
	return *o.Anagrafica
}

// GetAnagraficaOk returns a tuple with the Anagrafica field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatiAnagraficiCedentePrestatore) GetAnagraficaOk() (*Anagrafica, bool) {
	if o == nil || IsNil(o.Anagrafica) {
		return nil, false
	}
	return o.Anagrafica, true
}

// HasAnagrafica returns a boolean if a field has been set.
func (o *DatiAnagraficiCedentePrestatore) HasAnagrafica() bool {
	if o != nil && !IsNil(o.Anagrafica) {
		return true
	}

	return false
}

// SetAnagrafica gets a reference to the given Anagrafica and assigns it to the Anagrafica field.
func (o *DatiAnagraficiCedentePrestatore) SetAnagrafica(v Anagrafica) {
	o.Anagrafica = &v
}

// GetAlboProfessionale returns the AlboProfessionale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiAnagraficiCedentePrestatore) GetAlboProfessionale() string {
	if o == nil || IsNil(o.AlboProfessionale.Get()) {
		var ret string
		return ret
	}
	return *o.AlboProfessionale.Get()
}

// GetAlboProfessionaleOk returns a tuple with the AlboProfessionale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiAnagraficiCedentePrestatore) GetAlboProfessionaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlboProfessionale.Get(), o.AlboProfessionale.IsSet()
}

// HasAlboProfessionale returns a boolean if a field has been set.
func (o *DatiAnagraficiCedentePrestatore) HasAlboProfessionale() bool {
	if o != nil && o.AlboProfessionale.IsSet() {
		return true
	}

	return false
}

// SetAlboProfessionale gets a reference to the given NullableString and assigns it to the AlboProfessionale field.
func (o *DatiAnagraficiCedentePrestatore) SetAlboProfessionale(v string) {
	o.AlboProfessionale.Set(&v)
}
// SetAlboProfessionaleNil sets the value for AlboProfessionale to be an explicit nil
func (o *DatiAnagraficiCedentePrestatore) SetAlboProfessionaleNil() {
	o.AlboProfessionale.Set(nil)
}

// UnsetAlboProfessionale ensures that no value is present for AlboProfessionale, not even an explicit nil
func (o *DatiAnagraficiCedentePrestatore) UnsetAlboProfessionale() {
	o.AlboProfessionale.Unset()
}

// GetProvinciaAlbo returns the ProvinciaAlbo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiAnagraficiCedentePrestatore) GetProvinciaAlbo() string {
	if o == nil || IsNil(o.ProvinciaAlbo.Get()) {
		var ret string
		return ret
	}
	return *o.ProvinciaAlbo.Get()
}

// GetProvinciaAlboOk returns a tuple with the ProvinciaAlbo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiAnagraficiCedentePrestatore) GetProvinciaAlboOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProvinciaAlbo.Get(), o.ProvinciaAlbo.IsSet()
}

// HasProvinciaAlbo returns a boolean if a field has been set.
func (o *DatiAnagraficiCedentePrestatore) HasProvinciaAlbo() bool {
	if o != nil && o.ProvinciaAlbo.IsSet() {
		return true
	}

	return false
}

// SetProvinciaAlbo gets a reference to the given NullableString and assigns it to the ProvinciaAlbo field.
func (o *DatiAnagraficiCedentePrestatore) SetProvinciaAlbo(v string) {
	o.ProvinciaAlbo.Set(&v)
}
// SetProvinciaAlboNil sets the value for ProvinciaAlbo to be an explicit nil
func (o *DatiAnagraficiCedentePrestatore) SetProvinciaAlboNil() {
	o.ProvinciaAlbo.Set(nil)
}

// UnsetProvinciaAlbo ensures that no value is present for ProvinciaAlbo, not even an explicit nil
func (o *DatiAnagraficiCedentePrestatore) UnsetProvinciaAlbo() {
	o.ProvinciaAlbo.Unset()
}

// GetNumeroIscrizioneAlbo returns the NumeroIscrizioneAlbo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiAnagraficiCedentePrestatore) GetNumeroIscrizioneAlbo() string {
	if o == nil || IsNil(o.NumeroIscrizioneAlbo.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroIscrizioneAlbo.Get()
}

// GetNumeroIscrizioneAlboOk returns a tuple with the NumeroIscrizioneAlbo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiAnagraficiCedentePrestatore) GetNumeroIscrizioneAlboOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroIscrizioneAlbo.Get(), o.NumeroIscrizioneAlbo.IsSet()
}

// HasNumeroIscrizioneAlbo returns a boolean if a field has been set.
func (o *DatiAnagraficiCedentePrestatore) HasNumeroIscrizioneAlbo() bool {
	if o != nil && o.NumeroIscrizioneAlbo.IsSet() {
		return true
	}

	return false
}

// SetNumeroIscrizioneAlbo gets a reference to the given NullableString and assigns it to the NumeroIscrizioneAlbo field.
func (o *DatiAnagraficiCedentePrestatore) SetNumeroIscrizioneAlbo(v string) {
	o.NumeroIscrizioneAlbo.Set(&v)
}
// SetNumeroIscrizioneAlboNil sets the value for NumeroIscrizioneAlbo to be an explicit nil
func (o *DatiAnagraficiCedentePrestatore) SetNumeroIscrizioneAlboNil() {
	o.NumeroIscrizioneAlbo.Set(nil)
}

// UnsetNumeroIscrizioneAlbo ensures that no value is present for NumeroIscrizioneAlbo, not even an explicit nil
func (o *DatiAnagraficiCedentePrestatore) UnsetNumeroIscrizioneAlbo() {
	o.NumeroIscrizioneAlbo.Unset()
}

// GetDataIscrizioneAlbo returns the DataIscrizioneAlbo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiAnagraficiCedentePrestatore) GetDataIscrizioneAlbo() time.Time {
	if o == nil || IsNil(o.DataIscrizioneAlbo.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DataIscrizioneAlbo.Get()
}

// GetDataIscrizioneAlboOk returns a tuple with the DataIscrizioneAlbo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiAnagraficiCedentePrestatore) GetDataIscrizioneAlboOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataIscrizioneAlbo.Get(), o.DataIscrizioneAlbo.IsSet()
}

// HasDataIscrizioneAlbo returns a boolean if a field has been set.
func (o *DatiAnagraficiCedentePrestatore) HasDataIscrizioneAlbo() bool {
	if o != nil && o.DataIscrizioneAlbo.IsSet() {
		return true
	}

	return false
}

// SetDataIscrizioneAlbo gets a reference to the given NullableTime and assigns it to the DataIscrizioneAlbo field.
func (o *DatiAnagraficiCedentePrestatore) SetDataIscrizioneAlbo(v time.Time) {
	o.DataIscrizioneAlbo.Set(&v)
}
// SetDataIscrizioneAlboNil sets the value for DataIscrizioneAlbo to be an explicit nil
func (o *DatiAnagraficiCedentePrestatore) SetDataIscrizioneAlboNil() {
	o.DataIscrizioneAlbo.Set(nil)
}

// UnsetDataIscrizioneAlbo ensures that no value is present for DataIscrizioneAlbo, not even an explicit nil
func (o *DatiAnagraficiCedentePrestatore) UnsetDataIscrizioneAlbo() {
	o.DataIscrizioneAlbo.Unset()
}

// GetRegimeFiscale returns the RegimeFiscale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiAnagraficiCedentePrestatore) GetRegimeFiscale() string {
	if o == nil || IsNil(o.RegimeFiscale.Get()) {
		var ret string
		return ret
	}
	return *o.RegimeFiscale.Get()
}

// GetRegimeFiscaleOk returns a tuple with the RegimeFiscale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiAnagraficiCedentePrestatore) GetRegimeFiscaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegimeFiscale.Get(), o.RegimeFiscale.IsSet()
}

// HasRegimeFiscale returns a boolean if a field has been set.
func (o *DatiAnagraficiCedentePrestatore) HasRegimeFiscale() bool {
	if o != nil && o.RegimeFiscale.IsSet() {
		return true
	}

	return false
}

// SetRegimeFiscale gets a reference to the given NullableString and assigns it to the RegimeFiscale field.
func (o *DatiAnagraficiCedentePrestatore) SetRegimeFiscale(v string) {
	o.RegimeFiscale.Set(&v)
}
// SetRegimeFiscaleNil sets the value for RegimeFiscale to be an explicit nil
func (o *DatiAnagraficiCedentePrestatore) SetRegimeFiscaleNil() {
	o.RegimeFiscale.Set(nil)
}

// UnsetRegimeFiscale ensures that no value is present for RegimeFiscale, not even an explicit nil
func (o *DatiAnagraficiCedentePrestatore) UnsetRegimeFiscale() {
	o.RegimeFiscale.Unset()
}

func (o DatiAnagraficiCedentePrestatore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatiAnagraficiCedentePrestatore) ToMap() (map[string]interface{}, error) {
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
	if o.AlboProfessionale.IsSet() {
		toSerialize["albo_professionale"] = o.AlboProfessionale.Get()
	}
	if o.ProvinciaAlbo.IsSet() {
		toSerialize["provincia_albo"] = o.ProvinciaAlbo.Get()
	}
	if o.NumeroIscrizioneAlbo.IsSet() {
		toSerialize["numero_iscrizione_albo"] = o.NumeroIscrizioneAlbo.Get()
	}
	if o.DataIscrizioneAlbo.IsSet() {
		toSerialize["data_iscrizione_albo"] = o.DataIscrizioneAlbo.Get()
	}
	if o.RegimeFiscale.IsSet() {
		toSerialize["regime_fiscale"] = o.RegimeFiscale.Get()
	}
	return toSerialize, nil
}

type NullableDatiAnagraficiCedentePrestatore struct {
	value *DatiAnagraficiCedentePrestatore
	isSet bool
}

func (v NullableDatiAnagraficiCedentePrestatore) Get() *DatiAnagraficiCedentePrestatore {
	return v.value
}

func (v *NullableDatiAnagraficiCedentePrestatore) Set(val *DatiAnagraficiCedentePrestatore) {
	v.value = val
	v.isSet = true
}

func (v NullableDatiAnagraficiCedentePrestatore) IsSet() bool {
	return v.isSet
}

func (v *NullableDatiAnagraficiCedentePrestatore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatiAnagraficiCedentePrestatore(val *DatiAnagraficiCedentePrestatore) *NullableDatiAnagraficiCedentePrestatore {
	return &NullableDatiAnagraficiCedentePrestatore{value: val, isSet: true}
}

func (v NullableDatiAnagraficiCedentePrestatore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatiAnagraficiCedentePrestatore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


