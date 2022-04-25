package main

import (
	"github.com/ebauman/hf-git/apis/v1alpha1"
	controllergen "github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"
	v1 "k8s.io/api/core/v1"
)

func main() {
	controllergen.Run(args.Options{
		OutputPackage: "github.com/ebauman/hf-git/generated",
		Boilerplate: "hack/boilerplate.txt",
		Groups: map[string]args.Group{
			"hobbyfarm.io": {
				Types: []interface{}{
					v1alpha1.GitRepo{},
				},
				GenerateTypes: true,
			},
			"": {
				Types: []interface{}{
					v1.Secret{},
				},
				ClientSetPackage: "k8s.io/client-go/kubernetes",
			},
		},
	})
}