/*
AUSF API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nausf_UEAuthentication

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nausf_UEAuthentication_DefaultApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test DefaultApiService EapAuthMethod", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var authCtxId string

        resp, httpRes, err := apiClient.DefaultApi.EapAuthMethod(context.Background(), authCtxId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ProseAuth", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var authCtxId string

        resp, httpRes, err := apiClient.DefaultApi.ProseAuth(context.Background(), authCtxId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ProseAuthenticationsPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.ProseAuthenticationsPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService RgAuthenticationsPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.RgAuthenticationsPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService UeAuthenticationsAuthCtxId5gAkaConfirmationPut", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var authCtxId string

        resp, httpRes, err := apiClient.DefaultApi.UeAuthenticationsAuthCtxId5gAkaConfirmationPut(context.Background(), authCtxId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService UeAuthenticationsDeregisterPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.UeAuthenticationsDeregisterPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService UeAuthenticationsPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.UeAuthenticationsPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}