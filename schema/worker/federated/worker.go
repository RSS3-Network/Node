package federated

import (
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/protocol-go/schema/tag"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Worker --linecomment --output worker_string.go --json --yaml --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.6 --type=Worker --linecomment --output ../../../docs/schemas/FederatedWorker.yaml -t ../../../docs/schemas/tmpl/Federated.yaml.tmpl
type Worker int

const (
	Mastodon Worker = iota + 1 // mastodon
	Bluesky                    // bluesky
)

func (w Worker) Component() string {
	return "federated"
}

func (w Worker) Name() string {
	return w.String()
}

var _ echo.BindUnmarshaler = (*Worker)(nil)

func (w *Worker) UnmarshalParam(param string) error {
	worker, err := WorkerString(param)
	if err != nil {
		return err
	}

	*w = worker

	return nil
}

func GetValueByWorkerStr(workerStr string) Worker {
	return _WorkerNameToValueMap[workerStr]
}

// ToTagsMap is a map of worker to tags
var ToTagsMap = map[Worker][]tag.Tag{
	Mastodon: {tag.Social},
	Bluesky:  {tag.Social},
}
