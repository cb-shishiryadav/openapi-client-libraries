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

// checks if the TaxEstimationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxEstimationResponse{}

// TaxEstimationResponse The response sent by the Tax Service Adapter to Chargebee for a tax estimation request.
type TaxEstimationResponse struct {
	Seller Seller `json:"seller"`
	Customer Customer `json:"customer"`
	// The time as of which the tax estimation is to be calculated. This can be a value in the past. For example, if the value is provided as 2022-10-28T15:36:28.129+05:30, then the tax rates applicable on October 28, 2022, at 15:36:28.129, with an offset of +05:30 ahead of UTC/GMT are used for calculations. In case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the 'Z' indicating that the time is in UTC
	EstimateDateTime time.Time `json:"estimateDateTime"`
	// The [currency](https://en.wikipedia.org/wiki/Currency) in the [ISO-4217 format](https://www.iso.org/iso-4217-currency-codes.html).
	Currency string `json:"currency"`
	// The amount after discounts. This is the sum of all `lineItems.subtotal`.
	Subtotal float64 `json:"subtotal"`
	// The part of the `subtotal` that is exempted from tax.
	ExemptAmount float64 `json:"exemptAmount"`
	// The total discount applied. This is the sum of all `lineItems.discount`.
	DiscountAmount float64 `json:"discountAmount"`
	// The part of the `subtotal` that is taxable.
	TaxableAmount float64 `json:"taxableAmount"`
	// The total tax payable. This is the sum of all `lineItems.taxAmount`.
	TaxAmount float64 `json:"taxAmount"`
	// The total after discounts and taxes. This is the same as `subtotal` if it is tax inclusive; otherwise it is `subtotal` + `taxAmount`. `total` can also be expressed as `exemptAmount` + `taxableAmount` + `taxAmount`.
	Total float64 `json:"total"`
	// List of line item details for the tax estimation response.
	LineItems []InvoiceLineItem `json:"lineItems"`
}

type _TaxEstimationResponse TaxEstimationResponse

// NewTaxEstimationResponse instantiates a new TaxEstimationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxEstimationResponse(seller Seller, customer Customer, estimateDateTime time.Time, currency string, subtotal float64, exemptAmount float64, discountAmount float64, taxableAmount float64, taxAmount float64, total float64, lineItems []InvoiceLineItem) *TaxEstimationResponse {
	this := TaxEstimationResponse{}
	this.Seller = seller
	this.Customer = customer
	this.EstimateDateTime = estimateDateTime
	this.Currency = currency
	this.Subtotal = subtotal
	this.ExemptAmount = exemptAmount
	this.DiscountAmount = discountAmount
	this.TaxableAmount = taxableAmount
	this.TaxAmount = taxAmount
	this.Total = total
	this.LineItems = lineItems
	return &this
}

// NewTaxEstimationResponseWithDefaults instantiates a new TaxEstimationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxEstimationResponseWithDefaults() *TaxEstimationResponse {
	this := TaxEstimationResponse{}
	return &this
}

// GetSeller returns the Seller field value
func (o *TaxEstimationResponse) GetSeller() Seller {
	if o == nil {
		var ret Seller
		return ret
	}

	return o.Seller
}

// GetSellerOk returns a tuple with the Seller field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationResponse) GetSellerOk() (*Seller, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seller, true
}

// SetSeller sets field value
func (o *TaxEstimationResponse) SetSeller(v Seller) {
	o.Seller = v
}

// GetCustomer returns the Customer field value
func (o *TaxEstimationResponse) GetCustomer() Customer {
	if o == nil {
		var ret Customer
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationResponse) GetCustomerOk() (*Customer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *TaxEstimationResponse) SetCustomer(v Customer) {
	o.Customer = v
}

// GetEstimateDateTime returns the EstimateDateTime field value
func (o *TaxEstimationResponse) GetEstimateDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EstimateDateTime
}

// GetEstimateDateTimeOk returns a tuple with the EstimateDateTime field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationResponse) GetEstimateDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimateDateTime, true
}

// SetEstimateDateTime sets field value
func (o *TaxEstimationResponse) SetEstimateDateTime(v time.Time) {
	o.EstimateDateTime = v
}

// GetCurrency returns the Currency field value
func (o *TaxEstimationResponse) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationResponse) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TaxEstimationResponse) SetCurrency(v string) {
	o.Currency = v
}

