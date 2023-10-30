package types

type ConvertData struct {
	Workflow    *Workflow `json:"workflow"`
	DeployMap   map[string]DeployContent
	DevContent  *DeployContent `json:"dev"`
	ProdContent *DeployContent `json:"prod"`
}

type ConvertOptions struct {
	Name            string
	Namespace       string
	WorkflowContent []byte
}
