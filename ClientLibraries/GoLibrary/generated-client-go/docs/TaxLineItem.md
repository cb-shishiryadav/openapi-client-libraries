# TaxLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** | Index or serial number of this tax line item. | 
**Jurisdiction** | [**TaxJurisdiction**](TaxJurisdiction.md) |  | 
**Name** | **string** | The name of the tax applied. | 
**Rate** | **float64** | The tax rate expressed in percentage. | 
**TaxableAmount** | **float64** | The part of the line item&#39;s &#x60;subtotal&#x60; that is taxable under this jurisdiction. | 
**TaxAmount** | **float64** | The tax payable for the line item under this jurisdiction. | 

## Methods

### NewTaxLineItem

`func NewTaxLineItem(number int32, jurisdiction TaxJurisdiction, name string, rate float64, taxableAmount float64, taxAmount float64, ) *TaxLineItem`

NewTaxLineItem instantiates a new TaxLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxLineItemWithDefaults

`func NewTaxLineItemWithDefaults() *TaxLineItem`

NewTaxLineItemWithDefaults instantiates a new TaxLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *TaxLineItem) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TaxLineItem) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TaxLineItem) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetJurisdiction

`func (o *TaxLineItem) GetJurisdiction() TaxJurisdiction`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *TaxLineItem) GetJurisdictionOk() (*TaxJurisdiction, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *TaxLineItem) SetJurisdiction(v TaxJurisdiction)`

SetJurisdiction sets Jurisdiction field to given value.


### GetName

`func (o *TaxLineItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxLineItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxLineItem) SetName(v string)`

SetName sets Name field to given value.


### GetRate

`func (o *TaxLineItem) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *TaxLineItem) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *TaxLineItem) SetRate(v float64)`

SetRate sets Rate field to given value.


### GetTaxableAmount

`func (o *TaxLineItem) GetTaxableAmount() float64`

GetTaxableAmount returns the TaxableAmount field if non-nil, zero value otherwise.

### GetTaxableAmountOk

`func (o *TaxLineItem) GetTaxableAmountOk() (*float64, bool)`

GetTaxableAmountOk returns a tuple with the TaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount

`func (o *TaxLineItem) SetTaxableAmount(v float64)`

SetTaxableAmount sets TaxableAmount field to given value.


### GetTaxAmount

`func (o *TaxLineItem) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *TaxLineItem) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *TaxLineItem) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


