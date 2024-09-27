# TaxEstimationRequest

Defines the parameters of a tax estimation request. This is sent to the Tax Service Adapter by Chargebee to estimate taxes for one or more line items.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**seller** | [**Seller**](Seller.md) |  | 
**customer** | [**Customer**](Customer.md) |  | 
**estimate_date_time** | **datetime** | The time as of which the tax estimation is to be calculated. This can be a value in the past. For example, if the value is provided as 2022-10-28T15:36:28.129+05:30, then the tax rates applicable on October 28, 2022, at 15:36:28.129, with an offset of +05:30 ahead of UTC/GMT are used for calculations. In case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the &#39;Z&#39; indicating that the time is in UTC | 
**currency** | **str** | The [currency](https://en.wikipedia.org/wiki/Currency) in the [ISO-4217 format](https://www.iso.org/iso-4217-currency-codes.html). | 
**line_items** | [**List[TaxEstimationLineItemRequest]**](TaxEstimationLineItemRequest.md) | Contains the details of each line item in the tax estimation request. | 

## Example

```python
from openapi_client.models.tax_estimation_request import TaxEstimationRequest

# TODO update the JSON string below
json = "{}"
# create an instance of TaxEstimationRequest from a JSON string
tax_estimation_request_instance = TaxEstimationRequest.from_json(json)
# print the JSON string representation of the object
print(TaxEstimationRequest.to_json())

# convert the object into a dict
tax_estimation_request_dict = tax_estimation_request_instance.to_dict()
# create an instance of TaxEstimationRequest from a dict
tax_estimation_request_from_dict = TaxEstimationRequest.from_dict(tax_estimation_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


