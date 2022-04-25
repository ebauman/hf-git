/*

© 2022 The HobbyFarm Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ebauman/hf-git/apis/v1alpha1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1alpha1.AddToScheme)
}

type Interface interface {
	GitRepo() GitRepoController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) GitRepo() GitRepoController {
	return NewGitRepoController(schema.GroupVersionKind{Group: "hobbyfarm.io", Version: "v1alpha1", Kind: "GitRepo"}, "gitrepos", true, c.controllerFactory)
}
