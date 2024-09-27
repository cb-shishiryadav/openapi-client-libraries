# ValidationErrorResponseErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to **string** | The target entity that has the invalid field or value. | [optional] 
**EntityField** | Pointer to **string** | The field of an entity that has the invalid value. | [optional] 
**Code** | [**ErrorCode**](ErrorCode.md) |  | 
**Message** | **string** | A short message describing the reason for the error. | 
**HelpUrl** | Pointer to **string** | The link to the documentation for more information about the error and the corrective action. | [optional] 

## Methods

### NewValidationErrorResponseErrorsInner

`func NewValidationErrorResponseErrorsInner(code ErrorCode, message string, ) *ValidationErrorResponseErrorsInner`

NewValidationErrorResponseErrorsInner instantiates a new ValidationErrorResponseErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorResponseErrorsInnerWithDefaults

`func NewValidationErrorResponseErrorsInnerWithDefaults() *ValidationErrorResponseErrorsInner`

NewValidationErrorResponseErrorsInnerWithDefaults instantiates a new ValidationErrorResponseErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *ValidationErrorResponseErrorsInner) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ValidationErrorResponseErrorsInner) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ValidationErrorResponseErrorsInner) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *ValidationErrorResponseErrorsInner) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetEntityField

`func (o *ValidationErrorResponseErrorsInner) GetEntityField() string`

GetEntityField returns the EntityField field if non-nil, zero value otherwise.

### GetEntityFieldOk

`func (o *ValidationErrorResponseErrorsInner) GetEntityFieldOk() (*string, bool)`

GetEntityFieldOk returns a tuple with the EntityField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityField

`func (o *ValidationErrorResponseErrorsInner) SetEntityField(v string)`

SetEntityField sets EntityField field to given value.

### HasEntityField

`func (o *ValidationErrorResponseErrorsInner) HasEntityField() bool`

HasEntityField returns a boolean if a field has been set.

### GetCode

`func (o *ValidationErrorResponseErrorsInner) GetCode() ErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ValidationErrorResponseErrorsInner) GetCodeOk() (*ErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ValidationErrorResponseErrorsInner) SetCode(v ErrorCode)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ValidationErrorResponseErrorsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationErrorResponseErrorsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationErrorResponseErrorsInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetHelpUrl

`func (o *ValidationErrorResponseErrorsInner) GetHelpUrl() string`

GetHelpUrl returns the HelpUrl field if non-nil, zero value otherwise.

### GetHelpUrlOk

`func (o *ValidationErrorResponseErrorsInner) GetHelpUrlOk() (*string, bool)`

GetHelpUrlOk returns a tuple with the HelpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpUrl

`func (o *ValidationErrorResponseErrorsInner) SetHelpUrl(v string)`

SetHelpUrl sets HelpUrl field to given value.

### HasHelpUrl

`func (o *ValidationErrorResponseErrorsInner) HasHelpUrl() bool`

HasHelpUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


