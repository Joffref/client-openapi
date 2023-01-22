/*
Nudm_UECM

Testing SMSFDeregistrationFor3GPPAccessApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nudm_UECM

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nudm_UECM_SMSFDeregistrationFor3GPPAccessApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SMSFDeregistrationFor3GPPAccessApiService Call3GppSmsfDeregistration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string

        resp, httpRes, err := apiClient.SMSFDeregistrationFor3GPPAccessApi.Call3GppSmsfDeregistration(context.Background(), ueId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}