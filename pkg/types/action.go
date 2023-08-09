package types

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"strings"
)

type Workflow struct {
	Name string `yaml:"name"`
	On   struct {
		Push struct {
			Branches []string `yaml:"branches"`
		} `yaml:"push"`
		PullRequest struct {
			Branches []string `yaml:"branches"`
		} `yaml:"pull_request"`
	} `yaml:"on"`
	Env struct {
		ImageName string `yaml:"IMAGE_NAME,omitempty"`
	} `yaml:"env,omitempty"`
	Jobs Jobs `yaml:"jobs"`
}

type Jobs struct {
	Build      Job `yaml:"build,omitempty"`
	DeployDev  Job `yaml:"deploy-dev,omitempty"`
	DeployProd Job `yaml:"deploy-prod,omitempty"`
}

type Job struct {
	RunsOn  string   `yaml:"runs-on,omitempty"`
	Needs   []string `yaml:"needs,omitempty"`
	Steps   []Step   `yaml:"steps,omitempty"`
	Step    `yaml:",inline,omitempty"`
	Secrets map[string]string `yaml:"secrets,omitempty"`
}

type Step struct {
	Name string            `yaml:"name,omitempty"`
	Uses string            `yaml:"uses,omitempty"`
	With map[string]string `yaml:"with,omitempty"`
	// ... other fields for steps ...
}

func (w Workflow) String() string {
	var buffer bytes.Buffer
	encoder := yaml.NewEncoder(&buffer)
	encoder.SetIndent(2)
	_ = encoder.Encode(&w)
	return strings.ReplaceAll(buffer.String(), `"on"`, `on`)
}
