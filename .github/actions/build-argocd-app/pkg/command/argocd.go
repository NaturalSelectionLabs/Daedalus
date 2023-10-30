package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
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

type App struct {
	Name      string
	Namespace string
	Cluster   string
	Project   string
	Repo      string
	Revision  string
	Helm      HelmApp
	Kustomize KustomizeApp
	Image     Image
}

type Image struct {
	Name string
	Tag  string
}

type HelmApp struct {
	ValueFiles  []string
	ReleaseName string
	Chart       HelmChart
}

type HelmChart struct {
	RepoUrl string
	Name    string
	Version string
}

type KustomizeApp struct {
	Directory string
}

var sch = runtime.NewScheme()

func ValueFilePathRemovePrefix(inputPath string) string {
	if strings.HasPrefix(inputPath, "./") {
		// 去除 "./" 开头
		return strings.TrimPrefix(inputPath, "./")
	} else {
		// 字符串不以 "./" 开头，不进行修改
		return inputPath
	}
}

func (a *App) Build() *v1alpha1.Application {

	ref := v1alpha1.ApplicationSource{
		Ref:            "values",
		RepoURL:        a.Repo,
		TargetRevision: a.Revision,
	}

	helm := v1alpha1.ApplicationSource{
		Chart: a.Helm.Chart.Name,
		Helm: &v1alpha1.ApplicationSourceHelm{
			ReleaseName: a.Helm.ReleaseName,
			ValueFiles: lo.Map(a.Helm.ValueFiles, func(item string, index int) string {
				return fmt.Sprintf("$values/%s", ValueFilePathRemovePrefix(item))
			}),
			Parameters: []v1alpha1.HelmParameter{
				{
					Name:  "image.tag",
					Value: a.Image.Tag,
				},
				{
					Name:  "repoUrl",
					Value: a.Repo,
				},
			},
		},
		RepoURL:        a.Helm.Chart.RepoUrl,
		TargetRevision: a.Helm.Chart.Version,
	}

	kustomize := v1alpha1.ApplicationSource{
		RepoURL:        a.Repo,
		TargetRevision: a.Revision,
		Path:           a.Kustomize.Directory,
		Kustomize: &v1alpha1.ApplicationSourceKustomize{
			Images: v1alpha1.KustomizeImages{
				v1alpha1.KustomizeImage(fmt.Sprint(a.Image)),
			},
			CommonAnnotations: map[string]string{
				"github.com/url": a.Repo,
			},
		},
	}

	var applicationSources v1alpha1.ApplicationSources

	if len(a.Helm.ValueFiles) > 0 {
		applicationSources = append(applicationSources, helm)
	}

	if a.Kustomize.Directory != "" {
		applicationSources = append(applicationSources, kustomize)
	}

	applicationSources = append(applicationSources, ref)

	return &v1alpha1.Application{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      a.AppName(),
			Namespace: "guardian",
		},
		Spec: v1alpha1.ApplicationSpec{
			Destination: v1alpha1.ApplicationDestination{
				Namespace: a.Namespace,
				Name:      a.Cluster,
			},
			Project: a.Project,
			Sources: applicationSources,
		},
	}
}

func (a *App) Yaml() []byte {
	codec := serializer.NewCodecFactory(sch).LegacyCodec(v1alpha1.SchemeGroupVersion)
	jsonData, _ := runtime.Encode(codec, a.Build())
	data := make(map[string]interface{})
	_ = json.Unmarshal(jsonData, &data)

	delete((data["metadata"]).(map[string]interface{}), "creationTimestamp")
	delete(data, "status")

	var b bytes.Buffer
	encoder := yaml.NewEncoder(&b)
	encoder.SetIndent(2)
	_ = encoder.Encode(data)
	return b.Bytes()
}

func (i Image) String() string {
	return fmt.Sprintf("%s:%s", i.Name, i.Tag)
}

func (a *App) AppName() string {
	return fmt.Sprintf("%s-%s", a.Name, a.Cluster)
}

func init() {
	err := v1alpha1.AddToScheme(sch)
	if err != nil {
		panic(err)
	}
}

var argocdCmd = &cobra.Command{
	Use:   "argocd",
	Short: "Argo Deploy",
	Run: func(cmd *cobra.Command, args []string) {
		app := App{
			Name:      name,
			Namespace: namespace,
			Cluster:   cluster,
			Project:   getProject(),
			Repo:      repo,
			Revision:  "main",
			Helm: HelmApp{
				ReleaseName: getReleaseName(),
				Chart: HelmChart{
					RepoUrl: helmChartUrl,
					Name:    helmChartName,
					Version: helmChartVersion,
				},
				ValueFiles: helmValues,
			},
			Kustomize: KustomizeApp{},
			Image:     Image{},
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
}
