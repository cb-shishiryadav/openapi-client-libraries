# openapi_client.AddressApi

All URIs are relative to *https://rest.taxes.provider.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**check_address_taxability**](AddressApi.md#check_address_taxability) | **POST** /address/check-taxability | Check taxability
[**validate_address**](AddressApi.md#validate_address) | **POST** /address/validate | Address validation


# **check_address_taxability**
> CheckAddressTaxabilityResponse check_address_taxability(check_address_taxability_request=check_address_taxability_request)

Check taxability

Checks whether the tax address is valid in terms of tax calculation. This endpoint checks whether the address information of the customer is sufficient for the tax provider to return a tax rate. It does not consider the nexus status of the merchant and is mandatory to integrate for the tax provider.

### Example

* Api Key Authentication (apiKey):
* Bearer Authentication (bearerAuth):

```python
import openapi_client
from openapi_client.models.check_address_taxability_request import CheckAddressTaxabilityRequest
from openapi_client.models.check_address_taxability_response import CheckAddressTaxabilityResponse
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
    api_instance = openapi_client.AddressApi(api_client)
    check_address_taxability_request = openapi_client.CheckAddressTaxabilityRequest() # CheckAddressTaxabilityRequest |  (optional)

    try:
        # Check taxability
        api_response = api_instance.check_address_taxability(check_address_taxability_request=check_address_taxability_request)
        print("The response of AddressApi->check_address_taxability:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AddressApi->check_address_taxability: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **check_address_taxability_request** | [**CheckAddressTaxabilityRequest**](CheckAddressTaxabilityRequest.md)|  | [optional] 

### Return type

[**CheckAddressTaxabilityResponse**](CheckAddressTaxabilityResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Tax can be calculated for the address provided. |  -  |
**400** | Bad request. |  -  |
**401** | Unauthenticated request. |  -  |
**403** | Unauthorized request. |  -  |
**429** | Too many requests. |  -  |
**500** | Unexpected error while processing request. |  -  |
**503** | Service is unhealthy. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **validate_address**
> AddressValidationResponse validate_address(address_validation_request=address_validation_request)

Address validation

Checks whether a given address is a valid delivery address for shipping purposes. The tax provider can decide whether to mention the full or valid address depending on their requirement.

### Example

* Api Key Authentication (apiKey):
* Bearer Authentication (bearerAuth):

```python
import openapi_client
from openapi_client.models.address_validation_request import AddressValidationRequest
from openapi_client.models.address_validation_response import AddressValidationResponse
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
    api_instance = openapi_client.AddressApi(api_client)
    address_validation_request = openapi_client.AddressValidationRequest() # AddressValidationRequest |  (optional)

    try:
        # Address validation
        api_response = api_instance.validate_address(address_validation_request=address_validation_request)
        print("The response of AddressApi->validate_address:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AddressApi->validate_address: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address_validation_request** | [**AddressValidationRequest**](AddressValidationRequest.md)|  | [optional] 

### Return type

[**AddressValidationResponse**](AddressValidationResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Address is validated successfully |  -  |
**400** | Bad request. |  -  |
**401** | Unauthenticated request. |  -  |
**403** | Unauthorized request. |  -  |
**429** | Too many requests. |  -  |
**500** | Unexpected error while processing request. |  -  |
**503** | Service is unhealthy. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

