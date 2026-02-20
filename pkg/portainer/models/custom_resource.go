package models

// CustomResourceDefinition represents a Kubernetes Custom Resource Definition.
type CustomResourceDefinition struct {
	Name             string `json:"name"`
	Group            string `json:"group"`
	Scope            string `json:"scope"`
	CreationDate     string `json:"creationDate"`
	ReleaseName      string `json:"releaseName,omitempty"`
	ReleaseNamespace string `json:"releaseNamespace,omitempty"`
	ReleaseVersion   string `json:"releaseVersion,omitempty"`
}

// CustomResource represents a Kubernetes Custom Resource instance.
type CustomResource struct {
	Name           string `json:"name"`
	Namespace      string `json:"namespace,omitempty"`
	DefinitionName string `json:"definitionName"`
	UID            string `json:"uid"`
	CreationDate   string `json:"creationDate"`
}
