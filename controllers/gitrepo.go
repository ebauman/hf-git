package controllers

import (
	"github.com/ebauman/hf-git/apis/v1alpha1"
	grc "github.com/ebauman/hf-git/generated/controllers/hobbyfarm.io/v1alpha1"
	"k8s.io/client-go/kubernetes"
)

type Handler struct {
	kclient kubernetes.Clientset
	gitRepos grc.GitRepoClient
}

func (h *Handler) OnGitRepoChanged(key string, repo *v1alpha1.GitRepo) (*v1alpha1.GitRepo, error) {
	if repo == nil {
		return nil, nil
	}


}