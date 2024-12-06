/*
Italian eInvoice API

Testing LogAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package invoicesdk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	invoicesdk "github.com/invoicetronic/invoice-go-sdk"
)

func Test_invoicesdk_LogAPIService(t *testing.T) {

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)

	t.Run("Test LogAPIService InvoiceV1LogGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LogAPI.InvoiceV1LogGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogAPIService InvoiceV1LogIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.LogAPI.InvoiceV1LogIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
