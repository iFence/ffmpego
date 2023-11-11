package main

import (
	"ffmpego/src/cmd"
	"fmt"
)

func main() {
	aa  := fmt.Sprintf("d = %.2f, s = %s\n", 60/23.0, "hello")
	print(aa)
	cmd.Execute()
}
