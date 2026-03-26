# RequestDeviceRequest

## Properties

| Name          | Type                                                 | Description | Notes                             |
| ------------- | ---------------------------------------------------- | ----------- | --------------------------------- |
| **device_id** | **string**                                           |             | [default to undefined]            |
| **is_active** | **boolean**                                          |             | [optional] [default to undefined] |
| **readings**  | [**Array&lt;RequestReading&gt;**](RequestReading.md) |             | [default to undefined]            |

## Example

```typescript
import { RequestDeviceRequest } from "./api";

const instance: RequestDeviceRequest = {
  device_id,
  is_active,
  readings,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
