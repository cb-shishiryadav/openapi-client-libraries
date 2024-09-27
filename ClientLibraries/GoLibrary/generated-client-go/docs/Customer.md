# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the Customer in Chargebee. | [optional] 
**CustomerCode** | **string** | The unique identifier for the Customer in Chargebee. | 
**Address** | [**Address**](Address.md) |  | 
**TaxRegistrationNumber** | Pointer to **NullableString** | The tax registration number of a business in a country. For example, this is the GSTIN for India or the VAT number for EU or Australia. | [optional] 
**TaxIdentifiers** | Pointer to [**[]FieldItem**](FieldItem.md) | It represents the information related to the customer&#39;s tax identifiers. This includes details such as exemption status etc. | [optional] 
**HasNexus** | Pointer to **bool** | Determines whether a tax nexus exists between the Seller and the tax authority at the address provided. | [optional] 
**LocationEvidence** | Pointer to [**CustomerLocationEvidence**](CustomerLocationEvidence.md) |  | [optional] 

## Methods

### NewCustomer

`func NewCustomer(customerCode string, address Address, ) *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Customer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Customer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Customer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Customer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCustomerCode

`func (o *Customer) GetCustomerCode() string`

GetCustomerCode returns the CustomerCode field if non-nil, zero value otherwise.

### GetCustomerCodeOk

`func (o *Customer) GetCustomerCodeOk() (*string, bool)`

GetCustomerCodeOk returns a tuple with the CustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCode

`func (o *Customer) SetCustomerCode(v string)`

SetCustomerCode sets CustomerCode field to given value.


### GetAddress

`func (o *Customer) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Customer) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Customer) SetAddress(v Address)`

SetAddress sets Address field to given value.


### GetTaxRegistrationNumber

`func (o *Customer) GetTaxRegistrationNumber() string`

GetTaxRegistrationNumber returns the TaxRegistrationNumber field if non-nil, zero value otherwise.

### GetTaxRegistrationNumberOk

`func (o *Customer) GetTaxRegistrationNumberOk() (*string, bool)`

GetTaxRegistrationNumberOk returns a tuple with the TaxRegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRegistrationNumber

`func (o *Customer) SetTaxRegistrationNumber(v string)`

SetTaxRegistrationNumber sets TaxRegistrationNumber field to given value.

### HasTaxRegistrationNumber

`func (o *Customer) HasTaxRegistrationNumber() bool`

HasTaxRegistrationNumber returns a boolean if a field has been set.

### SetTaxRegistrationNumberNil

`func (o *Customer) SetTaxRegistrationNumberNil(b bool)`

 SetTaxRegistrationNumberNil sets the value for TaxRegistrationNumber to be an explicit nil

### UnsetTaxRegistrationNumber
`func (o *Customer) UnsetTaxRegistrationNumber()`

UnsetTaxRegistrationNumber ensures that no value is present for TaxRegistrationNumber, not even an explicit nil
### GetTaxIdentifiers

`func (o *Customer) GetTaxIdentifiers() []FieldItem`

GetTaxIdentifiers returns the TaxIdentifiers field if non-nil, zero value otherwise.

### GetTaxIdentifiersOk

`func (o *Customer) GetTaxIdentifiersOk() (*[]FieldItem, bool)`

GetTaxIdentifiersOk returns a tuple with the TaxIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentifiers

`func (o *Customer) SetTaxIdentifiers(v []FieldItem)`

SetTaxIdentifiers sets TaxIdentifiers field to given value.

### HasTaxIdentifiers

`func (o *Customer) HasTaxIdentifiers() bool`

HasTaxIdentifiers returns a boolean if a field has been set.

### GetHasNexus

`func (o *Customer) GetHasNexus() bool`

GetHasNexus returns the HasNexus field if non-nil, zero value otherwise.

### GetHasNexusOk

`func (o *Customer) GetHasNexusOk() (*bool, bool)`

GetHasNexusOk returns a tuple with the HasNexus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNexus

`func (o *Customer) SetHasNexus(v bool)`

SetHasNexus sets HasNexus field to given value.

### HasHasNexus

`func (o *Customer) HasHasNexus() bool`

HasHasNexus returns a boolean if a field has been set.

### GetLocationEvidence

`func (o *Customer) GetLocationEvidence() CustomerLocationEvidence`

GetLocationEvidence returns the LocationEvidence field if non-nil, zero value otherwise.

### GetLocationEvidenceOk

`func (o *Customer) GetLocationEvidenceOk() (*CustomerLocationEvidence, bool)`

GetLocationEvidenceOk returns a tuple with the LocationEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEvidence

`func (o *Customer) SetLocationEvidence(v CustomerLocationEvidence)`

SetLocationEvidence sets LocationEvidence field to given value.

### HasLocationEvidence

`func (o *Customer) HasLocationEvidence() bool`

HasLocationEvidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


