# \TaxEstimateAPI

All URIs are relative to *https://rest.taxes.provider.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EstimateTaxes**](TaxEstimateAPI.md#EstimateTaxes) | **Post** /tax-estimate | Estimate tax



## EstimateTaxes

> TaxEstimationResponse EstimateTaxes(ctx).TaxEstimationRequest(taxEstimationRequest).Execute()

Estimate tax



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
	taxEstimationRequest := *openapiclient.NewTaxEstimationRequest(*openapiclient.NewSeller(*openapiclient.NewAddress()), *openapiclient.NewCustomer("CustomerCode_example", *openapiclient.NewAddress()), time.Now(), "Currency_example", []openapiclient.TaxEstimationLineItemRequest{*openapiclient.NewTaxEstimationLineItemRequest(int32(123), float64(123), false)}) // TaxEstimationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxEstimateAPI.EstimateTaxes(context.Background()).TaxEstimationRequest(taxEstimationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxEstimateAPI.EstimateTaxes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EstimateTaxes`: TaxEstimationResponse
	fmt.Fprintf(os.Stdout, "Response from `TaxEstimateAPI.EstimateTaxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEstimateTaxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxEstimationRequest** | [**TaxEstimationRequest**](TaxEstimationRequest.md) |  | 

### Return type

[**TaxEstimationResponse**](TaxEstimationResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

