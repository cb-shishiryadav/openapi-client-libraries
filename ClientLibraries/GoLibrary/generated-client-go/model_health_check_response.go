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

// checks if the HealthCheckResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckResponse{}

// HealthCheckResponse struct for HealthCheckResponse
type HealthCheckResponse struct {
	Version *string `json:"version,omitempty"`
	// The description of the health status returned by the Service Adapter.
	Description *string `json:"description,omitempty"`
	Status HealthStatus `json:"status"`
	// List of health status details for each component reported by the Service Adapter.
	Components []HealthCheckComponent `json:"components"`
	// The timestamp of the health status reported by the Service Adapter.
	Time time.Time `json:"time"`
}

type _HealthCheckResponse HealthCheckResponse

// NewHealthCheckResponse instantiates a new HealthCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckResponse(status HealthStatus, components []HealthCheckComponent, time time.Time) *HealthCheckResponse {
	this := HealthCheckResponse{}
	this.Status = status
	this.Components = components
	this.Time = time
	return &this
}

// NewHealthCheckResponseWithDefaults instantiates a new HealthCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckResponseWithDefaults() *HealthCheckResponse {
	this := HealthCheckResponse{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HealthCheckResponse) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HealthCheckResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HealthCheckResponse) SetVersion(v string) {
	o.Version = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HealthCheckResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HealthCheckResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HealthCheckResponse) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
func (o *HealthCheckResponse) GetStatus() HealthStatus {
	if o == nil {
		var ret HealthStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetStatusOk() (*HealthStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *HealthCheckResponse) SetStatus(v HealthStatus) {
	o.Status = v
}

// GetComponents returns the Components field value
func (o *HealthCheckResponse) GetComponents() []HealthCheckComponent {
	if o == nil {
		var ret []HealthCheckComponent
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetComponentsOk() ([]HealthCheckComponent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// SetComponents sets field value
func (o *HealthCheckResponse) SetComponents(v []HealthCheckComponent) {
	o.Components = v
}

// GetTime returns the Time field value
func (o *HealthCheckResponse) GetTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *HealthCheckResponse) SetTime(v time.Time) {
	o.Time = v
}

func (o HealthCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["status"] = o.Status
	toSerialize["components"] = o.Components
	toSerialize["time"] = o.Time
	return toSerialize, nil
}

func (o *HealthCheckResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"components",
		"time",
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

	varHealthCheckResponse := _HealthCheckResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHealthCheckResponse)

	if err != nil {
		return err
	}

	*o = HealthCheckResponse(varHealthCheckResponse)

	return err
}

type NullableHealthCheckResponse struct {
	value *HealthCheckResponse
	isSet bool
}

func (v NullableHealthCheckResponse) Get() *HealthCheckResponse {
	return v.value
}

func (v *NullableHealthCheckResponse) Set(val *HealthCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckResponse(val *HealthCheckResponse) *NullableHealthCheckResponse {
	return &NullableHealthCheckResponse{value: val, isSet: true}
}

func (v NullableHealthCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


