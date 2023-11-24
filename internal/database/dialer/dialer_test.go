package dialer_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/database/dialer"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/cockroachdb"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name   string
		driver database.Driver
	}{
		{
			name:   "cockroach",
			driver: database.DriverCockroach,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			container, dataSourceName, err := createContainer(context.Background(), testcase.driver)
			require.NoError(t, err)

			t.Cleanup(func() {
				require.NoError(t, gnomock.Stop(container))
			})

			client, err := dialer.Dial(context.Background(), database.Config{
				Driver: testcase.driver,
				URI:    dataSourceName,
				Mode:   database.ModeSingle,
			})

			require.NoError(t, err)
			require.NotNil(t, client)
		})
	}
}

func createContainer(ctx context.Context, driver database.Driver) (container *gnomock.Container, dataSourceName string, err error) {
	switch driver {
	case database.DriverCockroach:
		preset := cockroachdb.Preset(
			cockroachdb.WithDatabase("test"),
			cockroachdb.WithVersion("v23.1.8"),
		)

		container, err = gnomock.Start(preset, gnomock.WithContext(ctx))
		if err != nil {
			return nil, "", err
		}

		return container, fmt.Sprintf("postgres://root@%s:%d/%s?sslmode=disable", container.Host, container.DefaultPort(), "test"), nil
	default:
		return nil, "", fmt.Errorf("unsupported driver: %s", driver)
	}
}
