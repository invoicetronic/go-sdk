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
	"time"
)

// checks if the DettaglioLinee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DettaglioLinee{}

// DettaglioLinee struct for DettaglioLinee
type DettaglioLinee struct {
	NumeroLinea *int32 `json:"numero_linea,omitempty"`
	TipoCessionePrestazione NullableString `json:"tipo_cessione_prestazione,omitempty"`
	CodiceArticolo []CodiceArticolo `json:"codice_articolo,omitempty"`
	Descrizione NullableString `json:"descrizione,omitempty"`
	Quantita NullableFloat64 `json:"quantita,omitempty"`
	UnitaMisura NullableString `json:"unita_misura,omitempty"`
	DataInizioPeriodo NullableTime `json:"data_inizio_periodo,omitempty"`
	DataFinePeriodo NullableTime `json:"data_fine_periodo,omitempty"`
	PrezzoUnitario *float64 `json:"prezzo_unitario,omitempty"`
	ScontoMaggiorazione []ScontoMaggiorazione `json:"sconto_maggiorazione,omitempty"`
	PrezzoTotale *float64 `json:"prezzo_totale,omitempty"`
	AliquotaIva *float64 `json:"aliquota_iva,omitempty"`
	Ritenuta NullableString `json:"ritenuta,omitempty"`
	Natura NullableString `json:"natura,omitempty"`
	RiferimentoAmministrazione NullableString `json:"riferimento_amministrazione,omitempty"`
	AltriDatiGestionali []AltriDatiGestionali `json:"altri_dati_gestionali,omitempty"`
}

// NewDettaglioLinee instantiates a new DettaglioLinee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDettaglioLinee() *DettaglioLinee {
	this := DettaglioLinee{}
	return &this
}

// NewDettaglioLineeWithDefaults instantiates a new DettaglioLinee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDettaglioLineeWithDefaults() *DettaglioLinee {
	this := DettaglioLinee{}
	return &this
}

// GetNumeroLinea returns the NumeroLinea field value if set, zero value otherwise.
func (o *DettaglioLinee) GetNumeroLinea() int32 {
	if o == nil || IsNil(o.NumeroLinea) {
		var ret int32
		return ret
	}
	return *o.NumeroLinea
}

// GetNumeroLineaOk returns a tuple with the NumeroLinea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DettaglioLinee) GetNumeroLineaOk() (*int32, bool) {
	if o == nil || IsNil(o.NumeroLinea) {
		return nil, false
	}
	return o.NumeroLinea, true
}

// HasNumeroLinea returns a boolean if a field has been set.
func (o *DettaglioLinee) HasNumeroLinea() bool {
	if o != nil && !IsNil(o.NumeroLinea) {
		return true
	}

	return false
}

// SetNumeroLinea gets a reference to the given int32 and assigns it to the NumeroLinea field.
func (o *DettaglioLinee) SetNumeroLinea(v int32) {
	o.NumeroLinea = &v
}

// GetTipoCessionePrestazione returns the TipoCessionePrestazione field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioLinee) GetTipoCessionePrestazione() string {
	if o == nil || IsNil(o.TipoCessionePrestazione.Get()) {
		var ret string
		return ret
	}
	return *o.TipoCessionePrestazione.Get()
}

// GetTipoCessionePrestazioneOk returns a tuple with the TipoCessionePrestazione field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioLinee) GetTipoCessionePrestazioneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TipoCessionePrestazione.Get(), o.TipoCessionePrestazione.IsSet()
}

// HasTipoCessionePrestazione returns a boolean if a field has been set.
func (o *DettaglioLinee) HasTipoCessionePrestazione() bool {
	if o != nil && o.TipoCessionePrestazione.IsSet() {
		return true
	}

	return false
}

// SetTipoCessionePrestazione gets a reference to the given NullableString and assigns it to the TipoCessionePrestazione field.
func (o *DettaglioLinee) SetTipoCessionePrestazione(v string) {
	o.TipoCessionePrestazione.Set(&v)
}
// SetTipoCessionePrestazioneNil sets the value for TipoCessionePrestazione to be an explicit nil
func (o *DettaglioLinee) SetTipoCessionePrestazioneNil() {
	o.TipoCessionePrestazione.Set(nil)
}

