package converter

import (
	"github.com/naturalSelectionLabs/daedalus/pkg/types"
	"github.com/samber/lo"
	"github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefikcontainous/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ValuesOption func(value *types.HelmValue) *types.HelmValue

func ObjectsToValue(objSlice []runtime.Object, valuesOptions ...ValuesOption) types.HelmValue {

	name := getAppName(objSlice)

	deployment := getResource(
		objSlice,
		func(item *appsv1.Deployment) bool {
			return item.Name == name
		},
		func(gvk schema.GroupVersionKind) bool {
			return gvk.Kind == "Deployment"
		},
	)

	service := getResource(
		objSlice,
		func(item *v1.Service) bool {
			return item.Name == name
		},
		func(gvk schema.GroupVersionKind) bool {
			return gvk.Kind == "Service"
		},
	)

	ingressRoute := getResource(
		objSlice,
		func(item *v1alpha1.IngressRoute) bool {
			return item.Name == name
		},
		func(gvk schema.GroupVersionKind) bool {
			return gvk.Kind == "IngressRoute"
		},
	)

	sa := getResource(
		objSlice,
		func(item *v1.ServiceAccount) bool {
			return true
		},
		func(gvk schema.GroupVersionKind) bool {
			return gvk.Kind == "ServiceAccount"
		},
	)
	pod := deployment.Spec.Template
	container := pod.Spec.Containers[0]

	values := types.HelmValue{
		Image: types.Image{
			Repository: container.Image,
		},

		Workload: types.Workload{
			Name:           name,
			Type:           "Deployment",
			InitContainers: pod.Spec.InitContainers,
			Replicas:       deployment.Spec.Replicas,
			Command:        container.Command,
			Args:           container.Args,
			Deployment: types.Deployment{
				ProgressDeadlineSeconds: deployment.Spec.ProgressDeadlineSeconds,
				RevisionHistoryLimit:    deployment.Spec.RevisionHistoryLimit,
				Strategy:                deployment.Spec.Strategy,
			},
		},

		Affinity:           pod.Spec.Affinity,
		Env:                container.Env,
		EnvFrom:            container.EnvFrom,
		Resources:          container.Resources,
		ReadinessProbe:     *container.ReadinessProbe,
		LivenessProbe:      *container.LivenessProbe,
		Tolerations:        pod.Spec.Tolerations,
		PodAnnotations:     pod.Annotations,
		PodSecurityContext: pod.Spec.SecurityContext,
		NodeSelector:       pod.Spec.NodeSelector,

		Service: types.Service{
			Type: service.Spec.Type,
			Port: service.Spec.Ports[0].Port,
		},
	}

	if ingressRoute != nil && ingressRoute.Spec.Routes[0].Middlewares == nil {
		values.IngressRoute = types.IngressRoute{
			Enabled:     true,
			Annotations: ingressRoute.Annotations,
			EntryPoint:  ingressRoute.Spec.EntryPoints,
			Routes: lo.Map(ingressRoute.Spec.Routes, func(item v1alpha1.Route, index int) string {
				return item.Match
			}),
		}
	} else {
		values.IngressRoute = types.IngressRoute{
			Enabled: false,
		}
	}
	var _sa = types.ServiceAccount{
		Create: true,
	}

	if sa != nil && sa.Name != name {
		_sa.Annotations = sa.Annotations
		_sa.Name = sa.Name
	}
	values.ServiceAccount = _sa

	for _, option := range valuesOptions {
		option(&values)
	}

	return values

}

func ExcludeObjects(objSlice []runtime.Object, values types.HelmValue) []runtime.Object {

	name := getAppName(objSlice)

	return lo.Filter(objSlice, func(item runtime.Object, index int) bool {
		accessor, _ := meta.Accessor(item)

		gvk := item.GetObjectKind().GroupVersionKind()

		if gvk.Kind == "ServiceAccount" || gvk.Kind == "Secret" {
			return false
		}

		if accessor.GetName() == name {

			if gvk.Kind == "Deployment" || gvk.Kind == "Service" {
				return false
			}

			if values.IngressRoute.Enabled && gvk.Kind == "IngressRoute" {
				return false
			}
		}

		return true
	})

}

func getAppName(objSlice []runtime.Object) string {
	namesMap := lo.CountValuesBy(objSlice, func(item runtime.Object) string {
		accessor, _ := meta.Accessor(item)
		return accessor.GetName()
	})
	var maxKey string
	maxValue := -1
	for k, v := range namesMap {
		if v > maxValue {
			maxKey = k
			maxValue = v
		}
	}
	return maxKey
}

func getResource[T runtime.Object](objSlice []runtime.Object, predicate func(item T) bool, matchGVK func(gvk schema.GroupVersionKind) bool) T {
	var result T
	for _, obj := range objSlice {
		if matchGVK(obj.GetObjectKind().GroupVersionKind()) {
			res := obj.(T)
			if predicate(res) {
				return res
			}
		}
	}
	return result
}

func ValuesWithImageName(name string) ValuesOption {
	return func(value *types.HelmValue) *types.HelmValue {
		value.Image.Repository = name
		return value
	}
}
