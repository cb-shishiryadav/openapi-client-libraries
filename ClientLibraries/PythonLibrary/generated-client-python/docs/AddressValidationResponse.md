# AddressValidationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**status** | [**AddressValidationStatus**](AddressValidationStatus.md) |  | 

## Example

```python
from openapi_client.models.address_validation_response import AddressValidationResponse

# TODO update the JSON string below
json = "{}"
# create an instance of AddressValidationResponse from a JSON string
address_validation_response_instance = AddressValidationResponse.from_json(json)
# print the JSON string representation of the object
print(AddressValidationResponse.to_json())

# convert the object into a dict
address_validation_response_dict = address_validation_response_instance.to_dict()
# create an instance of AddressValidationResponse from a dict
address_validation_response_from_dict = AddressValidationResponse.from_dict(address_validation_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