// GetSubtotal returns the Subtotal field value
func (o *TaxEstimationResponse) GetSubtotal() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Subtotal
}

// GetSubtotalOk returns a tuple with the Subtotal field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationResponse) GetSubtotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtotal, true
}

// SetSubtotal sets field value
func (o *TaxEstimationResponse) SetSubtotal(v float64) {
	o.Subtotal = v
}

// GetExemptAmount returns the ExemptAmount field value
func (o *TaxEstimationResponse) GetExemptAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ExemptAmount
}

// GetExemptAmountOk returns a tuple with the ExemptAmount field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationResponse) GetExemptAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExemptAmount, true
}

// SetExemptAmount sets field value
func (o *TaxEstimationResponse) SetExemptAmount(v float64) {
	o.ExemptAmount = v
}

// GetDiscountAmount returns the DiscountAmount field value
func (o *TaxEstimationResponse) GetDiscountAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationResponse) GetDiscountAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountAmount, true
}

// SetDiscountAmount sets field value
func (o *TaxEstimationResponse) SetDiscountAmount(v float64) {
	o.DiscountAmount = v
}

// GetTaxableAmount returns the TaxableAmount field value
func (o *TaxEstimationResponse) GetTaxableAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.TaxableAmount
}

// GetTaxableAmountOk returns a tuple with the TaxableAmount field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationResponse) GetTaxableAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxableAmount, true
}

// SetTaxableAmount sets field value
func (o *TaxEstimationResponse) SetTaxableAmount(v float64) {
	o.TaxableAmount = v
}

// GetTaxAmount returns the TaxAmount field value
func (o *TaxEstimationResponse) GetTaxAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationResponse) GetTaxAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxAmount, true
}

// SetTaxAmount sets field value
func (o *TaxEstimationResponse) SetTaxAmount(v float64) {
	o.TaxAmount = v
}

// GetTotal returns the Total field value
func (o *TaxEstimationResponse) GetTotal() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationResponse) GetTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *TaxEstimationResponse) SetTotal(v float64) {
	o.Total = v
}

// GetLineItems returns the LineItems field value
func (o *TaxEstimationResponse) GetLineItems() []InvoiceLineItem {
	if o == nil {
		var ret []InvoiceLineItem
		return ret
	}

	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value
// and a boolean to check if the value has been set.
func (o *TaxEstimationResponse) GetLineItemsOk() ([]InvoiceLineItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineItems, true
}

// SetLineItems sets field value
func (o *TaxEstimationResponse) SetLineItems(v []InvoiceLineItem) {
	o.LineItems = v
}

func (o TaxEstimationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxEstimationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["seller"] = o.Seller
	toSerialize["customer"] = o.Customer
	toSerialize["estimateDateTime"] = o.EstimateDateTime
	toSerialize["currency"] = o.Currency
	toSerialize["subtotal"] = o.Subtotal
	toSerialize["exemptAmount"] = o.ExemptAmount
	toSerialize["discountAmount"] = o.DiscountAmount
	toSerialize["taxableAmount"] = o.TaxableAmount
	toSerialize["taxAmount"] = o.TaxAmount
	toSerialize["total"] = o.Total
	toSerialize["lineItems"] = o.LineItems
	return toSerialize, nil
}

func (o *TaxEstimationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"seller",
		"customer",
		"estimateDateTime",
		"currency",
		"subtotal",
		"exemptAmount",
		"discountAmount",
		"taxableAmount",
		"taxAmount",
		"total",
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

	varTaxEstimationResponse := _TaxEstimationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTaxEstimationResponse)

	if err != nil {
		return err
	}

	*o = TaxEstimationResponse(varTaxEstimationResponse)

	return err
}

type NullableTaxEstimationResponse struct {
	value *TaxEstimationResponse
	isSet bool
}

func (v NullableTaxEstimationResponse) Get() *TaxEstimationResponse {
	return v.value
}

func (v *NullableTaxEstimationResponse) Set(val *TaxEstimationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxEstimationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxEstimationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxEstimationResponse(val *TaxEstimationResponse) *NullableTaxEstimationResponse {
	return &NullableTaxEstimationResponse{value: val, isSet: true}
}

func (v NullableTaxEstimationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxEstimationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


