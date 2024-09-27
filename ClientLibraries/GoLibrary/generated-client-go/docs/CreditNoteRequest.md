# CreditNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditNoteCode** | **string** | The unique identifier of the credit note in Chargebee. | 
**InvoiceCode** | Pointer to **string** | The unique identifier of the invoice in Chargebee to which this credit note belongs. | [optional] 
**InvoiceId** | Pointer to **string** | The unique identifier of the invoice in the Tax Service Adapter or the Tax Service Provider. | [optional] 
**CreditNoteType** | [**CreditNoteType**](CreditNoteType.md) |  | 
**DocumentDateTime** | **time.Time** | The date and time at which the credit note was created in Chargebee. For example, if the value is 2022-10-28T15:36:28.129+05:30, then the timestamp represents October 28, 2022, at 15:36:28.129, with an offset of +05:30. This means that the time represented is 5 hours and 30 minutes ahead of UTC/GMT. In the case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the &#39;Z&#39; indicating that the time is in UTC. | 
**TaxDateTime** | Pointer to **time.Time** | The date and time at which the tax was applicable in Chargebee. For example, if the value is 2022-10-28T15:36:28.129+05:30, then the timestamp represents October 28, 2022, at 15:36:28.129, with an offset of +05:30. This means that the time represented is 5 hours and 30 minutes ahead of UTC/GMT.In the case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the &#39;Z&#39; indicating that the time is in UTC. | [optional] 
**Currency** | **string** | The [currency](https://en.wikipedia.org/wiki/Currency) in the [ISO-4217 format](https://www.iso.org/iso-4217-currency-codes.html). | 
**Seller** | [**Seller**](Seller.md) |  | 
**Customer** | [**Customer**](Customer.md) |  | 
**Total** | **float64** | The total amount of the credit note. &#x60;total&#x60; can be expressed as &#x60;exemptAmount&#x60; + &#x60;taxableAmount&#x60; + &#x60;taxAmount&#x60;. | 
**ExemptAmount** | **float64** | The amount exempted from tax. | 
**DiscountAmount** | **float64** | The total discount applied. This is the sum of all &#x60;lineItems.discount&#x60;. | 
**TaxableAmount** | **float64** | The amount upon which the tax is calculated. | 
**TaxAmount** | **float64** | The total tax payable. This is the sum of all &#x60;lineItems.taxAmount&#x60;. | 
**RoundingAmount** | Pointer to **float64** | The rounding amount added to the total amount to account for fractional correction. | [optional] 
**LineItems** | Pointer to [**[]InvoiceLineItem**](InvoiceLineItem.md) |  | [optional] 

## Methods

### NewCreditNoteRequest

`func NewCreditNoteRequest(creditNoteCode string, creditNoteType CreditNoteType, documentDateTime time.Time, currency string, seller Seller, customer Customer, total float64, exemptAmount float64, discountAmount float64, taxableAmount float64, taxAmount float64, ) *CreditNoteRequest`

NewCreditNoteRequest instantiates a new CreditNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteRequestWithDefaults

`func NewCreditNoteRequestWithDefaults() *CreditNoteRequest`

NewCreditNoteRequestWithDefaults instantiates a new CreditNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditNoteCode

`func (o *CreditNoteRequest) GetCreditNoteCode() string`

GetCreditNoteCode returns the CreditNoteCode field if non-nil, zero value otherwise.

### GetCreditNoteCodeOk

`func (o *CreditNoteRequest) GetCreditNoteCodeOk() (*string, bool)`

GetCreditNoteCodeOk returns a tuple with the CreditNoteCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteCode

`func (o *CreditNoteRequest) SetCreditNoteCode(v string)`

SetCreditNoteCode sets CreditNoteCode field to given value.


### GetInvoiceCode

`func (o *CreditNoteRequest) GetInvoiceCode() string`

GetInvoiceCode returns the InvoiceCode field if non-nil, zero value otherwise.

### GetInvoiceCodeOk

`func (o *CreditNoteRequest) GetInvoiceCodeOk() (*string, bool)`

GetInvoiceCodeOk returns a tuple with the InvoiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCode

`func (o *CreditNoteRequest) SetInvoiceCode(v string)`

SetInvoiceCode sets InvoiceCode field to given value.

### HasInvoiceCode

`func (o *CreditNoteRequest) HasInvoiceCode() bool`

HasInvoiceCode returns a boolean if a field has been set.

### GetInvoiceId

`func (o *CreditNoteRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreditNoteRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreditNoteRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *CreditNoteRequest) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetCreditNoteType

`func (o *CreditNoteRequest) GetCreditNoteType() CreditNoteType`

GetCreditNoteType returns the CreditNoteType field if non-nil, zero value otherwise.

### GetCreditNoteTypeOk

`func (o *CreditNoteRequest) GetCreditNoteTypeOk() (*CreditNoteType, bool)`

GetCreditNoteTypeOk returns a tuple with the CreditNoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteType

`func (o *CreditNoteRequest) SetCreditNoteType(v CreditNoteType)`

SetCreditNoteType sets CreditNoteType field to given value.


### GetDocumentDateTime

`func (o *CreditNoteRequest) GetDocumentDateTime() time.Time`

GetDocumentDateTime returns the DocumentDateTime field if non-nil, zero value otherwise.

### GetDocumentDateTimeOk

`func (o *CreditNoteRequest) GetDocumentDateTimeOk() (*time.Time, bool)`

GetDocumentDateTimeOk returns a tuple with the DocumentDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDateTime

`func (o *CreditNoteRequest) SetDocumentDateTime(v time.Time)`

SetDocumentDateTime sets DocumentDateTime field to given value.


### GetTaxDateTime

`func (o *CreditNoteRequest) GetTaxDateTime() time.Time`

GetTaxDateTime returns the TaxDateTime field if non-nil, zero value otherwise.

### GetTaxDateTimeOk

`func (o *CreditNoteRequest) GetTaxDateTimeOk() (*time.Time, bool)`

GetTaxDateTimeOk returns a tuple with the TaxDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDateTime

`func (o *CreditNoteRequest) SetTaxDateTime(v time.Time)`

SetTaxDateTime sets TaxDateTime field to given value.

### HasTaxDateTime

`func (o *CreditNoteRequest) HasTaxDateTime() bool`

HasTaxDateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *CreditNoteRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreditNoteRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreditNoteRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSeller

`func (o *CreditNoteRequest) GetSeller() Seller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *CreditNoteRequest) GetSellerOk() (*Seller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *CreditNoteRequest) SetSeller(v Seller)`

SetSeller sets Seller field to given value.


### GetCustomer

`func (o *CreditNoteRequest) GetCustomer() Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CreditNoteRequest) GetCustomerOk() (*Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CreditNoteRequest) SetCustomer(v Customer)`

SetCustomer sets Customer field to given value.


### GetTotal

`func (o *CreditNoteRequest) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CreditNoteRequest) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CreditNoteRequest) SetTotal(v float64)`

SetTotal sets Total field to given value.


### GetExemptAmount

`func (o *CreditNoteRequest) GetExemptAmount() float64`

GetExemptAmount returns the ExemptAmount field if non-nil, zero value otherwise.

### GetExemptAmountOk

`func (o *CreditNoteRequest) GetExemptAmountOk() (*float64, bool)`

GetExemptAmountOk returns a tuple with the ExemptAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptAmount

`func (o *CreditNoteRequest) SetExemptAmount(v float64)`

SetExemptAmount sets ExemptAmount field to given value.


### GetDiscountAmount

`func (o *CreditNoteRequest) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *CreditNoteRequest) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *CreditNoteRequest) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.


