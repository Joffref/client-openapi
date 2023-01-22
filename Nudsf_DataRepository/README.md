# Go API client for Nudsf_DataRepository

Nudsf Data Repository Service.  
© 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).  
All rights reserved.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import Nudsf_DataRepository "//"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), Nudsf_DataRepository.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), Nudsf_DataRepository.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), Nudsf_DataRepository.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), Nudsf_DataRepository.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/nudsf-dr/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BlockCRUDApi* | [**CreateOrModifyBlock**](docs/BlockCRUDApi.md#createormodifyblock) | **Put** /{realmId}/{storageId}/records/{recordId}/blocks/{blockId} | Create or Update a specific Block in a Record.
*BlockCRUDApi* | [**DeleteBlock**](docs/BlockCRUDApi.md#deleteblock) | **Delete** /{realmId}/{storageId}/records/{recordId}/blocks/{blockId} | Delete a specific Block. Then update the Record
*BlockCRUDApi* | [**GetBlock**](docs/BlockCRUDApi.md#getblock) | **Get** /{realmId}/{storageId}/records/{recordId}/blocks/{blockId} | Retrieve a specific Block
*BlockCRUDApi* | [**GetBlockList**](docs/BlockCRUDApi.md#getblocklist) | **Get** /{realmId}/{storageId}/records/{recordId}/blocks | Record&#39;s Blocks access
*MetaSchemaCRUDApi* | [**CreateOrModifyMetaSchema**](docs/MetaSchemaCRUDApi.md#createormodifymetaschema) | **Put** /{realmId}/{storageId}/meta-schemas/{schemaId} | Create/Modify Meta Schema
*MetaSchemaCRUDApi* | [**DeleteMetaSchema**](docs/MetaSchemaCRUDApi.md#deletemetaschema) | **Delete** /{realmId}/{storageId}/meta-schemas/{schemaId} | Delete a Meta Schema with an user provided SchemaId
*MetaSchemaCRUDApi* | [**GetMetaSchema**](docs/MetaSchemaCRUDApi.md#getmetaschema) | **Get** /{realmId}/{storageId}/meta-schemas/{schemaId} | Meta Schema access
*NotificationSubscriptionCRUDApi* | [**CreateAndUpdateNotificationSubscription**](docs/NotificationSubscriptionCRUDApi.md#createandupdatenotificationsubscription) | **Put** /{realmId}/{storageId}/subs-to-notify/{subscriptionId} | NotificationSubscription Create/Update
*NotificationSubscriptionCRUDApi* | [**DeleteNotificationSubscription**](docs/NotificationSubscriptionCRUDApi.md#deletenotificationsubscription) | **Delete** /{realmId}/{storageId}/subs-to-notify/{subscriptionId} | Delete a Notification Subscription of the storage
*NotificationSubscriptionCRUDApi* | [**GetNotificationSubscription**](docs/NotificationSubscriptionCRUDApi.md#getnotificationsubscription) | **Get** /{realmId}/{storageId}/subs-to-notify/{subscriptionId} | Notification subscription retrieval
*NotificationSubscriptionCRUDApi* | [**UpdateNotificationSubscription**](docs/NotificationSubscriptionCRUDApi.md#updatenotificationsubscription) | **Patch** /{realmId}/{storageId}/subs-to-notify/{subscriptionId} | NotificationSubscription update
*NotificationSubscriptionsCRUDApi* | [**GetNotificationSubscriptions**](docs/NotificationSubscriptionsCRUDApi.md#getnotificationsubscriptions) | **Get** /{realmId}/{storageId}/subs-to-notify | Notification subscription retrieval
*RecordCRUDApi* | [**BulkDeleteRecords**](docs/RecordCRUDApi.md#bulkdeleterecords) | **Delete** /{realmId}/{storageId}/records | Bulk Deletion of Records
*RecordCRUDApi* | [**CreateOrModifyRecord**](docs/RecordCRUDApi.md#createormodifyrecord) | **Put** /{realmId}/{storageId}/records/{recordId} | Create/Modify Record
*RecordCRUDApi* | [**DeleteRecord**](docs/RecordCRUDApi.md#deleterecord) | **Delete** /{realmId}/{storageId}/records/{recordId} | Delete a Record with an user provided RecordId
*RecordCRUDApi* | [**GetMeta**](docs/RecordCRUDApi.md#getmeta) | **Get** /{realmId}/{storageId}/records/{recordId}/meta | Record&#39;s meta access
*RecordCRUDApi* | [**GetRecord**](docs/RecordCRUDApi.md#getrecord) | **Get** /{realmId}/{storageId}/records/{recordId} | Record access
*RecordCRUDApi* | [**SearchRecord**](docs/RecordCRUDApi.md#searchrecord) | **Get** /{realmId}/{storageId}/records | Records search with get
*RecordCRUDApi* | [**UpdateMeta**](docs/RecordCRUDApi.md#updatemeta) | **Patch** /{realmId}/{storageId}/records/{recordId}/meta | Record&#39;s meta update


## Documentation For Models

 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [ClientId](docs/ClientId.md)
 - [ComparisonOperator](docs/ComparisonOperator.md)
 - [ComparisonOperatorAnyOf](docs/ComparisonOperatorAnyOf.md)
 - [ConditionOperator](docs/ConditionOperator.md)
 - [ConditionOperatorAnyOf](docs/ConditionOperatorAnyOf.md)
 - [GetBlockList200Response](docs/GetBlockList200Response.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [KeyType](docs/KeyType.md)
 - [KeyTypeAnyOf](docs/KeyTypeAnyOf.md)
 - [MetaSchema](docs/MetaSchema.md)
 - [NFType](docs/NFType.md)
 - [NFTypeAnyOf](docs/NFTypeAnyOf.md)
 - [NotificationDescription](docs/NotificationDescription.md)
 - [NotificationInfo](docs/NotificationInfo.md)
 - [NotificationSubscription](docs/NotificationSubscription.md)
 - [PatchItem](docs/PatchItem.md)
 - [PatchOperation](docs/PatchOperation.md)
 - [PatchOperationAnyOf](docs/PatchOperationAnyOf.md)
 - [PatchResult](docs/PatchResult.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnIdNid](docs/PlmnIdNid.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [Record](docs/Record.md)
 - [RecordIdList](docs/RecordIdList.md)
 - [RecordMeta](docs/RecordMeta.md)
 - [RecordNotification](docs/RecordNotification.md)
 - [RecordOperation](docs/RecordOperation.md)
 - [RecordOperationAnyOf](docs/RecordOperationAnyOf.md)
 - [RecordSearchResult](docs/RecordSearchResult.md)
 - [ReportItem](docs/ReportItem.md)
 - [RetrieveRecords](docs/RetrieveRecords.md)
 - [RetrieveRecordsAnyOf](docs/RetrieveRecordsAnyOf.md)
 - [SearchComparison](docs/SearchComparison.md)
 - [SearchCondition](docs/SearchCondition.md)
 - [SearchExpression](docs/SearchExpression.md)
 - [Snssai](docs/Snssai.md)
 - [SubscriptionFilter](docs/SubscriptionFilter.md)
 - [TagType](docs/TagType.md)


## Documentation For Authorization



### oAuth2ClientCredentials


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **nudsf-dr**: Access to the nudsf-dr API

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


