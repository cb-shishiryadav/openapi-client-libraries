# HealthCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | The description of the health status returned by the Service Adapter. | [optional] 
**Status** | [**HealthStatus**](HealthStatus.md) |  | 
**Components** | [**[]HealthCheckComponent**](HealthCheckComponent.md) | List of health status details for each component reported by the Service Adapter. | 
**Time** | **time.Time** | The timestamp of the health status reported by the Service Adapter. | 

## Methods

### NewHealthCheckResponse

`func NewHealthCheckResponse(status HealthStatus, components []HealthCheckComponent, time time.Time, ) *HealthCheckResponse`

NewHealthCheckResponse instantiates a new HealthCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckResponseWithDefaults

`func NewHealthCheckResponseWithDefaults() *HealthCheckResponse`

NewHealthCheckResponseWithDefaults instantiates a new HealthCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *HealthCheckResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HealthCheckResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HealthCheckResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HealthCheckResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDescription

`func (o *HealthCheckResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HealthCheckResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HealthCheckResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HealthCheckResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *HealthCheckResponse) GetStatus() HealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheckResponse) GetStatusOk() (*HealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheckResponse) SetStatus(v HealthStatus)`

SetStatus sets Status field to given value.


### GetComponents

`func (o *HealthCheckResponse) GetComponents() []HealthCheckComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *HealthCheckResponse) GetComponentsOk() (*[]HealthCheckComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *HealthCheckResponse) SetComponents(v []HealthCheckComponent)`

SetComponents sets Components field to given value.


### GetTime

`func (o *HealthCheckResponse) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *HealthCheckResponse) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *HealthCheckResponse) SetTime(v time.Time)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


