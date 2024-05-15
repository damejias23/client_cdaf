# \SubscriptionsCollectionCollectionApi

All URIs are relative to *https://example.com/ncdaf-evts/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](SubscriptionsCollectionCollectionApi.md#CreateSubscription) | **Post** /subscriptions | Ncdaf_EventExposure Subscribe service Operation



## CreateSubscription

> CdafCreatedEventSubscription CreateSubscription(ctx).CdafCreateEventSubscription(cdafCreateEventSubscription).Execute()

Ncdaf_EventExposure Subscribe service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cdafCreateEventSubscription := *openapiclient.NewCdafCreateEventSubscription(*openapiclient.NewCdafEventSubscription(*openapiclient.NewCdafEvent(*openapiclient.NewCdafEventType()), "EventNotifyUri_example")) // CdafCreateEventSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsCollectionCollectionApi.CreateSubscription(context.Background()).CdafCreateEventSubscription(cdafCreateEventSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionCollectionApi.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscription`: CdafCreatedEventSubscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionCollectionApi.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cdafCreateEventSubscription** | [**CdafCreateEventSubscription**](CdafCreateEventSubscription.md) |  | 

### Return type

[**CdafCreatedEventSubscription**](CdafCreatedEventSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

