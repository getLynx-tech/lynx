# AnchorApi

All URIs are relative to _http://localhost_

| Method                              | HTTP request      | Description   |
| ----------------------------------- | ----------------- | ------------- |
| [**getAllAnchors**](#getallanchors) | **GET** /anchors  | GetAllAnchors |
| [**upsertAnchors**](#upsertanchors) | **POST** /anchors | UpsertAnchors |

# **getAllAnchors**

> Array<ResponseAnchor> getAllAnchors()

### Example

```typescript
import { AnchorApi, Configuration } from "./api";

const configuration = new Configuration();
const apiInstance = new AnchorApi(configuration);

const { status, data } = await apiInstance.getAllAnchors();
```

### Parameters

This endpoint does not have any parameters.

### Return type

**Array<ResponseAnchor>**

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

# **upsertAnchors**

> upsertAnchors(data)

### Example

```typescript
import { AnchorApi, Configuration, RequestAnchorsRequest } from "./api";

const configuration = new Configuration();
const apiInstance = new AnchorApi(configuration);

let data: RequestAnchorsRequest; //Anchors Request

const { status, data } = await apiInstance.upsertAnchors(data);
```

### Parameters

| Name     | Type                      | Description     | Notes |
| -------- | ------------------------- | --------------- | ----- |
| **data** | **RequestAnchorsRequest** | Anchors Request |       |

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
