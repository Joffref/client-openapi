/*
3gpp-asti

Testing IndividualASTIConfigurationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ASTI

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_ASTI_IndividualASTIConfigurationApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualASTIConfigurationApiService DeleteAnConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var configId string

        resp, httpRes, err := apiClient.IndividualASTIConfigurationApi.DeleteAnConfiguration(context.Background(), afId, configId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualASTIConfigurationApiService FullyModifyAnConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var configId string

        resp, httpRes, err := apiClient.IndividualASTIConfigurationApi.FullyModifyAnConfiguration(context.Background(), afId, configId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualASTIConfigurationApiService ReadAnConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var configId string

        resp, httpRes, err := apiClient.IndividualASTIConfigurationApi.ReadAnConfiguration(context.Background(), afId, configId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}