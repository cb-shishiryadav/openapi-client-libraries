# \InvoiceAPI

All URIs are relative to *https://rest.taxes.provider.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommitInvoice**](InvoiceAPI.md#CommitInvoice) | **Post** /invoices/{invoiceId}/commit | Commit Invoice
[**CreateInvoice**](InvoiceAPI.md#CreateInvoice) | **Post** /invoices | Create Invoice
[**FetchInvoice**](InvoiceAPI.md#FetchInvoice) | **Get** /invoices/{invoiceId} | Retrieve Invoice
[**VoidInvoice**](InvoiceAPI.md#VoidInvoice) | **Post** /invoices/{invoiceId}/void | Void Invoice



## CommitInvoice

> CommitInvoice(ctx, invoiceId).Execute()

Commit Invoice



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	invoiceId := "invoiceId_example" // string | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoiceAPI.CommitInvoice(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.CommitInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommitInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoice

> Invoice CreateInvoice(ctx).InvoiceRequest(invoiceRequest).Execute()

Create Invoice



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	invoiceRequest := *openapiclient.NewInvoiceRequest("InvoiceCode_example", time.Now(), "Currency_example", *openapiclient.NewSeller(*openapiclient.NewAddress()), *openapiclient.NewCustomer("CustomerCode_example", *openapiclient.NewAddress()), float64(123), float64(123), float64(123), float64(123), float64(123), float64(123), []openapiclient.InvoiceLineItem{*openapiclient.NewInvoiceLineItem(int32(123), float64(123), float64(123), false, false, float64(123), float64(123), float64(123), float64(123), float64(123), []openapiclient.TaxLineItem{*openapiclient.NewTaxLineItem(int32(123), *openapiclient.NewTaxJurisdiction("Code_example", openapiclient.TaxJurisdictionType("COUNTRY"), "Name_example"), "GST", float64(123), float64(123), float64(123))})}) // InvoiceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceAPI.CreateInvoice(context.Background()).InvoiceRequest(invoiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.CreateInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.CreateInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceRequest** | [**InvoiceRequest**](InvoiceRequest.md) |  | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInvoice

> Invoice FetchInvoice(ctx, invoiceId).Execute()

Retrieve Invoice



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	invoiceId := "invoiceId_example" // string | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceAPI.FetchInvoice(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.FetchInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.FetchInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidInvoice

> VoidInvoice(ctx, invoiceId).Execute()

Void Invoice



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	invoiceId := "invoiceId_example" // string | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoiceAPI.VoidInvoice(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.VoidInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

