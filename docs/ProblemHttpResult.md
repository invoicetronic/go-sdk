# ProblemHttpResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProblemDetails** | Pointer to [**ProblemDetails**](ProblemDetails.md) |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] [readonly] 
**StatusCode** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewProblemHttpResult

`func NewProblemHttpResult() *ProblemHttpResult`

NewProblemHttpResult instantiates a new ProblemHttpResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemHttpResultWithDefaults

`func NewProblemHttpResultWithDefaults() *ProblemHttpResult`

NewProblemHttpResultWithDefaults instantiates a new ProblemHttpResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProblemDetails

`func (o *ProblemHttpResult) GetProblemDetails() ProblemDetails`

GetProblemDetails returns the ProblemDetails field if non-nil, zero value otherwise.

### GetProblemDetailsOk

`func (o *ProblemHttpResult) GetProblemDetailsOk() (*ProblemDetails, bool)`

GetProblemDetailsOk returns a tuple with the ProblemDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemDetails

`func (o *ProblemHttpResult) SetProblemDetails(v ProblemDetails)`

SetProblemDetails sets ProblemDetails field to given value.

### HasProblemDetails

`func (o *ProblemHttpResult) HasProblemDetails() bool`

HasProblemDetails returns a boolean if a field has been set.

### GetContentType

`func (o *ProblemHttpResult) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ProblemHttpResult) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ProblemHttpResult) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ProblemHttpResult) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *ProblemHttpResult) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *ProblemHttpResult) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetStatusCode

`func (o *ProblemHttpResult) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ProblemHttpResult) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ProblemHttpResult) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ProblemHttpResult) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


