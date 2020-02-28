package testmod
import (
		"fmt"
       )
func SayHello(name, str string) string {
	return fmt.Sprintf("你好, %s, %s", name, str)
}
