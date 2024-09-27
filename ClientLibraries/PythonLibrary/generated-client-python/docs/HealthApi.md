# openapi_client.HealthApi

All URIs are relative to *https://rest.taxes.provider.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**fetch_health**](HealthApi.md#fetch_health) | **GET** /health | Fetch the health status of the Service Provider.


# **fetch_health**
> HealthCheckResponse fetch_health()

Fetch the health status of the Service Provider.

This endpoint is used to fetch the health status of the Service Provider.

### Example

* Api Key Authentication (apiKey):
* Bearer Authentication (bearerAuth):

```python
import openapi_client
from openapi_client.models.health_check_response import HealthCheckResponse
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
    api_instance = openapi_client.HealthApi(api_client)

    try:
        # Fetch the health status of the Service Provider.
        api_response = api_instance.fetch_health()
        print("The response of HealthApi->fetch_health:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling HealthApi->fetch_health: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**HealthCheckResponse**](HealthCheckResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Service is healthy. |  -  |
**429** | Too many requests. |  -  |
**500** | Unexpected error while processing request. |  -  |
**503** | Service is unhealthy. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

