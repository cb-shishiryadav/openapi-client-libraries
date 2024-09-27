# CreditNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditNoteId** | **string** | The unique identifier of the credit note at the Tax Service Provider or Tax Service Adapter. | 
**CreditNoteCode** | **string** | The unique identifier of the credit note in Chargebee. | 
**InvoiceCode** | Pointer to **string** | The unique identifier of the invoice in Chargebee to which this credit note belongs. | [optional] 
**InvoiceId** | Pointer to **string** | The unique identifier of the invoice in the Tax Service Adapter or the Tax Service Provider. | [optional] 
**CreditNoteType** | [**CreditNoteType**](CreditNoteType.md) |  | 
**DocumentDateTime** | **time.Time** | The date and time at which the credit note was created in Chargebee. For example, if the value is 2022-10-28T15:36:28.129+05:30, then the timestamp represents October 28, 2022, at 15:36:28.129, with an offset of +05:30. This means that the time represented is 5 hours and 30 minutes ahead of UTC/GMT. In the case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the &#39;Z&#39; indicating that the time is in UTC. | 
**TaxDateTime** | Pointer to **time.Time** | The date and time at which the tax was applicable in Chargebee. For example, if the value is 2022-10-28T15:36:28.129+05:30, then the timestamp represents October 28, 2022, at 15:36:28.129, with an offset of +05:30. This means that the time represented is 5 hours and 30 minutes ahead of UTC/GMT.In the case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the &#39;Z&#39; indicating that the time is in UTC. | [optional] 
**Status** | [**DocumentStatus**](DocumentStatus.md) |  | 
**Currency** | **string** | The [currency](https://en.wikipedia.org/wiki/Currency) in the [ISO-4217 format](https://www.iso.org/iso-4217-currency-codes.html). | 
**Seller** | [**Seller**](Seller.md) |  | 
**Customer** | [**Customer**](Customer.md) |  | 
**DiscountAmount** | **float64** | The total discount applied. This is the sum of all &#x60;lineItems.discount&#x60;. | 
**Subtotal** | Pointer to **float64** | The amount after discounts. This is the sum of all &#x60;lineItems.subtotal&#x60;. | [optional] 
**ExemptAmount** | **float64** | The amount exempted from tax. | 
**TaxableAmount** | **float64** | The amount upon which the tax is calculated. | 
**TaxAmount** | **float64** | The total tax payable. This is the sum of all &#x60;lineItems.taxAmount&#x60;. | 
**Total** | **float64** | The total amount of the credit note. &#x60;total&#x60; can be expressed as &#x60;exemptAmount&#x60; + &#x60;taxableAmount&#x60; + &#x60;taxAmount&#x60;. | 
**RoundingAmount** | Pointer to **float64** | The rounding amount added to the total amount to account for fractional correction. | [optional] 
**LineItems** | [**[]InvoiceLineItem**](InvoiceLineItem.md) |  | 

## Methods

### NewCreditNote

`func NewCreditNote(creditNoteId string, creditNoteCode string, creditNoteType CreditNoteType, documentDateTime time.Time, status DocumentStatus, currency string, seller Seller, customer Customer, discountAmount float64, exemptAmount float64, taxableAmount float64, taxAmount float64, total float64, lineItems []InvoiceLineItem, ) *CreditNote`

NewCreditNote instantiates a new CreditNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteWithDefaults

`func NewCreditNoteWithDefaults() *CreditNote`

NewCreditNoteWithDefaults instantiates a new CreditNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditNoteId

`func (o *CreditNote) GetCreditNoteId() string`

GetCreditNoteId returns the CreditNoteId field if non-nil, zero value otherwise.

### GetCreditNoteIdOk

`func (o *CreditNote) GetCreditNoteIdOk() (*string, bool)`

GetCreditNoteIdOk returns a tuple with the CreditNoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteId

`func (o *CreditNote) SetCreditNoteId(v string)`

SetCreditNoteId sets CreditNoteId field to given value.


### GetCreditNoteCode

`func (o *CreditNote) GetCreditNoteCode() string`

GetCreditNoteCode returns the CreditNoteCode field if non-nil, zero value otherwise.

### GetCreditNoteCodeOk

`func (o *CreditNote) GetCreditNoteCodeOk() (*string, bool)`

GetCreditNoteCodeOk returns a tuple with the CreditNoteCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteCode

`func (o *CreditNote) SetCreditNoteCode(v string)`

SetCreditNoteCode sets CreditNoteCode field to given value.


### GetInvoiceCode

`func (o *CreditNote) GetInvoiceCode() string`

GetInvoiceCode returns the InvoiceCode field if non-nil, zero value otherwise.

### GetInvoiceCodeOk

`func (o *CreditNote) GetInvoiceCodeOk() (*string, bool)`

GetInvoiceCodeOk returns a tuple with the InvoiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCode

`func (o *CreditNote) SetInvoiceCode(v string)`

SetInvoiceCode sets InvoiceCode field to given value.

### HasInvoiceCode

`func (o *CreditNote) HasInvoiceCode() bool`

HasInvoiceCode returns a boolean if a field has been set.

### GetInvoiceId

`func (o *CreditNote) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreditNote) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreditNote) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *CreditNote) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetCreditNoteType

`func (o *CreditNote) GetCreditNoteType() CreditNoteType`

GetCreditNoteType returns the CreditNoteType field if non-nil, zero value otherwise.

### GetCreditNoteTypeOk

`func (o *CreditNote) GetCreditNoteTypeOk() (*CreditNoteType, bool)`

GetCreditNoteTypeOk returns a tuple with the CreditNoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteType

`func (o *CreditNote) SetCreditNoteType(v CreditNoteType)`

SetCreditNoteType sets CreditNoteType field to given value.


### GetDocumentDateTime

`func (o *CreditNote) GetDocumentDateTime() time.Time`

GetDocumentDateTime returns the DocumentDateTime field if non-nil, zero value otherwise.

### GetDocumentDateTimeOk

`func (o *CreditNote) GetDocumentDateTimeOk() (*time.Time, bool)`

GetDocumentDateTimeOk returns a tuple with the DocumentDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDateTime

`func (o *CreditNote) SetDocumentDateTime(v time.Time)`

SetDocumentDateTime sets DocumentDateTime field to given value.


### GetTaxDateTime

`func (o *CreditNote) GetTaxDateTime() time.Time`

GetTaxDateTime returns the TaxDateTime field if non-nil, zero value otherwise.

### GetTaxDateTimeOk

`func (o *CreditNote) GetTaxDateTimeOk() (*time.Time, bool)`

GetTaxDateTimeOk returns a tuple with the TaxDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDateTime

`func (o *CreditNote) SetTaxDateTime(v time.Time)`

SetTaxDateTime sets TaxDateTime field to given value.

### HasTaxDateTime

`func (o *CreditNote) HasTaxDateTime() bool`

HasTaxDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *CreditNote) GetStatus() DocumentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreditNote) GetStatusOk() (*DocumentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreditNote) SetStatus(v DocumentStatus)`

SetStatus sets Status field to given value.


### GetCurrency

`func (o *CreditNote) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreditNote) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreditNote) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSeller

`func (o *CreditNote) GetSeller() Seller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *CreditNote) GetSellerOk() (*Seller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *CreditNote) SetSeller(v Seller)`

SetSeller sets Seller field to given value.


### GetCustomer

`func (o *CreditNote) GetCustomer() Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CreditNote) GetCustomerOk() (*Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CreditNote) SetCustomer(v Customer)`

SetCustomer sets Customer field to given value.


### GetDiscountAmount

`func (o *CreditNote) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *CreditNote) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *CreditNote) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.


