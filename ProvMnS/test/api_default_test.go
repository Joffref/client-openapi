/*
Provisioning MnS

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ProvMnS

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_ProvMnS_DefaultApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test DefaultApiService ClassNameidDelete", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var className string
        var id string

        resp, httpRes, err := apiClient.DefaultApi.ClassNameidDelete(context.Background(), className, id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ClassNameidGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var className string
        var id string

        resp, httpRes, err := apiClient.DefaultApi.ClassNameidGet(context.Background(), className, id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ClassNameidPatch", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var className string
        var id string

        resp, httpRes, err := apiClient.DefaultApi.ClassNameidPatch(context.Background(), className, id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ClassNameidPut", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var className string
        var id string

        resp, httpRes, err := apiClient.DefaultApi.ClassNameidPut(context.Background(), className, id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}