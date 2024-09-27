# CheckAddressTaxabilityResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**is_taxable** | **bool** | The taxability of the address. | 

## Example

```python
from openapi_client.models.check_address_taxability_response import CheckAddressTaxabilityResponse

# TODO update the JSON string below
json = "{}"
# create an instance of CheckAddressTaxabilityResponse from a JSON string
check_address_taxability_response_instance = CheckAddressTaxabilityResponse.from_json(json)
# print the JSON string representation of the object
print(CheckAddressTaxabilityResponse.to_json())

# convert the object into a dict
check_address_taxability_response_dict = check_address_taxability_response_instance.to_dict()
# create an instance of CheckAddressTaxabilityResponse from a dict
check_address_taxability_response_from_dict = CheckAddressTaxabilityResponse.from_dict(check_address_taxability_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


