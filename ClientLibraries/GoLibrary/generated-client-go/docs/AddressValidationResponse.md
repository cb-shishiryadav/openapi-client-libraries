# AddressValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**AddressValidationStatus**](AddressValidationStatus.md) |  | 

## Methods

### NewAddressValidationResponse

`func NewAddressValidationResponse(status AddressValidationStatus, ) *AddressValidationResponse`

NewAddressValidationResponse instantiates a new AddressValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressValidationResponseWithDefaults

`func NewAddressValidationResponseWithDefaults() *AddressValidationResponse`

NewAddressValidationResponseWithDefaults instantiates a new AddressValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AddressValidationResponse) GetStatus() AddressValidationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddressValidationResponse) GetStatusOk() (*AddressValidationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddressValidationResponse) SetStatus(v AddressValidationStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


