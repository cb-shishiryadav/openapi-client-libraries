# HealthCheckComponent

The health status details of a specific component reported by the Service Adapter.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** | The id of the component. | 
**name** | **str** | The name of the component. | 
**type** | **str** | The type of component affected when &#x60;status&#x60; is &#x60;WARN&#x60; or &#x60;DOWN&#x60;. The possible values are: - &#x60;ADAPTER&#x60;: The reported status is for the Service Adapter. - &#x60;API&#x60;: The reported status is for the Service Provider. - &#x60;DATABASE&#x60;: The reported status is for a database dependency of the Service Provider. - &#x60;SYSTEM&#x60;: The reported status is for any other known system component such as cache or gateway. - &#x60;OTHER&#x60;: The reported status is either for a component that does not belong to the types described above or the source of the issue is unknown.  | 
**description** | **str** | The detailed status of the component. | [optional] 
**status** | [**HealthStatus**](HealthStatus.md) |  | 
**endpoints** | **List[str]** | When the &#x60;status&#x60; of the component is not &#x60;UP&#x60;, then the list of endpoints affected. | [optional] 

## Example

```python
from openapi_client.models.health_check_component import HealthCheckComponent

# TODO update the JSON string below
json = "{}"
# create an instance of HealthCheckComponent from a JSON string
health_check_component_instance = HealthCheckComponent.from_json(json)
# print the JSON string representation of the object
print(HealthCheckComponent.to_json())

# convert the object into a dict
health_check_component_dict = health_check_component_instance.to_dict()
# create an instance of HealthCheckComponent from a dict
health_check_component_from_dict = HealthCheckComponent.from_dict(health_check_component_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


