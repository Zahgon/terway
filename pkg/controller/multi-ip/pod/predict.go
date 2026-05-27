/*
Copyright 2022 Terway Authors.

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

package pod

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type predicateForPodEvent struct {
	predicate.Funcs
}

func (p *predicateForPodEvent) Create(e event.CreateEvent) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *predicateForPodEvent) Update(e event.UpdateEvent) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *predicateForPodEvent) Delete(e event.DeleteEvent) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *predicateForPodEvent) Generic(e event.GenericEvent) bool {
	_ = "STUB: not implemented"
	return false

	// needProcess filter pod which is ready to process
}

func needProcess(obj interface{}) bool { _ = "STUB: not implemented"; return false }
