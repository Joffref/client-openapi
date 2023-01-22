/*
Nadrf_DataManagement

Testing ADRFStorageSubscriptionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nadrf_DataManagement

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nadrf_DataManagement_ADRFStorageSubscriptionsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ADRFStorageSubscriptionsApiService CreateADRFStorageSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ADRFStorageSubscriptionsApi.CreateADRFStorageSubscription(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ADRFStorageSubscriptionsApiService DeleteADRFStorageSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ADRFStorageSubscriptionsApi.DeleteADRFStorageSubscription(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}