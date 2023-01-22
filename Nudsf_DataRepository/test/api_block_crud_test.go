/*
Nudsf_DataRepository

Testing BlockCRUDApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nudsf_DataRepository

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nudsf_DataRepository_BlockCRUDApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test BlockCRUDApiService CreateOrModifyBlock", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var realmId string
        var storageId string
        var recordId string
        var blockId string

        resp, httpRes, err := apiClient.BlockCRUDApi.CreateOrModifyBlock(context.Background(), realmId, storageId, recordId, blockId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BlockCRUDApiService DeleteBlock", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var realmId string
        var storageId string
        var recordId string
        var blockId string

        resp, httpRes, err := apiClient.BlockCRUDApi.DeleteBlock(context.Background(), realmId, storageId, recordId, blockId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BlockCRUDApiService GetBlock", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var realmId string
        var storageId string
        var recordId string
        var blockId string

        resp, httpRes, err := apiClient.BlockCRUDApi.GetBlock(context.Background(), realmId, storageId, recordId, blockId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BlockCRUDApiService GetBlockList", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var realmId string
        var storageId string
        var recordId string

        resp, httpRes, err := apiClient.BlockCRUDApi.GetBlockList(context.Background(), realmId, storageId, recordId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}