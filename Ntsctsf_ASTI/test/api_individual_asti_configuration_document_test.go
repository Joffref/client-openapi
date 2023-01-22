/*
Ntsctsf_ASTI Service API

Testing IndividualASTIConfigurationDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Ntsctsf_ASTI

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Ntsctsf_ASTI_IndividualASTIConfigurationDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualASTIConfigurationDocumentApiService DeleteIndividualASTIConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var configId string

        resp, httpRes, err := apiClient.IndividualASTIConfigurationDocumentApi.DeleteIndividualASTIConfiguration(context.Background(), configId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualASTIConfigurationDocumentApiService ModifyIndividualASTIConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var configId string

        resp, httpRes, err := apiClient.IndividualASTIConfigurationDocumentApi.ModifyIndividualASTIConfiguration(context.Background(), configId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}