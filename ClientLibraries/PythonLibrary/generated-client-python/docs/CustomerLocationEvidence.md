# CustomerLocationEvidence

Represent the properties for customer location evidence.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ip** | **str** | The customer&#39;s IP to determine which country the customer belongs to. | [optional] 
**bin** | **str** | The country associated with a card by using the first or last 6 digits of the Bank Identification Number. | [optional] 
**payment_country_code** | **str** | Identifies the country code associated with the payment method. | [optional] 

## Example

```python
from openapi_client.models.customer_location_evidence import CustomerLocationEvidence

# TODO update the JSON string below
json = "{}"
# create an instance of CustomerLocationEvidence from a JSON string
customer_location_evidence_instance = CustomerLocationEvidence.from_json(json)
# print the JSON string representation of the object
print(CustomerLocationEvidence.to_json())

# convert the object into a dict
customer_location_evidence_dict = customer_location_evidence_instance.to_dict()
# create an instance of CustomerLocationEvidence from a dict
customer_location_evidence_from_dict = CustomerLocationEvidence.from_dict(customer_location_evidence_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


