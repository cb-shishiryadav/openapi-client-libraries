# CredentialValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**CredentialStatus**](CredentialStatus.md) |  | 

## Methods

### NewCredentialValidationResponse

`func NewCredentialValidationResponse(status CredentialStatus, ) *CredentialValidationResponse`

NewCredentialValidationResponse instantiates a new CredentialValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialValidationResponseWithDefaults

`func NewCredentialValidationResponseWithDefaults() *CredentialValidationResponse`

NewCredentialValidationResponseWithDefaults instantiates a new CredentialValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CredentialValidationResponse) GetStatus() CredentialStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CredentialValidationResponse) GetStatusOk() (*CredentialStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CredentialValidationResponse) SetStatus(v CredentialStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


