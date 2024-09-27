/*
Taxes Service Adapter SPI

## Overview Tax Service Adapter is a Service Provider Interface (SPI) that helps Chargebee communicate with API platforms providing the following services: - Address validation - Estimating taxes for sales - Submitting invoices for tax filing  If you have an API that provides the services mentioned above and want Chargebee to be able to integrate with it, you can build an adapter service according to this specification.  ## Terminology Here's a list of terms we've used to describe this specification. ### Merchant    Chargebee’s customer who is carrying out the business of selling products. Also known as the Seller. The Seller has one or more Chargebee accounts. ### Chargebee   The SaaS which manages subscriptions and revenue operations on behalf of the Merchant. Chargebee makes API calls to the Tax Service Adapter for tax estimation and tax filing of transactions between the Merchant and the Customer. ### Customer  Merchant’s customer who purchases products from the Merchant. Their customer record(s) and other details are stored within Chargebee under the Merchant’s account. ### Tax Authority The institution, usually a government, that is responsible for collecting taxes for commercial transactions between Merchants and their Customers.   ### Tax Service Provider The service which provides tax estimation and tax filing services for merchants. ### Tax Service Adapter  The SPI for which the spec is defined in this document. This adapter serves the following purposes: - Translates API requests (as defined in this document) from Chargebee to the request format specified by the Tax Service Provider's API. - Translates API responses from the Tax Service Provider APIs to the format understood by Chargebee (as defined in this document). ### Authorization Chargebee uses HTTP header-based authorization for all the API endpoints associated with Tax Service Adapter. We dynamically pass this authorization key in the HTTP header. The parameter that holds this key is found in the JSON object `api_configuration` required for configuring your onboarding on Chargebee's marketplace. In the `api_configuration` object, our [Taxes Service Adapter SPI](https://chargebee.atlassian.net/l/cp/ca5aZ1mA) checks the authorization key parameter from `credential_configuration.id` and creates the HTTP header-based input query parameter for authorization. The `credential_configuration` is an array of objects with an `id` attribute, and the value of `id` is the parameter containing the authorization key.    Following are the JSON snippets for your reference.     ```json       \"api_configuration\": {           \"api_base_url\": \"https://xyz.abc.com/chargebee\",           \"credential_configuration\": [              {               \"id\":\"authorization_key\",               \"name\": \"Authorization Key\",                \"type\": \"text\",               \"is_sensitive\": true              },             {               \"id\": \"client_secret\",               \"name\": \"Client Secret\",               \"type\": \"text\",               \"is_sensitive\": true             }           ]         }    ```     `\"Authorization\":\"{\"<authorization_key>\":\"test_GykCvTykZFTBqHQAFzjMN9R1Thgdho0Q\"}\"` 

API version: 0.3.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AddressValidationStatus The validation status of an address.
type AddressValidationStatus string

// List of AddressValidationStatus
const (
	VALID AddressValidationStatus = "VALID"
	INVALID AddressValidationStatus = "INVALID"
)

// All allowed values of AddressValidationStatus enum
var AllowedAddressValidationStatusEnumValues = []AddressValidationStatus{
	"VALID",
	"INVALID",
}

func (v *AddressValidationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AddressValidationStatus(value)
	for _, existing := range AllowedAddressValidationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AddressValidationStatus", value)
}

// NewAddressValidationStatusFromValue returns a pointer to a valid AddressValidationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAddressValidationStatusFromValue(v string) (*AddressValidationStatus, error) {
	ev := AddressValidationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AddressValidationStatus: valid values are %v", v, AllowedAddressValidationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AddressValidationStatus) IsValid() bool {
	for _, existing := range AllowedAddressValidationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AddressValidationStatus value
func (v AddressValidationStatus) Ptr() *AddressValidationStatus {
	return &v
}

type NullableAddressValidationStatus struct {
	value *AddressValidationStatus
	isSet bool
}

func (v NullableAddressValidationStatus) Get() *AddressValidationStatus {
	return v.value
}

func (v *NullableAddressValidationStatus) Set(val *AddressValidationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressValidationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressValidationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressValidationStatus(val *AddressValidationStatus) *NullableAddressValidationStatus {
	return &NullableAddressValidationStatus{value: val, isSet: true}
}

func (v NullableAddressValidationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressValidationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

