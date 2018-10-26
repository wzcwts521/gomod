package main

import (
	"fmt"

	"github.com/patrickdappollonio/env"
)

func main() {
	environmentvar := env.GetDefault("PATH", "hello")
	abc := env.GetDefault("", "")
	fmt.Println(environmentvar)
	fmt.Println(abc)
}
