# CheckAddressTaxabilityRequest

The taxability request containing the address. Postal code & Country is mandatory.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | [**Address**](Address.md) |  | [optional] 

## Example

```python
from openapi_client.models.check_address_taxability_request import CheckAddressTaxabilityRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CheckAddressTaxabilityRequest from a JSON string
check_address_taxability_request_instance = CheckAddressTaxabilityRequest.from_json(json)
# print the JSON string representation of the object
print(CheckAddressTaxabilityRequest.to_json())

# convert the object into a dict
check_address_taxability_request_dict = check_address_taxability_request_instance.to_dict()
# create an instance of CheckAddressTaxabilityRequest from a dict
check_address_taxability_request_from_dict = CheckAddressTaxabilityRequest.from_dict(check_address_taxability_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


