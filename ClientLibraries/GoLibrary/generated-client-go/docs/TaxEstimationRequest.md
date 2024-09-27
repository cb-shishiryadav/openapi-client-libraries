# TaxEstimationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seller** | [**Seller**](Seller.md) |  | 
**Customer** | [**Customer**](Customer.md) |  | 
**EstimateDateTime** | **time.Time** | The time as of which the tax estimation is to be calculated. This can be a value in the past. For example, if the value is provided as 2022-10-28T15:36:28.129+05:30, then the tax rates applicable on October 28, 2022, at 15:36:28.129, with an offset of +05:30 ahead of UTC/GMT are used for calculations. In case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the &#39;Z&#39; indicating that the time is in UTC | 
**Currency** | **string** | The [currency](https://en.wikipedia.org/wiki/Currency) in the [ISO-4217 format](https://www.iso.org/iso-4217-currency-codes.html). | 
**LineItems** | [**[]TaxEstimationLineItemRequest**](TaxEstimationLineItemRequest.md) | Contains the details of each line item in the tax estimation request. | 

## Methods

### NewTaxEstimationRequest

`func NewTaxEstimationRequest(seller Seller, customer Customer, estimateDateTime time.Time, currency string, lineItems []TaxEstimationLineItemRequest, ) *TaxEstimationRequest`

NewTaxEstimationRequest instantiates a new TaxEstimationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxEstimationRequestWithDefaults

`func NewTaxEstimationRequestWithDefaults() *TaxEstimationRequest`

NewTaxEstimationRequestWithDefaults instantiates a new TaxEstimationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeller

`func (o *TaxEstimationRequest) GetSeller() Seller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *TaxEstimationRequest) GetSellerOk() (*Seller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *TaxEstimationRequest) SetSeller(v Seller)`

SetSeller sets Seller field to given value.


### GetCustomer

`func (o *TaxEstimationRequest) GetCustomer() Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *TaxEstimationRequest) GetCustomerOk() (*Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *TaxEstimationRequest) SetCustomer(v Customer)`

SetCustomer sets Customer field to given value.


### GetEstimateDateTime

`func (o *TaxEstimationRequest) GetEstimateDateTime() time.Time`

GetEstimateDateTime returns the EstimateDateTime field if non-nil, zero value otherwise.

### GetEstimateDateTimeOk

`func (o *TaxEstimationRequest) GetEstimateDateTimeOk() (*time.Time, bool)`

GetEstimateDateTimeOk returns a tuple with the EstimateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimateDateTime

`func (o *TaxEstimationRequest) SetEstimateDateTime(v time.Time)`

SetEstimateDateTime sets EstimateDateTime field to given value.


### GetCurrency

`func (o *TaxEstimationRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TaxEstimationRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TaxEstimationRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetLineItems

`func (o *TaxEstimationRequest) GetLineItems() []TaxEstimationLineItemRequest`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *TaxEstimationRequest) GetLineItemsOk() (*[]TaxEstimationLineItemRequest, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *TaxEstimationRequest) SetLineItems(v []TaxEstimationLineItemRequest)`

SetLineItems sets LineItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


