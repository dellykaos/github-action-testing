package hello

import "fmt"

//Add add a with b value
func Add(a, b int) int {
	return a + b
}

//Name will show 'Hello, #{name}'
func Name(n string) string {
	if len(n) > 0 {
		return fmt.Sprintf("Hello, %s!", n)
	}
	return "Hello, world!"
}