// UnsetTipoCessionePrestazione ensures that no value is present for TipoCessionePrestazione, not even an explicit nil
func (o *DettaglioLinee) UnsetTipoCessionePrestazione() {
	o.TipoCessionePrestazione.Unset()
}

// GetCodiceArticolo returns the CodiceArticolo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioLinee) GetCodiceArticolo() []CodiceArticolo {
	if o == nil {
		var ret []CodiceArticolo
		return ret
	}
	return o.CodiceArticolo
}

// GetCodiceArticoloOk returns a tuple with the CodiceArticolo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioLinee) GetCodiceArticoloOk() ([]CodiceArticolo, bool) {
	if o == nil || IsNil(o.CodiceArticolo) {
		return nil, false
	}
	return o.CodiceArticolo, true
}

// HasCodiceArticolo returns a boolean if a field has been set.
func (o *DettaglioLinee) HasCodiceArticolo() bool {
	if o != nil && !IsNil(o.CodiceArticolo) {
		return true
	}

	return false
}

// SetCodiceArticolo gets a reference to the given []CodiceArticolo and assigns it to the CodiceArticolo field.
func (o *DettaglioLinee) SetCodiceArticolo(v []CodiceArticolo) {
	o.CodiceArticolo = v
}

// GetDescrizione returns the Descrizione field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioLinee) GetDescrizione() string {
	if o == nil || IsNil(o.Descrizione.Get()) {
		var ret string
		return ret
	}
	return *o.Descrizione.Get()
}

// GetDescrizioneOk returns a tuple with the Descrizione field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioLinee) GetDescrizioneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Descrizione.Get(), o.Descrizione.IsSet()
}

// HasDescrizione returns a boolean if a field has been set.
func (o *DettaglioLinee) HasDescrizione() bool {
	if o != nil && o.Descrizione.IsSet() {
		return true
	}

	return false
}

// SetDescrizione gets a reference to the given NullableString and assigns it to the Descrizione field.
func (o *DettaglioLinee) SetDescrizione(v string) {
	o.Descrizione.Set(&v)
}
// SetDescrizioneNil sets the value for Descrizione to be an explicit nil
func (o *DettaglioLinee) SetDescrizioneNil() {
	o.Descrizione.Set(nil)
}

// UnsetDescrizione ensures that no value is present for Descrizione, not even an explicit nil
func (o *DettaglioLinee) UnsetDescrizione() {
	o.Descrizione.Unset()
}

// GetQuantita returns the Quantita field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioLinee) GetQuantita() float64 {
	if o == nil || IsNil(o.Quantita.Get()) {
		var ret float64
		return ret
	}
	return *o.Quantita.Get()
}

// GetQuantitaOk returns a tuple with the Quantita field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioLinee) GetQuantitaOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quantita.Get(), o.Quantita.IsSet()
}

// HasQuantita returns a boolean if a field has been set.
func (o *DettaglioLinee) HasQuantita() bool {
	if o != nil && o.Quantita.IsSet() {
		return true
	}

	return false
}

// SetQuantita gets a reference to the given NullableFloat64 and assigns it to the Quantita field.
func (o *DettaglioLinee) SetQuantita(v float64) {
	o.Quantita.Set(&v)
}
// SetQuantitaNil sets the value for Quantita to be an explicit nil
func (o *DettaglioLinee) SetQuantitaNil() {
	o.Quantita.Set(nil)
}

// UnsetQuantita ensures that no value is present for Quantita, not even an explicit nil
func (o *DettaglioLinee) UnsetQuantita() {
	o.Quantita.Unset()
}

// GetUnitaMisura returns the UnitaMisura field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioLinee) GetUnitaMisura() string {
	if o == nil || IsNil(o.UnitaMisura.Get()) {
		var ret string
		return ret
	}
	return *o.UnitaMisura.Get()
}

