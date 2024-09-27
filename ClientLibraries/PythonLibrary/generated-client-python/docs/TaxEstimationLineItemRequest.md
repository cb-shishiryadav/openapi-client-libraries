# TaxEstimationLineItemRequest

Represents the details of a line item in a tax estimation request.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**number** | **int** | Index or serial number of the line item. | 
**item_code** | **str** | The unique identifier (in Chargebee) of the product corresponding to the line item. If the line item corresponds to a one-off charge, then this identifier is not provided. | [optional] 
**description** | **str** | The description of the line item. | [optional] 
**quantity** | **float** | The quantity associated with this line item. | [optional] 
**unit_price** | **float** | The unit price for this line item. In case of [tiered pricing](https://www.chargebee.com/docs/1.0/plans.html#tiered-pricing) where the unit price varies for each quantity tier, this is the average unit price. | [optional] 
**amount** | **float** | The amount for this line item. This is &#x60;unitPrice&#x60; × &#x60;quantity&#x60;. | 
**discount_amount** | **float** | The discount applied to this line item. | [optional] 
**is_tax_inclusive** | **bool** | Indicates whether (&#x60;amount&#x60; - &#x60;discountAmount&#x60;)  is inclusive of taxes. | 
**tax_identifiers** | [**List[FieldItem]**](FieldItem.md) | The tax code fields of the product used for tax calculation. | [optional] 

## Example

```python
from openapi_client.models.tax_estimation_line_item_request import TaxEstimationLineItemRequest

# TODO update the JSON string below
json = "{}"
# create an instance of TaxEstimationLineItemRequest from a JSON string
tax_estimation_line_item_request_instance = TaxEstimationLineItemRequest.from_json(json)
# print the JSON string representation of the object
print(TaxEstimationLineItemRequest.to_json())

# convert the object into a dict
tax_estimation_line_item_request_dict = tax_estimation_line_item_request_instance.to_dict()
# create an instance of TaxEstimationLineItemRequest from a dict
tax_estimation_line_item_request_from_dict = TaxEstimationLineItemRequest.from_dict(tax_estimation_line_item_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


