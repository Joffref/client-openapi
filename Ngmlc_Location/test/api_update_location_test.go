/*
Ngmlc_Location

Testing UpdateLocationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Ngmlc_Location

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Ngmlc_Location_UpdateLocationApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test UpdateLocationApiService UpdateLocation", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.UpdateLocationApi.UpdateLocation(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}