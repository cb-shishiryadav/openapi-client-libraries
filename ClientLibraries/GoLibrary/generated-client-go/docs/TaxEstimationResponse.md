# TaxEstimationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seller** | [**Seller**](Seller.md) |  | 
**Customer** | [**Customer**](Customer.md) |  | 
**EstimateDateTime** | **time.Time** | The time as of which the tax estimation is to be calculated. This can be a value in the past. For example, if the value is provided as 2022-10-28T15:36:28.129+05:30, then the tax rates applicable on October 28, 2022, at 15:36:28.129, with an offset of +05:30 ahead of UTC/GMT are used for calculations. In case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the &#39;Z&#39; indicating that the time is in UTC | 
**Currency** | **string** | The [currency](https://en.wikipedia.org/wiki/Currency) in the [ISO-4217 format](https://www.iso.org/iso-4217-currency-codes.html). | 
**Subtotal** | **float64** | The amount after discounts. This is the sum of all &#x60;lineItems.subtotal&#x60;. | 
**ExemptAmount** | **float64** | The part of the &#x60;subtotal&#x60; that is exempted from tax. | 
**DiscountAmount** | **float64** | The total discount applied. This is the sum of all &#x60;lineItems.discount&#x60;. | 
**TaxableAmount** | **float64** | The part of the &#x60;subtotal&#x60; that is taxable. | 
**TaxAmount** | **float64** | The total tax payable. This is the sum of all &#x60;lineItems.taxAmount&#x60;. | 
**Total** | **float64** | The total after discounts and taxes. This is the same as &#x60;subtotal&#x60; if it is tax inclusive; otherwise it is &#x60;subtotal&#x60; + &#x60;taxAmount&#x60;. &#x60;total&#x60; can also be expressed as &#x60;exemptAmount&#x60; + &#x60;taxableAmount&#x60; + &#x60;taxAmount&#x60;. | 
**LineItems** | [**[]InvoiceLineItem**](InvoiceLineItem.md) | List of line item details for the tax estimation response. | 

## Methods

### NewTaxEstimationResponse

`func NewTaxEstimationResponse(seller Seller, customer Customer, estimateDateTime time.Time, currency string, subtotal float64, exemptAmount float64, discountAmount float64, taxableAmount float64, taxAmount float64, total float64, lineItems []InvoiceLineItem, ) *TaxEstimationResponse`

NewTaxEstimationResponse instantiates a new TaxEstimationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxEstimationResponseWithDefaults

`func NewTaxEstimationResponseWithDefaults() *TaxEstimationResponse`

NewTaxEstimationResponseWithDefaults instantiates a new TaxEstimationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeller

`func (o *TaxEstimationResponse) GetSeller() Seller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *TaxEstimationResponse) GetSellerOk() (*Seller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *TaxEstimationResponse) SetSeller(v Seller)`

SetSeller sets Seller field to given value.


### GetCustomer

`func (o *TaxEstimationResponse) GetCustomer() Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *TaxEstimationResponse) GetCustomerOk() (*Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *TaxEstimationResponse) SetCustomer(v Customer)`

SetCustomer sets Customer field to given value.


### GetEstimateDateTime

`func (o *TaxEstimationResponse) GetEstimateDateTime() time.Time`

GetEstimateDateTime returns the EstimateDateTime field if non-nil, zero value otherwise.

### GetEstimateDateTimeOk

`func (o *TaxEstimationResponse) GetEstimateDateTimeOk() (*time.Time, bool)`

GetEstimateDateTimeOk returns a tuple with the EstimateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimateDateTime

`func (o *TaxEstimationResponse) SetEstimateDateTime(v time.Time)`

SetEstimateDateTime sets EstimateDateTime field to given value.


### GetCurrency

`func (o *TaxEstimationResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TaxEstimationResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TaxEstimationResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSubtotal

`func (o *TaxEstimationResponse) GetSubtotal() float64`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *TaxEstimationResponse) GetSubtotalOk() (*float64, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *TaxEstimationResponse) SetSubtotal(v float64)`

SetSubtotal sets Subtotal field to given value.


### GetExemptAmount

`func (o *TaxEstimationResponse) GetExemptAmount() float64`

GetExemptAmount returns the ExemptAmount field if non-nil, zero value otherwise.

### GetExemptAmountOk

`func (o *TaxEstimationResponse) GetExemptAmountOk() (*float64, bool)`

GetExemptAmountOk returns a tuple with the ExemptAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptAmount

`func (o *TaxEstimationResponse) SetExemptAmount(v float64)`

SetExemptAmount sets ExemptAmount field to given value.


### GetDiscountAmount

`func (o *TaxEstimationResponse) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *TaxEstimationResponse) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *TaxEstimationResponse) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.


### GetTaxableAmount

`func (o *TaxEstimationResponse) GetTaxableAmount() float64`

GetTaxableAmount returns the TaxableAmount field if non-nil, zero value otherwise.

### GetTaxableAmountOk

`func (o *TaxEstimationResponse) GetTaxableAmountOk() (*float64, bool)`

GetTaxableAmountOk returns a tuple with the TaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount

`func (o *TaxEstimationResponse) SetTaxableAmount(v float64)`

SetTaxableAmount sets TaxableAmount field to given value.


### GetTaxAmount

`func (o *TaxEstimationResponse) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *TaxEstimationResponse) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *TaxEstimationResponse) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.


### GetTotal

`func (o *TaxEstimationResponse) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TaxEstimationResponse) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TaxEstimationResponse) SetTotal(v float64)`

SetTotal sets Total field to given value.


### GetLineItems

`func (o *TaxEstimationResponse) GetLineItems() []InvoiceLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *TaxEstimationResponse) GetLineItemsOk() (*[]InvoiceLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *TaxEstimationResponse) SetLineItems(v []InvoiceLineItem)`

SetLineItems sets LineItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


