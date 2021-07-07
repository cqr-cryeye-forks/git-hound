package main

import (
	_ "net/http/pprof"

	"github.com/cqr-cryeye-forks/git-hound/cmd"
)

func main() {
	cmd.Execute()
}
