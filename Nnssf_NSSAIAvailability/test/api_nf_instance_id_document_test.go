/*
NSSF NSSAI Availability

Testing NFInstanceIDDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nnssf_NSSAIAvailability

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nnssf_NSSAIAvailability_NFInstanceIDDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test NFInstanceIDDocumentApiService NSSAIAvailabilityDelete", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nfId string

        resp, httpRes, err := apiClient.NFInstanceIDDocumentApi.NSSAIAvailabilityDelete(context.Background(), nfId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test NFInstanceIDDocumentApiService NSSAIAvailabilityPatch", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nfId string

        resp, httpRes, err := apiClient.NFInstanceIDDocumentApi.NSSAIAvailabilityPatch(context.Background(), nfId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test NFInstanceIDDocumentApiService NSSAIAvailabilityPut", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var nfId string

        resp, httpRes, err := apiClient.NFInstanceIDDocumentApi.NSSAIAvailabilityPut(context.Background(), nfId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}