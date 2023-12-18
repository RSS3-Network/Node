package decentralized

import (
	"github.com/NaturalSelectionLabs/goapi"
	"github.com/naturalselectionlabs/rss3-node/common/http/response"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

var _ = goapi.Interface(
	new(ActivityResponseInterface),
	response.BadRequest{},
	response.InternalErrorResp{},
	response.NotFound{},
	ActivityResponse{},
)

var _ = goapi.Interface(
	new(ActivitiesResponseInterface),
	response.BadRequest{},
	response.InternalErrorResp{},
	ActivitiesResponse{},
)

type ActivityResponseInterface interface {
	goapi.Response
}

type ActivitiesResponseInterface interface {
	goapi.Response
}

type ActivityRequest struct {
	goapi.InURL

	ID          string `description:"Retrieve details for the specified activity ID" examples:"[\"0x5ffa607a127d63fb36827075493d1de06f58fc44710b9ffb887b2effe02d2b8b\"]"`
	ActionLimit int    `description:"Specify the number of actions within the activity to retrieve" examples:"[10]" default:"10" min:"1" max:"20"`
	ActionPage  int    `description:"Specify the pagination for actions" default:"1" min:"1"`
}

type AccountActivitiesRequest struct {
	goapi.InURL

	Account        string            `description:"Retrieve activities from the specified account" examples:"[\"vitalik.eth\",\"stani.lens\",\"diygod.csb\"]"`
	Limit          int               `description:"Specify the number of activities to retrieve" examples:"[20]" default:"100" min:"1" max:"100"`
	ActionLimit    int               `description:"Specify the number of actions within the activity to retrieve" examples:"[10]" default:"10" min:"1" max:"20"`
	Cursor         *string           `description:"Specify the cursor used for pagination"`
	SinceTimestamp *uint64           `description:"Retrieve activities starting from this timestamp" examples:"[1654000000]"`
	UntilTimestamp *uint64           `description:"Retrieve activities up to this timestamp" examples:"[1696000000]"`
	Status         *bool             `description:"Retrieve activities based on status"`
	Direction      *filter.Direction `description:"Retrieve activities based on direction"`
	Network        []filter.Network  `description:"Retrieve activities from the specified network(s)" examples:"[[\"ethereum\",\"polygon\"]]"`
	Tag            []filter.Tag      `description:"Retrieve activities from the specified tag(s)"`
	Type           []filter.Type     `description:"Retrieve activities from the specified type(s)"`
	Platform       []filter.Platform `description:"Retrieve activities from the specified platform(s)"`
}

type AccountsActivitiesRequest struct {
	Account        []string          `json:"account" description:"Retrieve activities from the specified list of accounts"`
	Limit          int               `json:"limit,omitempty" description:"Specify the number of activities to retrieve" default:"100" min:"1" max:"500"`
	ActionLimit    int               `json:"action_limit,omitempty" description:"Specify the number of actions within the activity to retrieve" default:"10" min:"1" max:"20"`
	Cursor         *string           `json:"cursor,omitempty" description:"Specify the cursor used for pagination"`
	SinceTimestamp *uint64           `json:"since_timestamp,omitempty" description:"Retrieve activities starting from this timestamp"`
	UntilTimestamp *uint64           `json:"until_timestamp,omitempty" description:"Retrieve activities up to this timestamp"`
	Status         *bool             `json:"status,omitempty" description:"Retrieve activities based on status"`
	Direction      *filter.Direction `json:"direction,omitempty" description:"Retrieve activities based on direction"`
	Tag            []filter.Tag      `json:"tag,omitempty" description:"Retrieve activities from the specified tag(s)"`
	Type           []filter.Type     `json:"type,omitempty" description:"Retrieve activities from the specified type(s)"`
	Network        []filter.Network  `json:"network,omitempty" description:"Retrieve activities from the specified network(s)"`
	Platform       []filter.Platform `json:"platform,omitempty" description:"Retrieve activities from the specified platform(s)"`
}

type PlatformActivitiesRequest struct {
	goapi.InURL

	Platform       filter.Platform   `description:"Retrieve activities from the specified platform" examples:"[\"Uniswap\"]"`
	Limit          int               `description:"Specify the number of activities to retrieve" examples:"[20]" default:"100" min:"1" max:"100"`
	ActionLimit    int               `description:"Specify the number of actions within the activity to retrieve" examples:"[10]" default:"10" min:"1" max:"20"`
	Cursor         *string           `description:"Specify the cursor used for pagination"`
	SinceTimestamp *uint64           `description:"Retrieve activities starting from this timestamp" examples:"[1654000000]"`
	UntilTimestamp *uint64           `description:"Retrieve activities up to this timestamp" examples:"[1696000000]"`
	Status         *bool             `description:"Retrieve activities based on status"`
	Direction      *filter.Direction `description:"Retrieve activities based on direction"`
	Network        []filter.Network  `description:"Retrieve activities from the specified network(s)"`
	Tag            []filter.Tag      `description:"Retrieve activities from the specified tag(s)"`
	Type           []filter.Type     `description:"Retrieve activities from the specified type(s)"`
}

type NetworkActivitiesRequest struct {
	goapi.InURL

	Network        filter.Network    `description:"Retrieve activities from the specified network"`
	Limit          int               `description:"Specify the number of activities to retrieve" examples:"[20]" default:"100" min:"1" max:"100"`
	ActionLimit    int               `description:"Specify the number of actions within the activity to retrieve" examples:"[10]" default:"10" min:"1" max:"20"`
	Cursor         *string           `description:"Specify the cursor used for pagination"`
	SinceTimestamp *uint64           `description:"Retrieve activities starting from this timestamp" examples:"[1654000000]"`
	UntilTimestamp *uint64           `description:"Retrieve activities up to this timestamp" examples:"[1696000000]"`
	Status         *bool             `description:"Retrieve activities based on status"`
	Direction      *filter.Direction `description:"Retrieve activities based on direction"`
}

type ActivityResponse struct {
	goapi.StatusOK
	Data *schema.Feed    `json:"data"`
	Meta *MetaTotalPages `json:"meta"`
}

type ActivitiesResponse struct {
	goapi.StatusOK
	Data []*schema.Feed `json:"data"`
	Meta *MetaCursor    `json:"meta,omitempty"`
}

type MetaTotalPages struct {
	TotalPages int `json:"totalPages"`
}

type MetaCursor struct {
	Cursor string `json:"cursor"`
}
