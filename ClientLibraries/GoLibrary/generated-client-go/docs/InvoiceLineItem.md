# InvoiceLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** | Index or serial number of the line item. | 
**ItemCode** | Pointer to **string** | The unique identifier (in Chargebee) of the product corresponding to the line item. If the line item corresponds to a one-off charge, then this identifier is not present. | [optional] 
**Description** | Pointer to **string** | The description of the line item. | [optional] 
**Quantity** | Pointer to **float64** | The quantity associated with this line item. | [optional] 
**UnitPrice** | Pointer to **float64** | The unit price for this line item. In case of [tiered pricing](https://www.chargebee.com/docs/1.0/plans.html#tiered-pricing) where the unit price varies for each quantity tier, this is the average unit price. | [optional] 
**Amount** | **float64** | The amount for this line item. This is &#x60;unitPrice&#x60; × &#x60;quantity&#x60;. | 
**Subtotal** | **float64** | The amount after discounts for this line item. This is &#x60;amount&#x60; - &#x60;discountAmount&#x60;. | 
**IsTaxInclusive** | **bool** | Indicates whether the &#x60;subtotal&#x60; for this line item is inclusive of taxes. | 
**IsTaxable** | **bool** | Indicates whether this line item is taxable. | 
**TaxIdentifiers** | Pointer to [**[]FieldItem**](FieldItem.md) | The tax code fields of the product used for tax calculation. | [optional] 
**TaxExemptType** | Pointer to [**TaxExemptType**](TaxExemptType.md) |  | [optional] 
**TaxExemptReason** | Pointer to **string** | The reason due to which a line item is exempted from tax. This is a mandatory parameter while applying tax exemption on any line-item. | [optional] 
**ExemptAmount** | **float64** | The part of this line item&#39;s &#x60;subtotal&#x60; that is exempted from tax. | 
**DiscountAmount** | **float64** | The discount applied to this line item. | 
**TaxableAmount** | **float64** | The part of this line item&#39;s &#x60;subtotal&#x60; that is taxable. | 
**TaxAmount** | **float64** | The tax payable for this line item. This is the sum of all &#x60;taxes.taxAmount&#x60; for this line item. | 
**Total** | **float64** | The total for this line item after discounts and taxes. This is the same as &#x60;subtotal&#x60; if it is tax inclusive; otherwise it is &#x60;subtotal&#x60; + &#x60;taxAmount&#x60;. &#x60;total&#x60; can also be expressed as &#x60;exemptAmount&#x60; + &#x60;taxableAmount&#x60; + &#x60;taxAmount&#x60;. | 
**IsPartialTax** | Pointer to **bool** | Indicates if taxes were applied only partially for this line item. | [optional] 
**Taxes** | [**[]TaxLineItem**](TaxLineItem.md) | List of taxes applied for this line item under each jurisdiction. | 

## Methods

### NewInvoiceLineItem

`func NewInvoiceLineItem(number int32, amount float64, subtotal float64, isTaxInclusive bool, isTaxable bool, exemptAmount float64, discountAmount float64, taxableAmount float64, taxAmount float64, total float64, taxes []TaxLineItem, ) *InvoiceLineItem`

NewInvoiceLineItem instantiates a new InvoiceLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineItemWithDefaults

`func NewInvoiceLineItemWithDefaults() *InvoiceLineItem`

NewInvoiceLineItemWithDefaults instantiates a new InvoiceLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *InvoiceLineItem) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InvoiceLineItem) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InvoiceLineItem) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetItemCode

`func (o *InvoiceLineItem) GetItemCode() string`

GetItemCode returns the ItemCode field if non-nil, zero value otherwise.

### GetItemCodeOk

`func (o *InvoiceLineItem) GetItemCodeOk() (*string, bool)`

GetItemCodeOk returns a tuple with the ItemCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCode

`func (o *InvoiceLineItem) SetItemCode(v string)`

SetItemCode sets ItemCode field to given value.

### HasItemCode

`func (o *InvoiceLineItem) HasItemCode() bool`

HasItemCode returns a boolean if a field has been set.

### GetDescription

`func (o *InvoiceLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceLineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *InvoiceLineItem) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceLineItem) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceLineItem) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceLineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitPrice

`func (o *InvoiceLineItem) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *InvoiceLineItem) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *InvoiceLineItem) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *InvoiceLineItem) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetAmount

`func (o *InvoiceLineItem) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceLineItem) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceLineItem) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetSubtotal

`func (o *InvoiceLineItem) GetSubtotal() float64`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *InvoiceLineItem) GetSubtotalOk() (*float64, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *InvoiceLineItem) SetSubtotal(v float64)`

SetSubtotal sets Subtotal field to given value.


### GetIsTaxInclusive

`func (o *InvoiceLineItem) GetIsTaxInclusive() bool`

GetIsTaxInclusive returns the IsTaxInclusive field if non-nil, zero value otherwise.

### GetIsTaxInclusiveOk

`func (o *InvoiceLineItem) GetIsTaxInclusiveOk() (*bool, bool)`

GetIsTaxInclusiveOk returns a tuple with the IsTaxInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaxInclusive

`func (o *InvoiceLineItem) SetIsTaxInclusive(v bool)`

