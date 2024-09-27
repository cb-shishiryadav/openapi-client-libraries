# CreditNote

The details of a credit note returned by the Tax Service Adapter. A credit note is used to reduce the amount due on an invoice. If the credit note is issued after payments have been made for the invoice, refunds can be issued to the Customer.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**credit_note_id** | **str** | The unique identifier of the credit note at the Tax Service Provider or Tax Service Adapter. | 
**credit_note_code** | **str** | The unique identifier of the credit note in Chargebee. | 
**invoice_code** | **str** | The unique identifier of the invoice in Chargebee to which this credit note belongs. | [optional] 
**invoice_id** | **str** | The unique identifier of the invoice in the Tax Service Adapter or the Tax Service Provider. | [optional] 
**credit_note_type** | [**CreditNoteType**](CreditNoteType.md) |  | 
**document_date_time** | **datetime** | The date and time at which the credit note was created in Chargebee. For example, if the value is 2022-10-28T15:36:28.129+05:30, then the timestamp represents October 28, 2022, at 15:36:28.129, with an offset of +05:30. This means that the time represented is 5 hours and 30 minutes ahead of UTC/GMT. In the case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the &#39;Z&#39; indicating that the time is in UTC. | 
**tax_date_time** | **datetime** | The date and time at which the tax was applicable in Chargebee. For example, if the value is 2022-10-28T15:36:28.129+05:30, then the timestamp represents October 28, 2022, at 15:36:28.129, with an offset of +05:30. This means that the time represented is 5 hours and 30 minutes ahead of UTC/GMT.In the case of a merchant site located in UTC, these data types would send a timestamp in the format 2022-11-11T15:40:44.65Z. This timestamp represents November 11, 2022, at 15:40:44.65, with the &#39;Z&#39; indicating that the time is in UTC. | [optional] 
**status** | [**DocumentStatus**](DocumentStatus.md) |  | 
**currency** | **str** | The [currency](https://en.wikipedia.org/wiki/Currency) in the [ISO-4217 format](https://www.iso.org/iso-4217-currency-codes.html). | 
**seller** | [**Seller**](Seller.md) |  | 
**customer** | [**Customer**](Customer.md) |  | 
**discount_amount** | **float** | The total discount applied. This is the sum of all &#x60;lineItems.discount&#x60;. | 
**subtotal** | **float** | The amount after discounts. This is the sum of all &#x60;lineItems.subtotal&#x60;. | [optional] 
**exempt_amount** | **float** | The amount exempted from tax. | 
**taxable_amount** | **float** | The amount upon which the tax is calculated. | 
**tax_amount** | **float** | The total tax payable. This is the sum of all &#x60;lineItems.taxAmount&#x60;. | 
**total** | **float** | The total amount of the credit note. &#x60;total&#x60; can be expressed as &#x60;exemptAmount&#x60; + &#x60;taxableAmount&#x60; + &#x60;taxAmount&#x60;. | 
**rounding_amount** | **float** | The rounding amount added to the total amount to account for fractional correction. | [optional] 
**line_items** | [**List[InvoiceLineItem]**](InvoiceLineItem.md) |  | 

## Example

```python
from openapi_client.models.credit_note import CreditNote

# TODO update the JSON string below
json = "{}"
# create an instance of CreditNote from a JSON string
credit_note_instance = CreditNote.from_json(json)
# print the JSON string representation of the object
print(CreditNote.to_json())

# convert the object into a dict
credit_note_dict = credit_note_instance.to_dict()
# create an instance of CreditNote from a dict
credit_note_from_dict = CreditNote.from_dict(credit_note_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


