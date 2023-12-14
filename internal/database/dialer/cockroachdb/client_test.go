package cockroachdb_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/database/dialer"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/cockroachdb"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name        string
		driver      database.Driver
		partition   bool
		feedCreated []*schema.Feed
		feedUpdated []*schema.Feed
	}{
		{
			name:      "cockroach",
			driver:    database.DriverCockroachDB,
			partition: true,
			feedCreated: []*schema.Feed{
				{
					ID:      "0xddc42d4de320638dda200a59938514f7230bf6022355c6a8a7c39b9903598ced",
					Network: filter.NetworkArweave,
					From:    "0x566b8087067638b0cb16311e0f05bee58186e787",
					To:      "0x9e05155e5d924c179b39a8b9b427c1bea06face3",
					Type:    filter.TypeTransactionTransfer,
					Actions: []*schema.Action{
						{
							Type: filter.TypeTransactionTransfer,
							From: "0x000000A52a03835517E9d193B3c27626e1Bc96b1",
							To:   "0xA1b2DCAC834117F38FB0356b5176B5693E165c90",
						},
						{
							Type: filter.TypeTransactionTransfer,
							From: "0xA1b2DCAC834117F38FB0356b5176B5693E165c90",
							To:   "0x000000A52a03835517E9d193B3c27626e1Bc96b1",
						},
					},
					Timestamp: uint64(time.Now().Unix()),
				},
			},
			feedUpdated: []*schema.Feed{
				{
					ID:      "0xddc42d4de320638dda200a59938514f7230bf6022355c6a8a7c39b9903598ced",
					Network: filter.NetworkArweave,
					From:    "0x566b8087067638b0cb16311e0f05bee58186e787",
					To:      "0x9e05155e5d924c179b39a8b9b427c1bea06face3",
					Type:    filter.TypeTransactionTransfer,
					Actions: []*schema.Action{
						{
							Type: filter.TypeTransactionTransfer,
							From: "0x566b8087067638b0cb16311e0f05bee58186e787",
							To:   "0x9e05155e5d924c179b39a8b9b427c1bea06face3",
						},
					},
					Timestamp: uint64(time.Now().Unix()),
				},
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			container, dataSourceName, err := createContainer(context.Background(), testcase.driver, testcase.partition)
			require.NoError(t, err)

			t.Cleanup(func() {
				require.NoError(t, gnomock.Stop(container))
			})

			// Dial the database.
			client, err := dialer.Dial(context.Background(), &database.Config{
				Driver:    testcase.driver,
				URI:       dataSourceName,
				Partition: testcase.partition,
			})

			require.NoError(t, err)
			require.NotNil(t, client)

			// Migrate the database.
			require.NoError(t, client.Migrate(context.Background()))

			// Begin a transaction.
			transaction, err := client.Begin(context.Background())
			require.NoError(t, err)

			// Insert feeds.
			require.NoError(t, transaction.SaveFeeds(context.Background(), testcase.feedCreated))

			// Update feeds.
			require.NoError(t, transaction.SaveFeeds(context.Background(), testcase.feedUpdated))

			// Commit the transaction.
			require.NoError(t, transaction.Commit())
		})
	}
}

func createContainer(ctx context.Context, driver database.Driver, partition bool) (container *gnomock.Container, dataSourceName string, err error) {
	config := database.Config{
		Driver:    driver,
		Partition: partition,
	}

	switch driver {
	case database.DriverCockroachDB:
		preset := cockroachdb.Preset(
			cockroachdb.WithDatabase("test"),
			cockroachdb.WithVersion("v23.1.8"),
		)

		// Use a health check function to wait for the database to be ready.
		healthcheckFunc := func(ctx context.Context, container *gnomock.Container) error {
			config.URI = formatContainerURI(container)

			client, err := dialer.Dial(ctx, &config)
			if err != nil {
				return err
			}

			transaction, err := client.Begin(ctx)
			if err != nil {
				return err
			}

			defer lo.Try(transaction.Rollback)

			return nil
		}

		container, err = gnomock.Start(preset, gnomock.WithContext(ctx), gnomock.WithHealthCheck(healthcheckFunc))
		if err != nil {
			return nil, "", err
		}

		return container, formatContainerURI(container), nil
	default:
		return nil, "", fmt.Errorf("unsupported driver: %s", driver)
	}
}

func formatContainerURI(container *gnomock.Container) string {
	return fmt.Sprintf(
		"postgres://root@%s:%d/%s?sslmode=disable",
		container.Host,
		container.DefaultPort(),
		"test",
	)
}