// GetUnitaMisuraOk returns a tuple with the UnitaMisura field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioLinee) GetUnitaMisuraOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitaMisura.Get(), o.UnitaMisura.IsSet()
}

// HasUnitaMisura returns a boolean if a field has been set.
func (o *DettaglioLinee) HasUnitaMisura() bool {
	if o != nil && o.UnitaMisura.IsSet() {
		return true
	}

	return false
}

// SetUnitaMisura gets a reference to the given NullableString and assigns it to the UnitaMisura field.
func (o *DettaglioLinee) SetUnitaMisura(v string) {
	o.UnitaMisura.Set(&v)
}
// SetUnitaMisuraNil sets the value for UnitaMisura to be an explicit nil
func (o *DettaglioLinee) SetUnitaMisuraNil() {
	o.UnitaMisura.Set(nil)
}

// UnsetUnitaMisura ensures that no value is present for UnitaMisura, not even an explicit nil
func (o *DettaglioLinee) UnsetUnitaMisura() {
	o.UnitaMisura.Unset()
}

// GetDataInizioPeriodo returns the DataInizioPeriodo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioLinee) GetDataInizioPeriodo() time.Time {
	if o == nil || IsNil(o.DataInizioPeriodo.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DataInizioPeriodo.Get()
}

// GetDataInizioPeriodoOk returns a tuple with the DataInizioPeriodo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioLinee) GetDataInizioPeriodoOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataInizioPeriodo.Get(), o.DataInizioPeriodo.IsSet()
}

// HasDataInizioPeriodo returns a boolean if a field has been set.
func (o *DettaglioLinee) HasDataInizioPeriodo() bool {
	if o != nil && o.DataInizioPeriodo.IsSet() {
		return true
	}

	return false
}

// SetDataInizioPeriodo gets a reference to the given NullableTime and assigns it to the DataInizioPeriodo field.
func (o *DettaglioLinee) SetDataInizioPeriodo(v time.Time) {
	o.DataInizioPeriodo.Set(&v)
}
// SetDataInizioPeriodoNil sets the value for DataInizioPeriodo to be an explicit nil
func (o *DettaglioLinee) SetDataInizioPeriodoNil() {
	o.DataInizioPeriodo.Set(nil)
}

// UnsetDataInizioPeriodo ensures that no value is present for DataInizioPeriodo, not even an explicit nil
func (o *DettaglioLinee) UnsetDataInizioPeriodo() {
	o.DataInizioPeriodo.Unset()
}

// GetDataFinePeriodo returns the DataFinePeriodo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioLinee) GetDataFinePeriodo() time.Time {
	if o == nil || IsNil(o.DataFinePeriodo.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DataFinePeriodo.Get()
}

// GetDataFinePeriodoOk returns a tuple with the DataFinePeriodo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioLinee) GetDataFinePeriodoOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataFinePeriodo.Get(), o.DataFinePeriodo.IsSet()
}

// HasDataFinePeriodo returns a boolean if a field has been set.
func (o *DettaglioLinee) HasDataFinePeriodo() bool {
	if o != nil && o.DataFinePeriodo.IsSet() {
		return true
	}

	return false
}

// SetDataFinePeriodo gets a reference to the given NullableTime and assigns it to the DataFinePeriodo field.
func (o *DettaglioLinee) SetDataFinePeriodo(v time.Time) {
	o.DataFinePeriodo.Set(&v)
}
// SetDataFinePeriodoNil sets the value for DataFinePeriodo to be an explicit nil
func (o *DettaglioLinee) SetDataFinePeriodoNil() {
	o.DataFinePeriodo.Set(nil)
}

// UnsetDataFinePeriodo ensures that no value is present for DataFinePeriodo, not even an explicit nil
func (o *DettaglioLinee) UnsetDataFinePeriodo() {
	o.DataFinePeriodo.Unset()
}

