package orderserviceprovider

import (
	"context"
	"fmt"
	"github.com/GowthamGirithar/contract-testing-demo/order-service-consumer/internal/domain"
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
	"time"

	message "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	protoPath, _ := filepath.Abs("../../../../proto/order-service/order.proto")
	pactDir, _ := filepath.Abs("../../../../pacts/orderserviceprovider/createorder")

	t.Run("happy flow", func(t *testing.T) {

		mockProvider, err := message.NewSynchronousPact(message.Config{
			Consumer: "create-order-consumer",
			Provider: "create-order-provider",
			PactDir:  pactDir,
		})
		assert.NoError(t, err)

		grpcInteraction := `{
		"pact:proto": "` + protoPath + `",
		"pact:proto-service": "Order/CreateOrder",
		"pact:content-type": "application/protobuf",
		"request": {
			"order_number" :   "matching(number, 3)",
			"customer_email":  "notEmpty('test@gmail.com')"
		},
		"response": {
			"order_number": "matching(number, 4)"
		}
	}`
		err = mockProvider.
			AddSynchronousMessage("create order request").
			UsingPlugin(message.PluginConfig{
				Plugin: "protobuf",
			}).
			WithContents(grpcInteraction, "application/grpc").
			StartTransport("grpc", "127.0.0.1", nil).
			ExecuteTest(t, func(transport message.TransportConfig, m message.SynchronousMessage) error {
				err := NewClient(transport.Address, fmt.Sprintf("%d", transport.Port), "test", 1*time.Second).
					CreateOrder(context.Background(), domain.Order{
						OrderID:       "12",
						CustomerEmail: "test@gmail.com",
					})
				assert.NoError(t, err)
				return err
			})
		assert.NoError(t, err)
	})

	t.Run("invalid request", func(t *testing.T) {

		mockProvider, err := message.NewSynchronousPact(message.Config{
			Consumer: "create-order-consumer",
			Provider: "create-order-provider",
			PactDir:  pactDir,
		})
		assert.NoError(t, err)

		grpcInteraction := `{
		"pact:proto": "` + protoPath + `",
		"pact:proto-service": "Order/CreateOrder",
		"pact:content-type": "application/protobuf",
		"request": {
			"order_number" :   "matching(number, 3)"
		},
		"responseMetadata" :{
          "grpc-status" : "INTERNAL",
          "grpc-message" : "matching(type, 'Invalid email format')"
        }
	}`

		err = mockProvider.
			AddSynchronousMessage("create order request - invalid").
			UsingPlugin(message.PluginConfig{
				Plugin: "protobuf",
			}).
			WithContents(grpcInteraction, "application/grpc").
			StartTransport("grpc", "127.0.0.1", nil).
			ExecuteTest(t, func(transport message.TransportConfig, m message.SynchronousMessage) error {
				err := NewClient(transport.Address, fmt.Sprintf("%d", transport.Port), "test", 1*time.Second).
					CreateOrder(context.Background(), domain.Order{
						OrderID:       "123",
						CustomerEmail: "",
					})
				require.Error(t, err)
				return nil
			})
		assert.NoError(t, err)
	})
}
