package filter

import "github.com/labstack/echo/v4"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Tag --transform=snake --trimprefix=Tag --output tag_string.go --json --sql
type Tag uint64

const (
	TagUnknown Tag = iota
	TagTransaction
	TagCollectible
	TagExchange
	TagSocial
	TagRSS
)

func TagAndTypeString(tagValue string, typeValue string) (Tag, Type, error) {
	tag, err := TagString(tagValue)
	if err != nil {
		return TagUnknown, nil, err
	}

	_type, err := TypeString(tag, typeValue)
	if err != nil {
		return TagUnknown, nil, err
	}

	return tag, _type, err
}

var _ echo.BindUnmarshaler = (*Tag)(nil)

func (t *Tag) UnmarshalParam(param string) error {
	tag, err := TagString(param)
	if err != nil {
		return err
	}

	*t = tag

	return nil
}
