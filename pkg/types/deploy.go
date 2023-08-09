package types

import (
	"bytes"
	"encoding/json"
	"gopkg.in/yaml.v3"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

type HelmValue struct {
	Image            Image                     `json:"image"`
	ImagePullSecrets []v1.LocalObjectReference `json:"imagePullSecrets,omitempty"`

	Workload Workload `json:"workload"`

	Env            []v1.EnvVar             `json:"env,omitempty"`
	EnvFrom        []v1.EnvFromSource      `json:"envFrom,omitempty"`
	Resources      v1.ResourceRequirements `json:"resources,omitempty"`
	LivenessProbe  v1.Probe                `json:"livenessProbe,omitempty"`
	ReadinessProbe v1.Probe                `json:"readinessProbe,omitempty"`

	PodAnnotations         map[string]string `json:"podAnnotations,omitempty"`
	*v1.PodSecurityContext `json:"podSecurityContext,omitempty" `
	*v1.SecurityContext    `json:"securityContext,omitempty"`
	*v1.Affinity           `json:"affinity,omitempty"`

	Tolerations  []v1.Toleration   `json:"tolerations,omitempty"`
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	Service        Service        `json:"service,omitempty"`
	IngressRoute   IngressRoute   `json:"ingressRoute,omitempty"`
	ServiceAccount ServiceAccount `json:"serviceAccout,omitempty"`
}

type Image struct {
	Repository string `json:"repository"`
	PullPolicy string `json:"pullPolicy,omitempty"`
	Tag        string `json:"tag,omitempty"`
}

type Workload struct {
	Name           string         `json:"name"`
	Type           string         `json:"type,omitempty"`
	InitContainers []v1.Container `json:"initContainers,omitempty"`
	Replicas       *int32         `json:"replicas,omitempty"`
	Command        []string       `json:"command,omitempty"`
	Args           []string       `json:"args,omitempty"`
	Deployment     Deployment     `json:"deployment,omitempty"`
}

type Deployment struct {
	ProgressDeadlineSeconds *int32                    `json:"progressDeadlineSeconds,omitempty"`
	RevisionHistoryLimit    *int32                    `json:"revisionHistoryLimit,omitempty"`
	Strategy                appsv1.DeploymentStrategy `json:"strategy,omitempty"`
}

type Service struct {
	Type v1.ServiceType `json:"type,omitempty"`
	Port int32          `json:"port,omitempty"`
}

type IngressRoute struct {
	Enabled     bool              `json:"enabled"`
	Annotations map[string]string `json:"annotations,omitempty"`
	EntryPoint  []string          `json:"entrypoint,omitempty"`
	Routes      []string          `json:"routes,omitempty"`
}

type ServiceAccount struct {
	Create      bool              `json:"create"`
	Annotations map[string]string `json:"annotations,omitempty"`
	Name        string            `json:"name"`
}

func (v HelmValue) String() string {
	var data map[string]interface{}

	jsonData, _ := json.Marshal(v)
	json.Unmarshal(jsonData, &data)

	var buffer bytes.Buffer
	encoder := yaml.NewEncoder(&buffer)
	encoder.SetIndent(2)
	encoder.Encode(&data)

	return buffer.String()
}
