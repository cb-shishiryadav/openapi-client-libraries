# openapi_client.TaxEstimateApi

All URIs are relative to *https://rest.taxes.provider.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**estimate_taxes**](TaxEstimateApi.md#estimate_taxes) | **POST** /tax-estimate | Estimate tax


# **estimate_taxes**
> TaxEstimationResponse estimate_taxes(tax_estimation_request=tax_estimation_request)

Estimate tax

This endpoint is used to estimate taxes for a set of line items being sold by the Merchant to a Customer.

### Example

* Api Key Authentication (apiKey):
* Bearer Authentication (bearerAuth):

```python
import openapi_client
from openapi_client.models.tax_estimation_request import TaxEstimationRequest
from openapi_client.models.tax_estimation_response import TaxEstimationResponse
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
    api_instance = openapi_client.TaxEstimateApi(api_client)
    tax_estimation_request = openapi_client.TaxEstimationRequest() # TaxEstimationRequest |  (optional)

    try:
        # Estimate tax
        api_response = api_instance.estimate_taxes(tax_estimation_request=tax_estimation_request)
        print("The response of TaxEstimateApi->estimate_taxes:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TaxEstimateApi->estimate_taxes: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tax_estimation_request** | [**TaxEstimationRequest**](TaxEstimationRequest.md)|  | [optional] 

### Return type

[**TaxEstimationResponse**](TaxEstimationResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Tax estimated successfully for given line items. |  -  |
**400** | Bad request. |  -  |
**401** | Unauthenticated request. |  -  |
**403** | Unauthorized request. |  -  |
**404** | Not found. |  -  |
**429** | Too many requests. |  -  |
**500** | Unexpected error while processing request. |  -  |
**503** | Service is unhealthy. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

