# DocumentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **NullableString** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDocumentData

`func NewDocumentData() *DocumentData`

NewDocumentData instantiates a new DocumentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentDataWithDefaults

`func NewDocumentDataWithDefaults() *DocumentData`

NewDocumentDataWithDefaults instantiates a new DocumentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *DocumentData) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *DocumentData) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *DocumentData) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *DocumentData) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *DocumentData) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *DocumentData) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetDate

`func (o *DocumentData) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DocumentData) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DocumentData) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *DocumentData) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


