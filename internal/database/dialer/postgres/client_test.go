package postgres_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/adrianbrad/psqldocker"
	_ "github.com/lib/pq"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/database/dialer"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/node/schema/worker/decentralized"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name                        string
		driver                      database.Driver
		partition                   *bool
		coreWorkerActivityCreated   []*activityx.Activity
		otherWorkerActivityCreated  []*activityx.Activity
		coreWorkerActivityUpdated   []*activityx.Activity
		otherWorkerActivityUpdated  []*activityx.Activity
		activityUpdatedLowPriority  []*activityx.Activity
		activityUpdatedHighPriority []*activityx.Activity
	}{
		{
			name:      "postgres",
			driver:    database.DriverPostgreSQL,
			partition: lo.ToPtr(true),
			coreWorkerActivityCreated: []*activityx.Activity{
				{
					ID:      "0x30182d4468ddc7001b897908203abb57939fc57663c491435a2f88cafd51d101",
					Network: network.Ethereum,
					From:    "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
					To:      "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
					Type:    typex.TransactionTransfer,
					Actions: []*activityx.Action{
						{
							Type: typex.TransactionTransfer,
							From: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
							To:   "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
						},
					},
					Timestamp: uint64(time.Now().Unix()),
				},
				{
					ID:      "0xedc029f7c7acce7b72939f8bfff44c5fdc7e64e3e2ba650d195799db8fec4c90",
					Network: network.Ethereum,
					From:    "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
					To:      "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
					Type:    typex.TransactionTransfer,
					Actions: []*activityx.Action{
						{
							Type: typex.TransactionTransfer,
							From: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
							To:   "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
						},
					},
					Timestamp: uint64(time.Now().Add(-3 * 31 * 24 * time.Hour).Unix()),
				},
			},
			otherWorkerActivityCreated: []*activityx.Activity{
				{
					ID:       "0xc5549085b553588eb6fe13878c761f3447a87df1e28ebdc9af1ff30340b764ab",
					Network:  network.Ethereum,
					From:     "0xd66570eDCdAde24eDda902B28F2172e9Ef8395dD",
					To:       "0x5B93D80DA1a359340d1F339FB574bDC56763f995",
					Platform: decentralized.PlatformUniswap.String(),
					Type:     typex.ExchangeSwap,
					Actions: []*activityx.Action{
						{
							Type: typex.ExchangeSwap,
							From: "0xd66570eDCdAde24eDda902B28F2172e9Ef8395dD",
							To:   "0x5B93D80DA1a359340d1F339FB574bDC56763f995",
						},
					},
					Timestamp: uint64(time.Now().Add(-3 * 31 * 24 * time.Hour).Unix()),
				},
			},
			coreWorkerActivityUpdated: []*activityx.Activity{
				{
					ID:      "0x30182d4468ddc7001b897908203abb57939fc57663c491435a2f88cafd51d101",
					Network: network.Ethereum,
					From:    "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
					To:      "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
					Type:    typex.TransactionTransfer,
					Actions: []*activityx.Action{
						{
							Type: typex.TransactionTransfer,
							From: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
							To:   "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
						},
						{
							Type: typex.TransactionTransfer,
							From: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
							To:   "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
						},
					},
					Timestamp: uint64(time.Now().Unix()),
				},
			},
			otherWorkerActivityUpdated: []*activityx.Activity{
				{
					ID:       "0xc5549085b553588eb6fe13878c761f3447a87df1e28ebdc9af1ff30340b764ab",
					Network:  network.Ethereum,
					From:     "0xd66570eDCdAde24eDda902B28F2172e9Ef8395dD",
					To:       "0x5B93D80DA1a359340d1F339FB574bDC56763f995",
					Platform: decentralized.PlatformUniswap.String(),
					Type:     typex.ExchangeLiquidity,
					Actions: []*activityx.Action{
						{
							Type: typex.ExchangeLiquidity,
							From: "0xd66570eDCdAde24eDda902B28F2172e9Ef8395dD",
							To:   "0x5B93D80DA1a359340d1F339FB574bDC56763f995",
						},
					},
					Timestamp: uint64(time.Now().Add(-3 * 31 * 24 * time.Hour).Unix()),
				},
			},
			activityUpdatedLowPriority: []*activityx.Activity{
				{
					ID:      "0xc5549085b553588eb6fe13878c761f3447a87df1e28ebdc9af1ff30340b764ab",
					Network: network.Ethereum,
					From:    "0xd66570eDCdAde24eDda902B28F2172e9Ef8395dD",
					To:      "0x5B93D80DA1a359340d1F339FB574bDC56763f995",
					Type:    typex.TransactionApproval,
					Actions: []*activityx.Action{
						{
							Type: typex.TransactionTransfer,
							From: "0x69Ed24795649c23F9C13BFE317432fe0e679f1F6",
							To:   "0x5B93D80DA1a359340d1F339FB574bDC56763f995",
						},
						{
							Type: typex.TransactionTransfer,
							From: "0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640",
							To:   "0x2cC846fFf0b08FB3bFfaD71f53a60B4b6E6d6482",
						},
					},
					Timestamp: uint64(time.Now().Add(-3 * 31 * 24 * time.Hour).Unix()),
				},
			},
			activityUpdatedHighPriority: []*activityx.Activity{
				{
					ID:       "0x30182d4468ddc7001b897908203abb57939fc57663c491435a2f88cafd51d101",
					Network:  network.Ethereum,
					From:     "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
					To:       "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
					Platform: decentralized.Platform1Inch.String(),
					Type:     typex.ExchangeSwap,
					Actions: []*activityx.Action{
						{
							Type: typex.ExchangeSwap,
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

			var (
				container      *psqldocker.Container
				dataSourceName string
				err            error
			)

			for {
				container, dataSourceName, err = createContainer(context.Background(), testcase.driver, *testcase.partition)
				if err == nil {
					break
				}
			}

			t.Cleanup(func() {
				require.NoError(t, container.Close())
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
			require.NoError(t, client.SaveActivities(context.Background(), testcase.coreWorkerActivityCreated, true))
			require.NoError(t, client.SaveActivities(context.Background(), testcase.otherWorkerActivityCreated, false))

			// Query first activity
			for _, activity := range append(testcase.coreWorkerActivityCreated, testcase.otherWorkerActivityCreated...) {
				data, page, err := client.FindActivity(context.Background(), model.ActivityQuery{ID: lo.ToPtr(activity.ID), ActionLimit: 10})
				require.NoError(t, err)
				require.NotNil(t, data)
				require.Greater(t, lo.FromPtr(page), 0)
				require.Equal(t, data.ID, activity.ID)
				require.Equal(t, data.From, activity.From)
				require.Equal(t, data.To, activity.To)
				require.Equal(t, data.Platform, activity.Platform)
			}

			// Query activities by account.
			accounts := make(map[string]int)

			for _, activity := range append(testcase.coreWorkerActivityCreated, testcase.otherWorkerActivityCreated...) {
				accounts[activity.From]++
				accounts[activity.To]++
			}

			for account, count := range accounts {
				activities, err := client.FindActivities(context.Background(), model.ActivitiesQuery{Owner: &account, Limit: 100})
				require.NoError(t, err)
				require.Len(t, activities, count)
			}

			// Update activities.
			require.NoError(t, client.SaveActivities(context.Background(), testcase.coreWorkerActivityUpdated, true))
			require.NoError(t, client.SaveActivities(context.Background(), testcase.otherWorkerActivityUpdated, false))

			// Query updated activities.
			for _, activity := range append(testcase.coreWorkerActivityUpdated, testcase.otherWorkerActivityUpdated...) {
				data, page, err := client.FindActivity(context.Background(), model.ActivityQuery{ID: lo.ToPtr(activity.ID), ActionLimit: 10})
				require.NoError(t, err)
				require.NotNil(t, data)
				require.Greater(t, lo.FromPtr(page), 0)
				require.Equal(t, data.ID, activity.ID)
				require.Equal(t, data.From, activity.From)
				require.Equal(t, data.To, activity.To)
				require.Equal(t, data.Platform, activity.Platform)
				require.Equal(t, data.Type, activity.Type)
			}

			// Update activities with low priority.
			require.NoError(t, client.SaveActivities(context.Background(), testcase.activityUpdatedLowPriority, true))

			// Query updated activities with low priority.
			for _, activity := range testcase.activityUpdatedLowPriority {
				data, page, err := client.FindActivity(context.Background(), model.ActivityQuery{ID: lo.ToPtr(activity.ID), ActionLimit: 10})
				require.NoError(t, err)
				require.NotNil(t, data)
				require.Greater(t, lo.FromPtr(page), 0)
				require.NotEqual(t, data.Platform, activity.Platform)
				require.NotEqual(t, data.Type, activity.Type)
			}

			// Check if the activities are not updated.
			shouldNotExist := []string{
				"0x69Ed24795649c23F9C13BFE317432fe0e679f1F6",
				"0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640",
				"0x2cC846fFf0b08FB3bFfaD71f53a60B4b6E6d6482",
			}

			for _, address := range shouldNotExist {
				addressTemp := address
				activities, err := client.FindActivities(context.Background(), model.ActivitiesQuery{Owner: &addressTemp, Limit: 100})
				require.NoError(t, err)
				require.Len(t, activities, 0)
			}

			// Update activities with high priority.
			require.NoError(t, client.SaveActivities(context.Background(), testcase.activityUpdatedHighPriority, false))

			// Query updated activities with low priority.
			for _, activity := range testcase.activityUpdatedHighPriority {
				data, page, err := client.FindActivity(context.Background(), model.ActivityQuery{ID: lo.ToPtr(activity.ID), ActionLimit: 10})
				require.NoError(t, err)
				require.NotNil(t, data)
				require.Greater(t, lo.FromPtr(page), 0)
				require.Equal(t, data.Platform, activity.Platform)
				require.Equal(t, data.Type, activity.Type)
			}
		})
	}
}

func createContainer(_ context.Context, driver database.Driver, _ bool) (container *psqldocker.Container, dataSourceName string, err error) {
	switch driver {
	case database.DriverPostgreSQL:
		c, err := psqldocker.NewContainer(
			"user",
			"password",
			"test",
		)
		if err != nil {
			return nil, "", fmt.Errorf("create psql container: %w", err)
		}

		return c, formatContainerURI(c), nil
	default:
		return nil, "", fmt.Errorf("unsupported driver: %s", driver)
	}
}

func formatContainerURI(container *psqldocker.Container) string {
	return fmt.Sprintf(
		"postgres://user:password@%s:%s/%s?sslmode=disable",
		"127.0.0.1",
		container.Port(),
		"test",
	)
}
