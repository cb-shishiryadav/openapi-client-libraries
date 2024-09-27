# Customer

The details of the Customer.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of the Customer in Chargebee. | [optional] 
**customer_code** | **str** | The unique identifier for the Customer in Chargebee. | 
**address** | [**Address**](Address.md) |  | 
**tax_registration_number** | **str** | The tax registration number of a business in a country. For example, this is the GSTIN for India or the VAT number for EU or Australia. | [optional] 
**tax_identifiers** | [**List[FieldItem]**](FieldItem.md) | It represents the information related to the customer&#39;s tax identifiers. This includes details such as exemption status etc. | [optional] 
**has_nexus** | **bool** | Determines whether a tax nexus exists between the Seller and the tax authority at the address provided. | [optional] 
**location_evidence** | [**CustomerLocationEvidence**](CustomerLocationEvidence.md) |  | [optional] 

## Example

```python
from openapi_client.models.customer import Customer

# TODO update the JSON string below
json = "{}"
# create an instance of Customer from a JSON string
customer_instance = Customer.from_json(json)
# print the JSON string representation of the object
print(Customer.to_json())

# convert the object into a dict
customer_dict = customer_instance.to_dict()
# create an instance of Customer from a dict
customer_from_dict = Customer.from_dict(customer_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


