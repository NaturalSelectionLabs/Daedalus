package convert

import (
	"fmt"
	"github.com/naturalSelectionLabs/daedalus/internal/converter"
	"github.com/naturalSelectionLabs/daedalus/internal/util"
	"github.com/naturalSelectionLabs/daedalus/pkg/types"
	"github.com/spf13/cobra"
	"github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefikcontainous/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"os"
	"path/filepath"
	"strings"
)

var (
	imageName string
	sch       *runtime.Scheme
)

type HandleResourceFunc func(context string) (runtime.Object, *schema.GroupVersionKind, error)

type HandleFileFunc func(filePath string, handleResource HandleResourceFunc) []runtime.Object

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy functions",
	Run: func(cmd *cobra.Command, args []string) {
		if oldFile == "" {
			fmt.Println("Please provide a configuration file using -f flag.")
			return
		}
		values, others := handleDeploy(oldFile)

		fmt.Println(values)

		fmt.Println(util.ObjectsContext(others, sch))
	},
}

func init() {
	deployCmd.Flags().StringVarP(&oldFile, "file", "f", "", "from old file")
	deployCmd.Flags().StringVarP(&imageName, "image-name", "I", "", "override image name")

	sch = runtime.NewScheme()

	v1alpha1.AddToScheme(sch)
	v1.AddToScheme(sch)
	appsv1.AddToScheme(sch)
}

func handleDeploy(path string, options ...converter.ValuesOption) (*types.HelmValue, []runtime.Object) {

	fileInfo, err := os.Stat(path)

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	var objSlice []runtime.Object
	if fileInfo.IsDir() {
		objSlice = ObjectsFromDirectory(path, ObjectsFromFile)

	} else {
		objSlice = ObjectsFromFile(path, ObjectFromString)
	}

	values := converter.ObjectsToValue(objSlice, options...)

	others := converter.ExcludeObjects(objSlice, values)

	return &values, others
}

func ObjectsFromDirectory(dirPath string, handleFile HandleFileFunc) (objSlice []runtime.Object) {

	objSlice = []runtime.Object{}

	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("Error reading the directory: %s\n", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := filepath.Join(dirPath, file.Name())
			objSlice = append(objSlice, handleFile(filePath, ObjectFromString)...)
		}
	}
	return
}

func ObjectsFromFile(filePath string, handleResource HandleResourceFunc) (objSlice []runtime.Object) {
	objSlice = []runtime.Object{}
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading the configuration file: %s\n", err)
		return
	}

	yamlContexts := strings.Split(string(file), "---")
	for _, context := range yamlContexts {
		if obj, _, err := handleResource(context); err == nil {
			objSlice = append(objSlice, obj)
		}
	}
	return
}

func ObjectFromString(context string) (runtime.Object, *schema.GroupVersionKind, error) {
	return serializer.NewCodecFactory(sch).UniversalDeserializer().Decode(
		[]byte(context), nil, nil)

}
