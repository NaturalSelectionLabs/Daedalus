package types

type FlowOption func(workflow Workflow) Workflow

func FlowWithName(name string) FlowOption {
	return func(workflow Workflow) Workflow {
		if workflow.Jobs.DeployDev != nil {
			workflow.Jobs.DeployDev.With["releaseName"] = name
		}
		if workflow.Jobs.DeployProd != nil {
			workflow.Jobs.DeployProd.With["releaseName"] = name
		}
		return workflow
	}
}

func FlowWithNamespace(namespace string) FlowOption {
	return func(workflow Workflow) Workflow {
		if workflow.Jobs.DeployDev != nil {
			workflow.Jobs.DeployDev.With["namespace"] = namespace
		}
		if workflow.Jobs.DeployProd != nil {
			workflow.Jobs.DeployProd.With["namespace"] = namespace
		}

		return workflow
	}
}
