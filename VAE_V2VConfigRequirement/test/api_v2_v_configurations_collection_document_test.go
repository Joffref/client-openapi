/*
VAE_V2VConfigRequirement

Testing V2VConfigurationsCollectionDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package VAE_V2VConfigRequirement

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_VAE_V2VConfigRequirement_V2VConfigurationsCollectionDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test V2VConfigurationsCollectionDocumentApiService Create", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.V2VConfigurationsCollectionDocumentApi.Create(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}