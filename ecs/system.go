package ecs

import (
	"reflect"
)

type System interface {
	Update(*Coordinator)
	Insert(Entity)
	Remove(Entity)
}

type SystemManager struct {
	// Map from system type string to a signature
	signatures map[string]Signature
	// Map from system type string to a system pointer
	systems map[string]System // TODO: use pointer or hack with types
}

func registerSystem(sm *SystemManager, s System) {
	name := reflect.TypeOf(s).String()

	if _, ok := sm.systems[name]; ok {
		return
	}

	sm.systems[name] = s
}

func setSignature[T System](sm *SystemManager, signature Signature) {
	name := reflect.TypeOf((*T)(nil)).String()
	sm.signatures[name] = signature
}

func (sm *SystemManager) EntitySignatureChanged(entity Entity, entitySignature Signature) {
	for name, system := range sm.systems {
		signature := sm.signatures[name]

		if (entitySignature & signature) == signature {
			system.Insert(entity)
		} else {
			system.Remove(entity)
		}
	}
}

func NewSystemManager() *SystemManager {
	return &SystemManager{
		signatures: make(map[string]Signature),
		systems:    make(map[string]System),
	}
}
