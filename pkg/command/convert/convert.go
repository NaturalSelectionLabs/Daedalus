package convert

import (
	"fmt"
	"github.com/naturalSelectionLabs/daedalus/internal/converter"
	"github.com/naturalSelectionLabs/daedalus/internal/util"
	"github.com/spf13/cobra"
	"path"
)

var (
	gitDir     string
	actionPath string
	namespace  string
	name       string
)

var Cmd = &cobra.Command{
	Use:   "convert",
	Short: "convert actions",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a git repo path.")
			return
		}
		rootPath := args[0]
		workflow := handleSingleWorkflow(
			path.Join(rootPath, actionPath),
			converter.FlowWithName(name),
			converter.FlowWithNamespace(namespace),
		)

		fmt.Println("##############################")
		fmt.Println("##  Workflow")
		fmt.Println("## ", actionPath)
		fmt.Println("##############################")
		fmt.Println(workflow)

		image := workflow.Jobs.Build.With["images"]

		if &workflow.Jobs.DeployDev != nil {
			devPath := path.Join(rootPath, "deploy", "dev")
			devValues, devOthers := handleDeploy(devPath, converter.ValuesWithImageName(image))
			fmt.Println("##############################")
			fmt.Println("## Dev Deploy")
			fmt.Println("##############################")
			fmt.Println("#### Values")
			fmt.Println("#### ", path.Join(devPath, "values.yaml"))
			fmt.Println("##############################")
			fmt.Println(devValues)
			fmt.Println("##############################")
			fmt.Println("#### Others")
			fmt.Println("##############################")
			fmt.Println(util.ObjectsContext(devOthers, sch))
		}

		if &workflow.Jobs.DeployProd != nil {
			prodPath := path.Join(rootPath, "deploy", "prod")
			prodValues, prodOthers := handleDeploy(prodPath, converter.ValuesWithImageName(image))
			fmt.Println("##############################")
			fmt.Println("## Prod Deploy")
			fmt.Println("##############################")
			fmt.Println("#### Values")
			fmt.Println("#### ", path.Join(prodPath, "values.yaml"))
			fmt.Println("##############################")
			fmt.Println(prodValues)
			fmt.Println("##############################")
			fmt.Println("#### Others")
			fmt.Println("##############################")
			fmt.Println(util.ObjectsContext(prodOthers, sch))
		}

		//
		//values, others := handleDeploy(path.Join(rootPath, "deploy", "prod"))
	},
}

func init() {
	Cmd.Flags().StringVarP(&actionPath, "action", "a", ".github/workflows/docker-build-push.yml", "Github Action path")
	Cmd.Flags().StringVarP(&name, "name", "", "<name>", "app name")
	Cmd.Flags().StringVarP(&namespace, "namespace", "n", "<namespace>", "namespace")

	Cmd.AddCommand(actionCmd, deployCmd)
}
