# \CreditNoteAPI

All URIs are relative to *https://rest.taxes.provider.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommitCreditNote**](CreditNoteAPI.md#CommitCreditNote) | **Post** /credit-notes/{creditNoteId}/commit | Commit credit note
[**CreateCreditNote**](CreditNoteAPI.md#CreateCreditNote) | **Post** /credit-notes | Create credit note
[**FetchCreditNote**](CreditNoteAPI.md#FetchCreditNote) | **Get** /credit-notes/{creditNoteId} | Retrieve credit note
[**VoidCreditNote**](CreditNoteAPI.md#VoidCreditNote) | **Post** /credit-notes/{creditNoteId}/void | Void credit note



## CommitCreditNote

> CommitCreditNote(ctx, creditNoteId).InvoiceId(invoiceId).Execute()

Commit credit note



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
	creditNoteId := "creditNoteId_example" // string | The unique identifier of the credit note at the Tax Service Adapter or Tax Service Provider.
	invoiceId := "invoiceId_example" // string | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CreditNoteAPI.CommitCreditNote(context.Background(), creditNoteId).InvoiceId(invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNoteAPI.CommitCreditNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteId** | **string** | The unique identifier of the credit note at the Tax Service Adapter or Tax Service Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommitCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceId** | **string** | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider. | 

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


## CreateCreditNote

> CreditNote CreateCreditNote(ctx).CreditNoteRequest(creditNoteRequest).Execute()

Create credit note



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
	creditNoteRequest := *openapiclient.NewCreditNoteRequest("CreditNoteCode_example", openapiclient.CreditNoteType("FULL"), time.Now(), "Currency_example", *openapiclient.NewSeller(*openapiclient.NewAddress()), *openapiclient.NewCustomer("CustomerCode_example", *openapiclient.NewAddress()), float64(123), float64(123), float64(123), float64(123), float64(123)) // CreditNoteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditNoteAPI.CreateCreditNote(context.Background()).CreditNoteRequest(creditNoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNoteAPI.CreateCreditNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCreditNote`: CreditNote
	fmt.Fprintf(os.Stdout, "Response from `CreditNoteAPI.CreateCreditNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creditNoteRequest** | [**CreditNoteRequest**](CreditNoteRequest.md) |  | 

### Return type

[**CreditNote**](CreditNote.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCreditNote

> CreditNote FetchCreditNote(ctx, creditNoteId).InvoiceId(invoiceId).Execute()

Retrieve credit note



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
	creditNoteId := "creditNoteId_example" // string | The unique identifier of the credit note at the Tax Service Adapter or Tax Service Provider.
	invoiceId := "invoiceId_example" // string | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditNoteAPI.FetchCreditNote(context.Background(), creditNoteId).InvoiceId(invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNoteAPI.FetchCreditNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchCreditNote`: CreditNote
	fmt.Fprintf(os.Stdout, "Response from `CreditNoteAPI.FetchCreditNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteId** | **string** | The unique identifier of the credit note at the Tax Service Adapter or Tax Service Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceId** | **string** | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider. | 

### Return type

[**CreditNote**](CreditNote.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidCreditNote

> VoidCreditNote(ctx, creditNoteId).InvoiceId(invoiceId).Execute()

Void credit note



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
	creditNoteId := "creditNoteId_example" // string | The unique identifier of the credit note at the Tax Service Adapter or Tax Service Provider.
	invoiceId := "invoiceId_example" // string | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CreditNoteAPI.VoidCreditNote(context.Background(), creditNoteId).InvoiceId(invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNoteAPI.VoidCreditNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditNoteId** | **string** | The unique identifier of the credit note at the Tax Service Adapter or Tax Service Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceId** | **string** | The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider. | 

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

