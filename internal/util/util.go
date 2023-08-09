package util

import (
	"bytes"
	"github.com/samber/lo"
	"k8s.io/apimachinery/pkg/runtime"
	sjson "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"strings"
)

func ObjectsContext(objSlice []runtime.Object, sch *runtime.Scheme) string {
	otherYaml := lo.Map(objSlice, func(item runtime.Object, index int) string {
		var b bytes.Buffer
		s := sjson.NewSerializerWithOptions(
			sjson.DefaultMetaFactory, sch, sch, sjson.SerializerOptions{Yaml: true})
		s.Encode(item, &b)
		return b.String()
	})

	return strings.Join(otherYaml, "---\n")
}
