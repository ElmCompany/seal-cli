package main

import (
	_ "embed"
	"fmt"

	"git.elm.sa/devops/seal-cli/pkg/seal"
	"git.elm.sa/devops/seal-cli/pkg/utils"
)

func main() {
	s := new(seal.Seal)
	s.Instance()
	script := s.Script()
	if s.DryRun {
		fmt.Println(script)
	} else {
		utils.RunCommand("/bin/sh", "-c", script)
	}

}
