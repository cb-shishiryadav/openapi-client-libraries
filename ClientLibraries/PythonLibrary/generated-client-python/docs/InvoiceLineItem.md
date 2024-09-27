# InvoiceLineItem

The details of a line item.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**number** | **int** | Index or serial number of the line item. | 
**item_code** | **str** | The unique identifier (in Chargebee) of the product corresponding to the line item. If the line item corresponds to a one-off charge, then this identifier is not present. | [optional] 
**description** | **str** | The description of the line item. | [optional] 
**quantity** | **float** | The quantity associated with this line item. | [optional] 
**unit_price** | **float** | The unit price for this line item. In case of [tiered pricing](https://www.chargebee.com/docs/1.0/plans.html#tiered-pricing) where the unit price varies for each quantity tier, this is the average unit price. | [optional] 
**amount** | **float** | The amount for this line item. This is &#x60;unitPrice&#x60; × &#x60;quantity&#x60;. | 
**subtotal** | **float** | The amount after discounts for this line item. This is &#x60;amount&#x60; - &#x60;discountAmount&#x60;. | 
**is_tax_inclusive** | **bool** | Indicates whether the &#x60;subtotal&#x60; for this line item is inclusive of taxes. | 
**is_taxable** | **bool** | Indicates whether this line item is taxable. | 
**tax_identifiers** | [**List[FieldItem]**](FieldItem.md) | The tax code fields of the product used for tax calculation. | [optional] 
**tax_exempt_type** | [**TaxExemptType**](TaxExemptType.md) |  | [optional] 
**tax_exempt_reason** | **str** | The reason due to which a line item is exempted from tax. This is a mandatory parameter while applying tax exemption on any line-item. | [optional] 
**exempt_amount** | **float** | The part of this line item&#39;s &#x60;subtotal&#x60; that is exempted from tax. | 
**discount_amount** | **float** | The discount applied to this line item. | 
**taxable_amount** | **float** | The part of this line item&#39;s &#x60;subtotal&#x60; that is taxable. | 
**tax_amount** | **float** | The tax payable for this line item. This is the sum of all &#x60;taxes.taxAmount&#x60; for this line item. | 
**total** | **float** | The total for this line item after discounts and taxes. This is the same as &#x60;subtotal&#x60; if it is tax inclusive; otherwise it is &#x60;subtotal&#x60; + &#x60;taxAmount&#x60;. &#x60;total&#x60; can also be expressed as &#x60;exemptAmount&#x60; + &#x60;taxableAmount&#x60; + &#x60;taxAmount&#x60;. | 
**is_partial_tax** | **bool** | Indicates if taxes were applied only partially for this line item. | [optional] 
**taxes** | [**List[TaxLineItem]**](TaxLineItem.md) | List of taxes applied for this line item under each jurisdiction. | 

## Example

```python
from openapi_client.models.invoice_line_item import InvoiceLineItem

# TODO update the JSON string below
json = "{}"
# create an instance of InvoiceLineItem from a JSON string
invoice_line_item_instance = InvoiceLineItem.from_json(json)
# print the JSON string representation of the object
print(InvoiceLineItem.to_json())

# convert the object into a dict
invoice_line_item_dict = invoice_line_item_instance.to_dict()
# create an instance of InvoiceLineItem from a dict
invoice_line_item_from_dict = InvoiceLineItem.from_dict(invoice_line_item_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