### GetSubtotal

`func (o *CreditNote) GetSubtotal() float64`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *CreditNote) GetSubtotalOk() (*float64, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *CreditNote) SetSubtotal(v float64)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *CreditNote) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### GetExemptAmount

`func (o *CreditNote) GetExemptAmount() float64`

GetExemptAmount returns the ExemptAmount field if non-nil, zero value otherwise.

### GetExemptAmountOk

`func (o *CreditNote) GetExemptAmountOk() (*float64, bool)`

GetExemptAmountOk returns a tuple with the ExemptAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptAmount

`func (o *CreditNote) SetExemptAmount(v float64)`

SetExemptAmount sets ExemptAmount field to given value.


### GetTaxableAmount

`func (o *CreditNote) GetTaxableAmount() float64`

GetTaxableAmount returns the TaxableAmount field if non-nil, zero value otherwise.

### GetTaxableAmountOk

`func (o *CreditNote) GetTaxableAmountOk() (*float64, bool)`

GetTaxableAmountOk returns a tuple with the TaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount

`func (o *CreditNote) SetTaxableAmount(v float64)`

SetTaxableAmount sets TaxableAmount field to given value.


### GetTaxAmount

`func (o *CreditNote) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *CreditNote) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *CreditNote) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.


### GetTotal

`func (o *CreditNote) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CreditNote) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CreditNote) SetTotal(v float64)`

SetTotal sets Total field to given value.


### GetRoundingAmount

`func (o *CreditNote) GetRoundingAmount() float64`

GetRoundingAmount returns the RoundingAmount field if non-nil, zero value otherwise.

### GetRoundingAmountOk

`func (o *CreditNote) GetRoundingAmountOk() (*float64, bool)`

GetRoundingAmountOk returns a tuple with the RoundingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundingAmount

`func (o *CreditNote) SetRoundingAmount(v float64)`

SetRoundingAmount sets RoundingAmount field to given value.

### HasRoundingAmount

`func (o *CreditNote) HasRoundingAmount() bool`

HasRoundingAmount returns a boolean if a field has been set.

### GetLineItems

`func (o *CreditNote) GetLineItems() []InvoiceLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CreditNote) GetLineItemsOk() (*[]InvoiceLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CreditNote) SetLineItems(v []InvoiceLineItem)`

SetLineItems sets LineItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


