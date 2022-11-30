// share elements
package share

import (
	"fmt"
)

var FyneTheme string = "dark"
var Bible string = "NET"
var BookName string = "John"
var BookAbb string = "John"
var Reference string = "John 3:16"
var Book int = 43
var Chapter int = 3
var Verse int = 16

var Mode string = ""
var DividerStr string = "--------------------"

func Divider() {
	fmt.Println(DividerStr)
}
