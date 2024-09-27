# Seller

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxRegistrationNumber** | Pointer to **NullableString** | The tax registration number of a business in a country. For example, this is the GSTIN for India or the VAT number for EU or Australia. | [optional] 
**Address** | [**Address**](Address.md) |  | 
**HasNexus** | Pointer to **bool** | Determines whether a tax nexus exists between the Seller and the tax authority at the address provided. | [optional] 

## Methods

### NewSeller

`func NewSeller(address Address, ) *Seller`

NewSeller instantiates a new Seller object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellerWithDefaults

`func NewSellerWithDefaults() *Seller`

NewSellerWithDefaults instantiates a new Seller object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxRegistrationNumber

`func (o *Seller) GetTaxRegistrationNumber() string`

GetTaxRegistrationNumber returns the TaxRegistrationNumber field if non-nil, zero value otherwise.

### GetTaxRegistrationNumberOk

`func (o *Seller) GetTaxRegistrationNumberOk() (*string, bool)`

GetTaxRegistrationNumberOk returns a tuple with the TaxRegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRegistrationNumber

`func (o *Seller) SetTaxRegistrationNumber(v string)`

SetTaxRegistrationNumber sets TaxRegistrationNumber field to given value.

### HasTaxRegistrationNumber

`func (o *Seller) HasTaxRegistrationNumber() bool`

HasTaxRegistrationNumber returns a boolean if a field has been set.

### SetTaxRegistrationNumberNil

`func (o *Seller) SetTaxRegistrationNumberNil(b bool)`

 SetTaxRegistrationNumberNil sets the value for TaxRegistrationNumber to be an explicit nil

### UnsetTaxRegistrationNumber
`func (o *Seller) UnsetTaxRegistrationNumber()`

UnsetTaxRegistrationNumber ensures that no value is present for TaxRegistrationNumber, not even an explicit nil
### GetAddress

`func (o *Seller) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Seller) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Seller) SetAddress(v Address)`

SetAddress sets Address field to given value.


### GetHasNexus

`func (o *Seller) GetHasNexus() bool`

GetHasNexus returns the HasNexus field if non-nil, zero value otherwise.

### GetHasNexusOk

`func (o *Seller) GetHasNexusOk() (*bool, bool)`

GetHasNexusOk returns a tuple with the HasNexus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNexus

`func (o *Seller) SetHasNexus(v bool)`

SetHasNexus sets HasNexus field to given value.

### HasHasNexus

`func (o *Seller) HasHasNexus() bool`

HasHasNexus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


