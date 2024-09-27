# openapi_client.AuthenticationApi

All URIs are relative to *https://rest.taxes.provider.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**validate_credentials**](AuthenticationApi.md#validate_credentials) | **POST** /credentials/validate | Validate credentials


# **validate_credentials**
> CredentialValidationResponse validate_credentials()

Validate credentials

This endpoint is used to validate the credentials used to call the Service Provider.

### Example

* Api Key Authentication (apiKey):
* Bearer Authentication (bearerAuth):

```python
import openapi_client
from openapi_client.models.credential_validation_response import CredentialValidationResponse
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://rest.taxes.provider.com/api/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://rest.taxes.provider.com/api/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: apiKey
configuration.api_key['apiKey'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apiKey'] = 'Bearer'

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.AuthenticationApi(api_client)

    try:
        # Validate credentials
        api_response = api_instance.validate_credentials()
        print("The response of AuthenticationApi->validate_credentials:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AuthenticationApi->validate_credentials: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**CredentialValidationResponse**](CredentialValidationResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Authentication succeeded. |  -  |
**401** | Authentication failed. |  -  |
**429** | Too many requests. |  -  |
**500** | Unexpected error while processing request. |  -  |
**503** | Service is unhealthy. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

