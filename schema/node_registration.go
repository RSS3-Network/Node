package schema

type Maintainer struct {
	OwnerEVMAddress string `json:"owner_evm_address"`
	Signature       string `json:"signature"`
	Website         string `json:"website"`
}

type Node struct {
	Endpoint  string `json:"endpoint"`
	AccessKey string `json:"access_key"`
}

type Decentralized struct {
	Type    string   `json:"type"`
	Network []string `json:"network"`
}

type RSSHub struct {
	InstanceURL string `json:"instance_url"`
}

type IndexerModule struct {
	RSS RSS `json:"rss"`
}

type RSS struct {
	RSSHub RSSHub `json:"rsshub"`
}

type Component struct {
	Data Data `json:"data"`
}

type Data struct {
	IndexerModule IndexerModule `json:"indexer_module"`
}

type RegistrationData struct {
	Version    string     `json:"version"`
	Maintainer Maintainer `json:"maintainer"`
	Node       Node       `json:"node"`
	Component  Component  `json:"component"`
}
