/*
Nhss_imsSDM

Testing ReachabilitySubscriptionDeletionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nhss_imsSDM

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nhss_imsSDM_ReachabilitySubscriptionDeletionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ReachabilitySubscriptionDeletionApiService UeReachUnsubscribe", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var imsUeId string
        var subscriptionId string

        resp, httpRes, err := apiClient.ReachabilitySubscriptionDeletionApi.UeReachUnsubscribe(context.Background(), imsUeId, subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}