// GetPrezzoUnitario returns the PrezzoUnitario field value if set, zero value otherwise.
func (o *DettaglioLinee) GetPrezzoUnitario() float64 {
	if o == nil || IsNil(o.PrezzoUnitario) {
		var ret float64
		return ret
	}
	return *o.PrezzoUnitario
}

// GetPrezzoUnitarioOk returns a tuple with the PrezzoUnitario field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DettaglioLinee) GetPrezzoUnitarioOk() (*float64, bool) {
	if o == nil || IsNil(o.PrezzoUnitario) {
		return nil, false
	}
	return o.PrezzoUnitario, true
}

// HasPrezzoUnitario returns a boolean if a field has been set.
func (o *DettaglioLinee) HasPrezzoUnitario() bool {
	if o != nil && !IsNil(o.PrezzoUnitario) {
		return true
	}

	return false
}

// SetPrezzoUnitario gets a reference to the given float64 and assigns it to the PrezzoUnitario field.
func (o *DettaglioLinee) SetPrezzoUnitario(v float64) {
	o.PrezzoUnitario = &v
}

// GetScontoMaggiorazione returns the ScontoMaggiorazione field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioLinee) GetScontoMaggiorazione() []ScontoMaggiorazione {
	if o == nil {
		var ret []ScontoMaggiorazione
		return ret
	}
	return o.ScontoMaggiorazione
}

// GetScontoMaggiorazioneOk returns a tuple with the ScontoMaggiorazione field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioLinee) GetScontoMaggiorazioneOk() ([]ScontoMaggiorazione, bool) {
	if o == nil || IsNil(o.ScontoMaggiorazione) {
		return nil, false
	}
	return o.ScontoMaggiorazione, true
}

// HasScontoMaggiorazione returns a boolean if a field has been set.
func (o *DettaglioLinee) HasScontoMaggiorazione() bool {
	if o != nil && !IsNil(o.ScontoMaggiorazione) {
		return true
	}

	return false
}

// SetScontoMaggiorazione gets a reference to the given []ScontoMaggiorazione and assigns it to the ScontoMaggiorazione field.
func (o *DettaglioLinee) SetScontoMaggiorazione(v []ScontoMaggiorazione) {
	o.ScontoMaggiorazione = v
}

// GetPrezzoTotale returns the PrezzoTotale field value if set, zero value otherwise.
func (o *DettaglioLinee) GetPrezzoTotale() float64 {
	if o == nil || IsNil(o.PrezzoTotale) {
		var ret float64
		return ret
	}
	return *o.PrezzoTotale
}

// GetPrezzoTotaleOk returns a tuple with the PrezzoTotale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DettaglioLinee) GetPrezzoTotaleOk() (*float64, bool) {
	if o == nil || IsNil(o.PrezzoTotale) {
		return nil, false
	}
	return o.PrezzoTotale, true
}

// HasPrezzoTotale returns a boolean if a field has been set.
func (o *DettaglioLinee) HasPrezzoTotale() bool {
	if o != nil && !IsNil(o.PrezzoTotale) {
		return true
	}

	return false
}

// SetPrezzoTotale gets a reference to the given float64 and assigns it to the PrezzoTotale field.
func (o *DettaglioLinee) SetPrezzoTotale(v float64) {
	o.PrezzoTotale = &v
}

// GetAliquotaIva returns the AliquotaIva field value if set, zero value otherwise.
func (o *DettaglioLinee) GetAliquotaIva() float64 {
	if o == nil || IsNil(o.AliquotaIva) {
		var ret float64
		return ret
	}
	return *o.AliquotaIva
}

// GetAliquotaIvaOk returns a tuple with the AliquotaIva field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DettaglioLinee) GetAliquotaIvaOk() (*float64, bool) {
	if o == nil || IsNil(o.AliquotaIva) {
		return nil, false
	}
	return o.AliquotaIva, true
}

// HasAliquotaIva returns a boolean if a field has been set.
func (o *DettaglioLinee) HasAliquotaIva() bool {
	if o != nil && !IsNil(o.AliquotaIva) {
		return true
	}

	return false
}

