# Invoice

The details of an invoice as returned by the Tax Service Adapter.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**invoice_id** | **str** | The unique identifier of the invoice in the Tax Service Adapter or the Tax Service Provider. | 
**invoice_code** | **str** | The unique identifier of the invoice in Chargebee. | 
**document_date_time** | **datetime** | The date and time at which the invoice was generated in Chargebee. For example, if the value is 2022-10-28T15:36:28.129+05:30, then the timestamp represents October 28, 2022, at 15:36:28.129, with an offset of +05:30. This means that the time represented is 5 hours and 30 minutes ahead of UTC/GMT.In the case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the &#39;Z&#39; indicating that the time is in UTC. | 
**tax_date_time** | **datetime** | The date and time at which the tax was applicable in Chargebee. For example, if the value is 2022-10-28T15:36:28.129+05:30, then the timestamp represents October 28, 2022, at 15:36:28.129, with an offset of +05:30. This means that the time represented is 5 hours and 30 minutes ahead of UTC/GMT.In the case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the &#39;Z&#39; indicating that the time is in UTC. | [optional] 
**status** | [**DocumentStatus**](DocumentStatus.md) |  | 
**currency** | **str** | The [currency](https://en.wikipedia.org/wiki/Currency) in the [ISO-4217 format](https://www.iso.org/iso-4217-currency-codes.html). | 
**seller** | [**Seller**](Seller.md) |  | 
**customer** | [**Customer**](Customer.md) |  | 
**subtotal** | **float** | The amount after discounts. This is the sum of all &#x60;lineItems.subtotal&#x60;. | 
**exempt_amount** | **float** | The part of the &#x60;subtotal&#x60; that is exempted from tax. | 
**discount_amount** | **float** | The total discount applied. This is the sum of all &#x60;lineItems.discount&#x60;. | 
**taxable_amount** | **float** | The part of the &#x60;subtotal&#x60; that is taxable. | 
**tax_amount** | **float** | The total tax payable. This is the sum of all &#x60;lineItems.taxAmount&#x60;. | 
**total** | **float** | The total after discounts and taxes. This is the same as &#x60;subtotal&#x60; if it is tax inclusive; otherwise it is &#x60;subtotal&#x60; + &#x60;taxAmount&#x60;. &#x60;total&#x60; can also be expressed as &#x60;exemptAmount&#x60; + &#x60;taxableAmount&#x60; + &#x60;taxAmount&#x60;. | 
**line_items** | [**List[InvoiceLineItem]**](InvoiceLineItem.md) |  | 

## Example

```python
from openapi_client.models.invoice import Invoice

# TODO update the JSON string below
json = "{}"
# create an instance of Invoice from a JSON string
invoice_instance = Invoice.from_json(json)
# print the JSON string representation of the object
print(Invoice.to_json())

# convert the object into a dict
invoice_dict = invoice_instance.to_dict()
# create an instance of Invoice from a dict
invoice_from_dict = Invoice.from_dict(invoice_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


