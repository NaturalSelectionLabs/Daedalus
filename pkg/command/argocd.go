package command

import (
	"fmt"
	"github.com/naturalSelectionLabs/daedalus/internal/argocd"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	name      string
	namespace string
	cluster   string
	project   string
	repo      string

	helmValues       []string
	helmRelease      string
	helmChartUrl     string
	helmChartName    string
	helmChartVersion string
)

var defaultEnv = map[string]string{
	"NAME":         "myapp",
	"NAMESPACE":    "default",
	"CLUSTER":      "prod",
	"REPO":         "https://github.com/my/repo",
	"HELM_VALUES":  "",
	"HELM_RELEASE": "",
}

var argocdCmd = &cobra.Command{
	Use:   "argocd",
	Short: "Argo Deploy",
	Run: func(cmd *cobra.Command, args []string) {
		app := argocd.App{
			Name:      name,
			Namespace: namespace,
			Cluster:   cluster,
			Project:   getProject(),
			Repo:      repo,
			Revision:  "main",
			Helm: argocd.HelmApp{
				ReleaseName: getReleaseName(),
				Chart: argocd.HelmChart{
					RepoUrl: helmChartUrl,
					Name:    helmChartName,
					Version: helmChartVersion,
				},
				ValueFiles: helmValues,
			},
			Kustomize: argocd.KustomizeApp{},
			Image:     argocd.Image{},
		}

		appString := string(app.Yaml())
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), appString)
		_ = appendToEnv("GITHUB_OUTPUT", "app", appString)
	},
}

func getProject() string {
	if project != "" {
		return project
	}
	return namespace
}

func getReleaseName() string {
	if helmRelease != "" {
		return helmRelease
	}
	return name
}

func appendToEnv(env, key, content string) error {
	currentEnv := os.Getenv(env)
	updateEnv := currentEnv + fmt.Sprintf("%s=%s\n", key, content)
	return os.Setenv(env, updateEnv)
}

func init() {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvPrefix("DAEDALUS")
	v.MustBindEnv(lo.MapToSlice(defaultEnv, func(key string, value string) string {
		return key
	})...)
	for _k, _v := range defaultEnv {
		v.SetDefault(_k, _v)
	}

	argocdCmd.Flags().StringVarP(&name, "name", "", v.GetString("NAME"), "App name")
	argocdCmd.Flags().StringVarP(&namespace, "namespace", "n", v.GetString("NAMESPACE"), "App namespace")
	argocdCmd.Flags().StringVarP(&cluster, "cluster", "", v.GetString("CLUSTER"), "App cluster")
	argocdCmd.Flags().StringVarP(&project, "project", "p", v.GetString("PROJECT"), "App project")

	argocdCmd.Flags().StringSliceVarP(&helmValues,
		"helm-values", "",
		lo.Filter(strings.Split(v.GetString("HELM_VALUES"), "\n"), func(item string, index int) bool {
			return item != ""
		}),
		"helm values path")
	argocdCmd.Flags().StringVarP(&helmRelease,
		"helm-release", "", v.GetString("HELM_RELEASE"), "Helm release name")
	//server.InitGithubClient(token)
}
