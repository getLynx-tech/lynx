# ScaleApi

All URIs are relative to _http://localhost_

| Method                          | HTTP request     | Description |
| ------------------------------- | ---------------- | ----------- |
| [**createScale**](#createscale) | **POST** /scales | CreateScale |
| [**getScale**](#getscale)       | **GET** /scales  | GetScale    |

# **createScale**

> createScale(data)

### Example

```typescript
import { ScaleApi, Configuration, RequestScaleRequest } from "./api";

const configuration = new Configuration();
const apiInstance = new ScaleApi(configuration);

let data: RequestScaleRequest; //Scale Request

const { status, data } = await apiInstance.createScale(data);
```

### Parameters

| Name     | Type                    | Description   | Notes |
| -------- | ----------------------- | ------------- | ----- |
| **data** | **RequestScaleRequest** | Scale Request |       |

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

# **getScale**

> ResponseScale getScale()

### Example

```typescript
import { ScaleApi, Configuration } from "./api";

const configuration = new Configuration();
const apiInstance = new ScaleApi(configuration);

const { status, data } = await apiInstance.getScale();
```

### Parameters

This endpoint does not have any parameters.

### Return type

**ResponseScale**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: _/_

### HTTP response details

| Status code | Description | Response headers |
| ----------- | ----------- | ---------------- |
| **200**     | OK          | -                |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
