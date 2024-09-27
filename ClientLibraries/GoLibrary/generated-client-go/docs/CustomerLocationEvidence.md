# CustomerLocationEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The customer&#39;s IP to determine which country the customer belongs to. | [optional] 
**Bin** | Pointer to **string** | The country associated with a card by using the first or last 6 digits of the Bank Identification Number. | [optional] 
**PaymentCountryCode** | Pointer to **string** | Identifies the country code associated with the payment method. | [optional] 

## Methods

### NewCustomerLocationEvidence

`func NewCustomerLocationEvidence() *CustomerLocationEvidence`

NewCustomerLocationEvidence instantiates a new CustomerLocationEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerLocationEvidenceWithDefaults

`func NewCustomerLocationEvidenceWithDefaults() *CustomerLocationEvidence`

NewCustomerLocationEvidenceWithDefaults instantiates a new CustomerLocationEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *CustomerLocationEvidence) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CustomerLocationEvidence) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CustomerLocationEvidence) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CustomerLocationEvidence) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetBin

`func (o *CustomerLocationEvidence) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *CustomerLocationEvidence) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *CustomerLocationEvidence) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *CustomerLocationEvidence) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetPaymentCountryCode

`func (o *CustomerLocationEvidence) GetPaymentCountryCode() string`

GetPaymentCountryCode returns the PaymentCountryCode field if non-nil, zero value otherwise.

### GetPaymentCountryCodeOk

`func (o *CustomerLocationEvidence) GetPaymentCountryCodeOk() (*string, bool)`

GetPaymentCountryCodeOk returns a tuple with the PaymentCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCountryCode

`func (o *CustomerLocationEvidence) SetPaymentCountryCode(v string)`

SetPaymentCountryCode sets PaymentCountryCode field to given value.

### HasPaymentCountryCode

`func (o *CustomerLocationEvidence) HasPaymentCountryCode() bool`

HasPaymentCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


