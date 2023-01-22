/*
Nudr_DataRepository API OpenAPI file

Testing SDMSubscriptionDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nudr_DR

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nudr_DR_SDMSubscriptionDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SDMSubscriptionDocumentApiService ModifysdmSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var subsId string

        resp, httpRes, err := apiClient.SDMSubscriptionDocumentApi.ModifysdmSubscription(context.Background(), ueId, subsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SDMSubscriptionDocumentApiService QuerysdmSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var subsId string

        resp, httpRes, err := apiClient.SDMSubscriptionDocumentApi.QuerysdmSubscription(context.Background(), ueId, subsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SDMSubscriptionDocumentApiService RemovesdmSubscriptions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var subsId string

        resp, httpRes, err := apiClient.SDMSubscriptionDocumentApi.RemovesdmSubscriptions(context.Background(), ueId, subsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SDMSubscriptionDocumentApiService Updatesdmsubscriptions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var subsId string

        resp, httpRes, err := apiClient.SDMSubscriptionDocumentApi.Updatesdmsubscriptions(context.Background(), ueId, subsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}