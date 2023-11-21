package database_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/cockroachdb"
	"github.com/orlangure/gnomock/preset/mysql"
	"github.com/orlangure/gnomock/preset/postgres"
	"github.com/stretchr/testify/require"
)

func TestDial(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name      string
		arguments database.Driver
	}{
		{
			name:      "cockroach",
			arguments: database.DriverCockroach,
		},
		{
			name:      "postgres",
			arguments: database.DriverPostgres,
		},
		{
			name:      "mysql",
			arguments: database.DriverMySQL,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			container, dataSourceName, err := createContainer(context.TODO(), testcase.arguments)
			require.NoError(t, err)

			t.Cleanup(func() {
				require.NoError(t, gnomock.Stop(container))
			})

			client, err := database.Dial(context.TODO(), database.Config{
				Driver: testcase.arguments,
				URI:    dataSourceName,
			})

			require.NoError(t, err)
			require.NoError(t, client.Close())
		})
	}
}

func createContainer(ctx context.Context, driver database.Driver) (container *gnomock.Container, dataSourceName string, err error) {
	switch driver {
	case database.DriverCockroach:
		preset := cockroachdb.Preset(cockroachdb.WithDatabase("test"))

		container, err = gnomock.Start(preset, gnomock.WithContext(ctx))
		if err != nil {
			return nil, "", err
		}

		return container, fmt.Sprintf("postgres://root@%s:%d/%s?sslmode=disable", container.Host, container.DefaultPort(), "test"), nil
	case database.DriverPostgres:
		preset := postgres.Preset(postgres.WithDatabase("test"))

		container, err = gnomock.Start(preset, gnomock.WithContext(ctx))
		if err != nil {
			return nil, "", err
		}

		return container, fmt.Sprintf("postgres://root@%s:%d/%s?sslmode=disable", container.Host, container.DefaultPort(), "test"), nil
	case database.DriverMySQL:
		preset := mysql.Preset(mysql.WithDatabase("test"))

		container, err = gnomock.Start(preset, gnomock.WithContext(ctx))
		if err != nil {
			return nil, "", err
		}

		return container, fmt.Sprintf("root:root@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", container.Host, container.DefaultPort(), "test"), nil
	default:
		return nil, "", fmt.Errorf("unsupported driver: %s", driver)
	}
}
