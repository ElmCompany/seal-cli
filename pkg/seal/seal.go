package seal

import (
	_ "embed"
	"flag"

	"git.elm.sa/devops/seal-cli/pkg/utils"
)

//go:embed files/script.sh
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

func (s *Seal) Instance() {
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
}

func (s *Seal) Script() string {
	return utils.ProcessString(script, &s)
}
