// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Taxes Service Adapter SPI
 *
 * ## Overview Tax Service Adapter is a Service Provider Interface (SPI) that helps Chargebee communicate with API platforms providing the following services: - Address validation - Estimating taxes for sales - Submitting invoices for tax filing  If you have an API that provides the services mentioned above and want Chargebee to be able to integrate with it, you can build an adapter service according to this specification.  ## Terminology Here's a list of terms we've used to describe this specification. ### Merchant    Chargebee’s customer who is carrying out the business of selling products. Also known as the Seller. The Seller has one or more Chargebee accounts. ### Chargebee   The SaaS which manages subscriptions and revenue operations on behalf of the Merchant. Chargebee makes API calls to the Tax Service Adapter for tax estimation and tax filing of transactions between the Merchant and the Customer. ### Customer  Merchant’s customer who purchases products from the Merchant. Their customer record(s) and other details are stored within Chargebee under the Merchant’s account. ### Tax Authority The institution, usually a government, that is responsible for collecting taxes for commercial transactions between Merchants and their Customers.   ### Tax Service Provider The service which provides tax estimation and tax filing services for merchants. ### Tax Service Adapter  The SPI for which the spec is defined in this document. This adapter serves the following purposes: - Translates API requests (as defined in this document) from Chargebee to the request format specified by the Tax Service Provider's API. - Translates API responses from the Tax Service Provider APIs to the format understood by Chargebee (as defined in this document). ### Authorization Chargebee uses HTTP header-based authorization for all the API endpoints associated with Tax Service Adapter. We dynamically pass this authorization key in the HTTP header. The parameter that holds this key is found in the JSON object `api_configuration` required for configuring your onboarding on Chargebee's marketplace. In the `api_configuration` object, our [Taxes Service Adapter SPI](https://chargebee.atlassian.net/l/cp/ca5aZ1mA) checks the authorization key parameter from `credential_configuration.id` and creates the HTTP header-based input query parameter for authorization. The `credential_configuration` is an array of objects with an `id` attribute, and the value of `id` is the parameter containing the authorization key.    Following are the JSON snippets for your reference.     ```json       \"api_configuration\": {           \"api_base_url\": \"https://xyz.abc.com/chargebee\",           \"credential_configuration\": [              {               \"id\":\"authorization_key\",               \"name\": \"Authorization Key\",                \"type\": \"text\",               \"is_sensitive\": true              },             {               \"id\": \"client_secret\",               \"name\": \"Client Secret\",               \"type\": \"text\",               \"is_sensitive\": true             }           ]         }    ```     `\"Authorization\":\"{\"<authorization_key>\":\"test_GykCvTykZFTBqHQAFzjMN9R1Thgdho0Q\"}\"` 
 *
 * API version: 0.3.7
 */

package openapi


import (
	"fmt"
)


// TaxExemptType : The tax exemption type for a line item. This is a mandatory parameter while applying tax exemption on any line-item.
type TaxExemptType string

// List of TaxExemptType
const (
	PRODUCT_EXEMPT TaxExemptType = "PRODUCT_EXEMPT"
	CUSTOMER_EXEMPT TaxExemptType = "CUSTOMER_EXEMPT"
	REGION_EXEMPT TaxExemptType = "REGION_EXEMPT"
	REVERSE_CHARGE TaxExemptType = "REVERSE_CHARGE"
	ZERO_RATE_TAX TaxExemptType = "ZERO_RATE_TAX"
	HIGH_VALUE_PHYSICAL_GOODS TaxExemptType = "HIGH_VALUE_PHYSICAL_GOODS"
	EXPORT TaxExemptType = "EXPORT"
	ZERO_VALUE_ITEM TaxExemptType = "ZERO_VALUE_ITEM"
	TAX_NOT_CONFIGURED TaxExemptType = "TAX_NOT_CONFIGURED"
)

// AllowedTaxExemptTypeEnumValues is all the allowed values of TaxExemptType enum
var AllowedTaxExemptTypeEnumValues = []TaxExemptType{
	"PRODUCT_EXEMPT",
	"CUSTOMER_EXEMPT",
	"REGION_EXEMPT",
	"REVERSE_CHARGE",
	"ZERO_RATE_TAX",
	"HIGH_VALUE_PHYSICAL_GOODS",
	"EXPORT",
	"ZERO_VALUE_ITEM",
	"TAX_NOT_CONFIGURED",
}

// validTaxExemptTypeEnumValue provides a map of TaxExemptTypes for fast verification of use input
var validTaxExemptTypeEnumValues = map[TaxExemptType]struct{}{
	"PRODUCT_EXEMPT": {},
	"CUSTOMER_EXEMPT": {},
	"REGION_EXEMPT": {},
	"REVERSE_CHARGE": {},
	"ZERO_RATE_TAX": {},
	"HIGH_VALUE_PHYSICAL_GOODS": {},
	"EXPORT": {},
	"ZERO_VALUE_ITEM": {},
	"TAX_NOT_CONFIGURED": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaxExemptType) IsValid() bool {
	_, ok := validTaxExemptTypeEnumValues[v]
	return ok
}

// NewTaxExemptTypeFromValue returns a pointer to a valid TaxExemptType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaxExemptTypeFromValue(v string) (TaxExemptType, error) {
	ev := TaxExemptType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for TaxExemptType: valid values are %v", v, AllowedTaxExemptTypeEnumValues)
}



// AssertTaxExemptTypeRequired checks if the required fields are not zero-ed
func AssertTaxExemptTypeRequired(obj TaxExemptType) error {
	return nil
}

// AssertTaxExemptTypeConstraints checks if the values respects the defined constraints
func AssertTaxExemptTypeConstraints(obj TaxExemptType) error {
	return nil
}
