package testdata

import "fmt"

func (c Celsius) String() string {
	return fmt.Sprintf("%v", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%v", f)
}
