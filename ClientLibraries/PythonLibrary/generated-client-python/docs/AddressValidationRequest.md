# AddressValidationRequest

The verification request containing the address. The following fields are mandatory -   - line1   - city   - postalCode   - state   - country

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | [**Address**](Address.md) |  | [optional] 

## Example

```python
from openapi_client.models.address_validation_request import AddressValidationRequest

# TODO update the JSON string below
json = "{}"
# create an instance of AddressValidationRequest from a JSON string
address_validation_request_instance = AddressValidationRequest.from_json(json)
# print the JSON string representation of the object
print(AddressValidationRequest.to_json())

# convert the object into a dict
address_validation_request_dict = address_validation_request_instance.to_dict()
# create an instance of AddressValidationRequest from a dict
address_validation_request_from_dict = AddressValidationRequest.from_dict(address_validation_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


