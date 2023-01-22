/*
Npcf_AMPolicyAuthorization Service API

Testing IndividualApplicationAMContextDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Npcf_AMPolicyAuthorization

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Npcf_AMPolicyAuthorization_IndividualApplicationAMContextDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualApplicationAMContextDocumentApiService DeleteAppAmContext", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var appAmContextId string

        resp, httpRes, err := apiClient.IndividualApplicationAMContextDocumentApi.DeleteAppAmContext(context.Background(), appAmContextId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualApplicationAMContextDocumentApiService GetAppAmContext", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var appAmContextId string

        resp, httpRes, err := apiClient.IndividualApplicationAMContextDocumentApi.GetAppAmContext(context.Background(), appAmContextId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualApplicationAMContextDocumentApiService ModAppAmContext", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var appAmContextId string

        resp, httpRes, err := apiClient.IndividualApplicationAMContextDocumentApi.ModAppAmContext(context.Background(), appAmContextId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}