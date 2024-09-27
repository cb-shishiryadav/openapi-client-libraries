# InvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceCode** | **string** | The unique identifier of the invoice in Chargebee. | 
**DocumentDateTime** | **time.Time** | The date and time at which the invoice was generated in Chargebee. For example, if the value is 2022-10-28T15:36:28.129+05:30, then the timestamp represents October 28, 2022, at 15:36:28.129, with an offset of +05:30. This means that the time represented is 5 hours and 30 minutes ahead of UTC/GMT.In the case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the &#39;Z&#39; indicating that the time is in UTC. | 
**TaxDateTime** | Pointer to **time.Time** | The date and time at which the tax was applicable in Chargebee. For example, if the value is 2022-10-28T15:36:28.129+05:30, then the timestamp represents October 28, 2022, at 15:36:28.129, with an offset of +05:30. This means that the time represented is 5 hours and 30 minutes ahead of UTC/GMT.In the case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the &#39;Z&#39; indicating that the time is in UTC. | [optional] 
**Currency** | **string** | The [currency](https://en.wikipedia.org/wiki/Currency) in the [ISO-4217 format](https://www.iso.org/iso-4217-currency-codes.html). | 
**Seller** | [**Seller**](Seller.md) |  | 
**Customer** | [**Customer**](Customer.md) |  | 
**Subtotal** | **float64** | The amount after discounts. This is the sum of all &#x60;lineItems.subtotal&#x60;. | 
**ExemptAmount** | **float64** | The part of the &#x60;subtotal&#x60; that is exempted from tax. | 
**DiscountAmount** | **float64** | The total discount applied. This is the sum of all &#x60;lineItems.discount&#x60;. | 
**TaxableAmount** | **float64** | The part of the &#x60;subtotal&#x60; that is taxable. | 
**TaxAmount** | **float64** | The total tax payable. This is the sum of all &#x60;lineItems.taxAmount&#x60;. | 
**Total** | **float64** | The total after discounts and taxes. This is the same as &#x60;subtotal&#x60; if it is tax inclusive; otherwise it is &#x60;subtotal&#x60; + &#x60;taxAmount&#x60;. &#x60;total&#x60; can also be expressed as &#x60;exemptAmount&#x60; + &#x60;taxableAmount&#x60; + &#x60;taxAmount&#x60;. | 
**LineItems** | [**[]InvoiceLineItem**](InvoiceLineItem.md) |  | 

## Methods

### NewInvoiceRequest

`func NewInvoiceRequest(invoiceCode string, documentDateTime time.Time, currency string, seller Seller, customer Customer, subtotal float64, exemptAmount float64, discountAmount float64, taxableAmount float64, taxAmount float64, total float64, lineItems []InvoiceLineItem, ) *InvoiceRequest`

NewInvoiceRequest instantiates a new InvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceRequestWithDefaults

`func NewInvoiceRequestWithDefaults() *InvoiceRequest`

NewInvoiceRequestWithDefaults instantiates a new InvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceCode

`func (o *InvoiceRequest) GetInvoiceCode() string`

GetInvoiceCode returns the InvoiceCode field if non-nil, zero value otherwise.

### GetInvoiceCodeOk

`func (o *InvoiceRequest) GetInvoiceCodeOk() (*string, bool)`

GetInvoiceCodeOk returns a tuple with the InvoiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCode

`func (o *InvoiceRequest) SetInvoiceCode(v string)`

SetInvoiceCode sets InvoiceCode field to given value.


### GetDocumentDateTime

`func (o *InvoiceRequest) GetDocumentDateTime() time.Time`

GetDocumentDateTime returns the DocumentDateTime field if non-nil, zero value otherwise.

### GetDocumentDateTimeOk

`func (o *InvoiceRequest) GetDocumentDateTimeOk() (*time.Time, bool)`

GetDocumentDateTimeOk returns a tuple with the DocumentDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDateTime

`func (o *InvoiceRequest) SetDocumentDateTime(v time.Time)`

SetDocumentDateTime sets DocumentDateTime field to given value.


### GetTaxDateTime

`func (o *InvoiceRequest) GetTaxDateTime() time.Time`

GetTaxDateTime returns the TaxDateTime field if non-nil, zero value otherwise.

### GetTaxDateTimeOk

`func (o *InvoiceRequest) GetTaxDateTimeOk() (*time.Time, bool)`

GetTaxDateTimeOk returns a tuple with the TaxDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDateTime

`func (o *InvoiceRequest) SetTaxDateTime(v time.Time)`

SetTaxDateTime sets TaxDateTime field to given value.

### HasTaxDateTime

`func (o *InvoiceRequest) HasTaxDateTime() bool`

HasTaxDateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *InvoiceRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSeller

`func (o *InvoiceRequest) GetSeller() Seller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *InvoiceRequest) GetSellerOk() (*Seller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *InvoiceRequest) SetSeller(v Seller)`

SetSeller sets Seller field to given value.


### GetCustomer

`func (o *InvoiceRequest) GetCustomer() Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InvoiceRequest) GetCustomerOk() (*Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InvoiceRequest) SetCustomer(v Customer)`

SetCustomer sets Customer field to given value.


### GetSubtotal

`func (o *InvoiceRequest) GetSubtotal() float64`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *InvoiceRequest) GetSubtotalOk() (*float64, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *InvoiceRequest) SetSubtotal(v float64)`

SetSubtotal sets Subtotal field to given value.


### GetExemptAmount

`func (o *InvoiceRequest) GetExemptAmount() float64`

GetExemptAmount returns the ExemptAmount field if non-nil, zero value otherwise.

### GetExemptAmountOk

`func (o *InvoiceRequest) GetExemptAmountOk() (*float64, bool)`

GetExemptAmountOk returns a tuple with the ExemptAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptAmount

`func (o *InvoiceRequest) SetExemptAmount(v float64)`

SetExemptAmount sets ExemptAmount field to given value.


### GetDiscountAmount

`func (o *InvoiceRequest) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *InvoiceRequest) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *InvoiceRequest) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.


### GetTaxableAmount

`func (o *InvoiceRequest) GetTaxableAmount() float64`

GetTaxableAmount returns the TaxableAmount field if non-nil, zero value otherwise.

### GetTaxableAmountOk

`func (o *InvoiceRequest) GetTaxableAmountOk() (*float64, bool)`

GetTaxableAmountOk returns a tuple with the TaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount

`func (o *InvoiceRequest) SetTaxableAmount(v float64)`

SetTaxableAmount sets TaxableAmount field to given value.


### GetTaxAmount

`func (o *InvoiceRequest) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *InvoiceRequest) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *InvoiceRequest) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.


### GetTotal

`func (o *InvoiceRequest) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvoiceRequest) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvoiceRequest) SetTotal(v float64)`

SetTotal sets Total field to given value.


### GetLineItems

`func (o *InvoiceRequest) GetLineItems() []InvoiceLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *InvoiceRequest) GetLineItemsOk() (*[]InvoiceLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *InvoiceRequest) SetLineItems(v []InvoiceLineItem)`

SetLineItems sets LineItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


