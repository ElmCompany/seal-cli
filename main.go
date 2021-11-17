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
	Secret               string `json: "secret"`
	Host                 string `json: "host"`
	Namespace            string `json: "Namespace"`
	Name                 string `json: "Name"`
	DryRun               bool
	DescriptionSecret    string
	DescriptionHost      string
	DescriptionNamespace string
	DescriptionName      string
	DescriptionDryRun    string
}

func main() {
	s := new(Seal)
	s.DescriptionSecret = "Your passowrd, your token,... to be encrypted"
	s.DescriptionHost = "Base URL of your sealed-secret API .i.e https://seal.apps.elm.sa"
	s.DescriptionNamespace = "k8s namespace where this secret can be decrypted. " +
		"If not specified, the scope of sealing will be cluster wide"
	s.DescriptionName = "k8s secret name where this secret can be decrypted." +
		"It requires to specify also the namespace of this k8s secret using -namespace option"
	s.DescriptionDryRun = "dry run - just print the command behind the scenes"
	flag.StringVar(&s.Secret, "secret", "", s.DescriptionSecret)
	flag.StringVar(&s.Secret, "s", "", s.DescriptionSecret+" (shorthand)")
	flag.StringVar(&s.Host, "host", "", s.DescriptionHost)
	flag.StringVar(&s.Host, "h", "", s.DescriptionHost+" (shorthand)")
	flag.StringVar(&s.Namespace, "namespace", "", s.DescriptionNamespace)
	flag.StringVar(&s.Namespace, "n", "", s.DescriptionNamespace+" (shorthand)")
	flag.StringVar(&s.Name, "name", "", s.DescriptionName)
	flag.BoolVar(&s.DryRun, "dry-run", false, s.DescriptionDryRun)
	flag.Parse()
	scriptProcessed := processString(script, &s)
	if s.DryRun {
		fmt.Println(scriptProcessed)
	} else {
		runCommand("/bin/sh", "-c", scriptProcessed)
	}

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

func processString(str string, vars interface{}) string {
	tmpl, err := template.New("tmpl").Parse(str)

	if err != nil {
		panic(err)
	}
	return process(tmpl, vars)
}

func process(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		panic(err)
	}
	return tmplBytes.String()
}
