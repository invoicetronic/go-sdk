# SubKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Human-readable label, e.g. the name of the tenant the key is for. | 
**Active** | Pointer to **bool** | Whether the key can authenticate. Defaults to true. | [optional] 
**Permissions** | Pointer to [**Permissions**](Permissions.md) |  | [optional] 
**CompanyIds** | Pointer to **[]int32** | Companies the key can access. When omitted or empty, the key can access all the companies of the account, including the ones created later. | [optional] 
**CorsOrigins** | Pointer to **[]string** | Browser origins allowed to call the API with this key (CORS), e.g. &#x60;https://app.example.com&#x60; or &#x60;*.example.com&#x60;. A key used from a browser is public: keep its permissions and companies minimal. | [optional] 

## Methods

### NewSubKeyRequest

`func NewSubKeyRequest(description string, ) *SubKeyRequest`

NewSubKeyRequest instantiates a new SubKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubKeyRequestWithDefaults

`func NewSubKeyRequestWithDefaults() *SubKeyRequest`

NewSubKeyRequestWithDefaults instantiates a new SubKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SubKeyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubKeyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubKeyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetActive

`func (o *SubKeyRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SubKeyRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SubKeyRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SubKeyRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPermissions

`func (o *SubKeyRequest) GetPermissions() Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SubKeyRequest) GetPermissionsOk() (*Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SubKeyRequest) SetPermissions(v Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SubKeyRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetCompanyIds

`func (o *SubKeyRequest) GetCompanyIds() []int32`

GetCompanyIds returns the CompanyIds field if non-nil, zero value otherwise.

### GetCompanyIdsOk

`func (o *SubKeyRequest) GetCompanyIdsOk() (*[]int32, bool)`

GetCompanyIdsOk returns a tuple with the CompanyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIds

`func (o *SubKeyRequest) SetCompanyIds(v []int32)`

SetCompanyIds sets CompanyIds field to given value.

### HasCompanyIds

`func (o *SubKeyRequest) HasCompanyIds() bool`

HasCompanyIds returns a boolean if a field has been set.

### SetCompanyIdsNil

`func (o *SubKeyRequest) SetCompanyIdsNil(b bool)`

 SetCompanyIdsNil sets the value for CompanyIds to be an explicit nil

### UnsetCompanyIds
`func (o *SubKeyRequest) UnsetCompanyIds()`

UnsetCompanyIds ensures that no value is present for CompanyIds, not even an explicit nil
### GetCorsOrigins

`func (o *SubKeyRequest) GetCorsOrigins() []string`

GetCorsOrigins returns the CorsOrigins field if non-nil, zero value otherwise.

### GetCorsOriginsOk

`func (o *SubKeyRequest) GetCorsOriginsOk() (*[]string, bool)`

GetCorsOriginsOk returns a tuple with the CorsOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsOrigins

`func (o *SubKeyRequest) SetCorsOrigins(v []string)`

SetCorsOrigins sets CorsOrigins field to given value.

### HasCorsOrigins

`func (o *SubKeyRequest) HasCorsOrigins() bool`

HasCorsOrigins returns a boolean if a field has been set.

### SetCorsOriginsNil

`func (o *SubKeyRequest) SetCorsOriginsNil(b bool)`

 SetCorsOriginsNil sets the value for CorsOrigins to be an explicit nil

### UnsetCorsOrigins
`func (o *SubKeyRequest) UnsetCorsOrigins()`

UnsetCorsOrigins ensures that no value is present for CorsOrigins, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


