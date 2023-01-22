/*
3gpp-mbs-session

Testing MBSParametersProvisioningApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package MBSSession

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_MBSSession_MBSParametersProvisioningApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MBSParametersProvisioningApiService CreateMBSParamsProvisioning", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MBSParametersProvisioningApi.CreateMBSParamsProvisioning(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}