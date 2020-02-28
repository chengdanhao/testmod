package testmod
import "fmt"
func SayHello(name, greeting string) string {
   return fmt.Sprintf("Hi, %s, %s", name, greeting)
}
