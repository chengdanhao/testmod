package testmod
import "fmt"
func SayHello(name string) string {
   return fmt.Sprintf("你好 %s", name)
}
package testmod
import (
"fmt"
)
func SayHello(name, str string) string {
return fmt.Sprintf("你好, %s, %s", name, str)
}
