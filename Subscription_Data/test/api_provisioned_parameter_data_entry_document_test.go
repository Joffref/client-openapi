/*
Unified Data Repository Service API file for subscription data

Testing ProvisionedParameterDataEntryDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Subscription_Data

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Subscription_Data_ProvisionedParameterDataEntryDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ProvisionedParameterDataEntryDocumentApiService CreatePPDataEntry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId GetPPDataEntryUeIdParameter
        var afInstanceId string

        resp, httpRes, err := apiClient.ProvisionedParameterDataEntryDocumentApi.CreatePPDataEntry(context.Background(), ueId, afInstanceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProvisionedParameterDataEntryDocumentApiService DeletePPDataEntry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId GetPPDataEntryUeIdParameter
        var afInstanceId string

        resp, httpRes, err := apiClient.ProvisionedParameterDataEntryDocumentApi.DeletePPDataEntry(context.Background(), ueId, afInstanceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ProvisionedParameterDataEntryDocumentApiService GetPPDataEntry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId GetPPDataEntryUeIdParameter
        var afInstanceId string

        resp, httpRes, err := apiClient.ProvisionedParameterDataEntryDocumentApi.GetPPDataEntry(context.Background(), ueId, afInstanceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}