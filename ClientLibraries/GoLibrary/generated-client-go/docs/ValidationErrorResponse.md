# ValidationErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ValidationErrorResponseErrorsInner**](ValidationErrorResponseErrorsInner.md) |  | 

## Methods

### NewValidationErrorResponse

`func NewValidationErrorResponse(errors []ValidationErrorResponseErrorsInner, ) *ValidationErrorResponse`

NewValidationErrorResponse instantiates a new ValidationErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorResponseWithDefaults

`func NewValidationErrorResponseWithDefaults() *ValidationErrorResponse`

NewValidationErrorResponseWithDefaults instantiates a new ValidationErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ValidationErrorResponse) GetErrors() []ValidationErrorResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidationErrorResponse) GetErrorsOk() (*[]ValidationErrorResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidationErrorResponse) SetErrors(v []ValidationErrorResponseErrorsInner)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


