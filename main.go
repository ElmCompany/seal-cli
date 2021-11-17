package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"os/exec"
	"text/template"
)

//go:embed scripts/script.sh
var script string

type Seal struct {
	Secret string `json: "secret"`
	Host   string `json: "host`
}

func main() {
	s := new(Seal)
	flag.StringVar(&s.Secret, "secret", "", "Your passowrd, your token,... that you want to encrypt")
	flag.StringVar(&s.Host, "host", "", "Base URL of your sealed-secret API .i.e https://seal.apps.elm.sa")
	flag.Parse()
	runCommand("/bin/sh", "-c", processString(script, &s))
	fmt.Printf("%+v\n", *s)
	fmt.Print(*s)

}

func runCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	// error output and standard output of command is connected to the same pipeline
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}
	// Get real-time output from the pipe to the terminal and print
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func process(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		panic(err)
	}
	return tmplBytes.String()
}

func processString(str string, vars interface{}) string {
	tmpl, err := template.New("tmpl").Parse(str)

	if err != nil {
		panic(err)
	}
	return process(tmpl, vars)
}
