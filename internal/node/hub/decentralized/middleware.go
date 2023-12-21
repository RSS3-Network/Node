package decentralized

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"reflect"
)

func (h *Hub) Bind(i interface{}, c echo.Context) error {
	req := reflect.ValueOf(i).Elem()
	tags := make([]filter.Tag, 0)

	// Parse enum values.
	queryParams := c.QueryParams()
	for k, v := range queryParams {
		for _, param := range v {
			var field, value reflect.Value

			switch k {
			case "platform":
				platform, err := filter.PlatformString(param)
				if err != nil {
					return fmt.Errorf("invalid platform: %s", param)
				}

				field = req.FieldByName("Platform")
				value = reflect.ValueOf(platform)
			case "network":
				network, err := filter.NetworkString(param)
				if err != nil {
					return fmt.Errorf("invalid network: %s", param)
				}

				field = req.FieldByName("Network")
				value = reflect.ValueOf(network)
			case "tag":
				tag, err := filter.TagString(param)
				if err != nil {
					return fmt.Errorf("invalid tag: %s", param)
				}

				tags = append(tags, tag)
				field = req.FieldByName("Tag")
				value = reflect.ValueOf(tag)
			case "type":
				for _, tag := range tags {
					typex, err := filter.TypeString(tag, param)
					if err == nil {
						field = req.FieldByName("Type")
						value = reflect.ValueOf(typex)
						break
					}
				}

				if field == (reflect.Value{}) {
					return fmt.Errorf("invalid type: %s", param)
				}
			case "direction":
				direction, err := filter.DirectionString(param)
				if err != nil {
					return fmt.Errorf("invalid direction: %s", param)
				}

				field = req.FieldByName("Direction")
				value = reflect.ValueOf(&direction)
			default:
				continue
			}

			if field.IsValid() {
				switch field.Kind() {
				case reflect.Slice:
					field.Set(reflect.Append(field, value))
				default:
					field.Set(value)
				}
				delete(queryParams, k)
			}
		}
	}

	// Use default binder.
	db := new(echo.DefaultBinder)
	if err := db.Bind(i, c); err != nil {
		return err
	}

	return nil
}
