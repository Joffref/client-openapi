/*
Nnssaaf_AIW

Testing AuthenticationsCollectionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nnssaaf_AIW

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nnssaaf_AIW_AuthenticationsCollectionsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test AuthenticationsCollectionsApiService CreateAuthenticationContext", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AuthenticationsCollectionsApi.CreateAuthenticationContext(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}