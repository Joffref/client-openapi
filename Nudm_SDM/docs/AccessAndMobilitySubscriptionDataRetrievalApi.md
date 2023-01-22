# \AccessAndMobilitySubscriptionDataRetrievalApi

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAmData**](AccessAndMobilitySubscriptionDataRetrievalApi.md#GetAmData) | **Get** /{supi}/am-data | retrieve a UE&#39;s Access and Mobility Subscription Data



## GetAmData

> AccessAndMobilitySubscriptionData GetAmData(ctx, supi).SupportedFeatures(supportedFeatures).PlmnId(plmnId).AdjacentPlmns(adjacentPlmns).DisasterRoamingInd(disasterRoamingInd).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

retrieve a UE's Access and Mobility Subscription Data

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
    supi := "supi_example" // string | Identifier of the UE
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    plmnId := map[string][]openapiclient.PlmnIdNid{ ... } // PlmnIdNid | Serving PLMN ID or SNPN ID (optional)
    adjacentPlmns := []openapiclient.PlmnId{*openapiclient.NewPlmnId("Mcc_example", "Mnc_example")} // []PlmnId | List of PLMNs adjacent to the UE's serving PLMN (optional)
    disasterRoamingInd := true // bool | Indication whether Disaster Roaming service is applied or not (optional) (default to false)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 7232, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessAndMobilitySubscriptionDataRetrievalApi.GetAmData(context.Background(), supi).SupportedFeatures(supportedFeatures).PlmnId(plmnId).AdjacentPlmns(adjacentPlmns).DisasterRoamingInd(disasterRoamingInd).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessAndMobilitySubscriptionDataRetrievalApi.GetAmData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmData`: AccessAndMobilitySubscriptionData
    fmt.Fprintf(os.Stdout, "Response from `AccessAndMobilitySubscriptionDataRetrievalApi.GetAmData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAmDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 
 **plmnId** | [**PlmnIdNid**](PlmnIdNid.md) | Serving PLMN ID or SNPN ID | 
 **adjacentPlmns** | [**[]PlmnId**](PlmnId.md) | List of PLMNs adjacent to the UE&#39;s serving PLMN | 
 **disasterRoamingInd** | **bool** | Indication whether Disaster Roaming service is applied or not | [default to false]
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 7232, 3.3 | 

### Return type

[**AccessAndMobilitySubscriptionData**](AccessAndMobilitySubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
