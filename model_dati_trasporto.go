/*
Italian eInvoice API

The Italian eInvoice API is a RESTful API that allows you to send and receive invoices through the Italian [Servizio di Interscambio (SDI)][1], or Interchange Service. The API is designed by Invoicetronic to be simple and easy to use, abstracting away SDI complexity while still providing complete control over the invoice send/receive process. The API also provides advanced features and a rich toolchain, such as invoice validation, multiple upload methods, webhooks, event logs, CORS support, client SDKs for commonly used languages, and CLI tools.  For more information, see  [Invoicetronic website][2]  [1]: https://www.fatturapa.gov.it/it/sistemainterscambio/cose-il-sdi/ [2]: https://invoicetronic.com/

API version: 1.0.0
Contact: support@invoicetronic.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package invoicesdk

import (
	"encoding/json"
	"time"
)

// checks if the DatiTrasporto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatiTrasporto{}

// DatiTrasporto struct for DatiTrasporto
type DatiTrasporto struct {
	DatiAnagraficiVettore *DatiAnagraficiVettore `json:"dati_anagrafici_vettore,omitempty"`
	MezzoTrasporto NullableString `json:"mezzo_trasporto,omitempty"`
	CausaleTrasporto NullableString `json:"causale_trasporto,omitempty"`
	NumeroColli NullableInt32 `json:"numero_colli,omitempty"`
	Descrizione NullableString `json:"descrizione,omitempty"`
	UnitaMisuraPeso NullableString `json:"unita_misura_peso,omitempty"`
	PesoLordo NullableFloat64 `json:"peso_lordo,omitempty"`
	PesoNetto NullableFloat64 `json:"peso_netto,omitempty"`
	DataOraRitiro NullableTime `json:"data_ora_ritiro,omitempty"`
	DataInizioTrasporto NullableTime `json:"data_inizio_trasporto,omitempty"`
	TipoResa NullableString `json:"tipo_resa,omitempty"`
	IndirizzoResa *IndirizzoResa `json:"indirizzo_resa,omitempty"`
	DataOraConsegna NullableTime `json:"data_ora_consegna,omitempty"`
}

// NewDatiTrasporto instantiates a new DatiTrasporto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatiTrasporto() *DatiTrasporto {
	this := DatiTrasporto{}
	return &this
}

// NewDatiTrasportoWithDefaults instantiates a new DatiTrasporto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatiTrasportoWithDefaults() *DatiTrasporto {
	this := DatiTrasporto{}
	return &this
}

// GetDatiAnagraficiVettore returns the DatiAnagraficiVettore field value if set, zero value otherwise.
func (o *DatiTrasporto) GetDatiAnagraficiVettore() DatiAnagraficiVettore {
	if o == nil || IsNil(o.DatiAnagraficiVettore) {
		var ret DatiAnagraficiVettore
		return ret
	}
	return *o.DatiAnagraficiVettore
}

// GetDatiAnagraficiVettoreOk returns a tuple with the DatiAnagraficiVettore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatiTrasporto) GetDatiAnagraficiVettoreOk() (*DatiAnagraficiVettore, bool) {
	if o == nil || IsNil(o.DatiAnagraficiVettore) {
		return nil, false
	}
	return o.DatiAnagraficiVettore, true
}

// HasDatiAnagraficiVettore returns a boolean if a field has been set.
func (o *DatiTrasporto) HasDatiAnagraficiVettore() bool {
	if o != nil && !IsNil(o.DatiAnagraficiVettore) {
		return true
	}

	return false
}

// SetDatiAnagraficiVettore gets a reference to the given DatiAnagraficiVettore and assigns it to the DatiAnagraficiVettore field.
func (o *DatiTrasporto) SetDatiAnagraficiVettore(v DatiAnagraficiVettore) {
	o.DatiAnagraficiVettore = &v
}

// GetMezzoTrasporto returns the MezzoTrasporto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiTrasporto) GetMezzoTrasporto() string {
	if o == nil || IsNil(o.MezzoTrasporto.Get()) {
		var ret string
		return ret
	}
	return *o.MezzoTrasporto.Get()
}

// GetMezzoTrasportoOk returns a tuple with the MezzoTrasporto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiTrasporto) GetMezzoTrasportoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MezzoTrasporto.Get(), o.MezzoTrasporto.IsSet()
}

// HasMezzoTrasporto returns a boolean if a field has been set.
func (o *DatiTrasporto) HasMezzoTrasporto() bool {
	if o != nil && o.MezzoTrasporto.IsSet() {
		return true
	}

	return false
}

// SetMezzoTrasporto gets a reference to the given NullableString and assigns it to the MezzoTrasporto field.
func (o *DatiTrasporto) SetMezzoTrasporto(v string) {
	o.MezzoTrasporto.Set(&v)
}
// SetMezzoTrasportoNil sets the value for MezzoTrasporto to be an explicit nil
func (o *DatiTrasporto) SetMezzoTrasportoNil() {
	o.MezzoTrasporto.Set(nil)
}

// UnsetMezzoTrasporto ensures that no value is present for MezzoTrasporto, not even an explicit nil
func (o *DatiTrasporto) UnsetMezzoTrasporto() {
	o.MezzoTrasporto.Unset()
}

// GetCausaleTrasporto returns the CausaleTrasporto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiTrasporto) GetCausaleTrasporto() string {
	if o == nil || IsNil(o.CausaleTrasporto.Get()) {
		var ret string
		return ret
	}
	return *o.CausaleTrasporto.Get()
}

// GetCausaleTrasportoOk returns a tuple with the CausaleTrasporto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiTrasporto) GetCausaleTrasportoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CausaleTrasporto.Get(), o.CausaleTrasporto.IsSet()
}

// HasCausaleTrasporto returns a boolean if a field has been set.
func (o *DatiTrasporto) HasCausaleTrasporto() bool {
	if o != nil && o.CausaleTrasporto.IsSet() {
		return true
	}

	return false
}

// SetCausaleTrasporto gets a reference to the given NullableString and assigns it to the CausaleTrasporto field.
func (o *DatiTrasporto) SetCausaleTrasporto(v string) {
	o.CausaleTrasporto.Set(&v)
}
// SetCausaleTrasportoNil sets the value for CausaleTrasporto to be an explicit nil
func (o *DatiTrasporto) SetCausaleTrasportoNil() {
	o.CausaleTrasporto.Set(nil)
}

// UnsetCausaleTrasporto ensures that no value is present for CausaleTrasporto, not even an explicit nil
func (o *DatiTrasporto) UnsetCausaleTrasporto() {
	o.CausaleTrasporto.Unset()
}

// GetNumeroColli returns the NumeroColli field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiTrasporto) GetNumeroColli() int32 {
	if o == nil || IsNil(o.NumeroColli.Get()) {
		var ret int32
		return ret
	}
	return *o.NumeroColli.Get()
}

// GetNumeroColliOk returns a tuple with the NumeroColli field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiTrasporto) GetNumeroColliOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroColli.Get(), o.NumeroColli.IsSet()
}

// HasNumeroColli returns a boolean if a field has been set.
func (o *DatiTrasporto) HasNumeroColli() bool {
	if o != nil && o.NumeroColli.IsSet() {
		return true
	}

	return false
}

// SetNumeroColli gets a reference to the given NullableInt32 and assigns it to the NumeroColli field.
func (o *DatiTrasporto) SetNumeroColli(v int32) {
	o.NumeroColli.Set(&v)
}
// SetNumeroColliNil sets the value for NumeroColli to be an explicit nil
func (o *DatiTrasporto) SetNumeroColliNil() {
	o.NumeroColli.Set(nil)
}

// UnsetNumeroColli ensures that no value is present for NumeroColli, not even an explicit nil
func (o *DatiTrasporto) UnsetNumeroColli() {
	o.NumeroColli.Unset()
}

// GetDescrizione returns the Descrizione field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiTrasporto) GetDescrizione() string {
	if o == nil || IsNil(o.Descrizione.Get()) {
		var ret string
		return ret
	}
	return *o.Descrizione.Get()
}

// GetDescrizioneOk returns a tuple with the Descrizione field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiTrasporto) GetDescrizioneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Descrizione.Get(), o.Descrizione.IsSet()
}

// HasDescrizione returns a boolean if a field has been set.
func (o *DatiTrasporto) HasDescrizione() bool {
	if o != nil && o.Descrizione.IsSet() {
		return true
	}

	return false
}

// SetDescrizione gets a reference to the given NullableString and assigns it to the Descrizione field.
func (o *DatiTrasporto) SetDescrizione(v string) {
	o.Descrizione.Set(&v)
}
// SetDescrizioneNil sets the value for Descrizione to be an explicit nil
func (o *DatiTrasporto) SetDescrizioneNil() {
	o.Descrizione.Set(nil)
}

// UnsetDescrizione ensures that no value is present for Descrizione, not even an explicit nil
func (o *DatiTrasporto) UnsetDescrizione() {
	o.Descrizione.Unset()
}

// GetUnitaMisuraPeso returns the UnitaMisuraPeso field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiTrasporto) GetUnitaMisuraPeso() string {
	if o == nil || IsNil(o.UnitaMisuraPeso.Get()) {
		var ret string
		return ret
	}
	return *o.UnitaMisuraPeso.Get()
}

// GetUnitaMisuraPesoOk returns a tuple with the UnitaMisuraPeso field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiTrasporto) GetUnitaMisuraPesoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitaMisuraPeso.Get(), o.UnitaMisuraPeso.IsSet()
}

// HasUnitaMisuraPeso returns a boolean if a field has been set.
func (o *DatiTrasporto) HasUnitaMisuraPeso() bool {
	if o != nil && o.UnitaMisuraPeso.IsSet() {
		return true
	}

	return false
}

// SetUnitaMisuraPeso gets a reference to the given NullableString and assigns it to the UnitaMisuraPeso field.
func (o *DatiTrasporto) SetUnitaMisuraPeso(v string) {
	o.UnitaMisuraPeso.Set(&v)
}
// SetUnitaMisuraPesoNil sets the value for UnitaMisuraPeso to be an explicit nil
func (o *DatiTrasporto) SetUnitaMisuraPesoNil() {
	o.UnitaMisuraPeso.Set(nil)
}

// UnsetUnitaMisuraPeso ensures that no value is present for UnitaMisuraPeso, not even an explicit nil
func (o *DatiTrasporto) UnsetUnitaMisuraPeso() {
	o.UnitaMisuraPeso.Unset()
}

// GetPesoLordo returns the PesoLordo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiTrasporto) GetPesoLordo() float64 {
	if o == nil || IsNil(o.PesoLordo.Get()) {
		var ret float64
		return ret
	}
	return *o.PesoLordo.Get()
}

// GetPesoLordoOk returns a tuple with the PesoLordo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiTrasporto) GetPesoLordoOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PesoLordo.Get(), o.PesoLordo.IsSet()
}

// HasPesoLordo returns a boolean if a field has been set.
func (o *DatiTrasporto) HasPesoLordo() bool {
	if o != nil && o.PesoLordo.IsSet() {
		return true
	}

	return false
}

// SetPesoLordo gets a reference to the given NullableFloat64 and assigns it to the PesoLordo field.
func (o *DatiTrasporto) SetPesoLordo(v float64) {
	o.PesoLordo.Set(&v)
}
// SetPesoLordoNil sets the value for PesoLordo to be an explicit nil
func (o *DatiTrasporto) SetPesoLordoNil() {
	o.PesoLordo.Set(nil)
}

// UnsetPesoLordo ensures that no value is present for PesoLordo, not even an explicit nil
func (o *DatiTrasporto) UnsetPesoLordo() {
	o.PesoLordo.Unset()
}

// GetPesoNetto returns the PesoNetto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiTrasporto) GetPesoNetto() float64 {
	if o == nil || IsNil(o.PesoNetto.Get()) {
		var ret float64
		return ret
	}
	return *o.PesoNetto.Get()
}

// GetPesoNettoOk returns a tuple with the PesoNetto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiTrasporto) GetPesoNettoOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PesoNetto.Get(), o.PesoNetto.IsSet()
}

// HasPesoNetto returns a boolean if a field has been set.
func (o *DatiTrasporto) HasPesoNetto() bool {
	if o != nil && o.PesoNetto.IsSet() {
		return true
	}

	return false
}

// SetPesoNetto gets a reference to the given NullableFloat64 and assigns it to the PesoNetto field.
func (o *DatiTrasporto) SetPesoNetto(v float64) {
	o.PesoNetto.Set(&v)
}
// SetPesoNettoNil sets the value for PesoNetto to be an explicit nil
func (o *DatiTrasporto) SetPesoNettoNil() {
	o.PesoNetto.Set(nil)
}

// UnsetPesoNetto ensures that no value is present for PesoNetto, not even an explicit nil
func (o *DatiTrasporto) UnsetPesoNetto() {
	o.PesoNetto.Unset()
}

// GetDataOraRitiro returns the DataOraRitiro field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiTrasporto) GetDataOraRitiro() time.Time {
	if o == nil || IsNil(o.DataOraRitiro.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DataOraRitiro.Get()
}

// GetDataOraRitiroOk returns a tuple with the DataOraRitiro field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiTrasporto) GetDataOraRitiroOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataOraRitiro.Get(), o.DataOraRitiro.IsSet()
}

// HasDataOraRitiro returns a boolean if a field has been set.
func (o *DatiTrasporto) HasDataOraRitiro() bool {
	if o != nil && o.DataOraRitiro.IsSet() {
		return true
	}

	return false
}

// SetDataOraRitiro gets a reference to the given NullableTime and assigns it to the DataOraRitiro field.
func (o *DatiTrasporto) SetDataOraRitiro(v time.Time) {
	o.DataOraRitiro.Set(&v)
}
// SetDataOraRitiroNil sets the value for DataOraRitiro to be an explicit nil
func (o *DatiTrasporto) SetDataOraRitiroNil() {
	o.DataOraRitiro.Set(nil)
}

// UnsetDataOraRitiro ensures that no value is present for DataOraRitiro, not even an explicit nil
func (o *DatiTrasporto) UnsetDataOraRitiro() {
	o.DataOraRitiro.Unset()
}

// GetDataInizioTrasporto returns the DataInizioTrasporto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiTrasporto) GetDataInizioTrasporto() time.Time {
	if o == nil || IsNil(o.DataInizioTrasporto.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DataInizioTrasporto.Get()
}

// GetDataInizioTrasportoOk returns a tuple with the DataInizioTrasporto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiTrasporto) GetDataInizioTrasportoOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataInizioTrasporto.Get(), o.DataInizioTrasporto.IsSet()
}

// HasDataInizioTrasporto returns a boolean if a field has been set.
func (o *DatiTrasporto) HasDataInizioTrasporto() bool {
	if o != nil && o.DataInizioTrasporto.IsSet() {
		return true
	}

	return false
}

// SetDataInizioTrasporto gets a reference to the given NullableTime and assigns it to the DataInizioTrasporto field.
func (o *DatiTrasporto) SetDataInizioTrasporto(v time.Time) {
	o.DataInizioTrasporto.Set(&v)
}
// SetDataInizioTrasportoNil sets the value for DataInizioTrasporto to be an explicit nil
func (o *DatiTrasporto) SetDataInizioTrasportoNil() {
	o.DataInizioTrasporto.Set(nil)
}

// UnsetDataInizioTrasporto ensures that no value is present for DataInizioTrasporto, not even an explicit nil
func (o *DatiTrasporto) UnsetDataInizioTrasporto() {
	o.DataInizioTrasporto.Unset()
}

// GetTipoResa returns the TipoResa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiTrasporto) GetTipoResa() string {
	if o == nil || IsNil(o.TipoResa.Get()) {
		var ret string
		return ret
	}
	return *o.TipoResa.Get()
}

// GetTipoResaOk returns a tuple with the TipoResa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiTrasporto) GetTipoResaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TipoResa.Get(), o.TipoResa.IsSet()
}

// HasTipoResa returns a boolean if a field has been set.
func (o *DatiTrasporto) HasTipoResa() bool {
	if o != nil && o.TipoResa.IsSet() {
		return true
	}

	return false
}

// SetTipoResa gets a reference to the given NullableString and assigns it to the TipoResa field.
func (o *DatiTrasporto) SetTipoResa(v string) {
	o.TipoResa.Set(&v)
}
// SetTipoResaNil sets the value for TipoResa to be an explicit nil
func (o *DatiTrasporto) SetTipoResaNil() {
	o.TipoResa.Set(nil)
}

// UnsetTipoResa ensures that no value is present for TipoResa, not even an explicit nil
func (o *DatiTrasporto) UnsetTipoResa() {
	o.TipoResa.Unset()
}

// GetIndirizzoResa returns the IndirizzoResa field value if set, zero value otherwise.
func (o *DatiTrasporto) GetIndirizzoResa() IndirizzoResa {
	if o == nil || IsNil(o.IndirizzoResa) {
		var ret IndirizzoResa
		return ret
	}
	return *o.IndirizzoResa
}

// GetIndirizzoResaOk returns a tuple with the IndirizzoResa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatiTrasporto) GetIndirizzoResaOk() (*IndirizzoResa, bool) {
	if o == nil || IsNil(o.IndirizzoResa) {
		return nil, false
	}
	return o.IndirizzoResa, true
}

// HasIndirizzoResa returns a boolean if a field has been set.
func (o *DatiTrasporto) HasIndirizzoResa() bool {
	if o != nil && !IsNil(o.IndirizzoResa) {
		return true
	}

	return false
}

// SetIndirizzoResa gets a reference to the given IndirizzoResa and assigns it to the IndirizzoResa field.
func (o *DatiTrasporto) SetIndirizzoResa(v IndirizzoResa) {
	o.IndirizzoResa = &v
}

// GetDataOraConsegna returns the DataOraConsegna field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiTrasporto) GetDataOraConsegna() time.Time {
	if o == nil || IsNil(o.DataOraConsegna.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DataOraConsegna.Get()
}

// GetDataOraConsegnaOk returns a tuple with the DataOraConsegna field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiTrasporto) GetDataOraConsegnaOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataOraConsegna.Get(), o.DataOraConsegna.IsSet()
}

// HasDataOraConsegna returns a boolean if a field has been set.
func (o *DatiTrasporto) HasDataOraConsegna() bool {
	if o != nil && o.DataOraConsegna.IsSet() {
		return true
	}

	return false
}

// SetDataOraConsegna gets a reference to the given NullableTime and assigns it to the DataOraConsegna field.
func (o *DatiTrasporto) SetDataOraConsegna(v time.Time) {
	o.DataOraConsegna.Set(&v)
}
// SetDataOraConsegnaNil sets the value for DataOraConsegna to be an explicit nil
func (o *DatiTrasporto) SetDataOraConsegnaNil() {
	o.DataOraConsegna.Set(nil)
}

// UnsetDataOraConsegna ensures that no value is present for DataOraConsegna, not even an explicit nil
func (o *DatiTrasporto) UnsetDataOraConsegna() {
	o.DataOraConsegna.Unset()
}

func (o DatiTrasporto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatiTrasporto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatiAnagraficiVettore) {
		toSerialize["dati_anagrafici_vettore"] = o.DatiAnagraficiVettore
	}
	if o.MezzoTrasporto.IsSet() {
		toSerialize["mezzo_trasporto"] = o.MezzoTrasporto.Get()
	}
	if o.CausaleTrasporto.IsSet() {
		toSerialize["causale_trasporto"] = o.CausaleTrasporto.Get()
	}
	if o.NumeroColli.IsSet() {
		toSerialize["numero_colli"] = o.NumeroColli.Get()
	}
	if o.Descrizione.IsSet() {
		toSerialize["descrizione"] = o.Descrizione.Get()
	}
	if o.UnitaMisuraPeso.IsSet() {
		toSerialize["unita_misura_peso"] = o.UnitaMisuraPeso.Get()
	}
	if o.PesoLordo.IsSet() {
		toSerialize["peso_lordo"] = o.PesoLordo.Get()
	}
	if o.PesoNetto.IsSet() {
		toSerialize["peso_netto"] = o.PesoNetto.Get()
	}
	if o.DataOraRitiro.IsSet() {
		toSerialize["data_ora_ritiro"] = o.DataOraRitiro.Get()
	}
	if o.DataInizioTrasporto.IsSet() {
		toSerialize["data_inizio_trasporto"] = o.DataInizioTrasporto.Get()
	}
	if o.TipoResa.IsSet() {
		toSerialize["tipo_resa"] = o.TipoResa.Get()
	}
	if !IsNil(o.IndirizzoResa) {
		toSerialize["indirizzo_resa"] = o.IndirizzoResa
	}
	if o.DataOraConsegna.IsSet() {
		toSerialize["data_ora_consegna"] = o.DataOraConsegna.Get()
	}
	return toSerialize, nil
}

type NullableDatiTrasporto struct {
	value *DatiTrasporto
	isSet bool
}

func (v NullableDatiTrasporto) Get() *DatiTrasporto {
	return v.value
}

func (v *NullableDatiTrasporto) Set(val *DatiTrasporto) {
	v.value = val
	v.isSet = true
}

func (v NullableDatiTrasporto) IsSet() bool {
	return v.isSet
}

func (v *NullableDatiTrasporto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatiTrasporto(val *DatiTrasporto) *NullableDatiTrasporto {
	return &NullableDatiTrasporto{value: val, isSet: true}
}

func (v NullableDatiTrasporto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatiTrasporto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


