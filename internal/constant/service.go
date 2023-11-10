package constant

import (
	"fmt"
	"github.com/naturalselectionlabs/framework"
)

var ServiceName string

func BuildServiceName() string {
	return fmt.Sprintf("%s.%s", framework.Name, ServiceName)
}