### GetTaxableAmount

`func (o *CreditNoteRequest) GetTaxableAmount() float64`

GetTaxableAmount returns the TaxableAmount field if non-nil, zero value otherwise.

### GetTaxableAmountOk

`func (o *CreditNoteRequest) GetTaxableAmountOk() (*float64, bool)`

GetTaxableAmountOk returns a tuple with the TaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount

`func (o *CreditNoteRequest) SetTaxableAmount(v float64)`

SetTaxableAmount sets TaxableAmount field to given value.


### GetTaxAmount

`func (o *CreditNoteRequest) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *CreditNoteRequest) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *CreditNoteRequest) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.


### GetRoundingAmount

`func (o *CreditNoteRequest) GetRoundingAmount() float64`

GetRoundingAmount returns the RoundingAmount field if non-nil, zero value otherwise.

### GetRoundingAmountOk

`func (o *CreditNoteRequest) GetRoundingAmountOk() (*float64, bool)`

GetRoundingAmountOk returns a tuple with the RoundingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundingAmount

`func (o *CreditNoteRequest) SetRoundingAmount(v float64)`

SetRoundingAmount sets RoundingAmount field to given value.

### HasRoundingAmount

`func (o *CreditNoteRequest) HasRoundingAmount() bool`

HasRoundingAmount returns a boolean if a field has been set.

### GetLineItems

`func (o *CreditNoteRequest) GetLineItems() []InvoiceLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CreditNoteRequest) GetLineItemsOk() (*[]InvoiceLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CreditNoteRequest) SetLineItems(v []InvoiceLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *CreditNoteRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


