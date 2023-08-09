package converter

import (
	"github.com/naturalSelectionLabs/daedalus/pkg/types"
	"github.com/samber/lo"
	"strings"
)

type FlowOption func(workflow *types.Workflow) *types.Workflow

func MigrateWorkflow(workflow types.Workflow, options ...FlowOption) (newWorkflow types.Workflow) {

	newWorkflow.Name = workflow.Name

	newWorkflow.On = workflow.On

	var (
		context    = "."
		dockerfile = "./Dockerfile"
	)

	if buildJob := workflow.Jobs.Build; &buildJob == nil {
		if pushStep, ok := lo.Find(buildJob.Steps, func(item types.Step) bool {
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

	newWorkflow.Jobs = types.Jobs{
		Build: types.Job{
			Step: types.Step{
				Uses: "NaturalSelectionLabs/Daedalus/.github/workflows/docker-tpl.yaml@main",
				With: map[string]string{
					"images":     workflow.Env.ImageName,
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

	if &workflow.Jobs.DeployDev != nil {
		newWorkflow.Jobs.DeployDev = types.Job{
			Step: types.Step{
				Uses: "NaturalSelectionLabs/Daedalus/.github/workflows/deploy-tpl.yaml@main",
				With: map[string]string{
					"images":      workflow.Env.ImageName,
					"tag":         "sha-${{ github.sha }}",
					"cluster":     "dev",
					"namespace":   "<namespace>",
					"chart":       "web-app",
					"releaseName": "<app-name>",
				},
			},
			Needs: workflow.Jobs.DeployDev.Needs,
		}
	}

	if &workflow.Jobs.DeployProd != nil {
		newWorkflow.Jobs.DeployProd = types.Job{
			Step: types.Step{
				Uses: "NaturalSelectionLabs/Daedalus/.github/workflows/deploy-tpl.yaml@main",
				With: map[string]string{
					"images":      workflow.Env.ImageName,
					"tag":         "sha-${{ github.sha }}",
					"cluster":     "dev",
					"namespace":   "<namespace>",
					"chart":       "web-app",
					"releaseName": "<app-name>",
				},
			},
			Needs: workflow.Jobs.DeployProd.Needs,
		}
	}

	for _, option := range options {
		option(&newWorkflow)
	}

	return
}

func FlowWithName(name string) FlowOption {
	return func(workflow *types.Workflow) *types.Workflow {
		if &workflow.Jobs.DeployDev != nil {
			workflow.Jobs.DeployDev.With["releaseName"] = name
		}
		if &workflow.Jobs.DeployProd != nil {
			workflow.Jobs.DeployProd.With["releaseName"] = name
		}
		return workflow
	}
}

func FlowWithNamespace(namespace string) FlowOption {
	return func(workflow *types.Workflow) *types.Workflow {
		if &workflow.Jobs.DeployDev != nil {
			workflow.Jobs.DeployDev.With["namespace"] = namespace
		}
		if &workflow.Jobs.DeployProd != nil {
			workflow.Jobs.DeployProd.With["namespace"] = namespace
		}

		return workflow
	}
}
