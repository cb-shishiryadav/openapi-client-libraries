# Seller

The details of the seller involved in the transaction including company code and address.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tax_registration_number** | **str** | The tax registration number of a business in a country. For example, this is the GSTIN for India or the VAT number for EU or Australia. | [optional] 
**address** | [**Address**](Address.md) |  | 
**has_nexus** | **bool** | Determines whether a tax nexus exists between the Seller and the tax authority at the address provided. | [optional] 

## Example

```python
from openapi_client.models.seller import Seller

# TODO update the JSON string below
json = "{}"
# create an instance of Seller from a JSON string
seller_instance = Seller.from_json(json)
# print the JSON string representation of the object
print(Seller.to_json())

# convert the object into a dict
seller_dict = seller_instance.to_dict()
# create an instance of Seller from a dict
seller_from_dict = Seller.from_dict(seller_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


