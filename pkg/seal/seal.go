package seal

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
)

//go:embed files/script.sh
var Script string

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
	// args are the positional (non-flag) command-line arguments.
	args []string
}

// Instance - get instance of seal struct from CLI flags
//    				exposed just for testing purposes
func Instance() *Seal {
	s := new(Seal)
	s.DescriptionSecret = "Your passowrd, your token,... to be encrypted"
	s.DescriptionHost = "Base URL of your sealed-secret API .i.e https://seal.apps.elm.sa"
	s.DescriptionNamespace = "k8s namespace where this secret can be decrypted. " +
		"If not specified, the scope of sealing will be cluster wide"
	s.DescriptionName = "k8s secret name where this secret can be decrypted." +
		"It requires to specify also the namespace of this k8s secret using -namespace option"
	s.DescriptionDryRun = "dry run - just print the command behind the scenes"
	return s
}

// ParseFlags parses the command-line arguments provided to the program.
// Typically os.Args[0] is provided as 'progname' and os.args[1:] as 'args'.
// Returns the Config in case parsing succeeded, or an error. In any case, the
// output of the flag.Parse is returned in output.
// A special case is usage requests with -h or -help: then the error
// flag.ErrHelp is returned and output will contain the usage message.
func ParseFlags(progname string, args []string) (s *Seal, output string, err error) {
	fs := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	fs.SetOutput(&buf)
	s = Instance()
	fs.StringVar(&s.Secret, "secret", "", s.DescriptionSecret)
	fs.StringVar(&s.Secret, "s", "", s.DescriptionSecret+" (shorthand)")
	fs.StringVar(&s.Host, "host", "", s.DescriptionHost)
	fs.StringVar(&s.Host, "h", "", s.DescriptionHost+" (shorthand)")
	fs.StringVar(&s.Namespace, "namespace", "", s.DescriptionNamespace)
	fs.StringVar(&s.Namespace, "n", "", s.DescriptionNamespace+" (shorthand)")
	fs.StringVar(&s.Name, "name", "", s.DescriptionName)
	fs.BoolVar(&s.DryRun, "dry-run", false, s.DescriptionDryRun)
	err = fs.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}
	s.args = fs.Args()
	return s, buf.String(), nil
}

func (s *Seal) Execute(executor func(n string, args ...string) error, script string) {
	if s.DryRun {
		fmt.Println(script)
	} else {
		executor("/bin/sh", "-c", script)
	}
}
