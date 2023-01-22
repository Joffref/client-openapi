/*
Nadrf_DataManagement

Testing ADRFStoredDataApiService

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

func Test_Nadrf_DataManagement_ADRFStoredDataApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ADRFStoredDataApiService DeleteADRFData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ADRFStoredDataApi.DeleteADRFData(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}