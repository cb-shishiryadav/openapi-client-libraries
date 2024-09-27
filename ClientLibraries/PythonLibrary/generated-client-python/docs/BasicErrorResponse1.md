# BasicErrorResponse1

The basic error response containing the error message and the help documentation link.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message** | **str** | The description of the error with details about it&#39;s cause. | 
**help_url** | **str** | The link to the documentation for more information about the error and the corrective action. | [optional] 

## Example

```python
from openapi_client.models.basic_error_response1 import BasicErrorResponse1

# TODO update the JSON string below
json = "{}"
# create an instance of BasicErrorResponse1 from a JSON string
basic_error_response1_instance = BasicErrorResponse1.from_json(json)
# print the JSON string representation of the object
print(BasicErrorResponse1.to_json())

# convert the object into a dict
basic_error_response1_dict = basic_error_response1_instance.to_dict()
# create an instance of BasicErrorResponse1 from a dict
basic_error_response1_from_dict = BasicErrorResponse1.from_dict(basic_error_response1_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