// SetAliquotaIva gets a reference to the given float64 and assigns it to the AliquotaIva field.
func (o *DettaglioLinee) SetAliquotaIva(v float64) {
	o.AliquotaIva = &v
}

// GetRitenuta returns the Ritenuta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioLinee) GetRitenuta() string {
	if o == nil || IsNil(o.Ritenuta.Get()) {
		var ret string
		return ret
	}
	return *o.Ritenuta.Get()
}

// GetRitenutaOk returns a tuple with the Ritenuta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioLinee) GetRitenutaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ritenuta.Get(), o.Ritenuta.IsSet()
}

// HasRitenuta returns a boolean if a field has been set.
func (o *DettaglioLinee) HasRitenuta() bool {
	if o != nil && o.Ritenuta.IsSet() {
		return true
	}

	return false
}

// SetRitenuta gets a reference to the given NullableString and assigns it to the Ritenuta field.
func (o *DettaglioLinee) SetRitenuta(v string) {
	o.Ritenuta.Set(&v)
}
// SetRitenutaNil sets the value for Ritenuta to be an explicit nil
func (o *DettaglioLinee) SetRitenutaNil() {
	o.Ritenuta.Set(nil)
}

// UnsetRitenuta ensures that no value is present for Ritenuta, not even an explicit nil
func (o *DettaglioLinee) UnsetRitenuta() {
	o.Ritenuta.Unset()
}

// GetNatura returns the Natura field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioLinee) GetNatura() string {
	if o == nil || IsNil(o.Natura.Get()) {
		var ret string
		return ret
	}
	return *o.Natura.Get()
}

// GetNaturaOk returns a tuple with the Natura field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioLinee) GetNaturaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Natura.Get(), o.Natura.IsSet()
}

// HasNatura returns a boolean if a field has been set.
func (o *DettaglioLinee) HasNatura() bool {
	if o != nil && o.Natura.IsSet() {
		return true
	}

	return false
}

// SetNatura gets a reference to the given NullableString and assigns it to the Natura field.
func (o *DettaglioLinee) SetNatura(v string) {
	o.Natura.Set(&v)
}
// SetNaturaNil sets the value for Natura to be an explicit nil
func (o *DettaglioLinee) SetNaturaNil() {
	o.Natura.Set(nil)
}

// UnsetNatura ensures that no value is present for Natura, not even an explicit nil
func (o *DettaglioLinee) UnsetNatura() {
	o.Natura.Unset()
}

// GetRiferimentoAmministrazione returns the RiferimentoAmministrazione field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioLinee) GetRiferimentoAmministrazione() string {
	if o == nil || IsNil(o.RiferimentoAmministrazione.Get()) {
		var ret string
		return ret
	}
	return *o.RiferimentoAmministrazione.Get()
}

// GetRiferimentoAmministrazioneOk returns a tuple with the RiferimentoAmministrazione field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioLinee) GetRiferimentoAmministrazioneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RiferimentoAmministrazione.Get(), o.RiferimentoAmministrazione.IsSet()
}

// HasRiferimentoAmministrazione returns a boolean if a field has been set.
func (o *DettaglioLinee) HasRiferimentoAmministrazione() bool {
	if o != nil && o.RiferimentoAmministrazione.IsSet() {
		return true
	}

	return false
}

// SetRiferimentoAmministrazione gets a reference to the given NullableString and assigns it to the RiferimentoAmministrazione field.
func (o *DettaglioLinee) SetRiferimentoAmministrazione(v string) {
	o.RiferimentoAmministrazione.Set(&v)
}
// SetRiferimentoAmministrazioneNil sets the value for RiferimentoAmministrazione to be an explicit nil
func (o *DettaglioLinee) SetRiferimentoAmministrazioneNil() {
	o.RiferimentoAmministrazione.Set(nil)
}

// UnsetRiferimentoAmministrazione ensures that no value is present for RiferimentoAmministrazione, not even an explicit nil
func (o *DettaglioLinee) UnsetRiferimentoAmministrazione() {
	o.RiferimentoAmministrazione.Unset()
}

