package testmod
import (
		"fmt"
       )
func SayHello(name, str string) string {
	return fmt.Sprintf("Hi, %s, %s", name, str)
}
