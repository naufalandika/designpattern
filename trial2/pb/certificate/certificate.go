package certificate

import "fmt"

type Certificate struct{}

func New() *Certificate {
	return &Certificate{}
}

func (c *Certificate) Generate() error {
	fmt.Println("generate certificate")
	return nil
}
