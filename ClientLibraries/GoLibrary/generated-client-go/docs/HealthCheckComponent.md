# HealthCheckComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | 
**Name** | **string** | The name of the component. | 
**Type** | **string** | The type of component affected when &#x60;status&#x60; is &#x60;WARN&#x60; or &#x60;DOWN&#x60;. The possible values are: - &#x60;ADAPTER&#x60;: The reported status is for the Service Adapter. - &#x60;API&#x60;: The reported status is for the Service Provider. - &#x60;DATABASE&#x60;: The reported status is for a database dependency of the Service Provider. - &#x60;SYSTEM&#x60;: The reported status is for any other known system component such as cache or gateway. - &#x60;OTHER&#x60;: The reported status is either for a component that does not belong to the types described above or the source of the issue is unknown.  | 
**Description** | Pointer to **string** | The detailed status of the component. | [optional] 
**Status** | [**HealthStatus**](HealthStatus.md) |  | 
**Endpoints** | Pointer to **[]string** | When the &#x60;status&#x60; of the component is not &#x60;UP&#x60;, then the list of endpoints affected. | [optional] 

## Methods

### NewHealthCheckComponent

`func NewHealthCheckComponent(id string, name string, type_ string, status HealthStatus, ) *HealthCheckComponent`

NewHealthCheckComponent instantiates a new HealthCheckComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckComponentWithDefaults

`func NewHealthCheckComponentWithDefaults() *HealthCheckComponent`

NewHealthCheckComponentWithDefaults instantiates a new HealthCheckComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HealthCheckComponent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HealthCheckComponent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HealthCheckComponent) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *HealthCheckComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthCheckComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthCheckComponent) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *HealthCheckComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthCheckComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthCheckComponent) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *HealthCheckComponent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HealthCheckComponent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HealthCheckComponent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HealthCheckComponent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *HealthCheckComponent) GetStatus() HealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheckComponent) GetStatusOk() (*HealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheckComponent) SetStatus(v HealthStatus)`

SetStatus sets Status field to given value.


### GetEndpoints

`func (o *HealthCheckComponent) GetEndpoints() []string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *HealthCheckComponent) GetEndpointsOk() (*[]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *HealthCheckComponent) SetEndpoints(v []string)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *HealthCheckComponent) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


