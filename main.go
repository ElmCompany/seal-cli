package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"git.elm.sa/devops/seal-cli/pkg/seal"
	"git.elm.sa/devops/seal-cli/pkg/utils"
)

var (
	version = "dev"
)

func main() {
	if strings.Contains(os.Args[1], "version") {
		fmt.Println(version)
		return
	}
	s, output, err := seal.ParseFlags(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		fmt.Println(output)
		os.Exit(2)
	} else if err != nil {
		fmt.Println("got error:", err)
		fmt.Println("output:\n", output)
		os.Exit(1)
	}

	script := utils.ProcessString(seal.Script, &s)
	s.Execute(utils.RunCommand, script)

}
