package argocd

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestValueFilePathRemovePrefix(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Name  string
		Input string
		Want  string
	}{
		{
			Name:  "Prefix with ./",
			Input: "./path/to/something",
			Want:  "path/to/something",
		},
		{
			Name:  "Prefix without ./",
			Input: "another/path",
			Want:  "another/path",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := ValueFilePathRemovePrefix(testCase.Input)
			assert.Equal(t,
				result,
				testCase.Want,
				fmt.Sprintf("Expected %s, but got %s", testCase.Want, result),
			)
		})
	}

}

func TestBuildApplication(t *testing.T) {
	t.Parallel()
	// 创建一个示例 App 对象以进行测试
	cases := []struct {
		Name string
		App  App
		Dest string
	}{
		{
			Name: "Full Object Field Test",
			App: App{
				Name:      "my-app",
				Namespace: "my-namespace",
				Cluster:   "my-cluster",
				Project:   "my-project",
				Repo:      "https://github.com/my/repo",
				Revision:  "main",
				Helm: HelmApp{
					ValueFiles:  []string{"./values.yaml"},
					ReleaseName: "my-release",
					Chart: HelmChart{
						RepoUrl: "https://github.com/helm/charts",
						Name:    "my-chart",
						Version: "1.0.0",
					},
				},
				Kustomize: KustomizeApp{
					Directory: "kustomize-dir",
				},
				Image: Image{
					Name: "my-image",
					Tag:  "latest",
				},
			},
			Dest: `apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: my-app-my-cluster
  namespace: guardian
spec:
  destination:
    name: my-cluster
    namespace: my-namespace
  project: my-project
  sources:
    - chart: my-chart
      helm:
        parameters:
          - name: image.tag
            value: latest
          - name: repoUrl
            value: https://github.com/my/repo
        releaseName: my-release
        valueFiles:
          - $values/values.yaml
      repoURL: https://github.com/helm/charts
      targetRevision: 1.0.0
    - kustomize:
        commonAnnotations:
          github.com/url: https://github.com/my/repo
        images:
          - my-image:latest
      path: kustomize-dir
      repoURL: https://github.com/my/repo
      targetRevision: main
    - ref: values
      repoURL: https://github.com/my/repo
      targetRevision: main
`,
		},
		{
			Name: "Helm Only Field Test",
			App: App{
				Name:      "my-app",
				Namespace: "my-namespace",
				Cluster:   "my-cluster",
				Project:   "my-project",
				Repo:      "https://github.com/my/repo",
				Revision:  "main",
				Helm: HelmApp{
					ValueFiles:  []string{"./values.yaml"},
					ReleaseName: "my-release",
					Chart: HelmChart{
						RepoUrl: "https://github.com/helm/charts",
						Name:    "my-chart",
						Version: "1.0.0",
					},
				},
				Image: Image{
					Name: "my-image",
					Tag:  "latest",
				},
			},
			Dest: `apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: my-app-my-cluster
  namespace: guardian
spec:
  destination:
    name: my-cluster
    namespace: my-namespace
  project: my-project
  sources:
    - chart: my-chart
      helm:
        parameters:
          - name: image.tag
            value: latest
          - name: repoUrl
            value: https://github.com/my/repo
        releaseName: my-release
        valueFiles:
          - $values/values.yaml
      repoURL: https://github.com/helm/charts
      targetRevision: 1.0.0
    - ref: values
      repoURL: https://github.com/my/repo
      targetRevision: main
`,
		},
		{
			Name: "Kustomize Only Field Test",
			App: App{
				Name:      "my-app",
				Namespace: "my-namespace",
				Cluster:   "my-cluster",
				Project:   "my-project",
				Repo:      "https://github.com/my/repo",
				Revision:  "main",
				Kustomize: KustomizeApp{
					Directory: "kustomize-dir",
				},
				Image: Image{
					Name: "my-image",
					Tag:  "latest",
				},
			},
			Dest: `apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: my-app-my-cluster
  namespace: guardian
spec:
  destination:
    name: my-cluster
    namespace: my-namespace
  project: my-project
  sources:
    - kustomize:
        commonAnnotations:
          github.com/url: https://github.com/my/repo
        images:
          - my-image:latest
      path: kustomize-dir
      repoURL: https://github.com/my/repo
      targetRevision: main
    - ref: values
      repoURL: https://github.com/my/repo
      targetRevision: main
`,
		},
	}

	for _, testCase := range cases {

		t.Run(testCase.Name, func(t *testing.T) {
			assert.Equal(t, string(testCase.App.Yaml()), testCase.Dest)
		})
	}

	// 添加适当的测试断言，例如检查返回的对象是否符合预期
	//if appDefinition.Name != "my-app" {
	//	t.Errorf("Expected app name to be 'my-app', but got '%s'", appDefinition.Name)
	//}

	// 添加其他断言以验证返回的对象的其他属性
	// ...

}
