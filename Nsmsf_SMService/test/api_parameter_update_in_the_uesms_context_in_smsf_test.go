/*
Nsmsf_SMService Service API

Testing ParameterUpdateInTheUESMSContextInSMSFApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nsmsf_SMService

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nsmsf_SMService_ParameterUpdateInTheUESMSContextInSMSFApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ParameterUpdateInTheUESMSContextInSMSFApiService SMSServiceParameterUpdate", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var supi string

        resp, httpRes, err := apiClient.ParameterUpdateInTheUESMSContextInSMSFApi.SMSServiceParameterUpdate(context.Background(), supi).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}