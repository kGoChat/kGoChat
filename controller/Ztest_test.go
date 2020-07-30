package controller

import (
	"fmt"
	"testing"
)

func TestController(t *testing.T) {
	var e = ExampleController{}
	fmt.Print(e.GetPing() + "\n")
	fmt.Print(e.GetHelloIris(), "\n")
}
