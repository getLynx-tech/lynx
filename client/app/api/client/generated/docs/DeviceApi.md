# DeviceApi

All URIs are relative to _http://localhost_

| Method                                            | HTTP request               | Description          |
| ------------------------------------------------- | -------------------------- | -------------------- |
| [**updateDevicePosition**](#updatedeviceposition) | **POST** /devices/position | UpdateDevicePosition |

# **updateDevicePosition**

> updateDevicePosition(data)

### Example

```typescript
import { DeviceApi, Configuration, RequestDeviceRequest } from "./api";

const configuration = new Configuration();
const apiInstance = new DeviceApi(configuration);

let data: RequestDeviceRequest; //Device Request

const { status, data } = await apiInstance.updateDevicePosition(data);
```

### Parameters

| Name     | Type                     | Description    | Notes |
| -------- | ------------------------ | -------------- | ----- |
| **data** | **RequestDeviceRequest** | Device Request |       |

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
| ----------- | ----------- | ---------------- |
| **200**     | OK          | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
