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

// checks if the WebHookHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebHookHistory{}

// WebHookHistory Webhook history.
type WebHookHistory struct {
	// Unique identifier. Leave it at 0 for new records as it will be set automatically.
	Id *int32 `json:"id,omitempty"`
	// Creation date. It is set automatically.
	Created *time.Time `json:"created,omitempty"`
	// Row version, for optimistic concurrency. It is set automatically.
	Version *int32 `json:"version,omitempty"`
	// Webhook id.
	WebHookId *int32 `json:"web_hook_id,omitempty"`
	// User id.
	UserId *int32 `json:"user_id,omitempty"`
	// Event name.
	Event NullableString `json:"event,omitempty"`
	// Status code.
	StatusCode *int32 `json:"status_code,omitempty"`
	// Webhook request body.
	RequestBody NullableString `json:"request_body,omitempty"`
	// Webhook response body.
	ResponseBody NullableString `json:"response_body,omitempty"`
	// Date and time of the request.
	DateTime *time.Time `json:"date_time,omitempty"`
	// Wether the request was successful.
	Success *bool `json:"success,omitempty"`
}

// NewWebHookHistory instantiates a new WebHookHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebHookHistory() *WebHookHistory {
	this := WebHookHistory{}
	return &this
}

// NewWebHookHistoryWithDefaults instantiates a new WebHookHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebHookHistoryWithDefaults() *WebHookHistory {
	this := WebHookHistory{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebHookHistory) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookHistory) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebHookHistory) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *WebHookHistory) SetId(v int32) {
	o.Id = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WebHookHistory) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookHistory) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WebHookHistory) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *WebHookHistory) SetCreated(v time.Time) {
	o.Created = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WebHookHistory) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookHistory) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WebHookHistory) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *WebHookHistory) SetVersion(v int32) {
	o.Version = &v
}

// GetWebHookId returns the WebHookId field value if set, zero value otherwise.
func (o *WebHookHistory) GetWebHookId() int32 {
	if o == nil || IsNil(o.WebHookId) {
		var ret int32
		return ret
	}
	return *o.WebHookId
}

// GetWebHookIdOk returns a tuple with the WebHookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookHistory) GetWebHookIdOk() (*int32, bool) {
	if o == nil || IsNil(o.WebHookId) {
		return nil, false
	}
	return o.WebHookId, true
}

// HasWebHookId returns a boolean if a field has been set.
func (o *WebHookHistory) HasWebHookId() bool {
	if o != nil && !IsNil(o.WebHookId) {
		return true
	}

	return false
}

// SetWebHookId gets a reference to the given int32 and assigns it to the WebHookId field.
func (o *WebHookHistory) SetWebHookId(v int32) {
	o.WebHookId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *WebHookHistory) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookHistory) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *WebHookHistory) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *WebHookHistory) SetUserId(v int32) {
	o.UserId = &v
}

// GetEvent returns the Event field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebHookHistory) GetEvent() string {
	if o == nil || IsNil(o.Event.Get()) {
		var ret string
		return ret
	}
	return *o.Event.Get()
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebHookHistory) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Event.Get(), o.Event.IsSet()
}

// HasEvent returns a boolean if a field has been set.
func (o *WebHookHistory) HasEvent() bool {
	if o != nil && o.Event.IsSet() {
		return true
	}

	return false
}

// SetEvent gets a reference to the given NullableString and assigns it to the Event field.
func (o *WebHookHistory) SetEvent(v string) {
	o.Event.Set(&v)
}
// SetEventNil sets the value for Event to be an explicit nil
func (o *WebHookHistory) SetEventNil() {
	o.Event.Set(nil)
}

// UnsetEvent ensures that no value is present for Event, not even an explicit nil
func (o *WebHookHistory) UnsetEvent() {
	o.Event.Unset()
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *WebHookHistory) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookHistory) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *WebHookHistory) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *WebHookHistory) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetRequestBody returns the RequestBody field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebHookHistory) GetRequestBody() string {
	if o == nil || IsNil(o.RequestBody.Get()) {
		var ret string
		return ret
	}
	return *o.RequestBody.Get()
}

