# CredentialValidationResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**status** | [**CredentialStatus**](CredentialStatus.md) |  | 

## Example

```python
from openapi_client.models.credential_validation_response import CredentialValidationResponse

# TODO update the JSON string below
json = "{}"
# create an instance of CredentialValidationResponse from a JSON string
credential_validation_response_instance = CredentialValidationResponse.from_json(json)
# print the JSON string representation of the object
print(CredentialValidationResponse.to_json())

# convert the object into a dict
credential_validation_response_dict = credential_validation_response_instance.to_dict()
# create an instance of CredentialValidationResponse from a dict
credential_validation_response_from_dict = CredentialValidationResponse.from_dict(credential_validation_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


