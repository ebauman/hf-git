package v1alpha1

import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GitRepo struct {
	v1.TypeMeta
	v1.ObjectMeta

	Spec   GitRepoSpec
	Status GitRepoStatus
}

type GitRepoSpec struct {
	Repo string
	Revision string
	Branch string
	Secret string
}

type GitRepoStatus struct {
	Status string
	Revision string
	Messages []string
}
