# PolygonApi

All URIs are relative to _http://localhost_

| Method                              | HTTP request      | Description    |
| ----------------------------------- | ----------------- | -------------- |
| [**createPolygon**](#createpolygon) | **POST** /polygon | Create Polygon |

# **createPolygon**

> createPolygon(data)

### Example

```typescript
import { PolygonApi, Configuration, RequestCreatePolygon } from "./api";

const configuration = new Configuration();
const apiInstance = new PolygonApi(configuration);

let xUserID: string; //User ID (default to undefined)
let data: RequestCreatePolygon; //Create Polygon

const { status, data } = await apiInstance.createPolygon(xUserID, data);
```

### Parameters

| Name        | Type                     | Description    | Notes                 |
| ----------- | ------------------------ | -------------- | --------------------- |
| **data**    | **RequestCreatePolygon** | Create Polygon |                       |
| **xUserID** | [**string**]             | User ID        | defaults to undefined |

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
| **201**     | Created     | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
