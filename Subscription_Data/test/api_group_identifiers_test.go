/*
Unified Data Repository Service API file for subscription data

Testing GroupIdentifiersApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Subscription_Data

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Subscription_Data_GroupIdentifiersApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test GroupIdentifiersApiService GetGroupIdentifiers", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.GroupIdentifiersApi.GetGroupIdentifiers(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}