// GetAltriDatiGestionali returns the AltriDatiGestionali field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DettaglioLinee) GetAltriDatiGestionali() []AltriDatiGestionali {
	if o == nil {
		var ret []AltriDatiGestionali
		return ret
	}
	return o.AltriDatiGestionali
}

// GetAltriDatiGestionaliOk returns a tuple with the AltriDatiGestionali field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DettaglioLinee) GetAltriDatiGestionaliOk() ([]AltriDatiGestionali, bool) {
	if o == nil || IsNil(o.AltriDatiGestionali) {
		return nil, false
	}
	return o.AltriDatiGestionali, true
}

// HasAltriDatiGestionali returns a boolean if a field has been set.
func (o *DettaglioLinee) HasAltriDatiGestionali() bool {
	if o != nil && !IsNil(o.AltriDatiGestionali) {
		return true
	}

	return false
}

// SetAltriDatiGestionali gets a reference to the given []AltriDatiGestionali and assigns it to the AltriDatiGestionali field.
func (o *DettaglioLinee) SetAltriDatiGestionali(v []AltriDatiGestionali) {
	o.AltriDatiGestionali = v
}

func (o DettaglioLinee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DettaglioLinee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumeroLinea) {
		toSerialize["numero_linea"] = o.NumeroLinea
	}
	if o.TipoCessionePrestazione.IsSet() {
		toSerialize["tipo_cessione_prestazione"] = o.TipoCessionePrestazione.Get()
	}
	if o.CodiceArticolo != nil {
		toSerialize["codice_articolo"] = o.CodiceArticolo
	}
	if o.Descrizione.IsSet() {
		toSerialize["descrizione"] = o.Descrizione.Get()
	}
	if o.Quantita.IsSet() {
		toSerialize["quantita"] = o.Quantita.Get()
	}
	if o.UnitaMisura.IsSet() {
		toSerialize["unita_misura"] = o.UnitaMisura.Get()
	}
	if o.DataInizioPeriodo.IsSet() {
		toSerialize["data_inizio_periodo"] = o.DataInizioPeriodo.Get()
	}
	if o.DataFinePeriodo.IsSet() {
		toSerialize["data_fine_periodo"] = o.DataFinePeriodo.Get()
	}
	if !IsNil(o.PrezzoUnitario) {
		toSerialize["prezzo_unitario"] = o.PrezzoUnitario
	}
	if o.ScontoMaggiorazione != nil {
		toSerialize["sconto_maggiorazione"] = o.ScontoMaggiorazione
	}
	if !IsNil(o.PrezzoTotale) {
		toSerialize["prezzo_totale"] = o.PrezzoTotale
	}
	if !IsNil(o.AliquotaIva) {
		toSerialize["aliquota_iva"] = o.AliquotaIva
	}
	if o.Ritenuta.IsSet() {
		toSerialize["ritenuta"] = o.Ritenuta.Get()
	}
	if o.Natura.IsSet() {
		toSerialize["natura"] = o.Natura.Get()
	}
	if o.RiferimentoAmministrazione.IsSet() {
		toSerialize["riferimento_amministrazione"] = o.RiferimentoAmministrazione.Get()
	}
	if o.AltriDatiGestionali != nil {
		toSerialize["altri_dati_gestionali"] = o.AltriDatiGestionali
	}
	return toSerialize, nil
}

type NullableDettaglioLinee struct {
	value *DettaglioLinee
	isSet bool
}

func (v NullableDettaglioLinee) Get() *DettaglioLinee {
	return v.value
}

func (v *NullableDettaglioLinee) Set(val *DettaglioLinee) {
	v.value = val
	v.isSet = true
}

func (v NullableDettaglioLinee) IsSet() bool {
	return v.isSet
}

func (v *NullableDettaglioLinee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDettaglioLinee(val *DettaglioLinee) *NullableDettaglioLinee {
	return &NullableDettaglioLinee{value: val, isSet: true}
}

func (v NullableDettaglioLinee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDettaglioLinee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


