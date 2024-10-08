/*
Taxes Service Adapter SPI

## Overview Tax Service Adapter is a Service Provider Interface (SPI) that helps Chargebee communicate with API platforms providing the following services: - Address validation - Estimating taxes for sales - Submitting invoices for tax filing  If you have an API that provides the services mentioned above and want Chargebee to be able to integrate with it, you can build an adapter service according to this specification.  ## Terminology Here's a list of terms we've used to describe this specification. ### Merchant    Chargebee’s customer who is carrying out the business of selling products. Also known as the Seller. The Seller has one or more Chargebee accounts. ### Chargebee   The SaaS which manages subscriptions and revenue operations on behalf of the Merchant. Chargebee makes API calls to the Tax Service Adapter for tax estimation and tax filing of transactions between the Merchant and the Customer. ### Customer  Merchant’s customer who purchases products from the Merchant. Their customer record(s) and other details are stored within Chargebee under the Merchant’s account. ### Tax Authority The institution, usually a government, that is responsible for collecting taxes for commercial transactions between Merchants and their Customers.   ### Tax Service Provider The service which provides tax estimation and tax filing services for merchants. ### Tax Service Adapter  The SPI for which the spec is defined in this document. This adapter serves the following purposes: - Translates API requests (as defined in this document) from Chargebee to the request format specified by the Tax Service Provider's API. - Translates API responses from the Tax Service Provider APIs to the format understood by Chargebee (as defined in this document). ### Authorization Chargebee uses HTTP header-based authorization for all the API endpoints associated with Tax Service Adapter. We dynamically pass this authorization key in the HTTP header. The parameter that holds this key is found in the JSON object `api_configuration` required for configuring your onboarding on Chargebee's marketplace. In the `api_configuration` object, our [Taxes Service Adapter SPI](https://chargebee.atlassian.net/l/cp/ca5aZ1mA) checks the authorization key parameter from `credential_configuration.id` and creates the HTTP header-based input query parameter for authorization. The `credential_configuration` is an array of objects with an `id` attribute, and the value of `id` is the parameter containing the authorization key.    Following are the JSON snippets for your reference.     ```json       \"api_configuration\": {           \"api_base_url\": \"https://xyz.abc.com/chargebee\",           \"credential_configuration\": [              {               \"id\":\"authorization_key\",               \"name\": \"Authorization Key\",                \"type\": \"text\",               \"is_sensitive\": true              },             {               \"id\": \"client_secret\",               \"name\": \"Client Secret\",               \"type\": \"text\",               \"is_sensitive\": true             }           ]         }    ```     `\"Authorization\":\"{\"<authorization_key>\":\"test_GykCvTykZFTBqHQAFzjMN9R1Thgdho0Q\"}\"` 

API version: 0.3.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HealthCheckComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckComponent{}

// HealthCheckComponent The health status details of a specific component reported by the Service Adapter.
type HealthCheckComponent struct {
	// The id of the component.
	Id string `json:"id"`
	// The name of the component.
	Name string `json:"name"`
	// The type of component affected when `status` is `WARN` or `DOWN`. The possible values are: - `ADAPTER`: The reported status is for the Service Adapter. - `API`: The reported status is for the Service Provider. - `DATABASE`: The reported status is for a database dependency of the Service Provider. - `SYSTEM`: The reported status is for any other known system component such as cache or gateway. - `OTHER`: The reported status is either for a component that does not belong to the types described above or the source of the issue is unknown. 
	Type string `json:"type"`
	// The detailed status of the component.
	Description *string `json:"description,omitempty"`
	Status HealthStatus `json:"status"`
	// When the `status` of the component is not `UP`, then the list of endpoints affected.
	Endpoints []string `json:"endpoints,omitempty"`
}

type _HealthCheckComponent HealthCheckComponent

// NewHealthCheckComponent instantiates a new HealthCheckComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckComponent(id string, name string, type_ string, status HealthStatus) *HealthCheckComponent {
	this := HealthCheckComponent{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Status = status
	return &this
}

// NewHealthCheckComponentWithDefaults instantiates a new HealthCheckComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckComponentWithDefaults() *HealthCheckComponent {
	this := HealthCheckComponent{}
	return &this
}

// GetId returns the Id field value
func (o *HealthCheckComponent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HealthCheckComponent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HealthCheckComponent) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *HealthCheckComponent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HealthCheckComponent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HealthCheckComponent) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *HealthCheckComponent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HealthCheckComponent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HealthCheckComponent) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HealthCheckComponent) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckComponent) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HealthCheckComponent) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HealthCheckComponent) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
func (o *HealthCheckComponent) GetStatus() HealthStatus {
	if o == nil {
		var ret HealthStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *HealthCheckComponent) GetStatusOk() (*HealthStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *HealthCheckComponent) SetStatus(v HealthStatus) {
	o.Status = v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *HealthCheckComponent) GetEndpoints() []string {
	if o == nil || IsNil(o.Endpoints) {
		var ret []string
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckComponent) GetEndpointsOk() ([]string, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *HealthCheckComponent) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []string and assigns it to the Endpoints field.
func (o *HealthCheckComponent) SetEndpoints(v []string) {
	o.Endpoints = v
}

func (o HealthCheckComponent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Endpoints) {
		toSerialize["endpoints"] = o.Endpoints
	}
	return toSerialize, nil
}

func (o *HealthCheckComponent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
		"status",
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

	varHealthCheckComponent := _HealthCheckComponent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHealthCheckComponent)

	if err != nil {
		return err
	}

	*o = HealthCheckComponent(varHealthCheckComponent)

	return err
}

type NullableHealthCheckComponent struct {
	value *HealthCheckComponent
	isSet bool
}

func (v NullableHealthCheckComponent) Get() *HealthCheckComponent {
	return v.value
}

func (v *NullableHealthCheckComponent) Set(val *HealthCheckComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckComponent(val *HealthCheckComponent) *NullableHealthCheckComponent {
	return &NullableHealthCheckComponent{value: val, isSet: true}
}

func (v NullableHealthCheckComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


