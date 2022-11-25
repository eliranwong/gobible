// share elements
package share

import (
	"fmt"
)

// TODO: save config settings
var Bible string = "NET"

var Mode string = ""
var DividerStr string = "--------------------"

func Divider() {
	fmt.Println(DividerStr)
}
