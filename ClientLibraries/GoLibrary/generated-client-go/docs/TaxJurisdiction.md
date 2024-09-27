# TaxJurisdiction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The jurisdiction code. | 
**Type** | [**TaxJurisdictionType**](TaxJurisdictionType.md) |  | 
**Name** | **string** | The jurisdiction name. | 

## Methods

### NewTaxJurisdiction

`func NewTaxJurisdiction(code string, type_ TaxJurisdictionType, name string, ) *TaxJurisdiction`

NewTaxJurisdiction instantiates a new TaxJurisdiction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxJurisdictionWithDefaults

`func NewTaxJurisdictionWithDefaults() *TaxJurisdiction`

NewTaxJurisdictionWithDefaults instantiates a new TaxJurisdiction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TaxJurisdiction) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaxJurisdiction) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaxJurisdiction) SetCode(v string)`

SetCode sets Code field to given value.


### GetType

`func (o *TaxJurisdiction) GetType() TaxJurisdictionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxJurisdiction) GetTypeOk() (*TaxJurisdictionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxJurisdiction) SetType(v TaxJurisdictionType)`

SetType sets Type field to given value.


### GetName

`func (o *TaxJurisdiction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxJurisdiction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxJurisdiction) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


