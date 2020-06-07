package hello

import "fmt"

//Name will show 'Hello, #{name}'
func Name(n string) string {
	if len(n) > 0 {
		return fmt.Sprintf("Hello, %s!", n)
	}
	return "Hello, world!"
}
