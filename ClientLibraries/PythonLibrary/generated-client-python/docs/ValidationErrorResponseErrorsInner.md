# ValidationErrorResponseErrorsInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**entity** | **str** | The target entity that has the invalid field or value. | [optional] 
**entity_field** | **str** | The field of an entity that has the invalid value. | [optional] 
**code** | [**ErrorCode**](ErrorCode.md) |  | 
**message** | **str** | A short message describing the reason for the error. | 
**help_url** | **str** | The link to the documentation for more information about the error and the corrective action. | [optional] 

## Example

```python
from openapi_client.models.validation_error_response_errors_inner import ValidationErrorResponseErrorsInner

# TODO update the JSON string below
json = "{}"
# create an instance of ValidationErrorResponseErrorsInner from a JSON string
validation_error_response_errors_inner_instance = ValidationErrorResponseErrorsInner.from_json(json)
# print the JSON string representation of the object
print(ValidationErrorResponseErrorsInner.to_json())

# convert the object into a dict
validation_error_response_errors_inner_dict = validation_error_response_errors_inner_instance.to_dict()
# create an instance of ValidationErrorResponseErrorsInner from a dict
validation_error_response_errors_inner_from_dict = ValidationErrorResponseErrorsInner.from_dict(validation_error_response_errors_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


