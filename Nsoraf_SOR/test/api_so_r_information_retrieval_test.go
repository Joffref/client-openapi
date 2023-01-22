/*
Nsoraf_SOR

Testing SoRInformationRetrievalApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nsoraf_SOR

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nsoraf_SOR_SoRInformationRetrievalApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SoRInformationRetrievalApiService GetSorInformation", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var supi string

        resp, httpRes, err := apiClient.SoRInformationRetrievalApi.GetSorInformation(context.Background(), supi).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}