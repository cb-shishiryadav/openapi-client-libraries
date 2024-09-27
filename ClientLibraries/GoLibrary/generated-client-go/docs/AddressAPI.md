# \AddressAPI

All URIs are relative to *https://rest.taxes.provider.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAddressTaxability**](AddressAPI.md#CheckAddressTaxability) | **Post** /address/check-taxability | Check taxability
[**ValidateAddress**](AddressAPI.md#ValidateAddress) | **Post** /address/validate | Address validation



## CheckAddressTaxability

> CheckAddressTaxabilityResponse CheckAddressTaxability(ctx).CheckAddressTaxabilityRequest(checkAddressTaxabilityRequest).Execute()

Check taxability



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
	checkAddressTaxabilityRequest := *openapiclient.NewCheckAddressTaxabilityRequest() // CheckAddressTaxabilityRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.CheckAddressTaxability(context.Background()).CheckAddressTaxabilityRequest(checkAddressTaxabilityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.CheckAddressTaxability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckAddressTaxability`: CheckAddressTaxabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.CheckAddressTaxability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAddressTaxabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkAddressTaxabilityRequest** | [**CheckAddressTaxabilityRequest**](CheckAddressTaxabilityRequest.md) |  | 

### Return type

[**CheckAddressTaxabilityResponse**](CheckAddressTaxabilityResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAddress

> AddressValidationResponse ValidateAddress(ctx).AddressValidationRequest(addressValidationRequest).Execute()

Address validation



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
	addressValidationRequest := *openapiclient.NewAddressValidationRequest() // AddressValidationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.ValidateAddress(context.Background()).AddressValidationRequest(addressValidationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.ValidateAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateAddress`: AddressValidationResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.ValidateAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addressValidationRequest** | [**AddressValidationRequest**](AddressValidationRequest.md) |  | 

### Return type

[**AddressValidationResponse**](AddressValidationResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

