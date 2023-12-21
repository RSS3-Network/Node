package decentralized

import (
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

type ActivityRequest struct {
	ID          string `param:"id" description:"Retrieve details for the specified activity ID" examples:"[\"0x5ffa607a127d63fb36827075493d1de06f58fc44710b9ffb887b2effe02d2b8b\"]"`
	ActionLimit int    `query:"action_limit" description:"Specify the number of actions within the activity to retrieve" examples:"[10]" default:"10" min:"1" max:"20"`
	ActionPage  int    `query:"action_page" description:"Specify the pagination for actions" default:"1" min:"1"`
}

type AccountActivitiesRequest struct {
	Account        string            `param:"account" description:"Retrieve activities from the specified account" examples:"[\"vitalik.eth\",\"stani.lens\",\"diygod.csb\"]"`
	Limit          int               `query:"limit" description:"Specify the number of activities to retrieve" examples:"[20]" default:"100" min:"1" max:"100"`
	ActionLimit    int               `query:"action_limit" description:"Specify the number of actions within the activity to retrieve" examples:"[10]" default:"10" min:"1" max:"20"`
	Cursor         *string           `query:"cursor" description:"Specify the cursor used for pagination"`
	SinceTimestamp *uint64           `query:"since_timestamp" description:"Retrieve activities starting from this timestamp" examples:"[1654000000]"`
	UntilTimestamp *uint64           `query:"until_timestamp" description:"Retrieve activities up to this timestamp" examples:"[1696000000]"`
	Status         *bool             `query:"success" description:"Retrieve activities based on status"`
	Direction      *filter.Direction `query:"direction" description:"Retrieve activities based on direction"`
	Network        []filter.Network  `query:"network" description:"Retrieve activities from the specified network(s)" examples:"[[\"ethereum\",\"polygon\"]]"`
	Tag            []filter.Tag      `query:"tag" description:"Retrieve activities from the specified tag(s)"`
	Type           []filter.Type     `query:"type" description:"Retrieve activities from the specified type(s)"`
	Platform       []filter.Platform `query:"platform" description:"Retrieve activities from the specified platform(s)"`
}

type ActivityResponse struct {
	Data *schema.Feed    `json:"data"`
	Meta *MetaTotalPages `json:"meta"`
}

type ActivitiesResponse struct {
	Data []*schema.Feed `json:"data"`
	Meta *MetaCursor    `json:"meta,omitempty"`
}

type MetaTotalPages struct {
	TotalPages int `json:"totalPages"`
}

type MetaCursor struct {
	Cursor string `json:"cursor"`
}
