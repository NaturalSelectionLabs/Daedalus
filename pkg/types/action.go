package types

import (
	"bytes"
	"github.com/samber/lo"
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
	Build      *Job `yaml:"build,omitempty"`
	DeployDev  *Job `yaml:"deploy-dev,omitempty"`
	DeployProd *Job `yaml:"deploy-prod,omitempty"`
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

func (w Workflow) Migrate(options ...FlowOption) Workflow {

	newWorkflow := w.migrate()

	for _, option := range options {
		newWorkflow = option(newWorkflow)
	}

	return newWorkflow
}

func (w Workflow) migrate() (newWorkflow Workflow) {

	var (
		context    = "."
		dockerfile = "./Dockerfile"
	)

	if buildJob := w.Jobs.Build; &buildJob == nil {
		if pushStep, ok := lo.Find(buildJob.Steps, func(item Step) bool {
			return strings.HasSuffix(item.Uses, "docker/build-push-action")
		}); ok {
			if _context, ok := pushStep.With["context"]; ok {
				context = _context
			}
			if _dockerfile, ok := pushStep.With["file"]; ok {
				dockerfile = _dockerfile
			}
		}

	}

	newWorkflow.Name = w.Name
	newWorkflow.On = w.On

	newWorkflow.Jobs = Jobs{
		Build: &Job{
			Step: Step{
				Uses: "NaturalSelectionLabs/Daedalus/.github/workflows/docker-tpl.yaml@main",
				With: map[string]string{
					"images":     w.Env.ImageName,
					"context":    context,
					"dockerfile": dockerfile,
				},
			},
			Secrets: map[string]string{
				"DOCKERHUB_USERNAME": "${{ secrets.DOCKERHUB_USERNAME }}",
				"DOCKERHUB_TOKEN":    "${{ secrets.DOCKERHUB_TOKEN }}",
			},
		},
	}

	if w.Jobs.DeployDev != nil {
		newWorkflow.Jobs.DeployDev = &Job{
			Step: Step{
				Uses: "NaturalSelectionLabs/Daedalus/.github/workflows/deploy-tpl.yaml@main",
				With: map[string]string{
					"images":        w.Env.ImageName,
					"tag":           "sha-${{ github.sha }}",
					"cluster":       "dev",
					"namespace":     "<namespace>",
					"chart":         "web-app",
					"releaseName":   "<app-name>",
					"overrideFiles": "deploy/dev/values.yaml",
				},
			},
			Needs: w.Jobs.DeployDev.Needs,
		}
	}

	if w.Jobs.DeployProd != nil {
		newWorkflow.Jobs.DeployProd = &Job{
			Step: Step{
				Uses: "NaturalSelectionLabs/Daedalus/.github/workflows/deploy-tpl.yaml@main",
				With: map[string]string{
					"images":        w.Env.ImageName,
					"tag":           "sha-${{ github.sha }}",
					"cluster":       "prod",
					"namespace":     "<namespace>",
					"chart":         "web-app",
					"releaseName":   "<app-name>",
					"overrideFiles": "deploy/prod/values.yaml",
				},
			},
			Needs: w.Jobs.DeployProd.Needs,
		}
	}
	return
}
