# BasicErrorResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The description of the error with details about it&#39;s cause. | 
**HelpUrl** | Pointer to **string** | The link to the documentation for more information about the error and the corrective action. | [optional] 

## Methods

### NewBasicErrorResponse1

`func NewBasicErrorResponse1(message string, ) *BasicErrorResponse1`

NewBasicErrorResponse1 instantiates a new BasicErrorResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicErrorResponse1WithDefaults

`func NewBasicErrorResponse1WithDefaults() *BasicErrorResponse1`

NewBasicErrorResponse1WithDefaults instantiates a new BasicErrorResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BasicErrorResponse1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BasicErrorResponse1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BasicErrorResponse1) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetHelpUrl

`func (o *BasicErrorResponse1) GetHelpUrl() string`

GetHelpUrl returns the HelpUrl field if non-nil, zero value otherwise.

### GetHelpUrlOk

`func (o *BasicErrorResponse1) GetHelpUrlOk() (*string, bool)`

GetHelpUrlOk returns a tuple with the HelpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpUrl

`func (o *BasicErrorResponse1) SetHelpUrl(v string)`

SetHelpUrl sets HelpUrl field to given value.

### HasHelpUrl

`func (o *BasicErrorResponse1) HasHelpUrl() bool`

HasHelpUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


