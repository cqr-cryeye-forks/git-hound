package main

import (
	_ "net/http/pprof"

	"git-hound/cmd"
)

func main() {
	cmd.Execute()
}
