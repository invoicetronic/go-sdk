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
)

// checks if the ProblemHttpResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProblemHttpResult{}

// ProblemHttpResult struct for ProblemHttpResult
type ProblemHttpResult struct {
	ProblemDetails *ProblemDetails `json:"problem_details,omitempty"`
	ContentType NullableString `json:"content_type,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewProblemHttpResult instantiates a new ProblemHttpResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemHttpResult() *ProblemHttpResult {
	this := ProblemHttpResult{}
	return &this
}

// NewProblemHttpResultWithDefaults instantiates a new ProblemHttpResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemHttpResultWithDefaults() *ProblemHttpResult {
	this := ProblemHttpResult{}
	return &this
}

// GetProblemDetails returns the ProblemDetails field value if set, zero value otherwise.
func (o *ProblemHttpResult) GetProblemDetails() ProblemDetails {
	if o == nil || IsNil(o.ProblemDetails) {
		var ret ProblemDetails
		return ret
	}
	return *o.ProblemDetails
}

// GetProblemDetailsOk returns a tuple with the ProblemDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemHttpResult) GetProblemDetailsOk() (*ProblemDetails, bool) {
	if o == nil || IsNil(o.ProblemDetails) {
		return nil, false
	}
	return o.ProblemDetails, true
}

// HasProblemDetails returns a boolean if a field has been set.
func (o *ProblemHttpResult) HasProblemDetails() bool {
	if o != nil && !IsNil(o.ProblemDetails) {
		return true
	}

	return false
}

// SetProblemDetails gets a reference to the given ProblemDetails and assigns it to the ProblemDetails field.
func (o *ProblemHttpResult) SetProblemDetails(v ProblemDetails) {
	o.ProblemDetails = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProblemHttpResult) GetContentType() string {
	if o == nil || IsNil(o.ContentType.Get()) {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProblemHttpResult) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *ProblemHttpResult) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *ProblemHttpResult) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *ProblemHttpResult) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *ProblemHttpResult) UnsetContentType() {
	o.ContentType.Unset()
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ProblemHttpResult) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemHttpResult) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ProblemHttpResult) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ProblemHttpResult) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ProblemHttpResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProblemHttpResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProblemDetails) {
		toSerialize["problem_details"] = o.ProblemDetails
	}
	if o.ContentType.IsSet() {
		toSerialize["content_type"] = o.ContentType.Get()
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	return toSerialize, nil
}

type NullableProblemHttpResult struct {
	value *ProblemHttpResult
	isSet bool
}

func (v NullableProblemHttpResult) Get() *ProblemHttpResult {
	return v.value
}

func (v *NullableProblemHttpResult) Set(val *ProblemHttpResult) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemHttpResult) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemHttpResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemHttpResult(val *ProblemHttpResult) *NullableProblemHttpResult {
	return &NullableProblemHttpResult{value: val, isSet: true}
}

func (v NullableProblemHttpResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemHttpResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


