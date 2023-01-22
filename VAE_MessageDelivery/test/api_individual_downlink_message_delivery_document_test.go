/*
VAE_MessageDelivery

Testing IndividualDownlinkMessageDeliveryDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package VAE_MessageDelivery

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_VAE_MessageDelivery_IndividualDownlinkMessageDeliveryDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualDownlinkMessageDeliveryDocumentApiService ReadIndividualDownlinkMessageDelivery", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string
        var dlDeliveryId string

        resp, httpRes, err := apiClient.IndividualDownlinkMessageDeliveryDocumentApi.ReadIndividualDownlinkMessageDelivery(context.Background(), subscriptionId, dlDeliveryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}