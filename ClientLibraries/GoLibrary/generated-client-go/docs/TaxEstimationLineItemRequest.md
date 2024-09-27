# TaxEstimationLineItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** | Index or serial number of the line item. | 
**ItemCode** | Pointer to **string** | The unique identifier (in Chargebee) of the product corresponding to the line item. If the line item corresponds to a one-off charge, then this identifier is not provided. | [optional] 
**Description** | Pointer to **string** | The description of the line item. | [optional] 
**Quantity** | Pointer to **float64** | The quantity associated with this line item. | [optional] 
**UnitPrice** | Pointer to **float64** | The unit price for this line item. In case of [tiered pricing](https://www.chargebee.com/docs/1.0/plans.html#tiered-pricing) where the unit price varies for each quantity tier, this is the average unit price. | [optional] 
**Amount** | **float64** | The amount for this line item. This is &#x60;unitPrice&#x60; × &#x60;quantity&#x60;. | 
**DiscountAmount** | Pointer to **float64** | The discount applied to this line item. | [optional] 
**IsTaxInclusive** | **bool** | Indicates whether (&#x60;amount&#x60; - &#x60;discountAmount&#x60;)  is inclusive of taxes. | 
**TaxIdentifiers** | Pointer to [**[]FieldItem**](FieldItem.md) | The tax code fields of the product used for tax calculation. | [optional] 

## Methods

### NewTaxEstimationLineItemRequest

`func NewTaxEstimationLineItemRequest(number int32, amount float64, isTaxInclusive bool, ) *TaxEstimationLineItemRequest`

NewTaxEstimationLineItemRequest instantiates a new TaxEstimationLineItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxEstimationLineItemRequestWithDefaults

`func NewTaxEstimationLineItemRequestWithDefaults() *TaxEstimationLineItemRequest`

NewTaxEstimationLineItemRequestWithDefaults instantiates a new TaxEstimationLineItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *TaxEstimationLineItemRequest) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TaxEstimationLineItemRequest) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TaxEstimationLineItemRequest) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetItemCode

`func (o *TaxEstimationLineItemRequest) GetItemCode() string`

GetItemCode returns the ItemCode field if non-nil, zero value otherwise.

### GetItemCodeOk

`func (o *TaxEstimationLineItemRequest) GetItemCodeOk() (*string, bool)`

GetItemCodeOk returns a tuple with the ItemCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCode

`func (o *TaxEstimationLineItemRequest) SetItemCode(v string)`

SetItemCode sets ItemCode field to given value.

### HasItemCode

`func (o *TaxEstimationLineItemRequest) HasItemCode() bool`

HasItemCode returns a boolean if a field has been set.

### GetDescription

`func (o *TaxEstimationLineItemRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaxEstimationLineItemRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaxEstimationLineItemRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaxEstimationLineItemRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *TaxEstimationLineItemRequest) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *TaxEstimationLineItemRequest) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *TaxEstimationLineItemRequest) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *TaxEstimationLineItemRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitPrice

`func (o *TaxEstimationLineItemRequest) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *TaxEstimationLineItemRequest) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *TaxEstimationLineItemRequest) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *TaxEstimationLineItemRequest) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetAmount

`func (o *TaxEstimationLineItemRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TaxEstimationLineItemRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TaxEstimationLineItemRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetDiscountAmount

`func (o *TaxEstimationLineItemRequest) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *TaxEstimationLineItemRequest) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *TaxEstimationLineItemRequest) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *TaxEstimationLineItemRequest) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetIsTaxInclusive

`func (o *TaxEstimationLineItemRequest) GetIsTaxInclusive() bool`

GetIsTaxInclusive returns the IsTaxInclusive field if non-nil, zero value otherwise.

### GetIsTaxInclusiveOk

`func (o *TaxEstimationLineItemRequest) GetIsTaxInclusiveOk() (*bool, bool)`

GetIsTaxInclusiveOk returns a tuple with the IsTaxInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaxInclusive

`func (o *TaxEstimationLineItemRequest) SetIsTaxInclusive(v bool)`

SetIsTaxInclusive sets IsTaxInclusive field to given value.


### GetTaxIdentifiers

`func (o *TaxEstimationLineItemRequest) GetTaxIdentifiers() []FieldItem`

GetTaxIdentifiers returns the TaxIdentifiers field if non-nil, zero value otherwise.

### GetTaxIdentifiersOk

`func (o *TaxEstimationLineItemRequest) GetTaxIdentifiersOk() (*[]FieldItem, bool)`

GetTaxIdentifiersOk returns a tuple with the TaxIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentifiers

`func (o *TaxEstimationLineItemRequest) SetTaxIdentifiers(v []FieldItem)`

SetTaxIdentifiers sets TaxIdentifiers field to given value.

### HasTaxIdentifiers

`func (o *TaxEstimationLineItemRequest) HasTaxIdentifiers() bool`

HasTaxIdentifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


