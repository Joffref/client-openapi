/*
Namf_MBSBroadcast

Testing IndividualBroadcastMBSSessionContextDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Namf_MBSBroadcast

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Namf_MBSBroadcast_IndividualBroadcastMBSSessionContextDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualBroadcastMBSSessionContextDocumentApiService ContextDelete", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var mbsContextRef string

        resp, httpRes, err := apiClient.IndividualBroadcastMBSSessionContextDocumentApi.ContextDelete(context.Background(), mbsContextRef).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualBroadcastMBSSessionContextDocumentApiService ContextUpdate", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var mbsContextRef string

        resp, httpRes, err := apiClient.IndividualBroadcastMBSSessionContextDocumentApi.ContextUpdate(context.Background(), mbsContextRef).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}