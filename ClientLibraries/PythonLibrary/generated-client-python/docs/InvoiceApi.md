# openapi_client.InvoiceApi

All URIs are relative to *https://rest.taxes.provider.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**commit_invoice**](InvoiceApi.md#commit_invoice) | **POST** /invoices/{invoiceId}/commit | Commit Invoice
[**create_invoice**](InvoiceApi.md#create_invoice) | **POST** /invoices | Create Invoice
[**fetch_invoice**](InvoiceApi.md#fetch_invoice) | **GET** /invoices/{invoiceId} | Retrieve Invoice
[**void_invoice**](InvoiceApi.md#void_invoice) | **POST** /invoices/{invoiceId}/void | Void Invoice


# **commit_invoice**
> commit_invoice(invoice_id)

Commit Invoice

This endpoint is used to commit an invoice for a given invoice id. Once committed, the invoice is considered to be finalized.

### Example

* Api Key Authentication (apiKey):
* Bearer Authentication (bearerAuth):

```python
import openapi_client
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
    api_instance = openapi_client.InvoiceApi(api_client)
    invoice_id = 'invoice_id_example' # str | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider.

    try:
        # Commit Invoice
        api_instance.commit_invoice(invoice_id)
    except Exception as e:
        print("Exception when calling InvoiceApi->commit_invoice: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoice_id** | **str**| The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider. | 

### Return type

void (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | Invoice committed successfully. |  -  |
**401** | Unauthenticated request. |  -  |
**403** | Unauthorized request. |  -  |
**404** | Not found. |  -  |
**429** | Too many requests. |  -  |
**500** | Unexpected error while processing request. |  -  |
**503** | Service is unhealthy. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_invoice**
> Invoice create_invoice(invoice_request=invoice_request)

Create Invoice

This endpoint is used to send an invoice to the Tax Service Provider. Invoices created in Chargebee are statements of amounts owed by the Customer to the Merchant for a specific purchase.

### Example

* Api Key Authentication (apiKey):
* Bearer Authentication (bearerAuth):

```python
import openapi_client
from openapi_client.models.invoice import Invoice
from openapi_client.models.invoice_request import InvoiceRequest
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
    api_instance = openapi_client.InvoiceApi(api_client)
    invoice_request = openapi_client.InvoiceRequest() # InvoiceRequest |  (optional)

    try:
        # Create Invoice
        api_response = api_instance.create_invoice(invoice_request=invoice_request)
        print("The response of InvoiceApi->create_invoice:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InvoiceApi->create_invoice: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoice_request** | [**InvoiceRequest**](InvoiceRequest.md)|  | [optional] 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Invoice created successfully. |  -  |
**400** | Bad request. |  -  |
**401** | Unauthenticated request. |  -  |
**403** | Unauthorized request. |  -  |
**429** | Too many requests. |  -  |
**500** | Unexpected error while processing request. |  -  |
**503** | Service is unhealthy. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **fetch_invoice**
> Invoice fetch_invoice(invoice_id)

Retrieve Invoice

This endpoint is used to retrieve an invoice for a given invoice id.

### Example

* Api Key Authentication (apiKey):
* Bearer Authentication (bearerAuth):

```python
import openapi_client
from openapi_client.models.invoice import Invoice
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
    api_instance = openapi_client.InvoiceApi(api_client)
    invoice_id = 'invoice_id_example' # str | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider.

    try:
        # Retrieve Invoice
        api_response = api_instance.fetch_invoice(invoice_id)
        print("The response of InvoiceApi->fetch_invoice:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InvoiceApi->fetch_invoice: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoice_id** | **str**| The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider. | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Invoice retrieved successfully. |  -  |
**401** | Unauthenticated request. |  -  |
**403** | Unauthorized request. |  -  |
**404** | Not found. |  -  |
**429** | Too many requests. |  -  |
**500** | Unexpected error while processing request. |  -  |
**503** | Service is unhealthy. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **void_invoice**
> void_invoice(invoice_id)

Void Invoice

This endpoint is used to mark a specific invoice as void. Voiding cancels the invoice without deleting it.

### Example

* Api Key Authentication (apiKey):
* Bearer Authentication (bearerAuth):

```python
import openapi_client
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
    api_instance = openapi_client.InvoiceApi(api_client)
    invoice_id = 'invoice_id_example' # str | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider.

    try:
        # Void Invoice
        api_instance.void_invoice(invoice_id)
    except Exception as e:
        print("Exception when calling InvoiceApi->void_invoice: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoice_id** | **str**| The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider. | 

### Return type

void (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | Invoice voided successfully. |  -  |
**401** | Unauthenticated request. |  -  |
**403** | Unauthorized request. |  -  |
**404** | Not found. |  -  |
**429** | Too many requests. |  -  |
**500** | Unexpected error while processing request. |  -  |
**503** | Service is unhealthy. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

