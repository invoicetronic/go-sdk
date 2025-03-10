/*
Invoicetronic API

Testing UpdateAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package invoicetronicsdk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	invoicetronicsdk "github.com/invoicetronic/go-sdk"
)

func Test_invoicetronicsdk_UpdateAPIService(t *testing.T) {

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)

	t.Run("Test UpdateAPIService UpdateGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UpdateAPI.UpdateGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UpdateAPIService UpdateIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.UpdateAPI.UpdateIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