SetIsTaxInclusive sets IsTaxInclusive field to given value.


### GetIsTaxable

`func (o *InvoiceLineItem) GetIsTaxable() bool`

GetIsTaxable returns the IsTaxable field if non-nil, zero value otherwise.

### GetIsTaxableOk

`func (o *InvoiceLineItem) GetIsTaxableOk() (*bool, bool)`

GetIsTaxableOk returns a tuple with the IsTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaxable

`func (o *InvoiceLineItem) SetIsTaxable(v bool)`

SetIsTaxable sets IsTaxable field to given value.


### GetTaxIdentifiers

`func (o *InvoiceLineItem) GetTaxIdentifiers() []FieldItem`

GetTaxIdentifiers returns the TaxIdentifiers field if non-nil, zero value otherwise.

### GetTaxIdentifiersOk

`func (o *InvoiceLineItem) GetTaxIdentifiersOk() (*[]FieldItem, bool)`

GetTaxIdentifiersOk returns a tuple with the TaxIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentifiers

`func (o *InvoiceLineItem) SetTaxIdentifiers(v []FieldItem)`

SetTaxIdentifiers sets TaxIdentifiers field to given value.

### HasTaxIdentifiers

`func (o *InvoiceLineItem) HasTaxIdentifiers() bool`

HasTaxIdentifiers returns a boolean if a field has been set.

### GetTaxExemptType

`func (o *InvoiceLineItem) GetTaxExemptType() TaxExemptType`

GetTaxExemptType returns the TaxExemptType field if non-nil, zero value otherwise.

### GetTaxExemptTypeOk

`func (o *InvoiceLineItem) GetTaxExemptTypeOk() (*TaxExemptType, bool)`

GetTaxExemptTypeOk returns a tuple with the TaxExemptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExemptType

`func (o *InvoiceLineItem) SetTaxExemptType(v TaxExemptType)`

SetTaxExemptType sets TaxExemptType field to given value.

### HasTaxExemptType

`func (o *InvoiceLineItem) HasTaxExemptType() bool`

HasTaxExemptType returns a boolean if a field has been set.

### GetTaxExemptReason

`func (o *InvoiceLineItem) GetTaxExemptReason() string`

GetTaxExemptReason returns the TaxExemptReason field if non-nil, zero value otherwise.

### GetTaxExemptReasonOk

`func (o *InvoiceLineItem) GetTaxExemptReasonOk() (*string, bool)`

GetTaxExemptReasonOk returns a tuple with the TaxExemptReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExemptReason

`func (o *InvoiceLineItem) SetTaxExemptReason(v string)`

SetTaxExemptReason sets TaxExemptReason field to given value.

### HasTaxExemptReason

`func (o *InvoiceLineItem) HasTaxExemptReason() bool`

HasTaxExemptReason returns a boolean if a field has been set.

### GetExemptAmount

`func (o *InvoiceLineItem) GetExemptAmount() float64`

GetExemptAmount returns the ExemptAmount field if non-nil, zero value otherwise.

### GetExemptAmountOk

`func (o *InvoiceLineItem) GetExemptAmountOk() (*float64, bool)`

GetExemptAmountOk returns a tuple with the ExemptAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptAmount

`func (o *InvoiceLineItem) SetExemptAmount(v float64)`

SetExemptAmount sets ExemptAmount field to given value.


### GetDiscountAmount

`func (o *InvoiceLineItem) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *InvoiceLineItem) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *InvoiceLineItem) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.


### GetTaxableAmount

`func (o *InvoiceLineItem) GetTaxableAmount() float64`

GetTaxableAmount returns the TaxableAmount field if non-nil, zero value otherwise.

### GetTaxableAmountOk

`func (o *InvoiceLineItem) GetTaxableAmountOk() (*float64, bool)`

GetTaxableAmountOk returns a tuple with the TaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount

`func (o *InvoiceLineItem) SetTaxableAmount(v float64)`

SetTaxableAmount sets TaxableAmount field to given value.


### GetTaxAmount

`func (o *InvoiceLineItem) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *InvoiceLineItem) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *InvoiceLineItem) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.


### GetTotal

`func (o *InvoiceLineItem) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvoiceLineItem) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvoiceLineItem) SetTotal(v float64)`

SetTotal sets Total field to given value.


### GetIsPartialTax

`func (o *InvoiceLineItem) GetIsPartialTax() bool`

GetIsPartialTax returns the IsPartialTax field if non-nil, zero value otherwise.

### GetIsPartialTaxOk

`func (o *InvoiceLineItem) GetIsPartialTaxOk() (*bool, bool)`

GetIsPartialTaxOk returns a tuple with the IsPartialTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartialTax

`func (o *InvoiceLineItem) SetIsPartialTax(v bool)`

SetIsPartialTax sets IsPartialTax field to given value.

### HasIsPartialTax

`func (o *InvoiceLineItem) HasIsPartialTax() bool`

HasIsPartialTax returns a boolean if a field has been set.

### GetTaxes

`func (o *InvoiceLineItem) GetTaxes() []TaxLineItem`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *InvoiceLineItem) GetTaxesOk() (*[]TaxLineItem, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *InvoiceLineItem) SetTaxes(v []TaxLineItem)`

SetTaxes sets Taxes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


