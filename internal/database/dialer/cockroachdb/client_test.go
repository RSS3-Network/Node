package cockroachdb_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/cockroachdb"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/database/dialer"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name        string
		driver      database.Driver
		partition   *bool
		feedCreated []*schema.Feed
		feedUpdated []*schema.Feed
	}{
		{
			name:      "cockroach",
			driver:    database.DriverCockroachDB,
			partition: lo.ToPtr(true),
			feedCreated: []*schema.Feed{
				{
					ID:      "0x30182d4468ddc7001b897908203abb57939fc57663c491435a2f88cafd51d101",
					Network: filter.NetworkEthereum,
					From:    "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
					To:      "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
					Type:    filter.TypeTransactionTransfer,
					Actions: []*schema.Action{
						{
							Type: filter.TypeTransactionTransfer,
							From: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
							To:   "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
						},
					},
					Timestamp: uint64(time.Now().Unix()),
				},
				{
					ID:      "0xedc029f7c7acce7b72939f8bfff44c5fdc7e64e3e2ba650d195799db8fec4c90",
					Network: filter.NetworkEthereum,
					From:    "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
					To:      "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
					Type:    filter.TypeTransactionTransfer,
					Actions: []*schema.Action{
						{
							Type: filter.TypeTransactionTransfer,
							From: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
							To:   "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
						},
					},
					Timestamp: uint64(time.Now().Add(-3 * 31 * 24 * time.Hour).Unix()),
				},
			},
			feedUpdated: []*schema.Feed{
				{
					ID:      "0x30182d4468ddc7001b897908203abb57939fc57663c491435a2f88cafd51d101",
					Network: filter.NetworkEthereum,
					From:    "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
					To:      "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
					Type:    filter.TypeTransactionTransfer,
					Actions: []*schema.Action{
						{
							Type: filter.TypeTransactionTransfer,
							From: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
							To:   "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
						},
						{
							Type: filter.TypeTransactionTransfer,
							From: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
							To:   "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
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

			var container *gnomock.Container
			var dataSourceName string
			var err error

			for {
				container, dataSourceName, err = createContainer(context.Background(), testcase.driver, *testcase.partition)
				if err == nil {
					break
				}
			}

			t.Cleanup(func() {
				require.NoError(t, gnomock.Stop(container))
			})

			// Dial the database.
			client, err := dialer.Dial(context.Background(), &config.Database{
				Driver:    testcase.driver,
				URI:       dataSourceName,
				Partition: testcase.partition,
			})

			require.NoError(t, err)
			require.NotNil(t, client)

			// Migrate the database.
			require.NoError(t, client.Migrate(context.Background()))

			// Begin a transaction.
			require.NoError(t, client.SaveFeeds(context.Background(), testcase.feedCreated))

			// Query first feed.
			for _, feed := range testcase.feedCreated {
				data, page, err := client.FindFeed(context.Background(), model.FeedQuery{ID: lo.ToPtr(feed.ID), ActionLimit: 10})
				require.NoError(t, err)
				require.NotNil(t, data)
				require.Greater(t, lo.FromPtr(page), 0)
				require.Equal(t, data.ID, feed.ID)
				require.Equal(t, data.From, feed.From)
				require.Equal(t, data.To, feed.To)
			}

			// Query feeds by account.
			accounts := make(map[string]int)

			for _, feed := range testcase.feedCreated {
				accounts[feed.From]++
				accounts[feed.To]++
			}

			for account, count := range accounts {
				feeds, err := client.FindFeeds(context.Background(), model.FeedsQuery{Owner: &account, Limit: 100})
				require.NoError(t, err)
				require.Len(t, feeds, count)
			}

			// Update feeds.
			require.NoError(t, client.SaveFeeds(context.Background(), testcase.feedUpdated))
		})
	}
}

func createContainer(ctx context.Context, driver database.Driver, partition bool) (container *gnomock.Container, dataSourceName string, err error) {
	config := config.Database{
		Driver:    driver,
		Partition: &partition,
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
