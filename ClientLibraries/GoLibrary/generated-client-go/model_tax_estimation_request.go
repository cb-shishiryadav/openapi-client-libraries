/*
Taxes Service Adapter SPI

## Overview Tax Service Adapter is a Service Provider Interface (SPI) that helps Chargebee communicate with API platforms providing the following services: - Address validation - Estimating taxes for sales - Submitting invoices for tax filing  If you have an API that provides the services mentioned above and want Chargebee to be able to integrate with it, you can build an adapter service according to this specification.  ## Terminology Here's a list of terms we've used to describe this specification. ### Merchant    Chargebee’s customer who is carrying out the business of selling products. Also known as the Seller. The Seller has one or more Chargebee accounts. ### Chargebee   The SaaS which manages subscriptions and revenue operations on behalf of the Merchant. Chargebee makes API calls to the Tax Service Adapter for tax estimation and tax filing of transactions between the Merchant and the Customer. ### Customer  Merchant’s customer who purchases products from the Merchant. Their customer record(s) and other details are stored within Chargebee under the Merchant’s account. ### Tax Authority The institution, usually a government, that is responsible for collecting taxes for commercial transactions between Merchants and their Customers.   ### Tax Service Provider The service which provides tax estimation and tax filing services for merchants. ### Tax Service Adapter  The SPI for which the spec is defined in this document. This adapter serves the following purposes: - Translates API requests (as defined in this document) from Chargebee to the request format specified by the Tax Service Provider's API. - Translates API responses from the Tax Service Provider APIs to the format understood by Chargebee (as defined in this document). ### Authorization Chargebee uses HTTP header-based authorization for all the API endpoints associated with Tax Service Adapter. We dynamically pass this authorization key in the HTTP header. The parameter that holds this key is found in the JSON object `api_configuration` required for configuring your onboarding on Chargebee's marketplace. In the `api_configuration` object, our [Taxes Service Adapter SPI](https://chargebee.atlassian.net/l/cp/ca5aZ1mA) checks the authorization key parameter from `credential_configuration.id` and creates the HTTP header-based input query parameter for authorization. The `credential_configuration` is an array of objects with an `id` attribute, and the value of `id` is the parameter containing the authorization key.    Following are the JSON snippets for your reference.     ```json       \"api_configuration\": {           \"api_base_url\": \"https://xyz.abc.com/chargebee\",           \"credential_configuration\": [              {               \"id\":\"authorization_key\",               \"name\": \"Authorization Key\",                \"type\": \"text\",               \"is_sensitive\": true              },             {               \"id\": \"client_secret\",               \"name\": \"Client Secret\",               \"type\": \"text\",               \"is_sensitive\": true             }           ]         }    ```     `\"Authorization\":\"{\"<authorization_key>\":\"test_GykCvTykZFTBqHQAFzjMN9R1Thgdho0Q\"}\"` 

API version: 0.3.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the TaxEstimationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxEstimationRequest{}

// TaxEstimationRequest Defines the parameters of a tax estimation request. This is sent to the Tax Service Adapter by Chargebee to estimate taxes for one or more line items.
type TaxEstimationRequest struct {
	Seller Seller `json:"seller"`
	Customer Customer `json:"customer"`
	// The time as of which the tax estimation is to be calculated. This can be a value in the past. For example, if the value is provided as 2022-10-28T15:36:28.129+05:30, then the tax rates applicable on October 28, 2022, at 15:36:28.129, with an offset of +05:30 ahead of UTC/GMT are used for calculations. In case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the 'Z' indicating that the time is in UTC
	EstimateDateTime time.Time `json:"estimateDateTime"`
	// The [currency](https://en.wikipedia.org/wiki/Currency) in the [ISO-4217 format](https://www.iso.org/iso-4217-currency-codes.html).
	Currency string `json:"currency"`
	// Contains the details of each line item in the tax estimation request.
	LineItems []TaxEstimationLineItemRequest `json:"lineItems"`
}

type _TaxEstimationRequest TaxEstimationRequest

// NewTaxEstimationRequest instantiates a new TaxEstimationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxEstimationRequest(seller Seller, customer Customer, estimateDateTime time.Time, currency string, lineItems []TaxEstimationLineItemRequest) *TaxEstimationRequest {
	this := TaxEstimationRequest{}
	this.Seller = seller
	this.Customer = customer
	this.EstimateDateTime = estimateDateTime
	this.Currency = currency
	this.LineItems = lineItems
	return &this
}

// NewTaxEstimationRequestWithDefaults instantiates a new TaxEstimationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxEstimationRequestWithDefaults() *TaxEstimationRequest {
	this := TaxEstimationRequest{}
	return &this
}

// GetSeller returns the Seller field value
func (o *TaxEstimationRequest) GetSeller() Seller {
	if o == nil {
		var ret Seller
		return ret
	}

	return o.Seller
}

// GetSellerOk returns a tuple with the Seller field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationRequest) GetSellerOk() (*Seller, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seller, true
}

// SetSeller sets field value
func (o *TaxEstimationRequest) SetSeller(v Seller) {
	o.Seller = v
}

// GetCustomer returns the Customer field value
func (o *TaxEstimationRequest) GetCustomer() Customer {
	if o == nil {
		var ret Customer
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationRequest) GetCustomerOk() (*Customer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *TaxEstimationRequest) SetCustomer(v Customer) {
	o.Customer = v
}

// GetEstimateDateTime returns the EstimateDateTime field value
func (o *TaxEstimationRequest) GetEstimateDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EstimateDateTime
}

// GetEstimateDateTimeOk returns a tuple with the EstimateDateTime field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationRequest) GetEstimateDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimateDateTime, true
}

// SetEstimateDateTime sets field value
func (o *TaxEstimationRequest) SetEstimateDateTime(v time.Time) {
	o.EstimateDateTime = v
}

// GetCurrency returns the Currency field value
func (o *TaxEstimationRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TaxEstimationRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetLineItems returns the LineItems field value
func (o *TaxEstimationRequest) GetLineItems() []TaxEstimationLineItemRequest {
	if o == nil {
		var ret []TaxEstimationLineItemRequest
		return ret
	}

	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationRequest) GetLineItemsOk() ([]TaxEstimationLineItemRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineItems, true
}

// SetLineItems sets field value
func (o *TaxEstimationRequest) SetLineItems(v []TaxEstimationLineItemRequest) {
	o.LineItems = v
}

func (o TaxEstimationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxEstimationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["seller"] = o.Seller
	toSerialize["customer"] = o.Customer
	toSerialize["estimateDateTime"] = o.EstimateDateTime
	toSerialize["currency"] = o.Currency
	toSerialize["lineItems"] = o.LineItems
	return toSerialize, nil
}

func (o *TaxEstimationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"seller",
		"customer",
		"estimateDateTime",
		"currency",
		"lineItems",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTaxEstimationRequest := _TaxEstimationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTaxEstimationRequest)

	if err != nil {
		return err
	}

	*o = TaxEstimationRequest(varTaxEstimationRequest)

	return err
}

type NullableTaxEstimationRequest struct {
	value *TaxEstimationRequest
	isSet bool
}

func (v NullableTaxEstimationRequest) Get() *TaxEstimationRequest {
	return v.value
}

func (v *NullableTaxEstimationRequest) Set(val *TaxEstimationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxEstimationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxEstimationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxEstimationRequest(val *TaxEstimationRequest) *NullableTaxEstimationRequest {
	return &NullableTaxEstimationRequest{value: val, isSet: true}
}

func (v NullableTaxEstimationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxEstimationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


