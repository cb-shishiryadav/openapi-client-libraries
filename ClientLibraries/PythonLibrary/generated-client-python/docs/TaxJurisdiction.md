# TaxJurisdiction

The tax jurisdiction details.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**code** | **str** | The jurisdiction code. | 
**type** | [**TaxJurisdictionType**](TaxJurisdictionType.md) |  | 
**name** | **str** | The jurisdiction name. | 

## Example

```python
from openapi_client.models.tax_jurisdiction import TaxJurisdiction

# TODO update the JSON string below
json = "{}"
# create an instance of TaxJurisdiction from a JSON string
tax_jurisdiction_instance = TaxJurisdiction.from_json(json)
# print the JSON string representation of the object
print(TaxJurisdiction.to_json())

# convert the object into a dict
tax_jurisdiction_dict = tax_jurisdiction_instance.to_dict()
# create an instance of TaxJurisdiction from a dict
tax_jurisdiction_from_dict = TaxJurisdiction.from_dict(tax_jurisdiction_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


