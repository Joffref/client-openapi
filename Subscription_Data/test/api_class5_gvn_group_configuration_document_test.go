/*
Unified Data Repository Service API file for subscription data

Testing Class5GVnGroupConfigurationDocumentApiService

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

func Test_Subscription_Data_Class5GVnGroupConfigurationDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test Class5GVnGroupConfigurationDocumentApiService Create5GVnGroup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var externalGroupId string

        resp, httpRes, err := apiClient.Class5GVnGroupConfigurationDocumentApi.Create5GVnGroup(context.Background(), externalGroupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}