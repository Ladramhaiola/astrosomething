package ecs

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

type BaseSystem struct {
	Entities map[Entity]struct{}
}

func (s *BaseSystem) Insert(entity Entity) {
	if s.Entities == nil {
		s.Entities = make(map[Entity]struct{})
	}
	s.Entities[entity] = struct{}{}
}

func (s *BaseSystem) Remove(entiry Entity) {
	if s.Entities == nil {
		s.Entities = make(map[Entity]struct{})
	}
	delete(s.Entities, entiry)
}

type System interface {
	Update(dt float64)
	Render(*ebiten.Image)

	Insert(Entity)
	Remove(Entity)
}

type SystemManager struct {
	// Map from system type string to a signature
	signatures map[string]Signature
	// Map from system type string to a system pointer
	systems map[string]System // TODO: use pointer or hack with types
}

func (sm *SystemManager) RegisterSystem(s System) {
	name := reflect.TypeOf(s).String()

	if _, ok := sm.systems[name]; ok {
		return
	}

	sm.systems[name] = s
}

func (sm *SystemManager) SetSystemSignature(s System, signature Signature) {
	name := reflect.TypeOf(s).String()
	sm.signatures[name] = signature
}

func (sm *SystemManager) EntityDestroyed(entity Entity) {
	for _, system := range sm.systems {
		// TODO: check signature
		system.Remove(entity)
	}
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
