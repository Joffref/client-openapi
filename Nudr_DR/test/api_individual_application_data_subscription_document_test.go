/*
Nudr_DataRepository API OpenAPI file

Testing IndividualApplicationDataSubscriptionDocumentApiService

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

func Test_Nudr_DR_IndividualApplicationDataSubscriptionDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualApplicationDataSubscriptionDocumentApiService DeleteIndividualApplicationDataSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subsId string

        resp, httpRes, err := apiClient.IndividualApplicationDataSubscriptionDocumentApi.DeleteIndividualApplicationDataSubscription(context.Background(), subsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualApplicationDataSubscriptionDocumentApiService ReadIndividualApplicationDataSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subsId string

        resp, httpRes, err := apiClient.IndividualApplicationDataSubscriptionDocumentApi.ReadIndividualApplicationDataSubscription(context.Background(), subsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualApplicationDataSubscriptionDocumentApiService ReplaceIndividualApplicationDataSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subsId string

        resp, httpRes, err := apiClient.IndividualApplicationDataSubscriptionDocumentApi.ReplaceIndividualApplicationDataSubscription(context.Background(), subsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}