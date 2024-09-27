# FieldItem


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** | The id of the field. | 
**value** | **str** | The value of the field. | 

## Example

```python
from openapi_client.models.field_item import FieldItem

# TODO update the JSON string below
json = "{}"
# create an instance of FieldItem from a JSON string
field_item_instance = FieldItem.from_json(json)
# print the JSON string representation of the object
print(FieldItem.to_json())

# convert the object into a dict
field_item_dict = field_item_instance.to_dict()
# create an instance of FieldItem from a dict
field_item_from_dict = FieldItem.from_dict(field_item_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


