/*
SS_NetworkResourceAdaptation

Testing IndividualUnicastSubscriptionDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package SS_NetworkResourceAdaptation

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_SS_NetworkResourceAdaptation_IndividualUnicastSubscriptionDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualUnicastSubscriptionDocumentApiService DeleteUnicastSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var uniSubId string

        resp, httpRes, err := apiClient.IndividualUnicastSubscriptionDocumentApi.DeleteUnicastSubscription(context.Background(), uniSubId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualUnicastSubscriptionDocumentApiService GetUnicastSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var uniSubId string

        resp, httpRes, err := apiClient.IndividualUnicastSubscriptionDocumentApi.GetUnicastSubscription(context.Background(), uniSubId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}