// GetRequestBodyOk returns a tuple with the RequestBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebHookHistory) GetRequestBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestBody.Get(), o.RequestBody.IsSet()
}

// HasRequestBody returns a boolean if a field has been set.
func (o *WebHookHistory) HasRequestBody() bool {
	if o != nil && o.RequestBody.IsSet() {
		return true
	}

	return false
}

// SetRequestBody gets a reference to the given NullableString and assigns it to the RequestBody field.
func (o *WebHookHistory) SetRequestBody(v string) {
	o.RequestBody.Set(&v)
}
// SetRequestBodyNil sets the value for RequestBody to be an explicit nil
func (o *WebHookHistory) SetRequestBodyNil() {
	o.RequestBody.Set(nil)
}

// UnsetRequestBody ensures that no value is present for RequestBody, not even an explicit nil
func (o *WebHookHistory) UnsetRequestBody() {
	o.RequestBody.Unset()
}

// GetResponseBody returns the ResponseBody field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebHookHistory) GetResponseBody() string {
	if o == nil || IsNil(o.ResponseBody.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseBody.Get()
}

// GetResponseBodyOk returns a tuple with the ResponseBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebHookHistory) GetResponseBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseBody.Get(), o.ResponseBody.IsSet()
}

// HasResponseBody returns a boolean if a field has been set.
func (o *WebHookHistory) HasResponseBody() bool {
	if o != nil && o.ResponseBody.IsSet() {
		return true
	}

	return false
}

// SetResponseBody gets a reference to the given NullableString and assigns it to the ResponseBody field.
func (o *WebHookHistory) SetResponseBody(v string) {
	o.ResponseBody.Set(&v)
}
// SetResponseBodyNil sets the value for ResponseBody to be an explicit nil
func (o *WebHookHistory) SetResponseBodyNil() {
	o.ResponseBody.Set(nil)
}

// UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil
func (o *WebHookHistory) UnsetResponseBody() {
	o.ResponseBody.Unset()
}

// GetDateTime returns the DateTime field value if set, zero value otherwise.
func (o *WebHookHistory) GetDateTime() time.Time {
	if o == nil || IsNil(o.DateTime) {
		var ret time.Time
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookHistory) GetDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateTime) {
		return nil, false
	}
	return o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *WebHookHistory) HasDateTime() bool {
	if o != nil && !IsNil(o.DateTime) {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given time.Time and assigns it to the DateTime field.
func (o *WebHookHistory) SetDateTime(v time.Time) {
	o.DateTime = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *WebHookHistory) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookHistory) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *WebHookHistory) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *WebHookHistory) SetSuccess(v bool) {
	o.Success = &v
}

func (o WebHookHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebHookHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.WebHookId) {
		toSerialize["web_hook_id"] = o.WebHookId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if o.Event.IsSet() {
		toSerialize["event"] = o.Event.Get()
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	if o.RequestBody.IsSet() {
		toSerialize["request_body"] = o.RequestBody.Get()
	}
	if o.ResponseBody.IsSet() {
		toSerialize["response_body"] = o.ResponseBody.Get()
	}
	if !IsNil(o.DateTime) {
		toSerialize["date_time"] = o.DateTime
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableWebHookHistory struct {
	value *WebHookHistory
	isSet bool
}

func (v NullableWebHookHistory) Get() *WebHookHistory {
	return v.value
}

func (v *NullableWebHookHistory) Set(val *WebHookHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableWebHookHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableWebHookHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebHookHistory(val *WebHookHistory) *NullableWebHookHistory {
	return &NullableWebHookHistory{value: val, isSet: true}
}

func (v NullableWebHookHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebHookHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


