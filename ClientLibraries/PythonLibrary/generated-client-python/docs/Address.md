# Address

Represents the address used for validation.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**line1** | **str** | First line of the street address | [optional] 
**line2** | **str** | Second line of the street address | [optional] 
**line3** | **str** | Third line of the street address | [optional] 
**city** | **str** | The city of the address | [optional] 
**state** | **str** | The state of the address following the ISO 3166-2 state/province code without the country prefix. | [optional] 
**postal_code** | **str** | Postal Code / Zip Code of the address. | [optional] 
**country** | **str** | The country of the address following the ISO 3166-1 alpha-2 standard. | [optional] 

## Example

```python
from openapi_client.models.address import Address

# TODO update the JSON string below
json = "{}"
# create an instance of Address from a JSON string
address_instance = Address.from_json(json)
# print the JSON string representation of the object
print(Address.to_json())

# convert the object into a dict
address_dict = address_instance.to_dict()
# create an instance of Address from a dict
address_from_dict = Address.from_dict(address_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


