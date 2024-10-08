/*
Taxes Service Adapter SPI

## Overview Tax Service Adapter is a Service Provider Interface (SPI) that helps Chargebee communicate with API platforms providing the following services: - Address validation - Estimating taxes for sales - Submitting invoices for tax filing  If you have an API that provides the services mentioned above and want Chargebee to be able to integrate with it, you can build an adapter service according to this specification.  ## Terminology Here's a list of terms we've used to describe this specification. ### Merchant    Chargebee’s customer who is carrying out the business of selling products. Also known as the Seller. The Seller has one or more Chargebee accounts. ### Chargebee   The SaaS which manages subscriptions and revenue operations on behalf of the Merchant. Chargebee makes API calls to the Tax Service Adapter for tax estimation and tax filing of transactions between the Merchant and the Customer. ### Customer  Merchant’s customer who purchases products from the Merchant. Their customer record(s) and other details are stored within Chargebee under the Merchant’s account. ### Tax Authority The institution, usually a government, that is responsible for collecting taxes for commercial transactions between Merchants and their Customers.   ### Tax Service Provider The service which provides tax estimation and tax filing services for merchants. ### Tax Service Adapter  The SPI for which the spec is defined in this document. This adapter serves the following purposes: - Translates API requests (as defined in this document) from Chargebee to the request format specified by the Tax Service Provider's API. - Translates API responses from the Tax Service Provider APIs to the format understood by Chargebee (as defined in this document). ### Authorization Chargebee uses HTTP header-based authorization for all the API endpoints associated with Tax Service Adapter. We dynamically pass this authorization key in the HTTP header. The parameter that holds this key is found in the JSON object `api_configuration` required for configuring your onboarding on Chargebee's marketplace. In the `api_configuration` object, our [Taxes Service Adapter SPI](https://chargebee.atlassian.net/l/cp/ca5aZ1mA) checks the authorization key parameter from `credential_configuration.id` and creates the HTTP header-based input query parameter for authorization. The `credential_configuration` is an array of objects with an `id` attribute, and the value of `id` is the parameter containing the authorization key.    Following are the JSON snippets for your reference.     ```json       \"api_configuration\": {           \"api_base_url\": \"https://xyz.abc.com/chargebee\",           \"credential_configuration\": [              {               \"id\":\"authorization_key\",               \"name\": \"Authorization Key\",                \"type\": \"text\",               \"is_sensitive\": true              },             {               \"id\": \"client_secret\",               \"name\": \"Client Secret\",               \"type\": \"text\",               \"is_sensitive\": true             }           ]         }    ```     `\"Authorization\":\"{\"<authorization_key>\":\"test_GykCvTykZFTBqHQAFzjMN9R1Thgdho0Q\"}\"` 

API version: 0.3.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CustomerLocationEvidence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerLocationEvidence{}

// CustomerLocationEvidence Represent the properties for customer location evidence.
type CustomerLocationEvidence struct {
	// The customer's IP to determine which country the customer belongs to.
	Ip *string `json:"ip,omitempty"`
	// The country associated with a card by using the first or last 6 digits of the Bank Identification Number.
	Bin *string `json:"bin,omitempty"`
	// Identifies the country code associated with the payment method.
	PaymentCountryCode *string `json:"paymentCountryCode,omitempty"`
}

// NewCustomerLocationEvidence instantiates a new CustomerLocationEvidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerLocationEvidence() *CustomerLocationEvidence {
	this := CustomerLocationEvidence{}
	return &this
}

// NewCustomerLocationEvidenceWithDefaults instantiates a new CustomerLocationEvidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerLocationEvidenceWithDefaults() *CustomerLocationEvidence {
	this := CustomerLocationEvidence{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *CustomerLocationEvidence) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerLocationEvidence) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *CustomerLocationEvidence) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *CustomerLocationEvidence) SetIp(v string) {
	o.Ip = &v
}

// GetBin returns the Bin field value if set, zero value otherwise.
func (o *CustomerLocationEvidence) GetBin() string {
	if o == nil || IsNil(o.Bin) {
		var ret string
		return ret
	}
	return *o.Bin
}

// GetBinOk returns a tuple with the Bin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerLocationEvidence) GetBinOk() (*string, bool) {
	if o == nil || IsNil(o.Bin) {
		return nil, false
	}
	return o.Bin, true
}

// HasBin returns a boolean if a field has been set.
func (o *CustomerLocationEvidence) HasBin() bool {
	if o != nil && !IsNil(o.Bin) {
		return true
	}

	return false
}

// SetBin gets a reference to the given string and assigns it to the Bin field.
func (o *CustomerLocationEvidence) SetBin(v string) {
	o.Bin = &v
}

// GetPaymentCountryCode returns the PaymentCountryCode field value if set, zero value otherwise.
func (o *CustomerLocationEvidence) GetPaymentCountryCode() string {
	if o == nil || IsNil(o.PaymentCountryCode) {
		var ret string
		return ret
	}
	return *o.PaymentCountryCode
}

// GetPaymentCountryCodeOk returns a tuple with the PaymentCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerLocationEvidence) GetPaymentCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentCountryCode) {
		return nil, false
	}
	return o.PaymentCountryCode, true
}

// HasPaymentCountryCode returns a boolean if a field has been set.
func (o *CustomerLocationEvidence) HasPaymentCountryCode() bool {
	if o != nil && !IsNil(o.PaymentCountryCode) {
		return true
	}

	return false
}

// SetPaymentCountryCode gets a reference to the given string and assigns it to the PaymentCountryCode field.
func (o *CustomerLocationEvidence) SetPaymentCountryCode(v string) {
	o.PaymentCountryCode = &v
}

func (o CustomerLocationEvidence) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerLocationEvidence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Bin) {
		toSerialize["bin"] = o.Bin
	}
	if !IsNil(o.PaymentCountryCode) {
		toSerialize["paymentCountryCode"] = o.PaymentCountryCode
	}
	return toSerialize, nil
}

type NullableCustomerLocationEvidence struct {
	value *CustomerLocationEvidence
	isSet bool
}

func (v NullableCustomerLocationEvidence) Get() *CustomerLocationEvidence {
	return v.value
}

func (v *NullableCustomerLocationEvidence) Set(val *CustomerLocationEvidence) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerLocationEvidence) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerLocationEvidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerLocationEvidence(val *CustomerLocationEvidence) *NullableCustomerLocationEvidence {
	return &NullableCustomerLocationEvidence{value: val, isSet: true}
}

func (v NullableCustomerLocationEvidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerLocationEvidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


