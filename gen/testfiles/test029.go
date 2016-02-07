package testfiles

import "fmt"

func (p *Person) SayHello(r *Person) string {
	return fmt.Sprintf("Hello, %v", r.Name)
}
