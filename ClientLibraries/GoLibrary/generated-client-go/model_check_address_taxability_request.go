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

// checks if the CheckAddressTaxabilityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckAddressTaxabilityRequest{}

// CheckAddressTaxabilityRequest The taxability request containing the address. Postal code & Country is mandatory.
type CheckAddressTaxabilityRequest struct {
	Address *Address `json:"address,omitempty"`
}

// NewCheckAddressTaxabilityRequest instantiates a new CheckAddressTaxabilityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckAddressTaxabilityRequest() *CheckAddressTaxabilityRequest {
	this := CheckAddressTaxabilityRequest{}
	return &this
}

// NewCheckAddressTaxabilityRequestWithDefaults instantiates a new CheckAddressTaxabilityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckAddressTaxabilityRequestWithDefaults() *CheckAddressTaxabilityRequest {
	this := CheckAddressTaxabilityRequest{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CheckAddressTaxabilityRequest) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAddressTaxabilityRequest) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CheckAddressTaxabilityRequest) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *CheckAddressTaxabilityRequest) SetAddress(v Address) {
	o.Address = &v
}

func (o CheckAddressTaxabilityRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckAddressTaxabilityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableCheckAddressTaxabilityRequest struct {
	value *CheckAddressTaxabilityRequest
	isSet bool
}

func (v NullableCheckAddressTaxabilityRequest) Get() *CheckAddressTaxabilityRequest {
	return v.value
}

func (v *NullableCheckAddressTaxabilityRequest) Set(val *CheckAddressTaxabilityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckAddressTaxabilityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckAddressTaxabilityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckAddressTaxabilityRequest(val *CheckAddressTaxabilityRequest) *NullableCheckAddressTaxabilityRequest {
	return &NullableCheckAddressTaxabilityRequest{value: val, isSet: true}
}

func (v NullableCheckAddressTaxabilityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckAddressTaxabilityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


