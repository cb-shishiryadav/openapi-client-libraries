# TaxLineItem

The details of tax applied under a specific jurisdiction.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**number** | **int** | Index or serial number of this tax line item. | 
**jurisdiction** | [**TaxJurisdiction**](TaxJurisdiction.md) |  | 
**name** | **str** | The name of the tax applied. | 
**rate** | **float** | The tax rate expressed in percentage. | 
**taxable_amount** | **float** | The part of the line item&#39;s &#x60;subtotal&#x60; that is taxable under this jurisdiction. | 
**tax_amount** | **float** | The tax payable for the line item under this jurisdiction. | 

## Example

```python
from openapi_client.models.tax_line_item import TaxLineItem

# TODO update the JSON string below
json = "{}"
# create an instance of TaxLineItem from a JSON string
tax_line_item_instance = TaxLineItem.from_json(json)
# print the JSON string representation of the object
print(TaxLineItem.to_json())

# convert the object into a dict
tax_line_item_dict = tax_line_item_instance.to_dict()
# create an instance of TaxLineItem from a dict
tax_line_item_from_dict = TaxLineItem.from_dict(tax_line_item_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


