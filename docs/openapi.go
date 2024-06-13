package docs

import (
	"embed"
	"fmt"
	"os"
	"sort"

	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/samber/lo"
	"github.com/tidwall/sjson"
	"go.uber.org/zap"
)

//go:embed openapi.json
var EmbedFS embed.FS

var FilePath = "./docs/openapi.json"

func Generate(endpoint string) ([]byte, error) {
	// Read the existing openapi.json file from EmbedFS
	file, err := EmbedFS.ReadFile("openapi.json")
	if err != nil {
		zap.L().Error("read file error", zap.Error(err))
		return nil, err
	}

	// Generate network, tag, platform, direction enum.
	if file, err = generateEnum(file); err != nil {
		return nil, err
	}

	// Write the updated file back to the file system
	_ = os.WriteFile(FilePath, file, 0600)

	// Generate servers.
	if file, err = generateServers(file, endpoint); err != nil {
		return nil, err
	}

	return file, nil
}

func generateServers(file []byte, endpoint string) ([]byte, error) {
	var err error

	file, err = sjson.SetBytes(file, "servers", []map[string]interface{}{
		{"url": endpoint, "description": "Node Endpoint"},
		{"url": "http://localhost", "description": "Localhost"},
	})
	if err != nil {
		zap.L().Error("sjson set servers error", zap.Error(err))
		return nil, err
	}

	return file, nil
}

func generateEnum(file []byte) ([]byte, error) {
	var err error

	// Generate network values.
	networks := lo.Filter(network.NetworkStrings(), func(s string, _ int) bool {
		return !lo.Contains([]string{
			network.Unknown.String(),
			network.Bitcoin.String(),
			network.SatoshiVM.String(),
			network.RSS.String(),
		}, s)
	})

	sort.Strings(networks)

	file, err = sjson.SetBytes(file, "components.schemas.Network.enum", networks)
	if err != nil {
		return nil, fmt.Errorf("sjson set network enum err: %w", err)
	}

	// Generate tag values.
	tags := tag.TagStrings()

	sort.Strings(tags)

	file, err = sjson.SetBytes(file, "components.schemas.Tag.enum", tags)
	if err != nil {
		return nil, fmt.Errorf("sjson set tag enum err: %w", err)
	}

	// Generate platform values.
	platforms := decentralized.PlatformStrings()

	sort.Strings(platforms)

	file, err = sjson.SetBytes(file, "components.schemas.Platform.enum", platforms)
	if err != nil {
		return nil, fmt.Errorf("sjson set platform enum err: %w", err)
	}

	// Generate direction values.
	file, err = sjson.SetBytes(file, "components.schemas.Direction.enum", activity.DirectionStrings())
	if err != nil {
		return nil, fmt.Errorf("sjson set direction enum err: %w", err)
	}

	// Generate type values.
	types := make([]string, 0)

	for _, v := range tag.TagValues() {
		for _, t := range schema.GetTypesByTag(v) {
			types = append(types, t.Name())
		}
	}

	types = lo.Uniq(types)

	sort.Strings(types)

	file, err = sjson.SetBytes(file, "components.schemas.Type.enum", types)
	if err != nil {
		return nil, fmt.Errorf("sjson set type enum err: %w", err)
	}

	return file, nil
}
