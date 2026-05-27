/*
Copyright 2024 Terway Authors.

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

package node

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type predicateForNodeEvent struct {
	predicate.Funcs

	supportEFLO        bool
	nodeLabelWhiteList map[string]string
}

// Create returns true if the Create event should be processed
func (p *predicateForNodeEvent) Create(e event.CreateEvent) bool {
	_ = "STUB: not implemented"
	return false
}

// Delete returns true if the Delete event should be processed
func (p *predicateForNodeEvent) Delete(e event.DeleteEvent) bool {
	_ = "STUB: not implemented"
	return false
}

// Update returns true if the Update event should be processed
func (p *predicateForNodeEvent) Update(e event.UpdateEvent) bool {
	_ = "STUB: not implemented"
	return false
}

// Generic returns true if the Generic event should be processed
func (p *predicateForNodeEvent) Generic(e event.GenericEvent) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *predicateForNodeEvent) predicateNode(o client.Object) bool {
	_ = "STUB: not implemented"
	return false
}
